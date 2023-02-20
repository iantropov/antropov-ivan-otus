package storage

import (
	"database/sql"
	"fmt"
)

func SetFriend(userId, friendId string) error {
	var setFriendError error
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO friends (user_id, friend_id) VALUES (?, ?)",
			userId,
			friendId,
		)
		if err != nil {
			setFriendError = fmt.Errorf("setFriend: %v", err)
		}
	})

	return setFriendError
}

func DeleteFriend(userId, friendId string) error {
	var deleteFriendError error
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"DELETE FROM friends WHERE user_id=? AND friend_id=?",
			userId,
			friendId,
		)
		if err != nil {
			deleteFriendError = fmt.Errorf("deleteFriend: %v", err)
		}
	})

	return deleteFriendError
}
