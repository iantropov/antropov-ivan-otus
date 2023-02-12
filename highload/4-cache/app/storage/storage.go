package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"social-network-4/config"
	"social-network-4/types"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Init() {
	queryDb(func(db *sql.DB) {
		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
		fmt.Println("Connected to Database!")
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

	if !checkPassword(userPassword, *userRecord.HashedPassword) {
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
			getUserError = &types.UserNotFoundError{
				UserId: userId,
			}
		}
	})

	if getUserError != nil {
		return nil, getUserError
	}

	return &userRecord, nil
}

func SearchUsers(firstName, lastName string) ([]types.UserRecord, error) {
	var getUserError error
	var users []types.UserRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query("SELECT * FROM users WHERE first_name LIKE CONCAT(?, '%') and second_name LIKE CONCAT(?, '%') ORDER BY id LIMIT 10", firstName, lastName)
		if err != nil {
			getUserError = fmt.Errorf("searchUsers: firstName=%q, lastName=%q: %v", firstName, lastName, err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user types.UserRecord
			if err := rows.Scan(
				&user.Id,
				&user.FirstName,
				&user.SecondName,
				&user.Age,
				&user.HashedPassword,
				&user.Biography,
				&user.City,
			); err != nil {
				getUserError = fmt.Errorf("searchUsers: firstName=%q, lastName=%q: %v", firstName, lastName, err)
				return
			}
			users = append(users, user)
		}
	})

	if getUserError != nil {
		return nil, getUserError
	}

	return users, nil
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
