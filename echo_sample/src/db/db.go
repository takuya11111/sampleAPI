package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.Printf("Connecting db...")
	DB, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/simpleapi")

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected: ", DB)
}
