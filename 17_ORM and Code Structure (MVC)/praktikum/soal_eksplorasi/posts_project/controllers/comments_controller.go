package controllers_posts

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"posts_project/config"
	errorHandler_posts "posts_project/errorHandler"
	model_post "posts_project/model"
)

func AddNewComment(ctx echo.Context) error {
	id := ctx.Param("id")

	var post model_post.Post
	if err := config.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	var request model_post.Comment
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newComment := model_post.Comment{
		PostId:  post.Id,
		Content: request.Content,
	}

	if err := config.DB.Create(&newComment).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.CommentResponse{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Create new comment successfully",
		Data:       newComment,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func UpdateComment(ctx echo.Context) error {
	id := ctx.Param("id")

	var comment model_post.Comment
	if err := config.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	var request model_post.Comment
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	comment.Content = request.Content

	if err := config.DB.Save(&comment).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.CommentResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Update comment successfully",
		Data:       comment,
	}
	return ctx.JSON(http.StatusOK, response)
}

func DeleteComment(ctx echo.Context) error {
	id := ctx.Param("id")

	var comment model_post.Comment
	if err := config.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.CommentResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete comment successfully",
		Data:       nil,
	}
	return ctx.JSON(http.StatusOK, response)
}
