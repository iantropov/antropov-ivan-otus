package storage

import (
	"database/sql"
	"fmt"
	"os"
	"social-network/params"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const DB_NAME = "social-network"
const CREATE_USERS_TABLE_QUERY = `
CREATE TABLE users (
	id        VARCHAR(36) NOT NULL,
	first_name      VARCHAR(128) NOT NULL,
	second_name     VARCHAR(128) NOT NULL,
	age     INT NOT NULL,
	password     VARCHAR(128) NOT NULL,
	biography     VARCHAR(255),
	city     VARCHAR(64),
	PRIMARY KEY (%sid%s)
);
`

func Init() {
	queryDb(func(db *sql.DB) {
		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
		fmt.Println("Connected to Database!")

		rows, err := db.Query("show tables")
		if err != nil {
			panic(pingErr)
		}
		defer rows.Close()

		if !rows.Next() {
			fmt.Println("No Users table found")
			_, err := db.Query(fmt.Sprintf(CREATE_USERS_TABLE_QUERY, "`", "`"))
			if err != nil {
				panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
			}
			fmt.Println("Created Users table")
		}
	})
}

func CreateUser(params params.UserParams) (string, error) {
	var userError error
	userId := uuid.New().String()
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO users (id, first_name, second_name, age, password, biography, city) VALUES (?, ?, ?, ?, ?, ?, ?)",
			userId,
			*params.FirstName,
			*params.SecondName,
			*params.Age,
			*params.Password,
			params.Biography,
			params.City,
		)
		if err != nil {
			userError = fmt.Errorf("createUser: %v", err)
		}
	})

	if userError != nil {
		return "", userError
	}

	return userId, nil
}

func queryDb(callback func(db *sql.DB)) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "social-network",
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
