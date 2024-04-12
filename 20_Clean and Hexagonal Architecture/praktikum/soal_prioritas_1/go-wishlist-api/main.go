package main

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api/config"
	"go-wishlist-api/routes"
)

func main() {
	config.LoadConfig()
	config.InitDatabase()

	e := echo.New()

	wishlists := e.Group("/wishlists")
	routes.WishlistRouter(wishlists)
	e.Logger.Fatal(e.Start(":1323"))
}
