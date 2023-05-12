package botsfwmodels

import "github.com/strongo/app/user"

var _ BotUserData = (*BotUserBaseData)(nil)

// BotUserBaseData hold common properties for bot user entities
type BotUserBaseData struct {
	BotBaseData
	user.LastLogin

	FirstName string `json:",omitempty" dalgo:",omitempty,noindex"`
	LastName  string `json:",omitempty" dalgo:",omitempty,noindex"`
	UserName  string `json:",omitempty" dalgo:",omitempty,noindex"`
}

func (v *BotUserBaseData) BaseData() *BotUserBaseData {
	return v
}
