package openai

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type OpenAIClient struct {
	Cfg OpenAIClientConfig
}

type OpenAIClientConfiger func(*OpenAIClient) error

func NewOpenAIClient(cfgs ...OpenAIClientConfiger) *OpenAIClient {
	cfg := DefaultConfig()
	client := OpenAIClient{Cfg: cfg}

	for _, cfgFn := range cfgs {
		cfgFn(&client)
	}

	return &client
}

func (c *OpenAIClient) _client() http.Client {
	return http.Client{
		Timeout: 30 * time.Second,
	}
}

func (c *OpenAIClient) Prompt(system string, message string) (string, error) {
	// Body
	r := NewRequestWithDefaults(c.Cfg)
	r.AddMessages(
		NewSystemMessage(system),
		NewUserMessage(message),
	)

	reqBody, err := r.ToBytes()
	if err != nil {
		zap.L().Error("failed to marshal request", zap.Error(err))
		return "", err
	}
	bodyReader := bytes.NewReader(reqBody)

	// Request
	req, err := http.NewRequest(http.MethodPost, c.Cfg.Endpoint, bodyReader)
	if err != nil {
		zap.L().Error("failed to create http request", zap.Error(err))
		return "", err
	}

	// Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Cfg.APIKey))

	client := c._client()

	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("failed to make request", zap.Error(err))
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("failed to read response body", zap.Error(err))
		return "", err
	}

	response, err := NewResponse(body)
	zap.L().Debug("received", zap.Any("response", response))
	if len(response.Choices) < 1 {
		return "", errors.New("RATE_LIMIT_EXCEEDED")
	}

	return response.Choices[0].Message.Content, nil
}
