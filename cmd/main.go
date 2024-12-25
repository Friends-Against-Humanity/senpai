package main

import (
	"fmt"
	"os"

	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
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

	response, err := mainSvc.Prompt("hello, what's your name")
	handleError(err)
	fmt.Println(response)
}
