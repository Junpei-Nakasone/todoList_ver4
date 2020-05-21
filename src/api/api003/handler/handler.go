package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/todoList_ver2/api/api003/domain"
	"github.com/todoList_ver2/environment/db"
)

// DeleteTodo returns todo
func DeleteTodo(c echo.Context) error {
	param := domain.RequestParam{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	db := db.CreateDBConnection()

	data := domain.Todo{}
	db.Table("todos").
		Where("id = ?", param.ID).
		Find(&data).
		Delete(data)

	return c.JSON(http.StatusAccepted, "Deleted")
}

func addTodo(param domain.Todo) (err error) {

	db := db.CreateDBConnection()

	err = db.Create(&param).Error
	if err != nil {
		return err
	}
	return err

}
