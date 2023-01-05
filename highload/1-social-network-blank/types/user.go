package types

type UserRegisterParams struct {
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
	Age        *int    `json:"age"`
	Biography  string  `json:"biography"`
	City       string  `json:"city"`
	Password   *string `json:"password"`
}

type UserLoginParams struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
}

type UserRecord struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Age            int    `json:"age"`
	Biography      string `json:"biography"`
	City           string `json:"city"`
	HashedPassword string `json:"-"`
	Id             string `json:"id"`
}
