package handler

import (
	"go-math-api/dto"
	"go-math-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculatorHandler struct {
	service service.CalculatorService
}

func NewCalculatorService(s service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{s}
}

func (h *CalculatorHandler) Add(c echo.Context) error {
	var CalculatorRequest dto.CalculatorRequest // num1 & num2

	if err := c.Bind(&CalculatorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	result := h.service.Add(&CalculatorRequest)

	return c.JSON(http.StatusOK, dto.CalculatorResponse{
		StatusCode: http.StatusOK,
		Message:    "Success to add two number",
		Result:     result,
	})
}

func (h *CalculatorHandler) Substract(c echo.Context) error {
	var CalculatorRequest dto.CalculatorRequest // num1 & num2
	if err := c.Bind(&CalculatorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	result := h.service.Substract(&CalculatorRequest)

	return c.JSON(http.StatusOK, dto.CalculatorResponse{
		StatusCode: http.StatusOK,
		Message:    "Success to substract two number",
		Result:     result,
	})
}

func (h *CalculatorHandler) Multiply(c echo.Context) error {
	var CalculatorRequest dto.CalculatorRequest // num1 & num2
	if err := c.Bind(&CalculatorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	result := h.service.Multiply(&CalculatorRequest)

	return c.JSON(http.StatusOK, dto.CalculatorResponse{
		StatusCode: http.StatusOK,
		Message:    "Success to multiply two number",
		Result:     result,
	})
}

func (h *CalculatorHandler) Divide(c echo.Context) error {
	var CalculatorRequest dto.CalculatorRequest // num1 & num2
	if err := c.Bind(&CalculatorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	result, err := h.service.Divide(&CalculatorRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CalculatorResponse{
		StatusCode: http.StatusOK,
		Message:    "Success to substract two number",
		Result:     result,
	})
}
