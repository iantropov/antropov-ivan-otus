package storage

import (
	"database/sql"
	"fmt"
	"users-app/config"
	"users-app/types"

	"github.com/go-sql-driver/mysql"
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

func CreateUser(params types.UserParams) (int64, error) {
	var createUserError error
	var userId int64
	queryDb(func(db *sql.DB) {
		result, err := db.Exec(
			"INSERT INTO users (username, first_name, last_name, email, phone) VALUES (?, ?, ?, ?, ?)",
			*params.Username,
			*params.FirstName,
			*params.LastName,
			*params.Email,
			*params.Phone,
		)
		if err != nil {
			createUserError = fmt.Errorf("createUser: %v", err)
			return
		}
		userId, err = result.LastInsertId()
		if err != nil {
			createUserError = fmt.Errorf("createUser: %v", err)
		}
	})

	return userId, createUserError
}

func UpdateUser(userId int64, params types.UserParams) error {
	var updateUserError error
	queryDb(func(db *sql.DB) {
		result, err := db.Exec(
			"UPDATE users SET username=?, first_name=?, last_name=?, email=?, phone=? WHERE id=?",
			*params.Username,
			*params.FirstName,
			*params.LastName,
			*params.Email,
			*params.Phone,
			userId,
		)
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			updateUserError = fmt.Errorf("deleteUser: %q: %v", userId, err)
		} else if rowsAffected == 0 {
			updateUserError = &types.UserNotFoundError{
				UserId: userId,
			}
		}
	})

	return updateUserError
}

func GetUser(userId int64) (*types.UserRecord, error) {
	var getUserError error
	var userRecord types.UserRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query("SELECT * FROM users WHERE id = ?", userId)
		if err != nil {
			getUserError = fmt.Errorf("getUser: %q: %v", userId, err)
			return
		}
		defer rows.Close()

		if rows.Next() {
			if err := rows.Scan(
				&userRecord.Id,
				&userRecord.Username,
				&userRecord.FirstName,
				&userRecord.LastName,
				&userRecord.Email,
				&userRecord.Phone,
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

func DeleteUser(userId int64) error {
	var deleteUserError error
	queryDb(func(db *sql.DB) {
		result, err := db.Exec("DELETE FROM users WHERE id = ?", userId)
		if err != nil {
			deleteUserError = fmt.Errorf("deleteUser: %q: %v", userId, err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			deleteUserError = fmt.Errorf("deleteUser: %q: %v", userId, err)
		} else if rowsAffected == 0 {
			deleteUserError = &types.UserNotFoundError{
				UserId: userId,
			}
		}
	})

	return deleteUserError
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
