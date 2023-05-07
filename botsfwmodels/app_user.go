package botsfwmodels

import (
	strongo "github.com/strongo/app"
)

//type AppUserIntID int64

// BotAppUser holds information about bot app user
type BotAppUser interface {
	strongo.AppUser
	SetBotUserID(platform, botID, botUserID string)
	GetFullName() string
}
