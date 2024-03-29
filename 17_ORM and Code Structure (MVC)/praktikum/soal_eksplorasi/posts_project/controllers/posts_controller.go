package controllers_posts

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"posts_project/config"
	errorHandler_posts "posts_project/errorHandler"
	model_post "posts_project/model"
)

func GetAllPosts(ctx echo.Context) error {
	var posts []model_post.Post
	if err := config.DB.Debug().Preload("Comment").Find(&posts).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.PostResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all posts successfully",
		Data:       posts,
	}

	return ctx.JSON(http.StatusOK, response)
}

func GetPostById(ctx echo.Context) error {
	id := ctx.Param("id")
	var posts model_post.Post
	if err := config.DB.Debug().Where("id = ?", id).Preload("Comment").
		Find(&posts).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.PostResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get post by id successfully",
		Data:       posts,
	}

	return ctx.JSON(http.StatusOK, response)
}

func AddNewPost(ctx echo.Context) error {
	var request model_post.PostRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newPost := model_post.Post{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := config.DB.Create(&newPost).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.PostResponse{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Create new post successfully",
		Data:       newPost,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func UpdatePost(ctx echo.Context) error {
	id := ctx.Param("id")
	var post model_post.Post
	if err := config.DB.Debug().Where("id = ?", id).Preload("Comment").
		Find(&post).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	var request model_post.PostRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	post.Title = request.Title
	post.Content = request.Content

	if err := config.DB.Save(&post).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.PostResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Update post successfully",
		Data:       post,
	}
	return ctx.JSON(http.StatusOK, response)
}

func DeletePost(ctx echo.Context) error {
	id := ctx.Param("id")
	var post model_post.Post
	if err := config.DB.Debug().Where("id = ?", id).Preload("Comment").
		Find(&post).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_posts.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_post.PostResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete post successfully",
		Data:       nil,
	}
	return ctx.JSON(http.StatusOK, response)
}
