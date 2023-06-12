package auth

import (
	"fmt"

	"github.com/Dev-Siri/gedis/models"
)

var SessionId string

func GetSession() models.Session {
	var authJson models.Auth

	if err := readAuthFile(&authJson); err != nil {
		fmt.Println("(Error) Failed to read auth session file.")
		return models.Session{}
	}

	return models.Session{
		SessionId: SessionId,
		Username: authJson.Username,
	}
}

func SetSession(sid string) {
	SessionId = sid
}