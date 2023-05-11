package botsfwmodels

type ChatKey struct {
	// BotID an id of a bot that owns this chat
	BotID string

	// ChatID is an id of a chat as was provided by a bot platform. Might be an integer converted to a string.
	// It's different from ChatInstanceID of Telegram - TODO: document what is the difference an why ChatInstanceID is needed.
	ChatID string
}

func (k ChatKey) String() string {
	return k.BotID + ":" + k.ChatID
}
