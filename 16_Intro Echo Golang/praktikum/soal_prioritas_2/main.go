package main

import (
	"soal_priortias_2/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/words", controllers.GetAllWords)
	e.POST("/words", controllers.AddNewWord)

	e.Logger.Fatal(e.Start(":5000"))
}
