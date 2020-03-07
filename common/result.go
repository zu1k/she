package common

type Result struct {
	Score int    `json:"score"`
	Hit   string `json:"hit"`
	Text  string `json:"text"`
}
