package params

type UserParams struct {
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
	Age        *int    `json:"age"`
	Biography  string  `json:"biography"`
	City       string  `json:"city"`
	Password   *string `json:"password"`
}
