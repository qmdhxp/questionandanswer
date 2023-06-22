package jwttoken

import (
	"QA_community/global"
	"QA_community/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// 解析Token
func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.SecretSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
