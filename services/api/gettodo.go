package api

import (
	"todoapplication/models"
	api_errors "todoapplication/services/api/errors"
)

func (api *TaskAPIImpl) GetTodoByIDAPI(id int64) (*models.TodoReponse, error) {
	todoResp, err := api.db.GetTodoByIDDB(id)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return todoResp, nil
}
