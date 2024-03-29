package main

import (
	"package-project/config"
	"package-project/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	r := echo.New()
	packages := r.Group("/api/v1/packages")
	routes.PackageRouter(packages)

	r.Logger.Fatal(r.Start(":5000"))
}
