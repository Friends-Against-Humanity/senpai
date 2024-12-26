package services

import (
	"strings"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/Friends-Against-Humanity/senpai/internal/core/ports"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/personas"
)

type Service struct {
	ports.ConversationalAgent
	ports.HistoryGateway

	Personas        []personas.PersonaHandler
	FallbackPersona personas.PersonaHandler
}

type MainServiceConfigurer func(*Service) error

func NewService(cfgs ...MainServiceConfigurer) Service {
	svc := Service{}

	for _, cfgFn := range cfgs {
		cfgFn(&svc)
	}

	return svc
}

func (s *Service) Prompt(persona string, metadata domain.Metadata, message string) (string, error) {
	p := s.getPersonaByName(persona)

	prompt := p.GetSystemPrompt()
	_chatHistory := ""
	_metadata := ""

	if p.Supports(personas.Metadata) {
		_metadata = dumpMetadata(metadata)
		prompt = p.AddMetadata(prompt)
	}

	if p.Supports(personas.ChatHistory) {
		chatHistory, err := s.HistoryGateway.GetHistory(metadata.ChatId)
		if err != nil {
			return "", s.formatError(err)
		}
		_chatHistory = dumpChatHistory(chatHistory)
		prompt = p.AddPermissions(prompt)
	}

	if p.Supports(personas.Permissions) {
		prompt = p.AddPermissions(prompt)
	}

	prompt = replace(
		prompt,
		"METADATA_JSON", _metadata,
		"CHAT_HISTORY", _chatHistory,
	)

	prompt += s.personasPrompt()

	result, err := s.ConversationalAgent.Prompt(prompt, message)
	if err != nil {
		return "", s.formatError(err)
	}

	return result, nil
}

func (s *Service) formatError(err error) error {
	switch err.Error() {
	case "RATE_LIMIT_EXCEEDED":
		return ErrQuotaExceeded
	}

	return ErrInternal
}

func (s *Service) personasPrompt() string {
	personas := []string{}
	for _, p := range s.Personas {
		personas = append(personas, p.GetPersonaName())
	}
	return "You support the following personas: " + strings.Join(personas, ",")
}
