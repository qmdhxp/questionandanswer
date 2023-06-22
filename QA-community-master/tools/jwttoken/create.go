package jwttoken

import (
	"QA_community/global"
	"QA_community/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(username string) (string, error) {
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    "pulsar",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.SecretSignKey)
	if err != nil {
		fmt.Println("token生成失败")
		return "", err
	}
	return tokenString, nil
}
