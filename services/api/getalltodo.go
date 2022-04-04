package api

import (
	"todoapplication/models"
	api_errors "todoapplication/services/api/errors"
)

func (api *TaskAPIImpl) GetAllTodoAPI() ([]*models.TodoReponse, error) {

	todoresp, err := api.db.GetAllTodoDB()
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return todoresp, nil
}
