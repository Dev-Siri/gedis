package routes

import (
	admin_handlers "github.com/Dev-Siri/gedis/handlers/admin"
	admin_login_handlers "github.com/Dev-Siri/gedis/handlers/admin/login"
	"github.com/buaazp/fasthttprouter"
)

func RegisterAdminRoutes(router *fasthttprouter.Router) {
	go router.GET("/admin", admin_handlers.AdminSession)
	go router.GET("/admin/login", admin_login_handlers.LoginPage)
	go router.POST("/admin/login", admin_login_handlers.FormDataLogin)
}
