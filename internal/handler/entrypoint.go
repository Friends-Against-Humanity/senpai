package handler

import (
	"strings"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) entrypoint(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if !h.Bot.IsTagged(h.Bot.BotTag(), message.Content) {
		return
	}

	prompt := strings.Replace(message.Content, h.Bot.BotTag(), h.Bot.Cfg.Name, -1)

	metadata := domain.Metadata{
		ChatId:       message.ChannelID,
		UserNickname: h.Bot.GetName(message.Author),
	}

	response, err := h.Service.Prompt(metadata, prompt)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "Ouch! I'm sorry, I'm not feeling well right now. Please try again later.")
		return
	}

	discord.ChannelMessageSend(message.ChannelID, response)
}
