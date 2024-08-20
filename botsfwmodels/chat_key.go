package botsfwmodels

import "github.com/strongo/validation"

// NewChatKey creates a new chat key
func NewChatKey(botID, chatID string) ChatKey {
	if botID == "" {
		panic("botID is required, got empty string")
	}
	if chatID == "" {
		panic("chatID is required, got empty string")
	}
	return ChatKey{BotID: botID, ChatID: chatID}
}

// ChatKey is a key for a chat that consists of bot ID and chat ID.
type ChatKey struct {
	// BotID an id of a bot that owns this chat
	BotID string

	// ChatID is an id of a chat as was provided by a bot platform. Might be an integer converted to a string.
	// It's different from ChatInstanceID of Telegram - TODO: document what is the difference an why ChatInstanceID is needed.
	ChatID string
}

// Validate returns error if key is invalid
// TODO(StackOverflow): is it better from performance point of view to use a pointer here and in String?
func (k ChatKey) Validate() error {
	if k.BotID == "" {
		return validation.NewErrRecordIsMissingRequiredField("BotID")
	}
	if k.ChatID == "" {
		return validation.NewErrRecordIsMissingRequiredField("ChatID")
	}
	return nil
}

// String returns string representation of a key
func (k ChatKey) String() string {
	return k.BotID + ":" + k.ChatID
}
