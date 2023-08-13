package botsfwmodels

import (
	"github.com/strongo/validation"
	"time"
)

// BotBaseData holds properties common to all bot entities
type BotBaseData struct {
	//user.OwnedByUserWithID
	AppUserID string // intentionally indexed & do NOT omitempty (so we can find records with empty AppUserID)

	// AppUserIntID is a strongly typed ID of the user
	// Deprecated: use AppUserID instead, remove once all bots are migrated
	//AppUserIntID int64 `json:",omitempty" datastore:",omitempty" firestore:",omitempty"`

	DtCreated time.Time `json:",omitempty" datastore:",omitempty" firestore:",omitempty"`
	DtUpdated time.Time `json:",omitempty" datastore:",omitempty" firestore:",omitempty"`

	// AccessGranted indicates if access to the bot has been granted
	AccessGranted bool
}

// Validate returns error if data is invalid
func (e *BotBaseData) Validate() error {
	if e.DtUpdated.Before(e.DtCreated) {
		return validation.NewErrBadRecordFieldValue("DtUpdated", "DtUpdated is before DtCreated")
	}
	//if e.AppUserID != "" && e.AppUserIntID != 0 && strconv.FormatInt(e.AppUserIntID, 10) != e.AppUserID {
	//	return validation.NewErrBadRecordFieldValue("AppUserIntID", "does not match AppUserID")
	//}
	return nil
}

// SetUpdatedTime sets updated time
func (e *BotBaseData) SetUpdatedTime(v time.Time) {
	e.DtUpdated = v
}

// GetAppUserID returns app user ID
func (e *BotBaseData) GetAppUserID() string {
	if e.AppUserID != "" {
		return e.AppUserID
	}
	//if e.AppUserIntID != 0 {
	//	return strconv.FormatInt(e.AppUserIntID, 10)
	//}
	return ""
}

// SetAppUserID sets app user ID
func (e *BotBaseData) SetAppUserID(s string) {
	e.AppUserID = s
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
