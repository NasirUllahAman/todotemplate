package server

import (
	"net/http"
	"todoapplication/models"
	_ "todoapplication/models"
	server_errors "todoapplication/services/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreateTodo(c echo.Context) error {
	todoData := models.NewTodo()
	if err := c.Bind(&todoData); err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	todoResp, err := s.api.CreateTodoAPI(todoData)
	if err != nil {
		return server_errors.NewUnauthorizedError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Your todo has been recorded",
		Data:    todoResp,
	}
	return c.JSON(http.StatusOK, result)
}
