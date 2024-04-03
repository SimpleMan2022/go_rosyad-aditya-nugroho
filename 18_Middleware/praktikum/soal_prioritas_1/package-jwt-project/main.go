package main

import (
	"github.com/labstack/echo/v4"
	"package-jwt-project/config"
	middleware_packages "package-jwt-project/middleware"
	routes_packages_jwt "package-jwt-project/routes"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	r := echo.New()
	middleware_packages.LogMiddleware(r)

	packages := r.Group("/api/v1/packages")
	routes_packages_jwt.PackageRouter(packages)

	users := r.Group("/api/v1")
	routes_packages_jwt.UserRouter(users)

	r.Logger.Fatal(r.Start(":5000"))
}
