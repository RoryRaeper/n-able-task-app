package clients

import (
	"context"
	"testing"

	"github.com/RoryRaeper/n-able-task-app/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateTask(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		task := models.Task{
			Title:       "Test Task",
			Description: "This is a test task",
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		createdTask, err := client.CreateTask(context.Background(), task)
		assert.NoError(t, err)
		assert.NotNil(t, createdTask.ID)
		assert.Equal(t, task.Title, createdTask.Title)
		assert.Equal(t, task.Description, createdTask.Description)
	})
}

func TestGetTaskByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		id := primitive.NewObjectID()
		task := models.Task{
			ID:          id,
			Title:       "Test Task",
			Description: "This is a test task",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "testdb.tasks", mtest.FirstBatch, primitive.D{{"_id", id}, {"title", task.Title}, {"description", task.Description}}))

		fetchedTask, err := client.GetTaskByID(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, task.ID, fetchedTask.ID)
		assert.Equal(t, task.Title, fetchedTask.Title)
		assert.Equal(t, task.Description, fetchedTask.Description)
	})

	mt.Run("not found", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		id := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "testdb.tasks", mtest.FirstBatch))

		fetchedTask, err := client.GetTaskByID(context.Background(), id)
		assert.Error(t, err)
		assert.Nil(t, fetchedTask)
	})
}

func TestUpdateTask(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		id := primitive.NewObjectID()
		updatedTask := models.Task{
			Title:       "Updated Task",
			Description: "This is an updated task",
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		result, err := client.UpdateTask(context.Background(), id, updatedTask)
		assert.NoError(t, err)
		assert.Equal(t, updatedTask.Title, result.Title)
		assert.Equal(t, updatedTask.Description, result.Description)
	})
}

func TestDeleteTask(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		id := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := client.DeleteTask(context.Background(), id)
		assert.NoError(t, err)
	})
}

func TestGetTasks(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		client := NewMongoDBClient(mt.Client, "testdb", "tasks")
		tasks := []models.Task{
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 1",
				Description: "This is task 1",
			},
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 2",
				Description: "This is task 2",
			},
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "testdb.tasks", mtest.FirstBatch,
			primitive.D{{"_id", tasks[0].ID}, {"title", tasks[0].Title}, {"description", tasks[0].Description}},
			primitive.D{{"_id", tasks[1].ID}, {"title", tasks[1].Title}, {"description", tasks[1].Description}}))

		fetchedTasks, err := client.GetTasks(context.Background(), 0, 0)
		assert.NoError(t, err)
		assert.Len(t, fetchedTasks, 2)
		assert.Equal(t, tasks[0].Title, fetchedTasks[0].Title)
		assert.Equal(t, tasks[1].Title, fetchedTasks[1].Title)
	})
}
