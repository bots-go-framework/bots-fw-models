package botsfwmodels

var _ BotUserData = (*BotUserBaseData)(nil)

// BotUserBaseData hold common properties for bot user entities
type BotUserBaseData struct {
	BotBaseData
	//user.LastLogin

	// FirstName is the first name of a user
	FirstName string `json:"firstName,omitempty" firestore:"firstName,omitempty,noindex" dalgo:"firstName,omitempty,noindex"`

	// LastName is the last name of a user
	LastName string `json:"lastName,omitempty" firestore:"lastName,omitempty,noindex" dalgo:"lastName,omitempty,noindex"`

	// UserName is login ID of a user
	UserName string `json:"userName,omitempty" firestore:"userName,omitempty,noindex" dalgo:"userName,omitempty,noindex"`
}

// BaseData returns base data of a user that should be included in all structs that implement BotUserData
func (v *BotUserBaseData) BaseData() *BotUserBaseData {
	return v
}
