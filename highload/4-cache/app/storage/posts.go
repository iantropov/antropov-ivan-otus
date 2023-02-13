package storage

import (
	"database/sql"
	"fmt"
	"social-network-4/types"

	"github.com/google/uuid"
)

func CreatePost(authorId, text string) (string, error) {
	var createPostError error
	postId := uuid.New().String()
	queryDb(func(db *sql.DB) {
		_, err := db.Exec(
			"INSERT INTO posts (id, author_id, text) VALUES (?, ?, ?)",
			postId,
			authorId,
			text,
		)
		if err != nil {
			createPostError = fmt.Errorf("createPost: %v", err)
		}
	})

	return postId, createPostError
}

func DeletePost(authorId, postId string) error {
	var deletePostError error
	queryDb(func(db *sql.DB) {
		result, err := db.Exec(
			"DELETE FROM posts WHERE author_id=? AND id=?",
			authorId,
			postId,
		)
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			deletePostError = fmt.Errorf("deletePost: %q: %v", postId, err)
		} else if rowsAffected == 0 {
			deletePostError = &types.NotFoundError{
				Id: postId,
			}
		}
	})

	return deletePostError
}

func UpdatePost(authorId, postId, text string) error {
	var updatePostError error
	queryDb(func(db *sql.DB) {
		result, err := db.Exec(
			"UPDATE posts SET text=? WHERE author_id=? AND id=?",
			text,
			authorId,
			postId,
		)
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			updatePostError = fmt.Errorf("updatePost: %q: %v", postId, err)
		} else if rowsAffected == 0 {
			updatePostError = &types.NotFoundError{
				Id: postId,
			}
		}
	})

	return updatePostError
}

func GetPost(postId string) (*types.PostRecord, error) {
	var getPostError error
	var postRecord types.PostRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query("SELECT id, author_id, text FROM posts WHERE id = ?", postId)
		if err != nil {
			getPostError = fmt.Errorf("getPost: %q: %v", postId, err)
			return
		}
		defer rows.Close()

		if rows.Next() {
			if err := rows.Scan(
				&postRecord.Id,
				&postRecord.AuthorId,
				&postRecord.Text,
			); err != nil {
				getPostError = fmt.Errorf("getPost: %q: %v", postId, err)
				return
			}
		} else {
			getPostError = &types.NotFoundError{
				Id: postId,
			}
		}
	})

	if getPostError != nil {
		return nil, getPostError
	}

	return &postRecord, nil
}
