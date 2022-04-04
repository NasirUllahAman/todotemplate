package db

import (
	"todoapplication/models"
	db_errors "todoapplication/services/db/errors"
)

func (db *TodoDBImpl) CreateTodoDB(tododata *models.Todo) (*models.TodoReponse, error) {
	todoresp := models.NewTodoReponse()
	tx := db.dbConn.MustBegin()
	_, err := tx.NamedQuery(`INSERT INTO userdata(,email,description,status)
		 VALUES(:name,:email,:order_no,:phone_no,:message);`, tododata)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	err = db.dbConn.Get(todoresp, `SELECT * FROM userdata WHERE userid=?;`, *tododata.Id)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return todoresp, nil
}
