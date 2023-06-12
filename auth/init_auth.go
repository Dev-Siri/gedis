package auth

import (
	"fmt"
	"os"

	"github.com/Dev-Siri/gedis/constants"
)

func InitAuth() error {
	if _, err := os.Stat(constants.AuthJsonPath); os.IsNotExist(err) {		
		if _, err := os.Create(constants.AuthJsonPath); err != nil {
			fmt.Println("(Error) Failed to create auth session file.")
			return err
		}

		os.WriteFile(constants.AuthJsonPath, []byte("{}"), 0644)
	}

	return nil
}