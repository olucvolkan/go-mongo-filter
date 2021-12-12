package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port  int
	Mongo struct {
		URL        string
		DB         string
		Collection string
	}
}

func main() {
	config := newConfig()

	mongoClient := NewMongoClient(config)
	mongoDatabase := mongoClient.Database(config.Mongo.DB)
	mongoRepo := NewMongoRepo(config, mongoClient, mongoDatabase)

	kvstore := NewInMemoryKVStore()
	http.HandleFunc("/mongo", buildMongoHandler(mongoRepo))
	http.HandleFunc("/in-memory/", buildInMemoryHandler(kvstore))

	log.Println("Starting Server")
	e := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", config.Port), nil)
	log.Fatal(e)
}

func newConfig() *Config {
	config := &Config{}

	if os.Getenv("PORT") == "" {
		log.Fatal("Wrong PORT info, failed to start the app")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Can't parse PORT, failed to start the app")
	}

	config.Port = port

	if os.Getenv("MONGO_URL") == "" {
		log.Fatal("MONGO_URL env variable is required")
	}

	config.Mongo.URL = os.Getenv("MONGO_URL")

	if os.Getenv("MONGO_DB") == "" {
		log.Fatal("MONGO_DB env variable is required")
	}

	config.Mongo.DB = os.Getenv("MONGO_DB")

	return config
}

func NewMongoClient(c *Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(c.Mongo.URL))
	_ = client.Ping(ctx, nil)
	return client
}
