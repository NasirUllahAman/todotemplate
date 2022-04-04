package models

type Todo struct {
	Id          *int64  `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
}

func NewTodo() *Todo {
	return &Todo{
		Id:          new(int64),
		Title:       new(string),
		Description: new(string),
		Status:      new(string),
	}
}
