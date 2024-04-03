package routes_posts_category_jwt

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	controllers_posts_category_jwt "post_category_project/controllers"
)

func init() {
	viper.AutomaticEnv()
}

func PostRouter(r *echo.Group) {
	r.Use(echojwt.JWT([]byte(viper.GetString("SECRET_TOKEN"))))
	r.GET("", controllers_posts_category_jwt.GetAllPosts)
	r.GET("/:id", controllers_posts_category_jwt.GetPostById)
	r.POST("", controllers_posts_category_jwt.AddNewPost)
	r.PUT("/:id", controllers_posts_category_jwt.UpdatePost)
	r.DELETE("/:id", controllers_posts_category_jwt.DeletePost)
}
