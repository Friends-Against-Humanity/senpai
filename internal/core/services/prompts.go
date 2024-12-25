package services

import (
	"strings"
)

type Prompt struct{}

// func NewPrompt(defaultPrompt string) *Prompt {
// 	return &Prompt{}
// }

// func NewMainPrompt() *Prompt {
// 	return NewPrompt(prompt.MAIN_PROMPT)
// }

func NewPrompt(prompt string, args ...string) string {
	if len(args)%2 == 1 {
		args = append(args, "")
	}

	i := 0
	for i < len(args) {
		key := args[i]
		value := args[i+1]

		prompt = strings.Replace(prompt, key, value, -1)
		i += 2
	}

	return prompt
}
