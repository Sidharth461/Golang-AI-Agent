package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type Client struct {
	api   *openai.Client
	model string
}

func NewClient() (*Client, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENROUTER_API_KEY is required")
	}

	baseurl := os.Getenv("OPENROUTER_URL")
	model := os.Getenv("MODEL_NAME")
	api := openai.NewClient(
		option.WithBaseURL(baseurl),
		option.WithAPIKey(apiKey),
	)

	return &Client{
		api:   &api,
		model: model,
	}, nil
}

func (c *Client) Chat(ctx context.Context, message []openai.ChatCompletionMessageParamUnion) (string, error) {
	params := openai.ChatCompletionNewParams{
		Model:    c.model,
		Messages: message,
	}

	res, err := c.api.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("OpenRouter request failed: %w", err)
	}
	if len(res.Choices) == 0 {
		return "", fmt.Errorf("OpenRouter returned an empty response")
	}

	return res.Choices[0].Message.Content, nil
}
