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

func (c *MongoDBClient) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	task.ID = primitive.NewObjectID()
	t := time.Now()
	task.CreatedAt = t
	task.UpdatedAt = t

	_, err := c.collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

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

func (c *MongoDBClient) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	_, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

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
