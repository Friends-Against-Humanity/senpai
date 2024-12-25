package ports

type ConversationalAgent interface {
	Prompt(system string, message string) (string, error)
}
