package root_handlers

import (
	"github.com/Dev-Siri/gedis/db"
	"github.com/valyala/fasthttp"
)

func IncrementOrDecrementValueHandler(ctx *fasthttp.RequestCtx) {
	searchParams := ctx.QueryArgs()

	key := string(searchParams.Peek("key"))
	action := string(searchParams.Peek("action"))
	amount := searchParams.GetUintOrZero("amount")

	if key == "" {
		ctx.Error("Key not provided", fasthttp.StatusBadRequest)
		return
	}

	appliedAmount := 1

	if amount > 0 {
		appliedAmount = amount
	}

	if action == "decrement" {
		db.Decrement(key, appliedAmount)
	} else {
		db.Increment(key, appliedAmount)
	}
}
