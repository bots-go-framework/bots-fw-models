package botsfwmodels

import (
	"time"
)

// BotUser interface provides information about bot user
// This should be implemented by bot user record struct.
type BotUser interface {
	// BaseData returns base data that should be implemented by all bot user structs
	BaseData() *BotUserBaseData

	// GetAppUserID returns app user ID if available
	GetAppUserID() string

	// SetAppUserID sets app user ID to associate bot user record with app user
	SetAppUserID(appUserID string)

	WithAccessGrantedFlag

	// SetUpdatedTime sets last updated time // TODO: document intended usage
	SetUpdatedTime(time.Time) //to satisfy github.com/strongo/app/user.UpdatedTimeSetter
}

type WithAccessGrantedFlag interface {

	// IsAccessGranted returns true if access is granted
	IsAccessGranted() bool

	// SetAccessGranted sets access granted flag
	SetAccessGranted(value bool) bool
}
