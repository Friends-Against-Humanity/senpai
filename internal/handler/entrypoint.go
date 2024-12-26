package handler

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func (h *Handler) entrypoint(discord *discordgo.Session, message *discordgo.MessageCreate) {
	zap.L().Debug("entrypoint", zap.String("persona", h.getPersona(message.ChannelID)))
	switch {

	case isSwitchPersonaMessage(message.Content):
		h.switchPersona(discord, message)
		break

	default:
		h.chat(discord, message)
	}
}
