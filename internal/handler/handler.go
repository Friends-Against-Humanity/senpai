package handler

import (
	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/pkg/channels/discord"
)

type Handler struct {
	Bot     *discord.Bot
	Service services.MainService
}

type HandlerConfigurer func(b *Handler)

func NewHandler(cfg ...HandlerConfigurer) (*Handler, error) {
	h := &Handler{}

	for _, c := range cfg {
		c(h)
	}

	h.Bot.Handler(h.entrypoint)

	return h, nil
}

// type Bot struct {
// 	Cfg     BotConfig
// 	session *discordgo.Session
// 	Service services.MainService

// 	// Workaround
// 	members                  map[string]string
// 	maximimMessagesInHistory int
// }

// type BotConfigurer func(b *Bot)

// func NewBot(cfg ...BotConfigurer) (*Bot, error) {
// 	bot := &Bot{
// 		maximimMessagesInHistory: 15,
// 		members:                  map[string]string{},
// 	}

// 	for _, c := range cfg {
// 		c(bot)
// 	}

// 	discord, err := discordgo.New("Bot " + bot.Cfg.APIToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	discord.AddHandler(bot.handler)

// 	bot.session = discord
// 	return bot, nil
// }

// func (b *Bot) Run() {
// 	b.session.Open()
// 	zap.L().Info("Bot is running")
// }

// func (b *Bot) Close() error {
// 	return b.session.Close()
// }
