package main

import (
	"net/http"

	"github.com/takuya11111/echo_sample/src/db"

	//"github.com/takuya11111/echo_sample/src/models"
	"github.com/takuya11111/echo_sample/src/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// DB接続
	db.Init()

	server := echo.New()
	routes.RegisterRoutes(server)
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//server.GET("/todo_items", models.GetAllUsers)
	//userModel := models.GetAllUsers(db)

	//userModel := models.NewUserModel(db)

	server.Logger.Fatal(server.Start(":8000"))
}
