package openai

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAIClientUtils(t *testing.T) {
	defaultCfg := OpenAIClientConfig{
		APIKey:      os.Getenv("GITHUB_TOKEN"),
		Endpoint:    "https://models.inference.ai.azure.com/chat/completions",
		Temperature: 1.0,
		TopP:        1.0,
		MaxTokens:   1000,
		Model:       GPT4O,
	}
	assert.Equal(t, defaultCfg, DefaultConfig(), "DefaultConfig() should return the default config")
	assert.Equal(t, defaultCfg, NewOpenAIClientConfig(
		os.Getenv("GITHUB_TOKEN"),
		"https://models.inference.ai.azure.com/chat/completions",
		1.0,
		1.0,
		1000,
		GPT4O,
	), "NewOpenAIClientConfig() should return the default config")

	assert.Equal(t, Message{Role: "user", Content: "Hello"}, NewMessage("user", "Hello"), "NewMessage() should return a new message")
	assert.Equal(t, Message{Role: "system", Content: "Hello"}, NewSystemMessage("Hello"), "NewSystemMessage() should return a new system message")
	assert.Equal(t, Message{Role: "user", Content: "Hello"}, NewUserMessage("Hello"), "NewUserMessage() should return a new user message")
	assert.Equal(t, Message{Role: "assistant", Content: "Hello"}, NewAssistantMessage("Hello"), "NewAssistantMessage() should return a new assistant message")

	req := Request{
		Messages:    []Message{},
		Temperature: 1.0,
		TopP:        1.0,
		MaxTokens:   1000,
		Model:       GPT4O,
	}

	assert.Equal(t, req, NewRequest([]Message{}, 1.0, 1.0, 1000, GPT4O), "NewRequest() should return a new request")
	assert.Equal(t, req, NewRequestWithDefaults(defaultCfg), "NewRequestWithDefaults() should return a new request with default config")

	req.AddMessages(Message{Role: "user", Content: "Hello"})
	assert.Equal(t, []Message{{Role: "user", Content: "Hello"}}, req.Messages, "AddMessages() should add messages to the request")
	b, err := req.ToBytes()
	assert.NoError(t, err, "ToBytes() should not return an error")
	assert.NotNil(t, b, "ToBytes() should return a byte slice")
}
