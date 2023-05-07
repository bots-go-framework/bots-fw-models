package botsfwmodels

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// BotChatBaseData hold common properties for bot chat entities not specific to any platform
type BotChatBaseData struct {
	BotBaseData
	AppUserIntIDs []int64
	BotID         string `datastore:",noindex"`
	//
	IsGroup bool   `datastore:",noindex,omitempty"`
	Type    string `datastore:",noindex,omitempty"`
	Title   string `datastore:",noindex,omitempty"`
	//
	AwaitingReplyTo   string    `datastore:",noindex,omitempty"`
	PreferredLanguage string    `datastore:",noindex,omitempty"`
	GaClientID        []byte    `datastore:",noindex,omitempty"`
	DtLastInteraction time.Time `datastore:",omitempty"`
	InteractionsCount int       `datastore:",omitempty"`
	DtForbidden       time.Time `datastore:",omitempty"`
	DtForbiddenLast   time.Time `datastore:",noindex,omitempty"`
	LanguageCodes     []string  `datastore:",noindex"` // UI languages
}

var _ BotChat = (*BotChatBaseData)(nil)

func (e *BotChatBaseData) Base() *BotChatBaseData {
	return e
}

// GetBotID returns bot ID
func (e *BotChatBaseData) GetBotID() string {
	return e.BotID
}

// IsGroupChat indicates if it is a group chat
func (e *BotChatBaseData) IsGroupChat() bool {
	return e.IsGroup
}

// SetIsGroupChat marks chat as a group one
func (e *BotChatBaseData) SetIsGroupChat(v bool) {
	e.IsGroup = v
}

// SetBotID sets bot ID
func (e *BotChatBaseData) SetBotID(botID string) {
	e.BotID = botID
}

// AddClientLanguage adds client UI language
func (e *BotChatBaseData) AddClientLanguage(languageCode string) (changed bool) {
	if languageCode == "" || languageCode == "root" {
		return false
	}
	for _, lc := range e.LanguageCodes {
		if lc == languageCode {
			return false
		}
	}
	e.LanguageCodes = append(e.LanguageCodes, languageCode)
	return false
}

// func (e *BotChatBaseData) GetBotUserIntID() int {
// 	panic("Should be overwritten in subclass")
// }
//
// func (e *BotChatBaseData) GetBotUserStringID() string {
// 	panic("Should be overwritten in subclass")
// }

// SetBotUserID sets bot user ID
func (e *BotChatBaseData) SetBotUserID(id interface{}) {
	panic(fmt.Sprintf("Should be overwritten in subclass, got: %T=%v", id, id))
}

// SetDtLastInteraction sets date time of last interaction
func (e *BotChatBaseData) SetDtLastInteraction(v time.Time) {
	e.DtLastInteraction = v
	e.InteractionsCount++
}

// GetGaClientID returns Google Analytics client UUID
// TODO: random implementation should not be here in this module so we do not have dep on random?
//func (e *BotChatBaseData) GetGaClientID() string {
//	if len(e.GaClientID) == 0 {
//		e.GaClientID = []byte(random.ID(32))
//	}
//	return string(e.GaClientID)
//}

// SetDtUpdateToNow mark entity updated with now
func (e *BotChatBaseData) SetDtUpdateToNow() {
	e.DtUpdated = time.Now()
}

// GetAwaitingReplyTo returns current state
func (e *BotChatBaseData) GetAwaitingReplyTo() string {
	return e.AwaitingReplyTo
}

// SetAwaitingReplyTo sets current state
func (e *BotChatBaseData) SetAwaitingReplyTo(value string) {
	e.AwaitingReplyTo = strings.TrimLeft(value, "/")
}

// GetPreferredLanguage returns preferred language
func (e *BotChatBaseData) GetPreferredLanguage() string {
	return e.PreferredLanguage
}

// SetPreferredLanguage sets preferred language
func (e *BotChatBaseData) SetPreferredLanguage(value string) {
	e.PreferredLanguage = value
}

// IsAwaitingReplyTo returns true if bot us awaiting reply to a specific command
func (e *BotChatBaseData) IsAwaitingReplyTo(code string) bool {
	awaitingReplyToPath := e.getAwaitingReplyToPath()
	return awaitingReplyToPath == code || strings.HasSuffix(awaitingReplyToPath, AwaitingReplyToPathSeparator+code)
}

func (e *BotChatBaseData) getAwaitingReplyToPath() string {
	pathAndQuery := strings.SplitN(e.AwaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	if len(pathAndQuery) > 1 {
		return pathAndQuery[0]
	}
	return e.AwaitingReplyTo
}

// PopStepsFromAwaitingReplyUpToSpecificParent go back in state
func (e *BotChatBaseData) PopStepsFromAwaitingReplyUpToSpecificParent(step string) {
	awaitingReplyTo := e.AwaitingReplyTo
	pathAndQuery := strings.SplitN(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	path := pathAndQuery[0]
	steps := strings.Split(path, AwaitingReplyToPathSeparator)
	for i := len(steps) - 1; i >= 0; i-- {
		if steps[i] == step {
			if i < len(steps)-1 {
				path = strings.Join(steps[:i+1], AwaitingReplyToPathSeparator)
				if len(pathAndQuery) > 1 {
					query := pathAndQuery[1]
					e.SetAwaitingReplyTo(path + AwaitingReplyToPath2QuerySeparator + query)
				} else {
					e.SetAwaitingReplyTo(path)
				}
			}
			//steps = steps[:i]
			break
			// } else {
			// log.Infof(c, "steps[%v]: %v != %v:", i, steps[i], step)
		}
	}
}

// PushStepToAwaitingReplyTo go down in state
func (e *BotChatBaseData) PushStepToAwaitingReplyTo(step string) {
	awaitingReplyTo := e.AwaitingReplyTo
	pathAndQuery := strings.SplitN(awaitingReplyTo, AwaitingReplyToPath2QuerySeparator, 2)
	if len(pathAndQuery) > 1 { // Has query part - something after "?" character
		if !e.IsAwaitingReplyTo(step) {
			path := pathAndQuery[0]
			query := pathAndQuery[1]
			awaitingReplyTo = strings.Join([]string{path, AwaitingReplyToPathSeparator, step, AwaitingReplyToPath2QuerySeparator, query}, "")
			e.SetAwaitingReplyTo(awaitingReplyTo)
		}
	} else { // Has no query - no "?" character
		if !e.IsAwaitingReplyTo(step) {
			awaitingReplyTo = awaitingReplyTo + AwaitingReplyToPathSeparator + step
			e.SetAwaitingReplyTo(awaitingReplyTo)
		}
	}
}

// AddWizardParam adds context param to state
func (e *BotChatBaseData) AddWizardParam(key, value string) {
	awaitingReplyTo := e.GetAwaitingReplyTo()
	awaitingURL, err := url.Parse(awaitingReplyTo)
	if err != nil {
		panic(fmt.Sprintf("Failed to call url.Parse(awaitingReplyTo=%v)", awaitingReplyTo))
	}
	query := awaitingURL.Query()
	query.Set(key, value)
	awaitingURL.RawQuery = query.Encode()
	e.SetAwaitingReplyTo(awaitingURL.String())
}

// GetWizardParam returns state param value
func (e *BotChatBaseData) GetWizardParam(key string) string {
	u, err := url.Parse(e.GetAwaitingReplyTo())
	if err != nil {
		return ""
	}
	return u.Query().Get(key)
}
