package services

import (
	"context"

	"github.com/RoryRaeper/n-able-task-app/models"
)

type Service interface {
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	ListTasks(ctx context.Context, offset, limit int64) ([]models.Task, error)
	CreateTask(ctx context.Context, task models.Task) (*models.Task, error)
	UpdateTask(ctx context.Context, id string, task models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id string) error
}
