package handler

import (
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) entrypoint(discord *discordgo.Session, message *discordgo.MessageCreate) {
	switch {

	case isSwitchPersonaMessage(message.Content):
		h.switchPersona(discord, message)
		break

	default:
		h.chat(discord, message)
	}
}
