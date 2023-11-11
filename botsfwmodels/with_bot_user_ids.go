package botsfwmodels

import (
	"fmt"
	"slices"
)

type WithBotUserIDs struct {
	BotUserIDs []string `json:"botUserIDs,omitempty" firestore:"botUserIDs,omitempty" dalgo:"botUserIDs,omitempty"`
}

func (v *WithBotUserIDs) SetBotUserID(platform, bot, userID string) {
	botUserID := fmt.Sprintf("%s:%s:%s", platform, bot, userID)
	if slices.Contains(v.BotUserIDs, botUserID) {
		return
	}
	v.BotUserIDs = append(v.BotUserIDs, botUserID)
}
