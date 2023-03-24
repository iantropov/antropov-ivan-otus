package lib

import "users-app/types"

func AreValidUserParams(params types.UserParams) bool {
	return params.Username != nil && params.FirstName != nil && params.LastName != nil && params.Phone != nil && params.Email != nil
}