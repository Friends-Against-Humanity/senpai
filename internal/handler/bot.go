package handler

import (
	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Bot struct {
	Cfg     BotConfig
	session *discordgo.Session
	Service services.MainService

	// Workaround
	members                  map[string]string
	latestMessages           []string
	maximimMessagesInHistory int
}

type BotConfigurer func(b *Bot)

func NewBot(cfg ...BotConfigurer) (*Bot, error) {
	bot := &Bot{
		latestMessages:           make([]string, 0),
		maximimMessagesInHistory: 10,
		members:                  map[string]string{},
	}

	for _, c := range cfg {
		c(bot)
	}

	discord, err := discordgo.New("Bot " + bot.Cfg.APIToken)
	if err != nil {
		return nil, err
	}

	discord.AddHandler(bot.handler)

	bot.session = discord
	return bot, nil
}

func (b *Bot) Run() {
	b.session.Open()
	zap.L().Info("Bot is running")
}

func (b *Bot) Close() error {
	return b.session.Close()
}
