package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/todoList_ver2/api/api001/domain"
	"github.com/todoList_ver2/environment/db"
)

// ShowTodos returns todo
func ShowTodos(c echo.Context) error {
	data, err := getTodos()
	if err != nil {
		return err
	}
	response := domain.ResponseData{
		Todos: data,
	}

	return c.JSON(http.StatusOK, response)
}

func getTodos() (tc []domain.Todo, err error) {

	db := db.CreateDBConnection()

	response := []domain.Todo{}

	err = db.Find(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil

}
