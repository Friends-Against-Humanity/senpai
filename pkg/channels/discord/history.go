package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func (b *Bot) GetHistory(channel string) ([]string, error) {
	return b.getLastestMessages(channel, b.maximimMessagesInHistory)
}

func (b *Bot) getNames() (map[string]string, error) {
	return b.members, nil
}

func (b *Bot) getLastestMessages(channelID string, count int) ([]string, error) {
	messages, err := b.session.ChannelMessages(channelID, count, "", "", "")
	if err != nil {
		zap.L().Error("error getting messages", zap.Error(err))
		return nil, err
	}

	members, err := b.getNames()
	if err != nil {
		return nil, err
	}

	history := make([]string, 0, len(messages))
	for _, message := range messages {
		content := message.Content
		tags := strings.Split("<@", content)
		for i := 0; i < len(tags); i++ {
			id := strings.Split(">", tags[i])[0]
			if name, ok := members[id]; ok {
				content = strings.Replace(content, fmt.Sprintf("<@%s>", id), name, 1)
			}
		}

		history = append(history, fmt.Sprintf("%s: %s", message.Author.GlobalName, content))
	}

	return history, nil
}

func (b *Bot) pre(message *discordgo.MessageCreate) error {
	b.members[message.Author.ID] = b.GetName(message.Author)
	return nil
}
