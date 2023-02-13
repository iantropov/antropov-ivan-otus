package types

import "fmt"

type UserNotFoundError struct {
	UserId int64
}

func (r *UserNotFoundError) Error() string {
	return fmt.Sprintf("UserNotFound %d", r.UserId)
}
