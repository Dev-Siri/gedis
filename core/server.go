package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dev-Siri/gedis/embeds"
	"github.com/Dev-Siri/gedis/routes"
)

func StartServer(port string) {
	staticFS := http.FS(embeds.StaticFiles)
	fs := http.FileServer(staticFS)

	http.Handle("/public/", fs)

	go http.HandleFunc("/", routes.Root)
	go http.HandleFunc("/admin", routes.Admin)
	go http.HandleFunc("/admin/login", routes.AdminLogin)

	addr := fmt.Sprintf(":%s", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
