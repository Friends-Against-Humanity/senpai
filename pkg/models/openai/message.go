package openai

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessage(role, content string) Message {
	return Message{
		Role:    role,
		Content: content,
	}
}

func NewSystemMessage(content string) Message {
	return NewMessage("system", content)
}

func NewUserMessage(content string) Message {
	return NewMessage("user", content)
}

func NewAssistantMessage(content string) Message {
	return NewMessage("assistant", content)
}
