package botsfwmodels

type ChatKey struct {
	BotID  string
	ChatID string
}

func (k ChatKey) String() string {
	return k.BotID + ":" + k.ChatID
}
