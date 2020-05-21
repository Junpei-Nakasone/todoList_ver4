package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/todoList_ver2/api/api004/domain"
	"github.com/todoList_ver2/environment/db"
)

// UpdateTodo returns todo
func UpdateTodo(c echo.Context) error {
	param := domain.Todo{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	db := db.CreateDBConnection()

	updateData := &domain.UpdateData{
		Name: param.Name,
	}

	err := db.Table("todos").
		Where("id = ?", param.ID).
		UpdateColumns(updateData).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "UPdated")
}
