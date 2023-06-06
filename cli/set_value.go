package cli

import (
	"fmt"

	"github.com/Dev-Siri/gedis/db"
)

func setValue(cmdChunks []string) {
	cmdChunksLen := len(cmdChunks)

	if cmdChunksLen < 2 {
		fmt.Println("Key and Value not provided for SET.")
		return
	} else if cmdChunksLen < 3 {
		fmt.Println("Value not provided for SET.")
		return
	}

	key := cmdChunks[1]
	value := cmdChunks[2]
	db.SetValue(key, value)
}