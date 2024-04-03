package routes_packages_jwt

import (
	"github.com/labstack/echo/v4"
	controller_packages_jwt "package-jwt-project/controller"
)

func UserRouter(r *echo.Group) {
	r.POST("/register", controller_packages_jwt.Register)
	r.POST("/login", controller_packages_jwt.Login)
}
