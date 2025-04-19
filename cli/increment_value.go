package cli

import (
	"fmt"
	"strconv"

	"github.com/Dev-Siri/gedis/db"
)

func incrementValue(cmdChunks []string) {
	cmdChunksLen := len(cmdChunks)

	var amount string = "1"

	if cmdChunksLen < 2 {
		fmt.Println("(Error) Key not provided for INCREMENT.")
		return
	}
	
	if cmdChunksLen > 2 {
		amount = cmdChunks[2]
	}

	key := cmdChunks[1]

	amountAsInt, err := strconv.Atoi(amount)

	if err != nil {
		fmt.Println("(Error) Amount must be an integer value")
		return
	}

	db.Increment(key, amountAsInt)
}