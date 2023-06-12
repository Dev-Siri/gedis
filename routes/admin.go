package routes

import (
	"net/http"

	admin_handlers "github.com/Dev-Siri/gedis/handlers/admin"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		admin_handlers.AdminSession(w, r)
	} else {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}