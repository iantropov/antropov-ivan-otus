package types

type UserParams struct {
	Username   *string `json:"username"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
	Email      *int    `json:"email"`
	Phone      *string `json:"phone"`
}

type UserRecord struct {
	Username   *string `json:"username"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
	Email      *int    `json:"email"`
	Phone      string  `json:"phone"`
	Id         int64   `json:"id"`
}

type UserCreatedResponse struct {
	UserId int64 `json:"id"`
}
