package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dev-Siri/gedis/db"
)

func SetValueHandler(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()

	key := searchParams.Get("key")
	value := searchParams.Get("value")
	ttlSearchParam := searchParams.Get("ttl")

	if key == "" {
		http.Error(w, "Key not provided", http.StatusBadRequest)
		return
	}

	if value == "" {
		http.Error(w, "Value not provided", http.StatusBadRequest)
		return
	}

	var ttl int = 0

	if ttlSearchParam == "" {
		timeToLive, err := strconv.Atoi(ttlSearchParam)

		if err != nil {
			http.Error(w, "Failed to convert time-to-live to int", http.StatusBadRequest)
			return
		}

		ttl = timeToLive
	}
	
	db.SetValue(key, value, ttl)

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("%s", value)
}