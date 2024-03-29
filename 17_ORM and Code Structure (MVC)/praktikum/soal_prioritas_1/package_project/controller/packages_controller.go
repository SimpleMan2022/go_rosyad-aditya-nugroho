package controller_packages

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"package-project/config"
	"package-project/errorHandler"
	model_package "package-project/model"
)

func GetAllPackages(ctx echo.Context) error {
	var packages []model_package.Package

	if err := config.DB.Find(&packages).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	if len(packages) == 0 {
		return ctx.JSON(http.StatusOK, model_package.PackageResponse{
			Status:     true,
			StatusCode: http.StatusOK,
			Message:    "Packages is empty",
			Data:       nil,
		})
	}

	response := model_package.PackageResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all packages successfully",
		Data:       packages,
	}

	return ctx.JSON(http.StatusOK, response)
}

func GetPackageById(ctx echo.Context) error {
	id := ctx.Param("id")

	var packages model_package.Package
	if err := config.DB.Where("id = ?", id).First(&packages).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Packages not found",
		})
	}
	response := model_package.PackageResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get packages by id successfully",
		Data:       packages,
	}

	return ctx.JSON(http.StatusOK, response)
}

func AddNewPackage(ctx echo.Context) error {
	var request model_package.PackageRequest

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	newPackage := model_package.Package{
		Name:           request.Name,
		Sender:         request.Sender,
		Receiver:       request.Receiver,
		SenderLocation: request.SenderLocation,
		SenderReceiver: request.SenderReceiver,
		Fee:            request.Fee,
		Weight:         request.Weight,
	}

	if err := config.DB.Create(&newPackage).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response := model_package.PackageResponse{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Create new packages successfully",
		Data:       newPackage,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func UpdatePackage(ctx echo.Context) error {
	id := ctx.Param("id")

	var packages model_package.Package
	if err := config.DB.Where("id = ?", id).First(&packages).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Packages not found",
		})
	}

	var request model_package.PackageRequest

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	packages.Name = request.Name
	packages.Sender = request.Sender
	packages.Receiver = request.Receiver
	packages.SenderLocation = request.SenderLocation
	packages.SenderReceiver = request.SenderReceiver
	packages.Fee = request.Fee
	packages.Weight = request.Weight

	if err := config.DB.Save(&packages).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response := model_package.PackageResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Update packages successfully",
		Data:       packages,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeletePackage(ctx echo.Context) error {
	id := ctx.Param("id")

	var packages model_package.Package
	if err := config.DB.Where("id = ?", id).First(&packages).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Packages not found",
		})
	}

	if err := config.DB.Delete(&packages).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorHandler.ErrorHandler{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	response := model_package.PackageResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Delete packages successfully",
		Data:       nil,
	}

	return ctx.JSON(http.StatusOK, response)
}
