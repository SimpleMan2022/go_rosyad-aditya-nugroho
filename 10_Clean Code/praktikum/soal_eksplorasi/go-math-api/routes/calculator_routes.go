package routes

import (
	"go-math-api/handler"

	"github.com/labstack/echo/v4"
)

func CalcutorRoutes(e *echo.Echo, handler *handler.CalculatorHandler) {
	e.POST("/api/add", handler.Add)
	e.POST("/api/substract", handler.Substract)
	e.POST("/api/multiply", handler.Multiply)
	e.POST("/api/div", handler.Divide)
}
