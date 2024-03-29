package routes_posts

import (
	"github.com/labstack/echo/v4"
	controllers_posts "posts_project/controllers"
)

func CommentsRouter(r *echo.Group) {

	r.POST("/:id", controllers_posts.AddNewComment)
	r.PUT("/:id", controllers_posts.UpdateComment)
	r.DELETE("/:id", controllers_posts.DeleteComment)
}
