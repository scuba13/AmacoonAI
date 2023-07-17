package services

import (
    "context"
    "github.com/sashabaranov/go-openai"
    "fmt"
    "github.com/scuba13/AmacoonAI/config"
    "strings"
)

func QueryOpenAI(prompt string, question string, config *config.Config) (string, error) {
    client := openai.NewClient(config.OpeinAIKey)

    // Gera um objeto Messages com o prompt como uma mensagem do sistema e a pergunta como uma mensagem do usuário
    messages := []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleSystem,
            Content: prompt,
        },
        {
            Role:    openai.ChatMessageRoleUser,
            Content: question,
        },
    }

    response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
        Model: openai.GPT4,  // Modifique o modelo conforme necessário
        Messages: messages,
    })

    if err != nil {
        return "", fmt.Errorf("Error querying OpenAI API: %v", err)
    }

    // A resposta é a última mensagem do chat, que é a resposta da IA
    answer := response.Choices[0].Message.Content
    answer = strings.TrimSpace(answer)

    // Imprime a quantidade de tokens da pergunta e da resposta
    fmt.Printf("Question length: %d tokens\n", len(question))
    fmt.Printf("Answer length: %d tokens\n", len(answer))

    return answer, nil
}
