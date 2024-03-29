package main

import (
	"fruits_project/config"
	controllers_fruits "fruits_project/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.LoadDb()

	r := echo.New()
	fruits := r.Group("/api/v1/fruits")
	fruits.GET("", controllers_fruits.GetAllFruits)
	fruits.GET("/:id", controllers_fruits.GetFruitById)
	fruits.POST("", controllers_fruits.AddNewFruit)
	fruits.DELETE("/:id", controllers_fruits.DeleteFruit)

	r.Logger.Fatal(r.Start(":5000"))
}
