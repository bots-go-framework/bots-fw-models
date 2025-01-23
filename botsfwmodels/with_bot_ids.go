package botsfwmodels

import (
	"fmt"
	"github.com/strongo/validation"
	"strings"
)

type WithBotIDs struct {
	BotIDs []string `json:"botIDs,omitempty" dalgo:"botIDs,omitempty,noindex" firestore:"botIDs,omitempty"`
}

func (v WithBotIDs) Validate() error {
	for i, botID := range v.BotIDs {
		if strings.TrimSpace(botID) == "" {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("botIDs[%d]", i), "is empty")
		}
	}
	return nil
}

type WithRequiredBotIDs WithBotIDs

func (v WithRequiredBotIDs) Validate() error {
	if len(v.BotIDs) == 0 {
		return validation.NewErrRecordIsMissingRequiredField("botIDs")
	}
	return WithBotIDs(v).Validate()
}
