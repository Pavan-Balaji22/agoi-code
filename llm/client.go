package llm

import (
	"context"
	"log"

	"github.com/ollama/ollama/api"
)

type OllamaClient struct {
	client    *api.Client
	modelName string
	messages  []api.Message
}

func NewOllamaClient(modelName string) *OllamaClient {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v", err)
	}
	return &OllamaClient{
		client:    client,
		modelName: modelName,
		messages:  []api.Message{},
	}
}

func (llm *OllamaClient) CallLLM(message string) (string, error) {
	ctx := context.Background()
	var response api.ChatResponse
	llm.messages = append(llm.messages, api.Message{
		Role:    "user",
		Content: message,
	})
	stream := false
	req := &api.ChatRequest{
		Model:    llm.modelName,
		Messages: llm.messages,
		Stream:   &stream,
	}
	respFunc := func(resp api.ChatResponse) error {
		response = resp
		return nil
	}
	err := llm.client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Println(err)
	}
	return response.Message.Content, err
}
