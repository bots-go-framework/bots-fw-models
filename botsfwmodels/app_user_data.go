package botsfwmodels

import (
	strongo "github.com/strongo/app"
)

// AppUserData holds information about bot app user
type AppUserData interface {
	strongo.AppUser
	SetBotUserID(platform, botID, botUserID string)
	GetFullName() string
}
