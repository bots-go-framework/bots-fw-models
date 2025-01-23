package botsfwmodels

import (
	"errors"
	"testing"
)

func TestWithBotUserIDs_SetBotUserID(t *testing.T) {
	type fields struct {
		BotUserIDs []string
	}
	type args struct {
		platform string
		bot      string
		userID   string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedPanic string
	}{
		{
			name:   "Should pass",
			fields: fields{BotUserIDs: []string{"platform1:bot1:user1"}},
			args:   args{platform: "platform2", bot: "bot2", userID: "user2"},
		},
		{
			name:          "Should panic when platform is empty",
			args:          args{platform: " ", bot: "bot2", userID: "user2"},
			expectedPanic: "value of `platform` argument is empty",
		},
		{
			name:          "Should panic when platform has space at start",
			args:          args{platform: " p1", bot: "bot2", userID: "user2"},
			expectedPanic: "value of `platform` argument is not trimmed",
		},
		{
			name:          "Should panic when platform has space at end",
			args:          args{platform: "p1 ", bot: "bot2", userID: "user2"},
			expectedPanic: "value of `platform` argument is not trimmed",
		},
		{
			name: "Should pass when bot is empty",
			args: args{platform: "platform", bot: "", userID: "user2"},
		},
		{
			name:          "Should panic when bot has only spaces",
			args:          args{platform: "platform", bot: " ", userID: "user2"},
			expectedPanic: "value of `bot` argument is not trimmed",
		},
		{
			name:          "Should panic when bot has space at start",
			args:          args{platform: "platform", bot: " bot2", userID: "user2"},
			expectedPanic: "value of `bot` argument is not trimmed",
		},
		{
			name:          "Should panic when bot has space at end",
			args:          args{platform: "bot", bot: "bot2 ", userID: "user2"},
			expectedPanic: "value of `bot` argument is not trimmed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &WithBotUserIDs{
				BotUserIDs: tt.fields.BotUserIDs,
			}
			if tt.expectedPanic != "" {
				defer func() {
					if r := recover(); r == nil || r.(string) != tt.expectedPanic {
						t.Errorf("SetBotUserID() panic = %v, expected: %v", r, tt.expectedPanic)
					}
				}()
			}
			v.SetBotUserID(tt.args.platform, tt.args.bot, tt.args.userID)
		})
	}
}

func TestWithBotUserIDs_Validate(t *testing.T) {
	type fields struct {
		BotUserIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Should pass",
			fields: fields{
				BotUserIDs: []string{"platform1:bot1:user1"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &WithBotUserIDs{
				BotUserIDs: tt.fields.BotUserIDs,
			}
			if err := v.Validate(); err == nil && tt.wantErr != nil || tt.wantErr == nil && err != nil || !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
