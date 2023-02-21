package storage

import (
	"database/sql"
	"fmt"
	"social-network-5/types"
)

func GetMessages(dialogId string) ([]types.MessageRecord, error) {
	var queryError error
	var messages []types.MessageRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query(
			"SELECT id, from_user_id, to_user_id, text FROM messages WHERE messages.dialog_id = ? LIMIT 100",
			dialogId,
		)
		if err != nil {
			queryError = fmt.Errorf("GetMessages: %w", err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var message types.MessageRecord
			if err := rows.Scan(
				&message.Id,
				&message.FromUserId,
				&message.ToUserId,
				&message.Text,
			); err != nil {
				queryError = fmt.Errorf("GetMessages: %w", err)
				return
			}
			messages = append(messages, message)
		}
	})

	if queryError != nil {
		return nil, queryError
	}

	return messages, nil
}

func CreateMessage(fromUserId, toUserId, dialogId, text string) error {
	var execError error
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO messages (from_user_id, to_user_id, dialog_id, text) VALUES (?, ?, ?, ?)",
			toUserId,
			fromUserId,
			dialogId,
			text,
		)
		if err != nil {
			execError = fmt.Errorf("CreateMessage: %w", err)
		}
	})

	return execError
}
