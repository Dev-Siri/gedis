package routes

import (
	root_handlers "github.com/Dev-Siri/gedis/handlers/root"
	"github.com/buaazp/fasthttprouter"
)

func RegisterRootRoutes(router *fasthttprouter.Router) {
	router.GET("/", root_handlers.GetValueHandler)
	router.POST("/", root_handlers.SetValueHandler)
	router.DELETE("/", root_handlers.DeleteValueHandler)
	router.PATCH("/", root_handlers.IncrementOrDecrementValueHandler)
}
