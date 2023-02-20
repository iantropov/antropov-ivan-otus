package types

import "fmt"

type NotFoundError struct {
	Id string
}

func (r *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound %s", r.Id)
}
