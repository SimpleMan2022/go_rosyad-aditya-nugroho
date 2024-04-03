package main

import (
	"github.com/labstack/echo/v4"
	"post_category_project/config"
	middleware_posts_category "post_category_project/middleware"
	routes_posts_category_jwt "post_category_project/routes"
)

func main() {
	config.LoadConfig()
	config.LoadDb()

	r := echo.New()
	middleware_posts_category.LogMiddleware(r)
	posts := r.Group("/api/v1/posts")
	routes_posts_category_jwt.PostRouter(posts)

	users := r.Group("/api/v1")
	routes_posts_category_jwt.UserRouter(users)

	categories := r.Group("/api/v1/categories")
	routes_posts_category_jwt.CategooryRoutes(categories)
	r.Logger.Fatal(r.Start(":5000"))
}
