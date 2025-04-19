package auth

import (
	"encoding/json"
	"os"

	"github.com/Dev-Siri/gedis/constants"
	"github.com/Dev-Siri/gedis/models"
)

func readAuthFile(result *models.Auth) error {
	authJsonBytes, err := os.ReadFile(constants.AuthJsonPath)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(authJsonBytes, result); err != nil {
		return err
	}

	return nil
}