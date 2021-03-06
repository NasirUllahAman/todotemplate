package api

import (
	"todoapplication/models"
	api_errors "todoapplication/services/api/errors"
)

func (api *TaskAPIImpl) CreateTodoAPI(tododata *models.Todo) (*models.TodoReponse, error) {
	todoresp, err := api.db.CreateTodoDB(tododata)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return todoresp, nil
}
