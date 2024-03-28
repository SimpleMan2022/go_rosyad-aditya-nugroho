package main

import (
	"soal_prioritas_1/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	food := e.Group("/api/v1/foods")

	food.GET("", controllers.GetAllFoods)
	food.POST("", controllers.AddNewFood)
	food.GET("/:id", controllers.GetDetailFoods)
	food.PUT("/:id", controllers.UpdateFood)
	food.DELETE("/:id", controllers.DeleteFood)

	e.Logger.Fatal(e.Start(":5000"))
}
