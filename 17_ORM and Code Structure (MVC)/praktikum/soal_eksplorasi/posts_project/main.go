package main

import (
	"github.com/labstack/echo/v4"
	"posts_project/config"
	routes_posts "posts_project/routes"
)

func main() {
	config.LoadConfig()
	config.LoadDb()

	r := echo.New()
	posts := r.Group("/api/v1/posts")
	routes_posts.PostRouter(posts)

	comments := r.Group("/api/v1/comments")
	routes_posts.CommentsRouter(comments)
	r.Logger.Fatal(r.Start(":5000"))
}
