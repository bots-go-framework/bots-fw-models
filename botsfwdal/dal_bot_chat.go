package botsfwdal

import (
	"context"
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
)

// BotChatStore is interface for DAL to store bot chat data
type BotChatStore interface {

	// GetBotChatData returns bot chat data from persistent storage.
	// It a record is not found it MUST return empty data and a NotFoundErr(err).
	GetBotChatData(c context.Context, key botsfwmodels.ChatKey) (chatData botsfwmodels.BotChatData, err error)

	// SaveBotChatData saves bot chat data to persistent storage
	SaveBotChatData(c context.Context, key botsfwmodels.ChatKey, chatData botsfwmodels.BotChatData) error

	// Close closes the store, e.g. commits sends a signal to commit transaction
	// TODO: Consider to remove this method if possible
	Close(c context.Context) error // TODO: Was io.Closer, should it?

	// NewBotChatEntity creates new bot chat record
	// NewBotChatEntity(c context.Context, botID string, botChat botsfw.WebhookChat, appUserID, botUserID string, isAccessGranted bool) botsfwmodels.BotChatData
}
