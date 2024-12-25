package domain

type Metadata struct {
	UserNickname string   `json:"user_nickname"`
	ChatHistory  []string `json:"chat_history"`
}
