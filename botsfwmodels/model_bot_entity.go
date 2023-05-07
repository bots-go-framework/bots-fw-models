package botsfwmodels

import "github.com/strongo/app/user"

// BotBaseData holds properties common to all bot entities
type BotBaseData struct {

	// Links bot data to a specific app user
	user.OwnedByUserWithID

	// AccessGranted indicates if access to the bot has been granted
	AccessGranted bool
}

// IsAccessGranted indicates if access to the bot has been granted
func (e *BotBaseData) IsAccessGranted() bool {
	return e.AccessGranted
}

// SetAccessGranted mark that access has been granted
func (e *BotBaseData) SetAccessGranted(value bool) bool {
	if e.AccessGranted != value {
		e.AccessGranted = value
		return true
	}
	return false
}
