package cli

import (
	"fmt"

	"github.com/Dev-Siri/gedis/db"
)

func createBackup(cmdChunks []string) {
	backupType := "json"

	if len(cmdChunks) > 1 {
		backupFlag := cmdChunks[1]

		if backupFlag == "--csv" {
			backupType = "csv"
		}
	}

	go db.MakeBackup(backupType)

	fmt.Println("(Info) Backup created successfully")
}