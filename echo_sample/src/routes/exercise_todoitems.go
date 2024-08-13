package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takuya11111/echo_sample/src/models"
)

func getAllUsers(context echo.Context) error {
	todoItems, err := models.GetAllUsers()
	// ※注意：エラーハンドリングはテキトーです
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, todoItems)
}

func getUserById(context echo.Context) error {
	id := context.Param("id")
	todoItem, err := models.GetUserById(id)
	// ※注意：エラーハンドリングはテキトーです
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, todoItem)
}
