package main

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var key string = "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"

func main() {
	db, err := SetupMongo()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}

	// loop principal do programa
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a question: ")
		question, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}
		question = strings.TrimSpace(question)

		// buscar todos os registros da coleção cat_ai
		collection := db.Database("amacoon").Collection("cats_ai")
		var catAIs []CatAI
		cur, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			fmt.Printf("Error querying cat_ai collection: %v\n", err)
			continue
		}

		defer cur.Close(context.Background())
		if err := cur.All(context.Background(), &catAIs); err != nil {
			fmt.Printf("Error querying cat_ai collection: %v\n", err)
			continue
		}

		// calcular a similaridade da pergunta com cada embedding
		var maxSimilarity float32 = -1
		var bestCatAI CatAI
		questionEmbedding := getEmbedding(question)
		for _, catAI := range catAIs {
			// calcular a similaridade do embedding com a pergunta
			//fmt.Println("Cat ID: ", catAI.CatID)
			embedding := catAI.Embedding
			
			//fmt.Printf("Question Embedding: %v\n", questionEmbedding)
			similarity := cosineSimilarity(embedding, questionEmbedding)

			// atualizar o melhor resultado se a similaridade for maior que a atual
			if similarity > maxSimilarity {
				maxSimilarity = similarity
				bestCatAI = catAI
			}
		}

		// gerar o prompt com o registro de gato mais similar encontrado
		fmt.Printf("Best match: Cat ID %d (Similarity: %f)\n", bestCatAI.CatID, maxSimilarity)
		prompt := generatePrompt(bestCatAI.Summary, question)
		fmt.Printf("Prompt: %s\n", prompt)

		// fazer a chamada da API OpenAI
		client := openai.NewClient(key)
		response, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
			Model:     openai.GPT3TextDavinci003,
			MaxTokens: 40,
			Prompt:    prompt,
		})
		if err != nil {
			fmt.Printf("Error querying OpenAI API: %v\n", err)
			continue
		}
		answer := response.Choices[0].Text
		answer = strings.TrimSpace(answer)

		// exibir o resultado
		fmt.Printf("Answer: %s\n", answer)
	}
}

func generatePrompt(context, question string) string {
	prompt := "Responda a pergunta abaixo, somente se você tiver 100% de certeza.\n"
	prompt += "Context: " + context + "\n"
	prompt += "Q: " + question + "\n"
	prompt += "A: "
	return prompt
}

func getEmbedding(summary string) []float32 {
	
	client := openai.NewClient(key)

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

// cosineSimilarity calculates the cosine similarity between two vectors.
func cosineSimilarity(a, b []float32) float32 {
	dotProduct := float32(0)
	magnitudeA := float32(0)
	magnitudeB := float32(0)

	for i := range a {
		dotProduct += a[i] * b[i]
		magnitudeA += a[i] * a[i]
		magnitudeB += b[i] * b[i]
	}

	magnitudeA = float32(math.Sqrt(float64(magnitudeA)))
	magnitudeB = float32(math.Sqrt(float64(magnitudeB)))

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return dotProduct / (magnitudeA * magnitudeB)
}

func vectorSimilarity(v1, v2 []float32) float32 {
	if len(v1) != len(v2) {
		panic("Vectors must have the same length")
	}
	var dotProduct float32
	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]
	}
	return dotProduct
}

func SetupMongo() (*mongo.Client, error) {
	// Define a string de conexão com o MongoDB
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		"amacoonservice",
		"2010mainecoon2010",
		"amacoon.vps-kinghost.net",
		"27017",
		"amacoon",
	)

	// Define as opções de conexão com o MongoDB
	opts := options.Client().ApplyURI(mongoURI)

	// Cria um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conecta ao MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Testa a conexão com o MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}

type CatAI struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CatID     uint               `bson:"cat_id"`
	Summary   string             `bson:"summary"`
	Embedding []float32          `bson:"embedding"`
}
