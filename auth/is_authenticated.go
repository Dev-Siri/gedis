package auth

import "github.com/Dev-Siri/gedis/models"

func IsAuthenticated() (bool, error) {
	var authJson models.Auth

	if err := readAuthFile(&authJson); err != nil {
		return false, err
	}

	isAuthenticated := authJson.SessionID == SessionId

	return isAuthenticated, nil
}