package routes

import (
	"net/http"

	admin_login_handlers "github.com/Dev-Siri/gedis/handlers/admin/login"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		admin_login_handlers.LoginPage(w, r)
	case http.MethodPost:
		admin_login_handlers.FormDataLogin(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}