package botsfwmodels

// BotRecordsMaker is an interface for making bot records
// This should be implemented by platform adapters
// (for example by https://github.com/bots-go-framework/bots-fw-telegram)
type BotRecordsMaker interface {
	Platform() string

	// MakeAppUserDto makes app user DTO for a given request sender
	MakeAppUserDto(botID string) (BotAppUser, error)

	// MakeBotUserDto makes bot user DTO for a given request sender
	MakeBotUserDto(botID string) (BotUser, error)

	// MakeBotChatDto makes bot chat DTO for a given request
	MakeBotChatDto(botID string) (BotChat, error)
}

func NewBotRecordsMaker(
	platform string,
	makeAppUserDto func(botID string) (appUser BotAppUser, err error),
	makeBotUserDto func(botID string) (botUser BotUser, err error),
	makeBotChatDto func(botID string) (botChat BotChat, err error),
) BotRecordsMaker {
	if makeAppUserDto == nil {
		panic("makeAppUserDto is nil")
	}
	if makeBotUserDto == nil {
		panic("makeBotUserDto is nil")
	}
	if makeBotChatDto == nil {
		panic("makeBotChatDto is nil")
	}
	return botRecordsMaker{
		platform:       platform,
		makeAppUserDto: makeAppUserDto,
		makeBotUserDto: makeBotUserDto,
		makeBotChatDto: makeBotChatDto,
	}
}

type botRecordsMaker struct {
	platform       string
	makeAppUserDto func(botID string) (BotAppUser, error)
	makeBotUserDto func(botID string) (BotUser, error)
	makeBotChatDto func(botID string) (botChat BotChat, err error)
}

func (b botRecordsMaker) Platform() string {
	return b.platform
}

func (b botRecordsMaker) MakeAppUserDto(botID string) (BotAppUser, error) {
	return b.makeAppUserDto(botID)
}

func (b botRecordsMaker) MakeBotUserDto(botID string) (BotUser, error) {
	return b.makeBotUserDto(botID)
}

func (b botRecordsMaker) MakeBotChatDto(botID string) (botChat BotChat, err error) {
	return b.makeBotChatDto(botID)
}