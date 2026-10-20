package models

type Question struct {
	ID         int      `json:"id"`
	Text       string   `json:"text"`
	Answer     string   `json:"answer"`
	Tags       []string `json:"tags"`
	Difficulty string   `json:"difficulty"`
}
