package botsfwmodels

import (
	"time"
)

// ChatData provides data about bot chat
type ChatData interface {
	Base() *ChatBaseData // TODO: Document why this is needed or remove

	GetBotID() (botID string)
	SetBotID(botID string)

	//GetChatID() (chatID string)
	//SetChatID(chatID string)

	// GetAppUserID returns app user ID
	GetAppUserID() (appUserID string)
	SetAppUserID(appUserID string)

	AddClientLanguage(languageCode string) (changed bool)

	/*
		GetBotUserIntID() int
		GetBotUserStringID() string
	*/

	SetBotUserID(id interface{})
	SetIsGroupChat(bool)

	IsAccessGranted() bool
	IsGroupChat() bool
	SetAccessGranted(value bool) bool

	GetPreferredLanguage() string
	SetPreferredLanguage(value string)

	SetUpdatedTime(time.Time) // github.com/strongo/user.UpdatedTimeSetter

	SetDtLastInteraction(time time.Time)

	GetAwaitingReplyTo() string
	SetAwaitingReplyTo(path string)
	IsAwaitingReplyTo(code string) bool
	AddWizardParam(key, value string)
	GetWizardParam(key string) string
	PopStepsFromAwaitingReplyUpToSpecificParent(code string)
	PushStepToAwaitingReplyTo(code string)
	//GetGaClientID() string
}

// NewChatID create a new bot chat ID, returns string
func NewChatID(botID, botChatID string) string {
	if botID == "" {
		panic("botID is empty")
	}
	if botChatID == "" {
		panic("botChatID is empty")
	}
	return botID + ":" + botChatID
}
