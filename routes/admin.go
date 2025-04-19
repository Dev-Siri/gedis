package routes

import (
	admin_handlers "github.com/Dev-Siri/gedis/handlers/admin"
	admin_login_handlers "github.com/Dev-Siri/gedis/handlers/admin/login"
	"github.com/buaazp/fasthttprouter"
)

func RegisterAdminRoutes(router *fasthttprouter.Router) {
	router.GET("/admin", admin_handlers.AdminSession)
	router.GET("/admin/login", admin_login_handlers.LoginPage)
	router.POST("/admin/login", admin_login_handlers.FormDataLogin)
}
