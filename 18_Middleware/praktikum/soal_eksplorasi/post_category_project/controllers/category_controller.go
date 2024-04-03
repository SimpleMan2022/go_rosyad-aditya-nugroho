package controllers_posts_category_jwt

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"post_category_project/config"
	errorHandler_posts_category_jwt "post_category_project/errorHandler"
	model_post_category_jwt "post_category_project/model"
)

func GetAllCategories(ctx echo.Context) error {
	var categories []model_post_category_jwt.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts_category_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post_category_jwt.CategoryResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all categories successfully",
		Data:       categories,
	}

	return ctx.JSON(http.StatusOK, response)
}

func AddNewCategory(ctx echo.Context) error {
	var request model_post_category_jwt.CategoryRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts_category_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newCategory := model_post_category_jwt.Category{
		Name: request.Name,
	}

	if err := config.DB.Create(&newCategory).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts_category_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post_category_jwt.CategoryResponse{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Create new category successfully",
		Data:       newCategory,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func DeleteCategory(ctx echo.Context) error {
	id := ctx.Param("id")
	var category model_post_category_jwt.Category
	if err := config.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts_category_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts_category_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post_category_jwt.PostResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete category successfully",
		Data:       nil,
	}
	return ctx.JSON(http.StatusOK, response)
}
