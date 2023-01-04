package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DB_USER = "user"
const DB_PASSWORD = "password"
const DB_NAME = "social-network"

func Init() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("Connected to Database!")
}
