package main

import (
	"os"
	"os/signal"

	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/internal/handler"
	"github.com/Friends-Against-Humanity/senpai/internal/utils/log"
	"github.com/Friends-Against-Humanity/senpai/pkg/models/openai"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	err := log.InitLogger("debug", "dev")
	handleError(err)

	// Model
	model := openai.NewOpenAIClient()

	// Services
	mainSvc := services.NewMainService(func(svc *services.MainService) error {
		svc.ConversationalAgent = model
		return nil
	})

	// Bot
	botCfg := handler.NewDefaultBotConfig()
	bot, err := handler.NewBot(func(b *handler.Bot) {
		b.Service = mainSvc
		b.Cfg = botCfg
	})
	handleError(err)
	bot.Run()
	defer bot.Close()

	wait()
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func wait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
