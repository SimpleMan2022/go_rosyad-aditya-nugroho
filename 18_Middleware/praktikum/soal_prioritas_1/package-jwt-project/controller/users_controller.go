package controller_packages_jwt

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"package-jwt-project/config"
	errorHandler_packages_jwt "package-jwt-project/errorHandler"
	middleware_packages "package-jwt-project/middleware"
	model_package_jwt "package-jwt-project/model"
)

func Register(ctx echo.Context) error {
	var request model_package_jwt.UserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	password, err := hashPassword(request.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newUser := model_package_jwt.User{
		Email:    request.Email,
		Password: password,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response := model_package_jwt.UserResponse{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Register user successfully",
		Data:       newUser,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func Login(ctx echo.Context) error {
	var request model_package_jwt.UserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	var user model_package_jwt.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Wrong email or password",
		})
	}

	if err := verifyPassword(request.Password, user.Password); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Wrong email or password",
		})
	}

	token, err := middleware_packages.GenerateToken(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler_packages_jwt.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	response := model_package_jwt.LoginResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Token:      token,
	}

	return ctx.JSON(http.StatusOK, response)

}

func hashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func verifyPassword(reqPass, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(reqPass))
	return err
}
