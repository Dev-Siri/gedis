package admin_login_handlers

import (
	"github.com/Dev-Siri/gedis/embeds"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

func LoginPage(ctx *fasthttp.RequestCtx) {
	template, err := embeds.Pages.ReadFile("pages/login.tmpl")

	if err != nil {
		ctx.Error("Failed to serve HTML", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("text/html")
	fasttemplate.Execute(string(template), "{{", "}}", ctx.Response.BodyWriter(), nil)
}
