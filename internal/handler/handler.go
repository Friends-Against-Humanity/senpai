package handler

import (
	"strings"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	b.auditMessage(message)
	if !b.isMentioned(message.Content) {
		return
	}

	latestMessages, err := b.getLastestMessages(message.ChannelID, 10)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "Ouch! I'm sorry, I'm not feeling well right now. Please try again later.")
		return
	}

	metadata := domain.Metadata{
		UserNickname: message.Author.GlobalName,
		ChatHistory:  latestMessages,
	}

	prompt := strings.Replace(message.Content, b.tag(), b.Cfg.Name, -1)
	response, err := b.Service.Prompt(metadata, prompt)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "Ouch! I'm sorry, I'm not feeling well right now. Please try again later.")
		return
	}

	discord.ChannelMessageSend(message.ChannelID, response)
}
