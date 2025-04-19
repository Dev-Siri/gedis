package root_handlers

import (
	"github.com/Dev-Siri/gedis/db"
	"github.com/valyala/fasthttp"
)

func SetValueHandler(ctx *fasthttp.RequestCtx) {
	searchParams := ctx.QueryArgs()

	key := string(searchParams.Peek("key"))
	value := string(searchParams.Peek("value"))
	ttl := searchParams.GetUintOrZero("ttl")

	if key == "" {
		ctx.Error("Key not provided", fasthttp.StatusBadRequest)
		return
	}

	if value == "" {
		ctx.Error("Value not provided", fasthttp.StatusBadRequest)
		return
	}

	db.SetValue(key, value, ttl)

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.Write([]byte(value))
}
