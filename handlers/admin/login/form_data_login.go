package admin_login_handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/Dev-Siri/gedis/auth"
	"github.com/Dev-Siri/gedis/embeds"
)

func FormDataLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		template, tmplParseError := template.ParseFS(embeds.Pages, "pages/error.tmpl")

		if tmplParseError != nil {
			http.Error(w, "Failed to serve HTML", http.StatusInternalServerError)
			return
		}

		template.Execute(w, err.Error())
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		template, tmplParseError := template.ParseFS(embeds.Pages, "pages/error.tmpl")

		if tmplParseError != nil {
			http.Error(w, "Failed to serve HTML", http.StatusInternalServerError)
			return
		}

		template.Execute(w, "Username or Password was not provided")
		return
	}

	if err := auth.Login(username, password); err != nil {
		template, tmplParseError := template.ParseFS(embeds.Pages, "pages/error.tmpl")

		if tmplParseError != nil {
			http.Error(w, "Failed to serve HTML", http.StatusInternalServerError)
			return
		}

		errorMessage := err.Error()
		formatedErrorMessage := strings.ToUpper(string(errorMessage[0])) + errorMessage[1:]

		template.Execute(w, formatedErrorMessage)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}