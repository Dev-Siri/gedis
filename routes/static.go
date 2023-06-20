package routes

import (
	"net/http"
	"path/filepath"

	"github.com/Dev-Siri/gedis/embeds"
	"github.com/valyala/fasthttp"
)

func detectContentTypeByExtension(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".css":
		return "text/css"
	default:
		return ""
	}
}

func StaticRoutes(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("filepath").(string)
	data, err := embeds.StaticFiles.ReadFile("public" + path)

	if err != nil {
		ctx.Error("Not found", fasthttp.StatusNotFound)
		return
	}

	contentType := detectContentTypeByExtension(path)

	if contentType == "" {
		contentType = http.DetectContentType(data)
	}

	ctx.SetContentType(contentType)
	ctx.Write(data)
}
