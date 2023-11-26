package botsfwmodels

type AppUserAdapter interface {
	// SetBotUserID associates bot user ID with an app user record
	SetBotUserID(platform, botID, botUserID string)

	// UserNamesHolder is an interface to set and get user's names.
	UserNamesHolder

	// PreferredLocaleHolder is an interface to set and get preferred locale
	PreferredLocaleHolder
}
