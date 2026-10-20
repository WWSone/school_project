package models

type Theory struct {
	ID      int    `json: "id"`
	Section string `json: "section"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
