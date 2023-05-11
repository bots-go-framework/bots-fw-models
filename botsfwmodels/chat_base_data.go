package botsfwmodels

import (
	"fmt"
	"time"
)

// BotChatBaseData hold common properties for bot chat entities not specific to any platform
type BotChatBaseData struct {
	BotID  string `datastore:",noindex"`
	ChatID string `datastore:",noindex"` // New field that is not populated before
	BotBaseData
	AppUserIntIDs []int64
	//
	IsGroup bool   `datastore:",noindex,omitempty"`
	Type    string `datastore:",noindex,omitempty"`
	Title   string `datastore:",noindex,omitempty"`
	//
	chatState
	//
	PreferredLanguage string    `datastore:",noindex,omitempty"`
	GaClientID        []byte    `datastore:",noindex,omitempty"`
	DtLastInteraction time.Time `datastore:",omitempty"`
	InteractionsCount int       `datastore:",omitempty"`
	DtForbidden       time.Time `datastore:",omitempty"`
	DtForbiddenLast   time.Time `datastore:",noindex,omitempty"`
	LanguageCodes     []string  `datastore:",noindex"` // UI languages
}

var _ BotChatData = (*BotChatBaseData)(nil)

func (e *BotChatBaseData) Base() *BotChatBaseData {
	return e
}

// GetBotID returns bot ID
func (e *BotChatBaseData) GetBotID() string {
	return e.BotID
}

// SetBotID sets bot ID
func (e *BotChatBaseData) SetBotID(botID string) {
	e.BotID = botID
}

// GetChatID returns chat ID
func (e *BotChatBaseData) GetChatID() string {
	return e.ChatID
}

// SetChatID sets chat ID
func (e *BotChatBaseData) SetChatID(chatID string) {
	e.ChatID = chatID
}

// IsGroupChat indicates if it is a group chat
func (e *BotChatBaseData) IsGroupChat() bool {
	return e.IsGroup
}

// SetIsGroupChat marks chat as a group one
func (e *BotChatBaseData) SetIsGroupChat(v bool) {
	e.IsGroup = v
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

// GetPreferredLanguage returns preferred language
func (e *BotChatBaseData) GetPreferredLanguage() string {
	return e.PreferredLanguage
}

// SetPreferredLanguage sets preferred language
func (e *BotChatBaseData) SetPreferredLanguage(value string) {
	e.PreferredLanguage = value
}
