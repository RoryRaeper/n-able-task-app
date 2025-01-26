package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

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
	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.ListTasks)
	router.GET("/tasks/:id", handler.ListTasks)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	// Start the HTTP server
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
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
