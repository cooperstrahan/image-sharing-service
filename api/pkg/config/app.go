package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// db *gorm.DB
	db *sql.DB
)

func Connect() {
	d, err := sql.Open("mysql", "coop:Strahan1!@tcp(imagesharingserver.mysql.database.azure.com:3306)/imagesharingdatabase?parseTime=true")
	if err != nil {
		fmt.Println("Database error")
		fmt.Println(err.Error())
		// panic(err.Error())
	}

	db = d
}

// func GetDB() *gorm.DB {
// 	return db
// }

func GetDB() *sql.DB {
	return db
}
