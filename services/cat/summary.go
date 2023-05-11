package cat

import (
	"context"
	"fmt"

	"github.com/scuba13/AmacoonAI/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateCatSummary(cat model.Cat) string {
	fmt.Println("catSummaryName :", cat.Name)
	fmt.Println("catSummaryBreed :", cat.Breed.BreedName)
	fmt.Println("catSummaryCattery :", cat.CatteryID)

	summary := fmt.Sprintf("O gato chamado %s é da raça %s e tem a cor %s. ", cat.Name, cat.Breed.BreedName, cat.Color.Name)
	summary += fmt.Sprintf("O gênero é %s e nasceu em %s. ", cat.Gender, cat.Birthdate.Format("02/01/2006"))

	if cat.FatherName != "" && cat.MotherName != "" {
		summary += fmt.Sprintf("Seu pai é o gato chamado %s e sua mãe é a gata chamada %s. ", cat.FatherName, cat.MotherName)
	}

	if cat.Cattery != nil {
		if cat.Cattery.BreederName != "" && cat.Cattery.Name != "" {
			summary += fmt.Sprintf("O criador é %s e o gatil é chamado de %s. ", cat.Cattery.BreederName, cat.Cattery.Name)
		}
	}

	if cat.Owner != nil && cat.Country != nil {
		summary += fmt.Sprintf("O proprietário do gato é %s e o país de origem é %s. ", cat.Owner.Name, cat.Country.Name)
	}

	if cat.Federation != nil {
		summary += fmt.Sprintf("A federação a qual pertence é a %s. ", cat.Federation.Name)
	}

	if cat.Registration != "" {
		summary += fmt.Sprintf("O numero de registro é %s. ", cat.Registration)
	}

	if cat.Microchip != "" {
		summary += fmt.Sprintf("O microchip implantado é %s.", cat.Microchip)
	}

	return summary
}

func insertCatAI(mongo *mongo.Client, cat model.CatAI) error {

	collection := mongo.Database("amacoon").Collection("cats_ai")
	_, err := collection.InsertOne(context.Background(), cat)
	if err != nil {
		return fmt.Errorf("failed to insert CatAI: %v", err)
	}

	return nil
}
