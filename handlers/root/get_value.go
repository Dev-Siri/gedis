package root_handlers

import (
	"github.com/Dev-Siri/gedis/db"
	"github.com/valyala/fasthttp"
)

func GetValueHandler(ctx *fasthttp.RequestCtx) {
	key := string(ctx.QueryArgs().Peek("key"))

	if key == "" {
		ctx.Error("Key not provided", fasthttp.StatusBadRequest)
		return
	}

	value := db.GetValue(key).Value

	if value == "" {
		ctx.Write([]byte("<nil>"))
		return
	}

	ctx.Write([]byte(value))
}
