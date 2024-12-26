package rafraf

import "github.com/Friends-Against-Humanity/senpai/internal/core/services/personas"

type RafrafPersona struct{}

func NewRafrafPersona() *RafrafPersona {
	return &RafrafPersona{}
}

func (RafrafPersona) GetPersonaName() string {
	return "rafraf"
}

func (RafrafPersona) Supports(key personas.Feature) bool {
	switch key {
	case personas.Metadata, personas.ChatHistory, personas.Permissions:
		return true
	}

	return false
}
