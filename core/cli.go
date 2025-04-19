package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dev-Siri/gedis/cli"
)

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("(Info) Enter .exit to stop the server.")

	for {
		fmt.Print("gedis> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ".exit" {
			break
		}

		if strings.Trim(input, " ") != "" {
			cli.ParseCommand(input)
		}
	}

	fmt.Println("(Info) Shutting down Gedis server")
}
