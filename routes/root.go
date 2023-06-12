package routes

import (
	"net/http"

	root_handlers "github.com/Dev-Siri/gedis/handlers/root"
)

func Root(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		root_handlers.GetValueHandler(w, r)
	case http.MethodPost:
		root_handlers.SetValueHandler(w, r)
	case http.MethodDelete:
		root_handlers.DeleteValueHandler(w, r)
	case http.MethodPatch:
		root_handlers.IncrementOrDecrementValueHandler(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}