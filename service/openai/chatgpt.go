package openai

import (
	"context"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"github.com/sashabaranov/go-openai"
)

var chatGPTClient *openai.Client = nil

func initClient() {
	if chatGPTClient != nil {
		return
	}
	chatGPTConfig := openai.DefaultConfig(config.String("application.openai.token"))
	chatGPTConfig.BaseURL = config.String("application.openai.url")
	chatGPTClient = openai.NewClientWithConfig(chatGPTConfig)
}

func GetChatGPTAck(question string) string {
	initClient()
	resp, err := chatGPTClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		log.Infof("ChatGPT 处理错误: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}
