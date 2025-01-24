package botsfwmodels

import (
	"time"
)

// BotChatData provides data about bot chat
type BotChatData interface {
	// Base returns a base struct that should be included in all structs that implement BotChatData.
	Base() *ChatBaseData // TODO: Document why this is needed or remove

	// IsChanged returns true if data has been changed
	IsChanged() bool

	// GetBotID returns bot ID
	//GetBotID() (botID string)

	//SetBotID(botID string)

	// GetChatID returns chat ID
	//GetChatID() (chatID string)
	//SetChatID(chatID string)

	// Key returns chat key
	//Key() ChatKey

	// GetAppUserID returns app user ID
	GetAppUserID() (appUserID string)

	// SetAppUserID sets app user ID
	SetAppUserID(appUserID string)

	// AddClientLanguage adds client language
	AddClientLanguage(languageCode string) (changed bool)

	/*
		GetBotUserIntID() int
		GetBotUserStringID() string
	*/

	// SetBotUserID sets bot user ID
	SetBotUserID(id interface{})

	// SetIsGroupChat marks current chat as a group chat
	SetIsGroupChat(bool)

	// IsAccessGranted returns true if access is granted
	IsAccessGranted() bool

	// IsGroupChat returns true if current chat is a group chat
	IsGroupChat() bool

	// SetAccessGranted sets access granted flag
	SetAccessGranted(value bool) bool

	// GetPreferredLanguage returns preferred language for the chat
	GetPreferredLanguage() string

	// SetPreferredLanguage sets preferred language for the chat
	SetPreferredLanguage(value string)

	// SetUpdatedTime sets updated time
	SetUpdatedTime(time.Time) // github.com/strongo/user.UpdatedTimeSetter

	// SetDtLastInteraction sets last interaction time
	SetDtLastInteraction(time time.Time)

	// GetAwaitingReplyTo returns path of the step that is awaiting reply to
	GetAwaitingReplyTo() string

	// SetAwaitingReplyTo sets path of the step that is awaiting reply to
	SetAwaitingReplyTo(path string)

	// IsAwaitingReplyTo returns true if awaiting reply to specific step
	IsAwaitingReplyTo(code string) bool

	// AddWizardParam adds a parameter to wizard with a given key & value
	AddWizardParam(key, value string)

	// GetWizardParam returns a parameter from wizard for a given key
	GetWizardParam(key string) string

	// PopStepsFromAwaitingReplyUpToSpecificParent pops steps from awaiting reply up to specific parent
	PopStepsFromAwaitingReplyUpToSpecificParent(code string)

	// PushStepToAwaitingReplyTo pushes step to awaiting reply to
	PushStepToAwaitingReplyTo(code string)
	//GetGaClientID() string

	SetVar(key string, value string)
	GetVar(key string) string
	DelVar(key string)
	HasChangedVars() bool
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
