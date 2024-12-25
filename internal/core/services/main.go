package services

import "github.com/Friends-Against-Humanity/senpai/internal/core/ports"

type MainService struct {
	ConversationalAgent ports.ConversationalAgent
}

type MainServiceConfigurer func(*MainService) error

func NewMainService(cfgs ...MainServiceConfigurer) MainService {
	svc := MainService{}

	for _, cfgFn := range cfgs {
		cfgFn(&svc)
	}

	return svc
}

func (s *MainService) Prompt(message string) string {
	return s.ConversationalAgent.PromptWithoutContext(MAIN_PROMPT, message)
}
