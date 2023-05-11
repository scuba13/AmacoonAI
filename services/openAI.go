package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"fmt"
	"github.com/scuba13/AmacoonAI/config"
	"strings"
)


func QueryOpenAI(prompt string, config *config.Config) (string, error) {
	
	client := openai.NewClient(config.OpeinAIKey)

	response, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 40,
		Prompt:    prompt,
	})

	if err != nil {
		return "", fmt.Errorf("Error querying OpenAI API: %v", err)
	}

	answer := response.Choices[0].Text
	answer = strings.TrimSpace(answer)

	return answer, nil
}



func GeneratePrompt(context, question string) string {
	prompt := "Responda a pergunta abaixo, somente se você tiver 100% de certeza.\n"
	prompt += "Context: " + context + "\n"
	prompt += "Q: " + question + "\n"
	prompt += "A: "
	return prompt
}

func GetEmbedding(config *config.Config, summary string) []float32 {
	
	client := openai.NewClient(config.OpeinAIKey)

	// Configure os parâmetros da chamada de API
	params := openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{summary},
		User:  "Scuba13",
	}

	// Faça a chamada da API para obter os embeddings
	embeddings, err := client.CreateEmbeddings(context.Background(), params)
	if err != nil {
		fmt.Printf("Error fetching embeddings: %v\n", err)
		return nil
	}

	return embeddings.Data[0].Embedding
}