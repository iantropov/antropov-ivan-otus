package cache

import (
	"encoding/json"
	"fmt"
	"social-network-4/types"
	"time"

	"github.com/redis/go-redis/v9"
)

const PostsFeedTTL = time.Minute

func GetPostsFeed(userId string) ([]types.PostRecord, error) {
	key := postsFeedKey(userId)
	postsStr, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("Failed to get key %s: %w", key, err)
	}
	var posts []types.PostRecord
	err = json.Unmarshal([]byte(postsStr), &posts)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts %s: %w", postsStr, err)
	}
	return posts, nil
}

func SetPostsFeed(userId string, posts []types.PostRecord) error {
	key := postsFeedKey(userId)
	marshaledPosts, err := json.Marshal(posts)
	if err != nil {
		return fmt.Errorf("Failed to marshal posts: %w", err)
	}
	err = rdb.Set(ctx, key, marshaledPosts, PostsFeedTTL).Err()
	if err != nil {
		return fmt.Errorf("Failed to set posts %s: %w", key, err)
	}
	return nil
}

func postsFeedKey(userId string) string {
	return fmt.Sprintf("posts_feed:%s", userId)
}
