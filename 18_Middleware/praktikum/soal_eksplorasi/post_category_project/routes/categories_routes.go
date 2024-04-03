package routes_posts_category_jwt

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	controllers_posts_category_jwt "post_category_project/controllers"
	middleware_posts_category "post_category_project/middleware"
)

func CategooryRoutes(r *echo.Group) {
	r.Use(echojwt.JWT([]byte(viper.GetString("SECRET_TOKEN"))))
	r.Use(middleware_posts_category.AdminOnlyMiddleware)
	r.GET("", controllers_posts_category_jwt.GetAllCategories)
	r.POST("", controllers_posts_category_jwt.AddNewCategory)
	r.DELETE("/:id", controllers_posts_category_jwt.DeleteCategory)
}
