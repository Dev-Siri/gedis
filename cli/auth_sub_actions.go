package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Dev-Siri/gedis/auth"
)

func authSubActions(cmdChunks []string) {
	if len(cmdChunks) < 2 {
		fmt.Println("(Error) The `AUTH` action requires a sub action.")
		return
	}

	subAction := strings.ToLower(cmdChunks[1])

	switch subAction {
	case "login":
		port := os.Getenv("PORT")

		if port == "" {
			port = "5000"
		}

		fmt.Printf("Open http://localhost:%s/admin/login to Login & start a session.\n", port)
	case "session":
		session := auth.GetSession()
		fmt.Printf(
			"Session: {\n  username: %s,\n  session_id: %s\n}\n",
			session.Username,
			session.SessionId,
		)
	default:
		fmt.Println("(Error) Invalid sub action for `AUTH`")
	}
}
