package botsfwmodels

// AppUserData holds information about bot app user
type AppUserData interface {

	// LocaleSettings is an interface to set and get preferred locale
	LocaleSettings

	// SetNames sets user record name fields using information provided by bot platform
	SetNames(first, last string)

	// SetBotUserID associates bot user ID with an app user record
	SetBotUserID(platform, botID, botUserID string)

	// GetFullName returns full name of a user
	GetFullName() string
}
