package botsfwdal

import "context"

// DataAccess is an interface for data access layer to store bot & app data
type DataAccess interface {
	AppUserStore
	BotUserStore
	BotChatStore
	RunInTransaction(c context.Context, botID string, f func(c context.Context) error) error
}
