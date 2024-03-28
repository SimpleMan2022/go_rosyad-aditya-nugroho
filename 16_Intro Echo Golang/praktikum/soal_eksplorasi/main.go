package main

import (
	"soal_eksplorasi/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	fruits := r.Group("/api/v1/fruits")
	fruits.GET("", controllers.GetAllFruits)
	fruits.GET("/:id", controllers.GetFruitById)
	fruits.POST("", controllers.AddNewFruit)
	fruits.DELETE("/:id", controllers.DeleteFruit)

	r.Logger.Fatal(r.Start(":5000"))
}
