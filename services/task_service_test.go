package services

import (
	"context"
	"testing"

	"github.com/RoryRaeper/n-able-task-app/mocks"
	"github.com/RoryRaeper/n-able-task-app/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetTaskByID(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	taskID := primitive.NewObjectID()
	task := &models.Task{ID: taskID, Title: "Test Task"}

	mockDBClient.On("GetTaskByID", ctx, taskID).Return(task, nil)

	result, err := taskService.GetTaskByID(ctx, taskID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, task, result)

	mockDBClient.AssertExpectations(t)
}

func TestGetTaskByID_InvalidID(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()

	_, err := taskService.GetTaskByID(ctx, "invalidID")
	assert.Error(t, err)
	assert.Equal(t, "invalid ID format", err.Error())
}

func TestListTasks(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	tasks := []models.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockDBClient.On("GetTasks", ctx, int64(10), int64(0)).Return(tasks, nil)

	result, err := taskService.ListTasks(ctx, 0, 10)
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)

	mockDBClient.AssertExpectations(t)
}

func TestCreateTask(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	task := models.Task{Title: "New Task"}

	mockDBClient.On("CreateTask", ctx, task).Return(&task, nil)

	result, err := taskService.CreateTask(ctx, task)
	assert.NoError(t, err)
	assert.Equal(t, &task, result)

	mockDBClient.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	taskID := primitive.NewObjectID()
	updatedTask := models.Task{Title: "Updated Task"}

	mockDBClient.On("UpdateTask", ctx, taskID, updatedTask).Return(&updatedTask, nil)

	result, err := taskService.UpdateTask(ctx, taskID.Hex(), updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, &updatedTask, result)

	mockDBClient.AssertExpectations(t)
}

func TestUpdateTask_InvalidID(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	updatedTask := models.Task{Title: "Updated Task"}

	_, err := taskService.UpdateTask(ctx, "invalidID", updatedTask)
	assert.Error(t, err)
	assert.Equal(t, "invalid ID format", err.Error())
}

func TestDeleteTask(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()
	taskID := primitive.NewObjectID()

	mockDBClient.On("DeleteTask", ctx, taskID).Return(nil)

	err := taskService.DeleteTask(ctx, taskID.Hex())
	assert.NoError(t, err)

	mockDBClient.AssertExpectations(t)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	mockDBClient := new(mocks.MockDBClient)
	taskService := NewTaskService(mockDBClient)
	ctx := context.TODO()

	err := taskService.DeleteTask(ctx, "invalidID")
	assert.Error(t, err)
	assert.Equal(t, "invalid ID format", err.Error())
}
