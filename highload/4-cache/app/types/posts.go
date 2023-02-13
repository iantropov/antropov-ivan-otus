package types

type PostRecord struct {
	Text     string `json:"text"`
	AuthorId string `json:"user_id"`
	Id       string `json:"id"`
}

type PostCreateResponse struct {
	Id string `json:"id"`
}

type PostParams struct {
	Text string `json:"text"`
}
