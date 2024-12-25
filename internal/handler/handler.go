package handler

import (
	"fmt"
	"strings"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func (b *Bot) handler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	zap.L().Info("received message", zap.String("message", message.ChannelID), zap.String("author", message.Author.ID))
	b.auditMessage(message)
	if !b.isMentioned(message.Content) {
		return
	}

	fmt.Println("Hello", message.Author.GlobalName)

	latestMessages, err := b.getLastestMessages(message.ChannelID, b.maximimMessagesInHistory)
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
	fmt.Println("Hello", message.Author.GlobalName, response)

	discord.ChannelMessageSend(message.ChannelID, response)
}
