package storage

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func GetDialogId(userId1, userId2 string) (string, error) {
	var queryError error
	var dialogId string
	queryDb(func(db *sql.DB) {
		rows, err := db.Query(
			"SELECT dialogId FROM users_dialogs WHERE user_id_1 = ? AND user_id_2 = ? OR user_id_1 = ? AND user_id_2 = ?",
			userId1,
			userId2,
			userId2,
			userId1,
		)
		if err != nil {
			queryError = fmt.Errorf("GetDialog: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&dialogId); err != nil {
				queryError = fmt.Errorf("GetDialog: %w", err)
			}
		}
	})

	return dialogId, queryError
}

func CreateDialog(name string) (string, error) {
	dialogId := uuid.New().String()
	var createDialogError error
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO dialogs (id, name) VALUES (?, ?)",
			dialogId,
			name,
		)
		if err != nil {
			createDialogError = fmt.Errorf("CreateDialog: %w", err)
		}
	})

	return dialogId, createDialogError
}

func LinkUsersWithDialog(userId1, userId2, dialogId string) error {
	var execError error
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO users_dialogs (user_id_1, user_id_2, dialog_id) VALUES (?, ?, ?)",
			userId1,
			userId2,
			dialogId,
		)
		if err != nil {
			execError = fmt.Errorf("LinkUsersWithDialog: %w", err)
		}
	})

	return execError
}
