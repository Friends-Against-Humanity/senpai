package openai

import "context"

type OpenAIClient struct {
	cfg OpenAIClientConfig
}

type OpenAIClientConfiger func(*OpenAIClient) error

func NewOpenAIClient(cfgs ...OpenAIClientConfiger) *OpenAIClient {
	cfg := DefaultConfig()
	client := OpenAIClient{cfg: cfg}

	for _, cfgFn := range cfgs {
		cfgFn(&client)
	}

	return &client
}

func (c *OpenAIClient) PromptWithoutContext(system string, messages ...string) string {
	return c.Prompt(context.Background(), system, messages...)
}

func (c *OpenAIClient) Prompt(ctx context.Context, system string, messages ...string) string {
	return "Hello, World!"
}
