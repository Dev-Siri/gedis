package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Dev-Siri/gedis/constants"
	"github.com/Dev-Siri/gedis/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) error {
	var authJson models.Auth

	if err := readAuthFile(&authJson); err != nil {
		return fmt.Errorf("failed to read auth file")
	}

	if authJson.Username == username {
		if err := bcrypt.CompareHashAndPassword([]byte(authJson.Password), []byte(password)); err != nil {
			return fmt.Errorf("password is incorrect")
		}

		newSessionId := uuid.New().String()

		authJson.SessionID = newSessionId

		SetSession(newSessionId)

		credBytes, jsonError := json.Marshal(authJson)

		if jsonError != nil {
			return fmt.Errorf("failed to convert credentials to JSON")
		}

		if err := os.WriteFile(constants.AuthJsonPath, credBytes, 0644); err != nil {
			return fmt.Errorf("failed to write auth file")
		}
	} else if authJson.Username != "" && authJson.Username != username {
		return fmt.Errorf("you cannot create two accounts on one instance of gedis")
	} else {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

		if err != nil {
			return fmt.Errorf("failed to hash password")
		}

		sessionId := uuid.New().String()

		credentials := models.Auth{
			SessionID: sessionId,
			Username: username,
			Password: string(encryptedPassword),
		}

		SetSession(sessionId)

		credBytes, jsonError := json.Marshal(credentials)

		if jsonError != nil {
			return fmt.Errorf("failed to convert credentials to JSON")
		}

		if err := os.WriteFile(constants.AuthJsonPath, credBytes, 0644); err != nil {
			return fmt.Errorf("failed to write auth file")
		}
	}

	fmt.Printf(
		"\n(Info) Logged in successfully\n=> Run `AUTH session` to get info about the current logged-in session\ngedis> ",
	)
	return nil
}