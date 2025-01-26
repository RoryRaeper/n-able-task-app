package main

import (
	"context"
	"log"
	"os"

	clients "github.com/RoryRaeper/n-able-task-app/clients/mongodb"
	"github.com/RoryRaeper/n-able-task-app/handlers"
	"github.com/RoryRaeper/n-able-task-app/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialise MongoDB client
	client := initMongoDB()
	defer client.Disconnect(context.Background())

	// Initialise database client, service, and handler
	repo := clients.NewMongoDBClient(client, "TASK_STORE", "TASKS")
	service := services.NewTaskService(repo)
	handler := handlers.NewHandler(service)
	router := gin.Default()
	router.POST("/users", handler.CreateTask)
	router.GET("/users", handler.ListTasks)
	router.GET("/users/:id", handler.ListTasks)
	router.PUT("/users/:id", handler.UpdateTask)
	router.DELETE("/users/:id", handler.DeleteTask)
}

func initMongoDB() *mongo.Client {
	// MongoDB connection URI
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ensure MongoDB is connected
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
