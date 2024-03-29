package routes_posts

import (
	"github.com/labstack/echo/v4"
	controllers_posts "posts_project/controllers"
)

func PostRouter(r *echo.Group) {
	r.GET("", controllers_posts.GetAllPosts)
	r.GET("/:id", controllers_posts.GetPostById)
	r.POST("", controllers_posts.AddNewPost)
	r.PUT("/:id", controllers_posts.UpdatePost)
	r.DELETE("/:id", controllers_posts.DeletePost)
}
