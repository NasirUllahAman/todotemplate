package db

import (
	"todoapplication/models"
	db_errors "todoapplication/services/db/errors"
)

func (db *TodoDBImpl) GetTodoByIDDB(id int64) (*models.TodoReponse, error) {
	todoDB := models.NewTodoReponse()
	err := db.dbConn.Get(todoDB, `SELECT * FROM userdata WHERE user_id=?;`, id)
	if err != nil {
		return nil, db_errors.NewError("no record found")
	}
	return todoDB, nil
}

func (db *TodoDBImpl) GetAllTodoDB() ([]*models.TodoReponse, error) {
	todoResp := make([]*models.TodoReponse, 0)
	err := db.dbConn.Select(&todoResp, `SELECT * FROM userdata;`)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())

	}
	return todoResp, nil
}
