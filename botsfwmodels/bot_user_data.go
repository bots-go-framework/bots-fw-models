package botsfwmodels

var _ BotUserData = (*BotUserBaseData)(nil)

// BotUserBaseData hold common properties for bot user entities
type BotUserBaseData struct {
	BotBaseData
	//user.LastLogin

	// FirstName is first name of a user
	FirstName string `json:",omitempty" dalgo:",omitempty,noindex"`

	// LastName is last name of a user
	LastName string `json:",omitempty" dalgo:",omitempty,noindex"`

	// UserName is login ID of a user
	UserName string `json:",omitempty" dalgo:",omitempty,noindex"`
}

// BaseData returns base data of a user that should be included in all structs that implement BotUserData
func (v *BotUserBaseData) BaseData() *BotUserBaseData {
	return v
}
