package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) switchPersona(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if !h.Bot.IsTagged(h.Bot.BotTag(), message.Content) {
		return
	}

	persona := strings.Split(message.Content, "switch persona to ")[1]
	if len(strings.Split(persona, " ")) > 1 {
		persona = strings.Split(persona, " ")[0]
	}

	h.setPersona(message.ChannelID, persona)

	discord.ChannelMessageSend(message.ChannelID, "Hello, world, I now represent "+persona)
}

func isSwitchPersonaMessage(message string) bool {
	return strings.Contains(message, "switch persona to ")
}
