package server

import (
	"net/http"
	"strconv"
	"todoapplication/models"
	server_errors "todoapplication/services/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetTodoByID(c echo.Context) error {
	userNoStr := c.Param("user_id")
	idNo, err := strconv.ParseInt(userNoStr, 0, 64)
	if err != nil {
		server_errors.NewInternalServerError(err.Error())
	}
	todoRecordResp, err := s.api.GetTodoByIDAPI(idNo)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Successfully fetched the todo",
		Data:    todoRecordResp,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) GetAllTodo(c echo.Context) error {
	getTodoResp, err := s.api.GetAllTodoAPI()
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Successfully fetched all the todo",
		Data:    getTodoResp,
	}
	return c.JSON(http.StatusOK, result)
}
