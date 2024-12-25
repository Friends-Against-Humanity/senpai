package openai

import "os"

type OpenAIClientConfig struct {
	APIKey      string
	Endpoint    string
	Temperature float64
	TopP        float64
	MaxTokens   int
	Model       string
}

func NewOpenAIClientConfig(
	apiKey string,
	endpoint string,
	temperature,
	topP float64,
	maxTokens int,
	model string) OpenAIClientConfig {
	return OpenAIClientConfig{
		APIKey:      apiKey,
		Endpoint:    endpoint,
		Temperature: temperature,
		TopP:        topP,
		MaxTokens:   maxTokens,
		Model:       model,
	}
}

func DefaultConfig() OpenAIClientConfig {
	return NewOpenAIClientConfig(
		os.Getenv("GITHUB_TOKEN"),
		"https://models.inference.ai.azure.com/chat/completions",
		1.0,
		1.0,
		1000,
		GPT4O,
	)
}
