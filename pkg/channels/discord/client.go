package discord

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Bot struct {
	Cfg     BotConfig
	session *discordgo.Session

	// // Workaround
	members                  map[string]string
	maximimMessagesInHistory int
}

type BotConfigurer func(b *Bot)

func NewBot(cfg ...BotConfigurer) (*Bot, error) {
	bot := &Bot{
		maximimMessagesInHistory: 15,
		members:                  map[string]string{},
	}

	for _, c := range cfg {
		c(bot)
	}

	discord, err := discordgo.New("Bot " + bot.Cfg.Token)
	if err != nil {
		return nil, err
	}

	bot.session = discord
	return bot, nil
}

func (b *Bot) Handler(fn func(s *discordgo.Session, m *discordgo.MessageCreate)) {
	b.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		zap.L().Info("Message received",
			zap.String("message", m.Content),
			zap.String("author-id", m.Author.ID),
			zap.String("author", m.Author.GlobalName),
			zap.String("channel", m.ChannelID),
		)
		b.pre(m)
		fn(s, m)
	})
}

func (b *Bot) Run() {
	b.session.Open()
	zap.L().Info("Bot is running")
}

func (b *Bot) Close() error {
	zap.L().Info("Bot is closing")
	return b.session.Close()
}
