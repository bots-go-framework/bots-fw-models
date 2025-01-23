package botsfwmodels

import (
	"fmt"
	"github.com/strongo/validation"
	"slices"
	"strings"
)

type WithBotUserIDs struct {
	BotUserIDs []string `json:"botUserIDs,omitempty" firestore:"botUserIDs,omitempty" dalgo:"botUserIDs,omitempty"`
}

func (v *WithBotUserIDs) SetBotUserID(platform, bot, userID string) {
	if pid := strings.TrimSpace(platform); pid == "" {
		panic("value of `platform` argument is empty")
	} else if pid != platform {
		panic("value of `platform` argument is not trimmed")
	}
	if bid := strings.TrimSpace(bot); bid != bot {
		panic("value of `bot` argument is not trimmed")
	}
	if uid := strings.TrimSpace(userID); uid == "" {
		panic("value of `platform` argument is empty")
	} else if uid != userID {
		panic("value of `userID` argument is not trimmed")
	}
	botUserID := fmt.Sprintf("%s:%s:%s", platform, bot, userID)
	if slices.Contains(v.BotUserIDs, botUserID) {
		return
	}
	v.BotUserIDs = append(v.BotUserIDs, botUserID)
}

func (v *WithBotUserIDs) Validate() error {
	for i, botUserID := range v.BotUserIDs {
		if strings.TrimSpace(botUserID) == "" {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("botUserIDs[%d]", i), "is empty")
		}
	}
	return nil
}
