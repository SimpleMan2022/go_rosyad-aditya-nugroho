package routes

import (
	"github.com/labstack/echo/v4"
	"package-project/controller"
)

func PackageRouter(r *echo.Group) {
	r.GET("", controller_packages.GetAllPackages)
	r.GET("/:id", controller_packages.GetPackageById)
	r.POST("", controller_packages.AddNewPackage)
	r.PUT("/:id", controller_packages.UpdatePackage)
	r.DELETE("/:id", controller_packages.DeletePackage)
}
