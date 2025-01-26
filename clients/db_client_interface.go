package clients

import (
	"context"

	"github.com/RoryRaeper/n-able-task-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBClient interface {
	CreateTask(ctx context.Context, task models.Task) (*models.Task, error)
	GetTaskByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error)
	UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
	GetTasks(ctx context.Context, limit, offset int64) ([]models.Task, error)
}
