package api

import (
	"todoapplication/models"
	"todoapplication/services/db"
)

type TaskAPI interface {
	CreateTodoAPI(tododata *models.Todo) (*models.TodoReponse, error)
	GetAllTodoAPI() ([]*models.TodoReponse, error)
	GetTodoByIDAPI(id int64) (*models.TodoReponse, error)
}
type TaskAPIImpl struct {
	db *db.TaskDBImpl
}

func NewTaskAPIImpl() *TaskAPIImpl {
	dbImpl := db.NewTaskDBImpl()
	return &TaskAPIImpl{
		db: dbImpl,
	}
}

var _ TaskAPI = &TaskAPIImpl{}
