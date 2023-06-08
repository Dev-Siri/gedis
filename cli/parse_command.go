package cli

import (
	"fmt"
	"strings"
)

func ParseCommand(command string) {
	cmdChunks := strings.Split(command, " ")	
	keyword := strings.ToLower(cmdChunks[0])

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
	default:
		fmt.Println("Invalid command.")
	}
}