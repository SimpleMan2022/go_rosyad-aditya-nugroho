package main

import (
	"github.com/labstack/echo/v4"
	"post_jwt_project/config"
	middleware_posts "post_jwt_project/middleware"
	routes_posts_jwt "post_jwt_project/routes"
)

func main() {
	config.LoadConfig()
	config.LoadDb()

	r := echo.New()
	middleware_posts.LogMiddleware(r)
	posts := r.Group("/api/v1/posts")
	routes_posts_jwt.PostRouter(posts)

	users := r.Group("/api/v1")
	routes_posts_jwt.UserRouter(users)
	r.Logger.Fatal(r.Start(":5000"))
}
