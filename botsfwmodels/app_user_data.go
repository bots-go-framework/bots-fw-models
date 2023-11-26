package botsfwmodels

// AppUserData is a record where information about app user is stored.
// Bots can use it to  store information about a user like names, preferred locale.
type AppUserData interface {

	// BotsFwAdapter returns to bots framework an adapter to app user data record.
	// Using an adapter ensures there is no clashes between bots framework interfaces and app user struct.
	BotsFwAdapter() AppUserAdapter
}

type UserNamesHolder interface {

	// SetNames sets names of a user.
	SetNames(firstName, lastName, fullName string) error

	//// GetName returns a name of a user. It is used to store first name, last name, etc.
	//// Parameters:
	//// - field: name of a field to set: "firstName", "lastName", "nickName", "fullName"
	//GetName(field string) string
	//
	//// GetFullName returns full name of a user
	//GetFullName() string
}
