package models

type Auth struct {
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}