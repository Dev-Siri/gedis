package cli

import (
	"fmt"

	"github.com/Dev-Siri/gedis/db"
)

func deleteValue(cmdChunks []string) {
	if len(cmdChunks) < 2 {
		fmt.Println("Key not provided for DELETE.")
		return
	}

	key := cmdChunks[1]

	if key == "*" {
		db.DeleteAll()
	} else {
		db.DeleteValue(key)
	}
}