package botsfwdal

import (
	"context"
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
)

//type UserID interface {
//	int | string
//}

// BotUserStore provider to store information about bot user
type BotUserStore interface {

	// GetBotUserByID returns bot user data
	// The `botID` parameter is used to pass to a DB provider to get a database connection
	// Should always return a valid bot user data object, even if it is emptu or an error occured
	GetBotUserByID(c context.Context, botID, botUserID string) (botsfwmodels.BotUserData, error)

	// SaveBotUser saves bot user data
	// The `botID` parameter is used to pass to a DB provider to get a database connection
	SaveBotUser(c context.Context, botID, botUserID string, botUserData botsfwmodels.BotUserData) error

	// CreateBotUser creates new bot user in DB - moved to bots-fw-* packages
	// CreateBotUser(c context.Context, botID string, apiUser botsfw.WebhookActor) (botsfwmodels.BotUserData, error)

	//io.Closer
}
