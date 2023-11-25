package botsfwmodels

// AppUserData is a record where information about app user is stored.
// Bots can use it to  store information about a user like names, preferred locale.
type AppUserData interface {

	// SetBotUserID associates bot user ID with an app user record
	SetBotUserID(platform, botID, botUserID string)

	// UserNamesHolder is an interface to set and get user's names.
	UserNamesHolder

	// PreferredLocaleHolder is an interface to set and get preferred locale
	PreferredLocaleHolder
}

type UserNamesHolder interface {
	// SetName sets a name of a user. It is used to store first name, last name, etc.
	// Parameters:
	// - field: name of a field to set: "firstName", "lastName", "nickName", "fullName"
	SetName(field, value string) error

	// GetName returns a name of a user. It is used to store first name, last name, etc.
	// Parameters:
	// - field: name of a field to set: "firstName", "lastName", "nickName", "fullName"
	GetName(field string) string

	// GetFullName returns full name of a user
	GetFullName() string
}
