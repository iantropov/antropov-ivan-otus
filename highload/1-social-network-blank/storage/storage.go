package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"social-network/config"
	"social-network/types"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(params types.UserRegisterParams) (string, error) {
	hashedPassword, err := hashPassword(*params.Password)
	if err != nil {
		return "", err
	}

	var createUserError error
	userId := uuid.New().String()

	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO users (id, first_name, second_name, age, password, biography, city) VALUES (?, ?, ?, ?, ?, ?, ?)",
			userId,
			*params.FirstName,
			*params.SecondName,
			*params.Age,
			hashedPassword,
			params.Biography,
			params.City,
		)
		if err != nil {
			createUserError = fmt.Errorf("createUser: %v", err)
		}
	})

	if createUserError != nil {
		return "", createUserError
	}

	return userId, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func LoginUser(userId, userPassword string) (*types.UserRecord, error) {
	userRecord, err := GetUser(userId)
	if err != nil {
		return nil, err
	}

	if !checkPassword(userPassword, userRecord.HashedPassword) {
		return nil, errors.New("invalid credentials")
	}

	return userRecord, nil
}

func GetUser(userId string) (*types.UserRecord, error) {
	var getUserError error
	var userRecord types.UserRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query("SELECT * FROM users WHERE id = ? LIMIT 1", userId)
		if err != nil {
			getUserError = fmt.Errorf("getUser: %q: %v", userId, err)
			return
		}
		defer rows.Close()

		if rows.Next() {
			if err := rows.Scan(
				&userRecord.Id,
				&userRecord.FirstName,
				&userRecord.SecondName,
				&userRecord.Age,
				&userRecord.HashedPassword,
				&userRecord.Biography,
				&userRecord.City,
			); err != nil {
				getUserError = fmt.Errorf("getUser: %q: %v", userId, err)
				return
			}
		} else {
			getUserError = fmt.Errorf("getUser: %q: not_found", userId)
		}
	})

	if getUserError != nil {
		return nil, getUserError
	}

	return &userRecord, nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
