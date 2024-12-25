package services

import (
	"encoding/json"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/Friends-Against-Humanity/senpai/internal/core/ports"
)

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

func (s *MainService) Prompt(metadata domain.Metadata, message string) (string, error) {
	_metadata := s.makeMetadata(metadata)
	prompt := NewPrompt(MAIN_PROMPT_WITH_METADATA, "METADATA_JSON", _metadata)
	result, err := s.ConversationalAgent.Prompt(prompt, message)
	if err != nil {
		return "", s.formatError(err)
	}

	return result, nil
}

func (s *MainService) formatError(err error) error {
	return ErrInternal
}

func (s *MainService) makeMetadata(metadata domain.Metadata) string {
	metadataBytes, _ := json.Marshal(metadata)
	return string(metadataBytes)
}
