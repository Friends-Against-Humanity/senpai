package services

import (
	"encoding/json"
	"strings"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/Friends-Against-Humanity/senpai/internal/core/services/personas"
)

func dumpMetadata(metadata domain.Metadata) string {
	metadataBytes, _ := json.Marshal(metadata)
	return string(metadataBytes)
}

func dumpChatHistory(history domain.ChatHistory) string {
	historyBytes, _ := json.Marshal(history)
	return string(historyBytes)
}

func (s *Service) getPersonaByName(name string) personas.PersonaHandler {
	for _, p := range s.Personas {
		if p.GetPersonaName() == name {
			return p
		}
	}

	return s.FallbackPersona
}

func replace(prompt string, args ...string) string {
	if len(args)%2 == 1 {
		args = append(args, "")
	}

	i := 0
	for i < len(args) {
		key := args[i]
		value := args[i+1]

		prompt = strings.Replace(prompt, key, value, -1)
		i += 2
	}

	return prompt
}
