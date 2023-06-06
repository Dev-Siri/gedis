package handlers

import (
	"fmt"
	"net/http"

	"github.com/Dev-Siri/gedis/db"
)

func SetValueHandler(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()

	key := searchParams.Get("key")
	value := searchParams.Get("value")

	if key == "" {
		http.Error(w, "Key not provided", http.StatusBadRequest)
		return
	}

	if value == "" {
		http.Error(w, "Value not provided", http.StatusBadRequest)
		return
	}

	db.SetValue(key, value)

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("%s", value)
}