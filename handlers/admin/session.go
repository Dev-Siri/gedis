package admin_handlers

import (
	"net/http"

	"github.com/Dev-Siri/gedis/auth"
	"github.com/Dev-Siri/gedis/embeds"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

func AdminSession(ctx *fasthttp.RequestCtx) {
	template, err := embeds.Pages.ReadFile("pages/session.tmpl")

	if err != nil {
		ctx.Error("Failed to serve HTML", http.StatusInternalServerError)
		return
	}

	session := auth.GetSession()

	ctx.SetContentType("text/html")
	fasttemplate.Execute(string(template), "{{", "}}", ctx.Response.BodyWriter(), map[string]interface{}{
		"username":  session.Username,
		"sessionId": session.SessionId,
	})
}
