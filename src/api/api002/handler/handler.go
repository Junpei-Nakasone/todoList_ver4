package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/todoList_ver2/api/api002/domain"
	"github.com/todoList_ver2/environment/db"
)

// AddTodo returns todo
func AddTodo(c echo.Context) error {
	param := domain.Todo{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	err := addTodo(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Created")
}

func addTodo(param domain.Todo) (err error) {

	db := db.CreateDBConnection()

	err = db.Create(&param).Error
	if err != nil {
		return err
	}
	return err

}
