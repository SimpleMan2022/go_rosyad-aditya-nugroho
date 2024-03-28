package controllers

import (
	"net/http"
	"soal_prioritas_1/entities"
	"soal_prioritas_1/errorHandler"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var datas []entities.Food

func AddNewFood(c echo.Context) error {
	var request entities.FoodRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newFood := entities.Food{
		Id:          uuid.New(),
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
	}

	datas = append(datas, newFood)

	response := entities.FoodResponse{
		Status:  true,
		Message: "Create new food successfully",
		Data:    newFood,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetAllFoods(c echo.Context) error {
	if len(datas) == 0 {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusNotFound,
			Message:    "Food is empty",
		})
	}

	response := entities.FoodResponse{
		Status:  true,
		Message: "Get all foods successfully",
		Data:    datas,
	}

	return c.JSON(http.StatusOK, response)
}

func GetDetailFoods(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid food id",
		})
	}
	for _, data := range datas {
		if data.Id == uuid {
			response := entities.FoodResponse{
				Status:  true,
				Message: "Get detail food successfully",
				Data:    data,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
		StatusCode: http.StatusBadRequest,
		Message:    "Food not found",
	})
}

func UpdateFood(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid food id",
		})
	}
	for idx, data := range datas {
		if data.Id == uuid {
			var request entities.FoodRequest
			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
					Status:     false,
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				})
			}

			data.Name = request.Name
			data.Price = request.Price
			data.Description = request.Description

			datas[idx] = data

			response := entities.FoodResponse{
				Status:  true,
				Message: "Update food successfully",
				Data:    data,
			}
			return c.JSON(http.StatusOK, response)

		}
	}
	return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
		StatusCode: http.StatusBadRequest,
		Message:    "Food not found",
	})
}

func DeleteFood(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid food id",
		})
	}
	for idx, data := range datas {
		if data.Id == uuid {
			var request entities.FoodRequest
			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
					Status:     false,
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
				})
			}

			datas = append(datas[:idx], datas[idx+1:]...)

			response := entities.FoodResponse{
				Status:  true,
				Message: "Delete food successfully",
				Data:    nil,
			}

			return c.JSON(http.StatusOK, response)

		}
	}
	return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
		StatusCode: http.StatusBadRequest,
		Message:    "Food not found",
	})
}
