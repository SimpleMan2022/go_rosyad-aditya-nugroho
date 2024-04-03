package middleware_posts_category

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	model_post_category_jwt "post_category_project/model"
	"strings"
	"time"
)

type JWTClaims struct {
	Id      int
	Email   string
	IsAdmin bool
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateToken(user *model_post_category_jwt.User) (string, error) {

	secretToken := []byte(viper.GetString("SECRET_TOKEN"))
	claims := JWTClaims{
		Id:      user.Id,
		Email:   user.Email,
		IsAdmin: user.Email == "admin123@gmail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(20 * time.Hour).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(secretToken)

	if err != nil {
		return "", err
	}
	return signedString, nil
}

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is not provided")
		}

		if !strings.HasPrefix(authHeader, "bearer ") {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token format. Use Bearer token")
		}
		tokenStr := strings.TrimPrefix(authHeader, "bearer ")

		claims, err := parseJWT(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		isAdmin := claims.IsAdmin
		fmt.Println(isAdmin)
		if !isAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "Only admins are allowed")
		}

		return next(c)
	}
}

func parseJWT(tokenStr string) (*JWTClaims, error) {
	secretToken := []byte(viper.GetString("SECRET_TOKEN"))
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretToken, nil
	})
	if err != nil || !token.Valid {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New("Your token is expired")
	}

	claims := token.Claims.(*JWTClaims)
	fmt.Println(claims)
	if claims == nil {
		return nil, errors.New("Your token is expired")
	}

	return claims, nil
}
