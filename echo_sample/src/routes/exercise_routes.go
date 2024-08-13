package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.GET("/todo_items", getAllUsers)
	server.GET("/todo_items/:id", getUserById)
}
