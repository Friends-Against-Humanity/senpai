package openai

import (
	"encoding/json"
)

type Request struct {
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
	MaxTokens   int       `json:"max_tokens"`
	Model       string    `json:"model"`
}

func NewRequest(messages []Message, temperature, topP float64, maxTokens int, model string) Request {
	return Request{
		Messages:    messages,
		Temperature: temperature,
		TopP:        topP,
		MaxTokens:   maxTokens,
		Model:       model,
	}
}

func NewRequestWithDefaults(cfg OpenAIClientConfig) Request {
	return Request{
		Messages:    []Message{},
		Temperature: cfg.Temperature,
		TopP:        cfg.TopP,
		MaxTokens:   cfg.MaxTokens,
		Model:       cfg.Model,
	}
}

func (r *Request) AddMessages(messages ...Message) {
	r.Messages = append(r.Messages, messages...)
}

func (r *Request) ToBytes() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return b, nil
}
