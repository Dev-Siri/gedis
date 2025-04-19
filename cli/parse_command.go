package cli

import (
	"fmt"
	"strings"

	"github.com/Dev-Siri/gedis/auth"
)

func ParseCommand(command string) {
	cmdChunks := strings.Split(command, " ")	
	keyword := strings.ToLower(cmdChunks[0])

	isAuthenticated, err := auth.IsAuthenticated()

	if err != nil {
		fmt.Println("(Error) Failed to check current auth status.")
		return
	}

	if !isAuthenticated {
		if keyword != "auth" {
			fmt.Println(
				"(Unauthorized) You can't use any database actions if you are logged out\n",
				"=> Run `AUTH login` to log into a database session.",
			)
			return
		}

		authSubActions(cmdChunks)
		return
	}

	switch keyword {
	case "get":
		getValue(cmdChunks)	
	case "set":
		setValue(cmdChunks)	
	case "delete":
		deleteValue(cmdChunks)
	case "increment":
		incrementValue(cmdChunks)
	case "decrement":
		decrementValue(cmdChunks)
	case "create_backup":
		createBackup(cmdChunks)
	case "load_backup":
		loadBackup(cmdChunks)
	case "auth":
		authSubActions(cmdChunks)
	default:
		fmt.Println("(Error) Invalid command.")
	}
}