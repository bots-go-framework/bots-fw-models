package botsfwdal

import (
	"context"
	"github.com/bots-go-framework/bots-fw-models/botsfwmodels"
)

//type UserID interface {
//	int | string
//}

// BotUserStore provider to store information about bot user
type BotUserStore interface {

	// GetBotUserByID returns bot user data
	GetBotUserByID(c context.Context, botUserID string) (botsfwmodels.BotUser, error)

	// SaveBotUser saves bot user data
	SaveBotUser(c context.Context, botUserID string, botUserData botsfwmodels.BotUser) error

	// CreateBotUser creates new bot user in DB
	// TODO: should be moved to bots-fw-* package or documented why we need a dedicated method for this
	//CreateBotUser(c context.Context, botID string, apiUser botsfw.WebhookActor) (botsfwmodels.BotUser, error)

	//io.Closer
}
