package db

import (
	"fmt"
	"todoapplication/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TaskDB interface {
	CreateTodoDB(tododata *models.Todo) (*models.TodoReponse, error)
	GetTodoByIDDB(id int64) (*models.TodoReponse, error)
	GetAllTodoDB() ([]*models.TodoReponse, error)
}

type TodoDBImpl struct {
	dbConn *sqlx.DB
}

func NewTodoDBImpl() *TodoDBImpl {
	dbConn, err := sqlx.Connect("mysql", "root:Nasir@12345@tcp(127.0.0.1:3306)/abcDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	return &TodoDBImpl{
		dbConn: dbConn,
	}
}

var _ TaskDB = &TodoDBImpl{}
