package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dev-Siri/gedis/handlers"
)

func StartServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetValueHandler(w, r)
		case http.MethodPost:
			handlers.SetValueHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteValueHandler(w, r)
		default:
			http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		}
	})

	addr := fmt.Sprintf(":%s", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
