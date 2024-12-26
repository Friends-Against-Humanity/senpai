package domain

type ChatHistory []string

type Metadata struct {
	ChatId       string `json:"chat_id"`
	UserNickname string `json:"user_nickname"`
}
