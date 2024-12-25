package handler

import "os"

type BotConfig struct {
	APIToken string
	ID       string
	Name     string
	GuildID  string
}

func NewBotConfig(name string, apiToken string, id string, guildId string) BotConfig {
	return BotConfig{
		Name:     name,
		APIToken: apiToken,
		ID:       id,
		GuildID:  guildId,
	}
}

func NewDefaultBotConfig() BotConfig {
	return NewBotConfig(
		os.Getenv("DISCORD_BOT_NAME"),
		os.Getenv("DISCORD_API_TOKEN"),
		os.Getenv("DISCORD_BOT_ID"),
		os.Getenv("DISCORD_GUILD_ID"),
	)
}
