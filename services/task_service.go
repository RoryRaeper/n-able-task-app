package services

import (
	"context"
	"fmt"

	"github.com/RoryRaeper/n-able-task-app/clients"
	"github.com/RoryRaeper/n-able-task-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	dbClient clients.DBClient
}

func NewTaskService(dbClient clients.DBClient) *TaskService {
	return &TaskService{dbClient}
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}
	return s.dbClient.GetTaskByID(ctx, objectID)
}

func (s *TaskService) ListTasks(ctx context.Context, offset, limit int64) ([]models.Task, error) {
	return s.dbClient.GetTasks(ctx, limit, offset)
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	return s.dbClient.CreateTask(ctx, task)
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, task models.Task) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}
	return s.dbClient.UpdateTask(ctx, objectID, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}
	return s.dbClient.DeleteTask(ctx, objectID)
}
