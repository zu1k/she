package common

type Result struct {
	Source string `json:"source"`
	Score  int    `json:"score"`
	Hit    string `json:"hit"`
	Text   string `json:"text"`
}
