package middleware_posts

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	model_post_jwt "post_jwt_project/model"
	"time"
)

type JWTClaims struct {
	Id    int
	Email string
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateToken(user *model_post_jwt.User) (string, error) {
	secretToken := []byte(viper.GetString("SECRET_TOKEN"))
	claims := JWTClaims{
		Id:    user.Id,
		Email: user.Email,
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
