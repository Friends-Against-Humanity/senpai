package openai

import "os"

type OpenAIClientConfig struct {
	APIKey      string
	Temperature float64
	TopP        float64
	MaxTokens   int
	Model       string
}

func NewOpenAIClientConfig(apiKey string, temperature, topP float64, maxTokens int, model string) OpenAIClientConfig {
	return OpenAIClientConfig{
		APIKey:      apiKey,
		Temperature: temperature,
		TopP:        topP,
		MaxTokens:   maxTokens,
		Model:       model,
	}
}

func DefaultConfig() OpenAIClientConfig {
	return NewOpenAIClientConfig(
		os.Getenv("GITHUB_TOKEN"),
		1.0,
		1.0,
		1000,
		GPT4O,
	)
}
