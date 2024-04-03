package routes_posts_jwt

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	controllers_posts_jwt "post_jwt_project/controllers"
)

func init() {
	viper.AutomaticEnv()
}

func PostRouter(r *echo.Group) {
	r.Use(echojwt.JWT([]byte(viper.GetString("SECRET_TOKEN"))))
	r.GET("", controllers_posts_jwt.GetAllPosts)
	r.GET("/:id", controllers_posts_jwt.GetPostById)
	r.POST("", controllers_posts_jwt.AddNewPost)
	r.PUT("/:id", controllers_posts_jwt.UpdatePost)
	r.DELETE("/:id", controllers_posts_jwt.DeletePost)
}
