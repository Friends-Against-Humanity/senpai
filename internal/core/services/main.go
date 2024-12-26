package services

import (
	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/Friends-Against-Humanity/senpai/internal/core/ports"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/prompt"
)

type MainService struct {
	ports.ConversationalAgent
	ports.HistoryGateway
}

type MainServiceConfigurer func(*MainService) error

func NewMainService(cfgs ...MainServiceConfigurer) MainService {
	svc := MainService{}

	for _, cfgFn := range cfgs {
		cfgFn(&svc)
	}

	return svc
}

func (s *MainService) Prompt(metadata domain.Metadata, message string) (string, error) {
	_metadata := dumpMetadata(metadata)
	chatHistory, err := s.HistoryGateway.GetHistory(metadata.ChatId)
	if err != nil {
		return "", s.formatError(err)
	}
	_chatHistory := dumpChatHistory(chatHistory)

	prompt := NewPrompt(
		prompt.MAIN_PROMPT_WITH_METADATA,
		"METADATA_JSON", _metadata,
		"CHAT_HISTORY", _chatHistory,
	)
	result, err := s.ConversationalAgent.Prompt(prompt, message)
	if err != nil {
		return "", s.formatError(err)
	}

	return result, nil
}

func (s *MainService) formatError(err error) error {
	return ErrInternal
}
