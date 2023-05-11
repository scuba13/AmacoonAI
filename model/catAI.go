package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatAI struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CatID     uint               `bson:"cat_id"`
	Summary   string             `bson:"summary"`
	Embedding []float32          `bson:"embedding"`
}