package config

import (
	"context"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

type LLMClient struct {
	Client *openai.Client
}

func NewOpenAI() *LLMClient {
	return &LLMClient{
		Client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (l *LLMClient) SetPrompt(prompt string) (string, error) {
	// 5초 Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := l.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: "너는 끝말잇기 전문가야"}, // 모델 설정
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
			MaxTokens: 10,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
