package model

type User struct {
	UserID    int    `json:"user_id"`
	RequestID string `json:"request_id"`
	Amount    int    `json:"amount"`
}
