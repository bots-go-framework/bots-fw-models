package botsfwmodels

import "github.com/strongo/strongoapp/appuser"

// AppUserData holds information about bot app user
type AppUserData interface {
	appuser.BaseUserData

	// LocaleSettings is an interface to set and get preferred locale
	LocaleSettings

	// SetBotUserID associates bot user ID with an app user record
	SetBotUserID(platform, botID, botUserID string)

	// GetFullName returns full name of a user
	GetFullName() string
}
