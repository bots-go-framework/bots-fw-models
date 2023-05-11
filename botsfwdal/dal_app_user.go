package botsfwdal

import (
	"context"
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
)

// AppUserStore is an interface for data access layer to store app user data
type AppUserStore interface {

	// CreateAppUser creates a new app user record in a persistent data store
	// The `botID` parameter is used to pass to a DB provider to get a database connection
	CreateAppUser(c context.Context, botID string, appUser botsfwmodels.AppUserData) (appUserID string, err error)

	// SaveAppUser saves app user data into a persistent data store
	// The `botID` parameter is used to pass to a DB provider to get a database connection
	SaveAppUser(c context.Context, botID, appUserID string, appUserData botsfwmodels.AppUserData) error

	// GetAppUserByID retrieves app user data from a persistent data store
	// The `botID` parameter is used to pass to a DB provider to get a database connection
	GetAppUserByID(c context.Context, botID, appUserID string, appUser botsfwmodels.AppUserData) error
}

//// BotAppUserStore interface for storing user information to persistent store
//type BotAppUserStore interface {
//	GetAppUserByID(c context.Context, appUserID string, appUser AppUserData) error
//	//CreateAppUser(c context.Context, botID string, actor botsfw.WebhookActor) (appUserID string, appUserEntity AppUserData, err error)
//	//SaveAppUser(c context.Context, appUserId int64, appUserEntity AppUserData) error
//}
