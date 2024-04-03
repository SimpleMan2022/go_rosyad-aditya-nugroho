package routes_posts_category_jwt

import (
	"github.com/labstack/echo/v4"
	controllers_posts_category_jwt "post_category_project/controllers"
)

func UserRouter(r *echo.Group) {

	r.POST("/register", controllers_posts_category_jwt.Register)
	r.POST("/login", controllers_posts_category_jwt.Login)
}
