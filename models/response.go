package model

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TodoReponse struct {
	Id          *int64  `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
}

func NewTodoReponse() *TodoReponse {
	return &TodoReponse{
		Id:          new(int64),
		Title:       new(string),
		Description: new(string),
		Status:      new(string),
	}
}
