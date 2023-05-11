package botsfwmodels

import (
	"fmt"
	"time"
)

// ChatBaseData hold common properties for bot chat entities not specific to any platform
type ChatBaseData struct {
	ChatKey // Have it for convenience

	BotBaseData
	chatState
	chatSettings

	// AppUserIntIDs is kept for legacy reasons
	// Deprecated: replace with `AppUserIDs []string`
	AppUserIntIDs []int64 // Legacy

	// IsGroup indicates if bot is added/used in a group chat
	IsGroup bool `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",noindex,omitempty"`

	// Type - TODO: document what is it
	Type string `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",noindex,omitempty"`

	// Title stores a title of a chat if bot platforms supports named chats
	Title string `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",noindex,omitempty"`

	// GAClientID is Google Analytics client ID
	// Deprecated: use GAClientIDs AnalyticsClientIDs
	GaClientID []byte `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",noindex,omitempty"`
	// AnalyticsClientIDs stores IDs of analytics clients. For example {"GA": "1234567890.1234567890"}
	AnalyticsClientIDs map[string]string `dalgo:",noindex,omitempty" datastore:",noindex,omitempty" firestore:",noindex,omitempty"`

	DtLastInteraction time.Time `dalgo:",omitempty" datastore:",omitempty" firestore:",omitempty"`
	InteractionsCount int       `dalgo:",omitempty" datastore:",omitempty" firestore:",omitempty"`
	DtForbidden       time.Time `dalgo:",omitempty" datastore:",omitempty" firestore:",omitempty"`
	DtForbiddenLast   time.Time `dalgo:",omitempty" datastore:",omitempty" firestore:",omitempty"`
}

var _ ChatData = (*ChatBaseData)(nil)

func (e *ChatBaseData) Base() *ChatBaseData {
	return e
}

// GetBotID returns bot ID
func (e *ChatBaseData) GetBotID() string {
	return e.BotID
}

// SetBotID sets bot ID
func (e *ChatBaseData) SetBotID(botID string) {
	e.BotID = botID
}

// GetChatID returns chat ID
func (e *ChatBaseData) GetChatID() string {
	return e.ChatID
}

// SetChatID sets chat ID
func (e *ChatBaseData) SetChatID(chatID string) {
	e.ChatID = chatID
}

// IsGroupChat indicates if it is a group chat
func (e *ChatBaseData) IsGroupChat() bool {
	return e.IsGroup
}

// SetIsGroupChat marks chat as a group one
func (e *ChatBaseData) SetIsGroupChat(v bool) {
	e.IsGroup = v
}

// func (e *ChatBaseData) GetBotUserIntID() int {
// 	panic("Should be overwritten in subclass")
// }
//
// func (e *ChatBaseData) GetBotUserStringID() string {
// 	panic("Should be overwritten in subclass")
// }

// SetBotUserID sets bot user ID
func (e *ChatBaseData) SetBotUserID(id interface{}) {
	panic(fmt.Sprintf("Should be overwritten in subclass, got: %T=%v", id, id))
}

// SetDtLastInteraction sets date time of last interaction
func (e *ChatBaseData) SetDtLastInteraction(v time.Time) {
	e.DtLastInteraction = v
	e.InteractionsCount++
}

// GetGaClientID returns Google Analytics client UUID
// TODO: random implementation should not be here in this module so we do not have dep on random?
//func (e *ChatBaseData) GetGaClientID() string {
//	if len(e.GaClientID) == 0 {
//		e.GaClientID = []byte(random.ID(32))
//	}
//	return string(e.GaClientID)
//}

// SetDtUpdateToNow mark entity updated with now
func (e *ChatBaseData) SetDtUpdateToNow() {
	e.DtUpdated = time.Now()
}
