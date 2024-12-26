package normal

import "github.com/Friends-Against-Humanity/senpai/internal/core/services/personas"

type NormalPersona struct{}

func NewNormalPersona() *NormalPersona {
	return &NormalPersona{}
}

func (NormalPersona) GetPersonaName() string {
	return "normal"
}

func (NormalPersona) Supports(key personas.Feature) bool {
	switch key {
	case personas.Metadata, personas.ChatHistory, personas.Permissions:
		return true
	}

	return false
}
