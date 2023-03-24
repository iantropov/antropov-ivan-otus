package storage

import (
	"database/sql"
	"fmt"
	"social-network-5/cache"
	"social-network-5/types"

	"github.com/google/uuid"
)

const FeedLimit = 1000

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

func GetPostsFeed(userId string, limit, offset int) ([]types.PostRecord, error) {
	postsFeedFromCache, err := cache.GetPostsFeed(userId)
	if err != nil {
		return nil, err
	}
	if postsFeedFromCache != nil {
		return getPostsWithLimitAndOffet(postsFeedFromCache, limit, offset), nil
	}
	postsFeedFromDb, err := getPostsFeedFromDb(userId)
	if err != nil {
		return nil, err
	}
	err = cache.SetPostsFeed(userId, postsFeedFromDb)
	if err != nil {
		return nil, err
	}

	return getPostsWithLimitAndOffet(postsFeedFromDb, limit, offset), nil
}

func getPostsFeedFromDb(userId string) ([]types.PostRecord, error) {
	var getPostsFeedError error
	var posts []types.PostRecord
	queryDb(func(db *sql.DB) {
		rows, err := db.Query(
			"SELECT id, author_id, text FROM posts JOIN friends ON author_id = friends.friend_id WHERE friends.user_id = ? ORDER BY id DESC LIMIT ?",
			userId,
			FeedLimit,
		)
		if err != nil {
			getPostsFeedError = fmt.Errorf("getPostsFeedFromDb: userId=%s: %w", userId, err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var post types.PostRecord
			if err := rows.Scan(
				&post.Id,
				&post.AuthorId,
				&post.Text,
			); err != nil {
				getPostsFeedError = fmt.Errorf("getPostsFeedFromDb: userId=%s: %w", userId, err)
				return
			}
			posts = append(posts, post)
		}
	})

	if getPostsFeedError != nil {
		return nil, getPostsFeedError
	}

	return posts, nil
}

func getPostsWithLimitAndOffet(posts []types.PostRecord, limit, offset int) []types.PostRecord {
	sliceStart := offset
	if sliceStart > len(posts) {
		sliceStart = len(posts)
	}
	sliceFinish := limit + offset
	if sliceFinish > len(posts) {
		sliceFinish = len(posts)
	}
	return posts[sliceStart:sliceFinish]
}
