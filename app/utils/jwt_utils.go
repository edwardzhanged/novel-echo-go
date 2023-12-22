package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/labstack/echo/v4"
	"time"
)

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	})
	var jwtKey = []byte(conf.GbViper.GetString("jwt.secret"))

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		// Validate the token
		if token != "valid_token" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
