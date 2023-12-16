package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/edwardzhanged/novel-go/app/conf"
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
