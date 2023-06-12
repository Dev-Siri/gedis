package admin_login_handlers

import (
	"html/template"
	"net/http"

	"github.com/Dev-Siri/gedis/embeds"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFS(embeds.Pages, "pages/login.tmpl")

	if err != nil {
		http.Error(w, "Failed to serve HTML", http.StatusInternalServerError)
		return
	}

	template.Execute(w, nil)
}