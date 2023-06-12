package admin_handlers

import (
	"html/template"
	"net/http"

	"github.com/Dev-Siri/gedis/auth"
	"github.com/Dev-Siri/gedis/embeds"
)

func AdminSession(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFS(embeds.Pages, "pages/session.tmpl")

	if err != nil {
		http.Error(w, "Failed to serve HTML", http.StatusInternalServerError)
		return
	}

	session := auth.GetSession()

	template.Execute(w, session)
}