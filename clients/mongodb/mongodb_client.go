package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/RoryRaeper/n-able-task-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	collection *mongo.Collection
}

// Initialises the mongoDB client
func NewMongoDBClient(client *mongo.Client, dbName, collectionName string) *MongoDBClient {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBClient{collection}
}

// Creates a new task based on the provided task model
// If successful, the task object is returned with its generated ID
func (c *MongoDBClient) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	t := time.Now()
	task.CreatedAt = t
	task.UpdatedAt = t

	result, err := c.collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = result.InsertedID.(primitive.ObjectID)
	return &task, nil
}

// Fetches the task with the provided ID from the database
func (c *MongoDBClient) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	var task models.Task
	err := c.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("task with id %s was not found", id.Hex())
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Updates the task with the provided ID with the updated task information
func (c *MongoDBClient) UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask models.Task) (*models.Task, error) {
	updatedTask.UpdatedAt = time.Now()
	update := bson.M{
		"$set": updatedTask,
	}

	_, err := c.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

// Deletes the task with the provided ID
func (c *MongoDBClient) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	_, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Returns all the tasks in the database with optional pagination/limiting/offset
func (c *MongoDBClient) GetTasks(ctx context.Context, limit, offset int64) ([]models.Task, error) {
	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(limit)
	}

	if offset > 0 {
		opts.SetSkip(offset)
	}

	cursor, err := c.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
