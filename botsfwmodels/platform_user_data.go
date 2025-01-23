package botsfwmodels

var _ PlatformUserData = (*PlatformUserBaseDbo)(nil)

// PlatformUserBaseDbo hold common properties for bot user entities
type PlatformUserBaseDbo struct {
	BotBaseData

	WithRequiredBotIDs

	// FirstName is the first name of a user
	FirstName string `json:"firstName,omitempty" dalgo:"firstName,omitempty,noindex" firestore:"firstName,omitempty"`

	// LastName is the last name of a user
	LastName string `json:"lastName,omitempty"  dalgo:"lastName,omitempty,noindex" firestore:"lastName,omitempty"`

	// UserName is login ID of a user
	UserName string `json:"userName,omitempty"  dalgo:"userName,omitempty,noindex" firestore:"userName,omitempty"`
}

// BaseData returns base data of a user that should be included in all structs that implement PlatformUserData
func (v *PlatformUserBaseDbo) BaseData() *PlatformUserBaseDbo {
	return v
}

func (v *PlatformUserBaseDbo) Validate() error {
	if err := v.BotBaseData.Validate(); err != nil {
		return err
	}
	if err := v.WithRequiredBotIDs.Validate(); err != nil {
		return err
	}
	return nil
}
