package models

type Session struct {
	Username  string `json:"username"`
	SessionId string `json:"session_id"`
}