package cat

import (
	"github.com/scuba13/AmacoonAI/model"
	"github.com/scuba13/AmacoonAI/services"
	"github.com/scuba13/AmacoonAI/config"
)

func CalculateSimilarities(question string, catAIs []model.CatAI, config *config.Config) (model.CatAI, float32) {
	var maxSimilarity float32 = -1
	var bestCatAI model.CatAI
	questionEmbedding := services.GetEmbedding(config, question)
	for _, catAI := range catAIs {
		similarity := services.CosineSimilarity(catAI.Embedding, questionEmbedding)

		if similarity > maxSimilarity {
			maxSimilarity = similarity
			bestCatAI = catAI
		}
	}
	return bestCatAI, maxSimilarity
}
