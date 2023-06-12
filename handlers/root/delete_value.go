package root_handlers

import (
	"fmt"
	"net/http"

	"github.com/Dev-Siri/gedis/db"
)

func DeleteValueHandler(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()

	key := searchParams.Get("key")

	if key == "*" {
		db.DeleteAll()
	} else {
		db.DeleteValue(key)
	}

	fmt.Fprintf(w, "Deleted successfully")
}