package routes

import (
	root_handlers "github.com/Dev-Siri/gedis/handlers/root"
	"github.com/buaazp/fasthttprouter"
)

func RegisterRootRoutes(router *fasthttprouter.Router) {
	go router.GET("/", root_handlers.GetValueHandler)
	go router.POST("/", root_handlers.SetValueHandler)
	go router.DELETE("/", root_handlers.DeleteValueHandler)
	go router.PATCH("/", root_handlers.IncrementOrDecrementValueHandler)
}
