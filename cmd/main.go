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

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {

	err := log.InitLogger("debug", "dev")
	handleError(err)

	// L1
	model := openai.NewOpenAIClient()

	mainSvc := services.NewMainService(func(svc *services.MainService) error {
		svc.ConversationalAgent = model
		return nil
	})

	// chatHistory := []string{
	// 	"adam: Hello, what's your name?",
	// 	"senpai: Hi, Adam. I'm Senpai.",
	// 	"adam: Yo, ill have a meeting mbaad maa Tom, thats the guy mtaa Company X from Japan, will show the game to my friend, and discuss the consulting plan, ill seek whats on his mind for first clients of the year, but stay tuned... U can apply GCP skills especially in Cloud Security and Kube, ill update but for sure it ll need more work force and not like a heavy one dw xD",
	// 	"nour: hey there, that was a lot",
	// }

	// metadata := domain.Metadata{
	// 	UserNickname: "nour",
	// 	ChatHistory:  chatHistory,
	// }

	// response, err := mainSvc.Prompt(metadata, "hey senpai, where is the company located?")
	// handleError(err)
	// fmt.Println(response)

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

func wait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
