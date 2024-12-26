package personas

type PersonaHandler interface {
	GetPersonaName() string
	GetSystemPrompt() string
	Supports(key Feature) bool
	AddMetadata(prompt string) string
	AddChatHistory(prompt string) string
	AddPermissions(prompt string) string
}

type Feature string

const (
	Metadata    = "metadata"
	ChatHistory = "chat-history"
	Permissions = "permissions"
)
