package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/scuba13/AmacoonAI/config"
	"github.com/scuba13/AmacoonAI/model"
	"github.com/scuba13/AmacoonAI/services/cat"
	"github.com/scuba13/AmacoonAI/services"
)


func main() {
	
	cfg := config.LoadConfig()
	
	mongo, err := config.SetupMongo()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}
	// Connection For Populate CatAI Mongo Collection
	// db, err := config.SetupDB(cfg)
	// if err != nil {
	// 	fmt.Printf("Error connecting to DB: %v\n", err)
	// 	return
	// }
	// // For Populate CatAIMongo Collection
	// cat.PopulateCatAI(db, mongo, cfg)



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
		collection := mongo.Database("amacoon").Collection("cats_ai")
		var catAIs []model.CatAI
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
		bestCatAI, maxSimilarity := cat.CalculateSimilarities(question, catAIs, cfg)

		// gerar o prompt com o registro de gato mais similar encontrado
		fmt.Printf("Best match: Cat ID %d (Similarity: %f)\n", bestCatAI.CatID, maxSimilarity)
		prompt := services.GeneratePrompt(bestCatAI.Summary, question)
		fmt.Printf("Prompt: %s\n", prompt)

		// fazer a chamada da API OpenAI
		answer, err := services.QueryOpenAI(prompt, cfg)
		if err != nil {
			fmt.Printf("Error querying OpenAI API: %v\n", err)
			continue
		}

		// exibir o resultado
		fmt.Printf("Answer: %s\n", answer)
	}
}
