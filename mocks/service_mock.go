package mocks

import (
	"context"

	"github.com/RoryRaeper/n-able-task-app/models"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) ListTasks(ctx context.Context, offset, limit int64) ([]models.Task, error) {
	args := m.Called(ctx, offset, limit)
	task := args.Get(0)
	if task == nil {
		return nil, args.Error(1)
	}
	return task.([]models.Task), args.Error(1)
}

func (m *MockService) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) UpdateTask(ctx context.Context, id string, task models.Task) (*models.Task, error) {
	args := m.Called(ctx, id, task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
