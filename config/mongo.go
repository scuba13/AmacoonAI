package config

import (
	"fmt"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo(config *Config) (*mongo.Client, error) {
	// Define a string de conexão com o MongoDB
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.MongoDBUsername,
		config.MongoDBPassword,
		config.MongoDBHost,
		config.MongoDBPort,
		config.MongoDBName,
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
