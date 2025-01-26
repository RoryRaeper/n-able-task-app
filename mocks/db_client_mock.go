package mocks

import (
	"context"

	"github.com/RoryRaeper/n-able-task-app/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockDBClient) GetTasks(ctx context.Context, limit, offset int64) ([]models.Task, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockDBClient) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockDBClient) UpdateTask(ctx context.Context, id primitive.ObjectID, task models.Task) (*models.Task, error) {
	args := m.Called(ctx, id, task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockDBClient) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
