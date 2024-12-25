package services

import "strings"

const (
	MAIN_PROMPT               = `You are senpai, an AI assistant. You are here to help and chat with me.`
	MAIN_PROMPT_WITH_METADATA = `You are senpai, an AI assistant. You are here to help and chat with me.

You do what a default AI assistant does, for now. You will be able to do more in the future, when Niemand develops the personnas.

In case you need metadata, here is the metadata in a JSON format: METADATA_JSON
`
)

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
