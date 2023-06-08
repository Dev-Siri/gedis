package handlers

import (
	"net/http"
	"strconv"

	"github.com/Dev-Siri/gedis/db"
)

func IncrementOrDecrementValueHandler(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()

	key := searchParams.Get("key")
	action := searchParams.Get("action")
	amountParam := searchParams.Get("amount")

	var amount string = "1"

	if key == "" {
		http.Error(w, "Key not provided", http.StatusBadRequest)
		return
	}
	
	if amountParam != "" {
		amount = amountParam
	}

	amountAsInt, err := strconv.Atoi(amount)

	if err != nil {
		http.Error(w, "Amount must be an integer value", http.StatusBadRequest)
		return
	}

	if action == "decrement" {	
		db.Decrement(key, amountAsInt)
	} else {
		db.Increment(key, amountAsInt)
	}
}