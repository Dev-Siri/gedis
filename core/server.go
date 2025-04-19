package core

import (
	"log"

	"github.com/Dev-Siri/gedis/routes"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func StartServer(port string) {
	router := fasthttprouter.New()

	router.GET("/public/*filepath", routes.StaticRoutes)

	routes.RegisterRootRoutes(router)
	routes.RegisterAdminRoutes(router)

	addr := ":" + port

	if err := fasthttp.ListenAndServe(addr, router.Handler); err != nil {
		log.Fatal(err)
	}
}
