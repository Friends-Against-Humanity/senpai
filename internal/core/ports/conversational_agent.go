package ports

import "context"

type ConversationalAgent interface {
	PromptWithoutContext(system string, messages ...string) string
	Prompt(ctx context.Context, system string, messages ...string) string
}
