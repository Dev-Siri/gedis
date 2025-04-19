package root_handlers

import (
	"github.com/Dev-Siri/gedis/db"
	"github.com/valyala/fasthttp"
)

func DeleteValueHandler(ctx *fasthttp.RequestCtx) {
	key := string(ctx.QueryArgs().Peek("key"))

	if key == "*" {
		db.DeleteAll()
	} else {
		db.DeleteValue(key)
	}

	ctx.Write([]byte("Deleted successfully"))
}
