package discord

import "os"

type BotConfig struct {
	Token string
	ID    string
	Name  string
}

func NewDefaultBotConfig() BotConfig {
	return BotConfig{
		Name:  os.Getenv("DISCORD_BOT_NAME"),
		Token: os.Getenv("DISCORD_API_TOKEN"),
		ID:    os.Getenv("DISCORD_BOT_ID"),
	}
}
