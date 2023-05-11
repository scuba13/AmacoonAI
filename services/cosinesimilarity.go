package services

import (
	"math"
)

// cosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(a, b []float32) float32 {
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

func VectorSimilarity(v1, v2 []float32) float32 {
	if len(v1) != len(v2) {
		panic("Vectors must have the same length")
	}
	var dotProduct float32
	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]
	}
	return dotProduct
}
