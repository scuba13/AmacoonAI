package cat

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonAI/model"
	"github.com/scuba13/AmacoonAI/services"
	"github.com/scuba13/AmacoonAI/config"
)

func PopulateCatAISummary(db *gorm.DB, mongo *mongo.Client, config *config.Config) {
	fmt.Println("Populating cat_ai table...")
	var cats []model.Cat
	db.Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner").
		Preload("Federation").
		Limit(100).Find(&cats)

	for _, cat := range cats {
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
		summary := generateCatSummary(cat)
		embedding := services.GetEmbedding(config, summary)
		

		catAI := model.CatAI{
			CatID:     cat.ID,
			Summary:   summary,
			Embedding: embedding,
		}

		if err := insertCatAI(mongo, catAI); err != nil {
			fmt.Printf("Error inserting cat_ai record for cat ID %d: %v\n", cat.ID, err)
		} else {
			fmt.Printf("Successfully inserted cat_ai record for cat ID %d\n", cat.ID)
		}
	}
}

func PopulateCatAI(db *gorm.DB, mongo *mongo.Client, config *config.Config) {
	fmt.Println("Populating cat_ai table...")
	var cats []model.Cat
	db.Preload("Breed").
		Preload("Color").
		Preload("Cattery").
		Preload("Country").
		Preload("Owner.Country").
		Preload("Federation").
		Preload("Titles.Titles").
		Limit(100).Find(&cats)

	for _, cat := range cats {
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
			continue
		}

		embedding := services.GetEmbedding(config, string(catJSON))

		catAI := model.CatAI{
			CatID:     cat.ID,
			Summary:   string(catJSON),
			Embedding: embedding,
		}

		if err := insertCatAI(mongo, catAI); err != nil {
			fmt.Printf("Error inserting cat_ai record for cat ID %d: %v\n", cat.ID, err)
		} else {
			fmt.Printf("Successfully inserted cat_ai record for cat ID %d\n", cat.ID)
		}
	}
}
