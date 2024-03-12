package main

import (
	"go-math-api/handler"
	"go-math-api/repository"
	"go-math-api/routes"
	"go-math-api/service"

	"github.com/labstack/echo/v4"
)

type R struct {
	A int
	B int
}

func main() {
	e := echo.New()

	repository := repository.NewCalculatorRepository()
	service := service.NewCalculatorService(repository)
	handler := handler.NewCalculatorService(service)

	routes.CalcutorRoutes(e, handler)

	e.Start(":1323")
}
