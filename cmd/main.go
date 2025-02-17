package main

import (
	"os"
	"os/signal"

	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/personas"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/personas/normal"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/personas/rafraf"
	"github.com/Friends-Against-Humanity/senpai/internal/handler"
	"github.com/Friends-Against-Humanity/senpai/internal/utils/log"
	"github.com/Friends-Against-Humanity/senpai/pkg/channels/discord"
	"github.com/Friends-Against-Humanity/senpai/pkg/models/openai"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := log.InitLogger("debug", "dev")
	handleError(err)

	// Model
	botCfg := discord.NewDefaultBotConfig()
	bot, err := discord.NewBot(func(b *discord.Bot) {
		b.Cfg = botCfg
	})
	model := openai.NewOpenAIClient(func(c *openai.OpenAIClient) error {
		c.Cfg = openai.DefaultConfig().WithModel(openai.GPT4OMini)
		return nil
	})

	// Services
	mainSvc := services.NewService(func(svc *services.Service) error {
		svc.ConversationalAgent = model
		svc.HistoryGateway = bot
		svc.FallbackPersona = normal.NewNormalPersona()
		svc.Personas = []personas.PersonaHandler{
			normal.NewNormalPersona(),
			rafraf.NewRafrafPersona(),
		}

		return nil
	})

	// Bot
	_, err = handler.NewHandler(func(b *handler.Handler) {
		b.Service = mainSvc
		b.Bot = bot
	})
	handleError(err)

	bot.Run()
	defer bot.Close()
	infinitLoop()
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func infinitLoop() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
