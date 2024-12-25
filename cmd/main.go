package main

import (
	"fmt"

	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/pkg/models/openai"
)

func main() {
	// L1
	model := openai.NewOpenAIClient()

	mainSvc := services.NewMainService(func(svc *services.MainService) error {
		svc.ConversationalAgent = model
		return nil
	})

	response := mainSvc.Prompt(
		"hello, what's your name",
	)

	fmt.Println(response)
}
