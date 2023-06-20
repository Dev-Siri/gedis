package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dev-Siri/gedis/db"
)

func setValue(cmdChunks []string) {
	cmdChunksLen := len(cmdChunks)

	if cmdChunksLen < 2 {
		fmt.Println("(Error) Key and Value not provided for SET.")
		return
	} else if cmdChunksLen < 3 {
		fmt.Println("(Error) Value not provided for SET.")
		return
	}

	key := cmdChunks[1]
	value := cmdChunks[2]

	var ttl int

	if cmdChunksLen > 3 {
		cmdTTL := strings.TrimPrefix(cmdChunks[3], "--ttl=")
		timeToLive, err := strconv.Atoi(cmdTTL)

		if err != nil {
			fmt.Println("(Error) Failed to convert time to live as int")
			return
		}

		ttl = timeToLive
	}

	db.SetValue(key, value, ttl)
}
