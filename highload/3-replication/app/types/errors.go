package types

import "fmt"

type UserNotFoundError struct {
	UserId string
}

func (r *UserNotFoundError) Error() string {
	return fmt.Sprintf("UserNotFound %s", r.UserId)
}
