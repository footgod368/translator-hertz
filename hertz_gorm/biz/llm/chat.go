package llm

import (
	"context"
	ark "github.com/sashabaranov/go-openai"
)

func Chat(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	resp, err := LLMClient.CreateChatCompletion(
		ctx,
		ark.ChatCompletionRequest{
			Model: modelID,
			Messages: []ark.ChatCompletionMessage{
				{
					Role:    ark.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    ark.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
