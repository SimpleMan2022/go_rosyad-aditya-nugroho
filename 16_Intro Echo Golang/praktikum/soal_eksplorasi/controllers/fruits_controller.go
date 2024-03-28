package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"soal_eksplorasi/entities"
	"soal_eksplorasi/errorHandler"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const apiUrl string = "https://www.fruityvice.com/api/fruit"

var fruitItems []entities.Data

func GetAllFruits(c echo.Context) error {
	if len(fruitItems) == 0 {
		return c.JSON(http.StatusNotFound, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusNotFound,
			Message:    "Fruit is empty",
		})
	}

	response := entities.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all fruits successfully",
		Data:       fruitItems,
	}

	return c.JSON(http.StatusOK, response)
}

func GetFruitById(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	for _, fruit := range fruitItems {
		if fruit.Id == uuid {
			response := entities.FruitResponse{
				Status:     true,
				StatusCode: http.StatusOK,
				Message:    "Get fruit by id successfully",
				Data:       fruit,
			}
			return c.JSON(http.StatusOK, response)
		}
	}

	return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
		Status:     false,
		StatusCode: http.StatusBadRequest,
		Message:    "Fruit not found",
	})
}

func AddNewFruit(c echo.Context) error {
	var request entities.FruitRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	resp, err := http.Get(apiUrl + "/" + request.Name)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var fruits entities.Fruit
	if err := json.NewDecoder(resp.Body).Decode(&fruits); err != nil {
		log.Fatal(err)
	}

	data := entities.Data{
		Id:         uuid.New(),
		Name:       request.Name,
		Price:      request.Price,
		Nutritions: fruits.Nutritions,
	}
	fruitItems = append(fruitItems, data)

	response := entities.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Create new fruit successfully",
		Data:       data,
	}

	return c.JSON(200, response)
}

func DeleteFruit(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	deleted := false
	for i, fruit := range fruitItems {
		if fruit.Id == uuid {
			fruitItems = append(fruitItems[:i], fruitItems[i+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		return c.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Fruit not found",
		})
	}

	response := entities.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete fruit successfully",
	}
	return c.JSON(http.StatusOK, response)
}
