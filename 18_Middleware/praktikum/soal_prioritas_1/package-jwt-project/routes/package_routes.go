package routes_packages_jwt

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	controller_packages_jwt "package-jwt-project/controller"
)

func init() {
	viper.AutomaticEnv()
}

func PackageRouter(r *echo.Group) {
	r.Use(echojwt.JWT([]byte(viper.GetString("SECRET_TOKEN"))))
	r.GET("", controller_packages_jwt.GetAllPackages)
	r.GET("/:id", controller_packages_jwt.GetPackageById)
	r.POST("", controller_packages_jwt.AddNewPackage)
	r.PUT("/:id", controller_packages_jwt.UpdatePackage)
	r.DELETE("/:id", controller_packages_jwt.DeletePackage)
}
