package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) BotTag() string {
	return b.GetTag(b.Cfg.ID)
}

func (b *Bot) GetTag(id string) string {
	return fmt.Sprintf("<@%s>", id)
}

func (b *Bot) IsTagged(id, message string) bool {
	return strings.Contains(message, id)
}

func (b *Bot) GetName(author *discordgo.User) string {
	name := author.GlobalName
	if name == "" {
		name = author.Username
	}
	return name
}
