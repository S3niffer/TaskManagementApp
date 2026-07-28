package models

type Task struct {
	ID          int    `json:"id"`
	User_ID     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Due_date    string `json:"due_date"`
	Created_at  string `json:"created_at"`
}
