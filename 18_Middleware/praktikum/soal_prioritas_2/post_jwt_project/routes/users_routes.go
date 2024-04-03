package routes_posts_jwt

import (
	controllers_posts_jwt "post_jwt_project/controllers"
	middleware_posts "post_jwt_project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRouter(r *echo.Group) {
	r.Use(middleware.RateLimiterWithConfig(middleware_posts.CustomeRateLimiter()))
	r.POST("/register", controllers_posts_jwt.Register)
	r.POST("/login", controllers_posts_jwt.Login)
}
