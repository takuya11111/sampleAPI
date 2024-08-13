package models

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya11111/echo_sample/src/db"
)

type TodoItem struct {
	// mysqlのカラム名をタグとして記載する
	ID       string `db:"id" json:"id"`
	TodoItem string `db:"todo" json:"todo"`
}

// 全ユーザー取得処理
func GetAllUsers() ([]TodoItem, error) {
	todoItems := []TodoItem{}
	if db.DB.Find(&todoItems).Error != nil {
		return nil, echo.ErrNotFound
	}
	return todoItems, nil
}

// ユーザー取得処理(id)
func GetUserById(id string) (*TodoItem, error) {
	todoItem := TodoItem{}
	if db.DB.Where("id = ?", id).First(&todoItem).Error != nil {
		return nil, echo.ErrNotFound
	}
	return &todoItem, nil
}
