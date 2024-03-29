package controllers_fruits

import (
	"encoding/json"
	"fruits_project/config"
	errorHandler_fruits "fruits_project/errorHandler"
	model_fruits "fruits_project/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

const apiUrl string = "https://www.fruityvice.com/api/fruit"

var fruitItems []model_fruits.Fruit

func GetAllFruits(c echo.Context) error {
	var fruits []model_fruits.Fruit
	if err := config.DB.Debug().Preload("Nutrition").Find(&fruits).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_fruits.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all fruits successfully",
		Data:       fruits,
	}

	return c.JSON(http.StatusOK, response)
}

func GetFruitById(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	var fruit model_fruits.Fruit

	if err := config.DB.Debug().Where("id = ?", uuid).Preload("Nutrition").
		First(&fruit).Error; err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, model_fruits.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get fruits by id successfully",
		Data:       fruit,
	})
}

func AddNewFruit(c echo.Context) error {
	var request model_fruits.FruitRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler_fruits.ErrorHandler{
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

	var fruits model_fruits.FruitApi
	if err := json.NewDecoder(resp.Body).Decode(&fruits); err != nil {
		log.Fatal(err)
	}

	data := model_fruits.Fruit{
		Id:          uuid.New(),
		Name:        request.Name,
		Price:       request.Price,
		NutritionId: fruits.Id,
		Nutrition:   fruits.Nutritions,
	}

	data.Nutrition.Id = fruits.Id
	if err := config.DB.Create(&data.Nutrition).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_fruits.FruitResponse{
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
		return c.JSON(http.StatusBadRequest, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	var fruit model_fruits.Fruit
	if err := config.DB.Debug().Where("id = ?", uuid).
		First(&fruit).Error; err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if err := config.DB.Delete(&fruit).Error; err != nil {
		return c.JSON(http.StatusBadRequest, errorHandler_fruits.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response := model_fruits.FruitResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete fruit successfully",
	}
	return c.JSON(http.StatusOK, response)
}
