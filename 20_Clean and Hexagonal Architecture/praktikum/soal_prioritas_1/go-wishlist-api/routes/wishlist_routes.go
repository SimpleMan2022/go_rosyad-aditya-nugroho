package routes

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api/config"
	"go-wishlist-api/handlers"
	"go-wishlist-api/repositories"
	"go-wishlist-api/usecases"
)

func WishlistRouter(wishlist *echo.Group) {
	repository := repositories.NewWishlistRepository(config.DB)
	usecase := usecases.NewWishlistUsecase(repository)
	handler := handlers.NewWishlistHandler(usecase)
	wishlist.GET("", handler.GetAll)
	wishlist.POST("", handler.Create)
}
