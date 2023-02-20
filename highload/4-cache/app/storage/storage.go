package storage

import (
	"database/sql"
	"fmt"
	"social-network-4/config"

	"github.com/go-sql-driver/mysql"
)

func init() {
	queryDb(func(db *sql.DB) {
		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
		fmt.Println("Connected to Database!")
	})
}

func queryDb(callback func(db *sql.DB)) {
	cfg := mysql.Config{
		User:   config.Config("DB_USER"),
		Passwd: config.Config("DB_PASS"),
		Net:    "tcp",
		Addr:   config.Config("DB_ADDR"),
		DBName: config.Config("DB_NAME"),
		Params: map[string]string{
			"charset":              "utf8mb4",
			"allowNativePasswords": "true",
		},
	}
	fmt.Println(cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	callback(db)
}
