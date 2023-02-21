package types

type MessageRecord struct {
	Text       string `json:"text"`
	FromUserId string `json:"from"`
	ToUserId   string `json:"to"`
	Id         string `json:"-"`
}
