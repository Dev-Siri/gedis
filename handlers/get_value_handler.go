package handlers

import (
	"fmt"
	"net/http"

	"github.com/Dev-Siri/gedis/db"
)

func GetValueHandler(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()

	key := searchParams.Get("key")

	if key == "" {
		http.Error(w, "Key not provided", http.StatusBadRequest)
		return
	}

	value := db.GetValue(key)

	if value == "" {
		fmt.Fprintf(w, "%v", nil)
		return
	}

	fmt.Fprintf(w, "%s", value)
}