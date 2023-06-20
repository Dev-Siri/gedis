package admin_login_handlers

import (
	"strings"

	"github.com/Dev-Siri/gedis/auth"
	"github.com/Dev-Siri/gedis/embeds"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

func FormDataLogin(ctx *fasthttp.RequestCtx) {
	formData := ctx.Request.PostArgs()

	username := string(formData.Peek("username"))
	password := string(formData.Peek("password"))

	if username == "" || password == "" {
		template, tmplParseError := embeds.Pages.ReadFile("pages/error.tmpl")

		if tmplParseError != nil {
			ctx.Error("Failed to serve HTML", fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetContentType("text/html")
		fasttemplate.Execute(string(template), "{{", "}}", ctx.Response.BodyWriter(), map[string]interface{}{
			"error": "Username or Password was not provided",
		})
		return
	}

	if err := auth.Login(username, password); err != nil {
		template, tmplParseError := embeds.Pages.ReadFile("pages/error.tmpl")

		if tmplParseError != nil {
			ctx.Error("Failed to serve HTML", fasthttp.StatusInternalServerError)
			return
		}

		errorMessage := err.Error()
		formatedErrorMessage := strings.ToUpper(string(errorMessage[0])) + errorMessage[1:]

		ctx.SetContentType("text/html")
		fasttemplate.Execute(string(template), "{{", "}}", ctx.Response.BodyWriter(), map[string]interface{}{
			"error": formatedErrorMessage,
		})
		return
	}

	ctx.Redirect("/admin", fasthttp.StatusSeeOther)
}
