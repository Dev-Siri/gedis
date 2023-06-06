package cli

import (
	"fmt"

	"github.com/Dev-Siri/gedis/db"
)

func loadBackup(cmdChunks []string) {
	if len(cmdChunks) < 2 {
		fmt.Println("Please specify backup file name")
		return
	}

	backupFile := cmdChunks[1]

	db.LoadBackup(backupFile)
}