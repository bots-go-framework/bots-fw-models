package botsfwdal

// DataAccess is an interface for data access layer to store bot & app data
type DataAccess interface {
	AppUserStore
	BotUserStore
	BotChatStore
}
