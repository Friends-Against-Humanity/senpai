package ports

import "context"

type ConversationalAgent interface {
	PromptWithoutContext(system string, message string) (string, error)
	Prompt(ctx context.Context, system string, message string) (string, error)
}
