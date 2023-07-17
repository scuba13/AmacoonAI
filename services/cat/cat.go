package cat

import (
	"encoding/json"
	"fmt"

	"github.com/scuba13/AmacoonAI/config"
	"github.com/scuba13/AmacoonAI/model"
	"gorm.io/gorm"
)

func FindCatAI(db *gorm.DB, config *config.Config, catID uint) (string, error) {
	fmt.Println("Populating cat_ai table...")

	var cat model.Cat
	result := db.Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner.Country").
		Preload("Federation").
		Preload("Titles.Titles").
		Where("id = ?", catID).First(&cat)

	if result.Error != nil {
		// Trate o erro aqui
		fmt.Printf("Erro ao buscar o gato: %v", result.Error)
		return "", result.Error
	}

	// Agora "cat" contém o gato com o ID fornecido
	fmt.Println("cat :", cat.Name)

	if cat.FatherID != nil {
		var father model.Cat
		db.Select("name").Where("id = ?", cat.FatherID).First(&father)
		cat.FatherName = father.Name
	}

	if cat.MotherID != nil {
		var mother model.Cat
		db.Select("name").Where("id = ?", cat.MotherID).First(&mother)
		cat.MotherName = mother.Name
	}

	catJSON, err := json.Marshal(cat)
	if err != nil {
		fmt.Printf("Error marshalling cat object to JSON for cat ID %d: %v\n", cat.ID, err)
		return "", err
	}

	fmt.Printf("Processed cat AI record for cat ID %d\n", cat.ID)

	// Retornar o catJSON e nenhum erro
	return string(catJSON), nil
}
