package server

import (
	"todoapplication/services/api"

	"github.com/labstack/echo/v4"
)

type ServerImpl interface {
	CreateTodo(c echo.Context) error
	GetTodoByID(c echo.Context) error
	GetAllTodo(c echo.Context) error
}

type Server struct {
	api *api.TaskAPIImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewTaskAPIImpl(),
	}
}

func NewServerImpl(e *echo.Echo) {
	server := NewServer()
	e.POST("/todo", server.CreateTodo)
	e.GET("/todo/:user_id", server.GetTodoByID)
	e.GET("/todo", server.GetAllTodo)
}

var _ ServerImpl = &Server{}
