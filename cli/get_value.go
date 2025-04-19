package cli

import (
	"fmt"

	"github.com/Dev-Siri/gedis/db"
)

func getValue(cmdChunks []string) {
	if len(cmdChunks) < 2 {
		fmt.Println("(Error) Key not provided for GET.")
		return
	}

	key := cmdChunks[1]
	value := db.GetValue(key).Value

	// Since map is a string -> string store, it returns ""
	// So it checks for an "" and then returns actual nil
	if value == "" {
		fmt.Println(nil)
		return
	}

	fmt.Println(value)	
}