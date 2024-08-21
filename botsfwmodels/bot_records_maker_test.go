package botsfwmodels

import (
	"testing"
)

func TestNewBotRecordsMaker(t *testing.T) {
	type args struct {
		platform       string
		makeAppUserDto func(botID string) (appUser AppUserData, err error)
		makeBotUserDto func(botID string) (botUser PlatformUserData, err error)
		makeBotChatDto func(botID string) (botChat BotChatData, err error)
	}

	makeAppUserDto := func(botID string) (appUser AppUserData, err error) {
		return nil, nil
	}

	makeBotUserDto := func(botID string) (botUser PlatformUserData, err error) {
		return nil, nil
	}

	makeBotChatDto := func(botID string) (botChat BotChatData, err error) {
		return nil, nil
	}

	tests := []struct {
		name string
		args args
		want BotRecordsMaker
	}{
		{
			name: "TestNewBotRecordsMaker",
			args: args{
				platform:       "test",
				makeAppUserDto: makeAppUserDto,
				makeBotUserDto: makeBotUserDto,
				makeBotChatDto: makeBotChatDto,
			},
			want: botRecordsMaker{
				platform:       "test",
				makeAppUserDto: makeAppUserDto,
				makeBotUserDto: makeBotUserDto,
				makeBotChatDto: makeBotChatDto,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBotRecordsMaker(tt.args.platform, tt.args.makeAppUserDto, tt.args.makeBotUserDto, tt.args.makeBotChatDto)
			if _, err := got.MakeAppUserDto("test"); err != nil {
				t.Error(err)
			}
			if _, err := got.MakeBotUserDto("test"); err != nil {
				t.Error(err)
			}
			if _, err := got.MakeBotChatDto("test"); err != nil {
				t.Error(err)
			}
		})
	}
}
