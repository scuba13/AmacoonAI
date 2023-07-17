package main

import (
	"bufio"
	//"context"
	"fmt"
	"os"
	"strings"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/scuba13/AmacoonAI/config"
	//"github.com/scuba13/AmacoonAI/model"
	"github.com/scuba13/AmacoonAI/services/cat"
	"github.com/scuba13/AmacoonAI/services"
	"strconv"
)


func main() {
	
	cfg := config.LoadConfig()
	db, _:= config.SetupDB(cfg)
	



// loop principal do programa
for {
    // Ler o ID do gato
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a cat ID: ")
    catIDString, err := reader.ReadString('\n')
    if err != nil {
        fmt.Printf("Error reading input: %v\n", err)
        continue
    }
    catIDString = strings.TrimSpace(catIDString)
    catID, err := strconv.ParseUint(catIDString, 10, 64)
    if err != nil {
        fmt.Printf("Error parsing cat ID: %v\n", err)
        continue
    }

    // Chamar FindCatAI para obter os dados do gato
    fmt.Println("Finding cat...")
    catJSON, err := cat.FindCatAI(db, cfg, uint(catID))
    if err != nil {
        fmt.Printf("Error populating cat AI: %v\n", err)
        continue
    }
    fmt.Println("Finding cat OK")

    // loop de perguntas
    for {
        // Solicitar a pergunta do usuário
        fmt.Print("Enter a question (or 'exit' to choose another cat): ")
        question, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Error reading input: %v\n", err)
            continue
        }
        question = strings.TrimSpace(question)

        // se o usuário inserir "exit", sair do loop de perguntas
        if strings.ToLower(question) == "exit" {
            break
        }

       
       // fmt.Printf("Prompt: %s\n", prompt)

        // Fazer a chamada da API OpenAI
        answer, err := services.QueryOpenAI(catJSON, question, cfg)
        if err != nil {
            fmt.Printf("Error querying OpenAI API: %v\n", err)
            continue
        }

        // Exibir o resultado
        fmt.Printf("Answer: %s\n", answer)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
    }
}


}
