package service

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type LocalClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var signKey = []byte("ichinisanshigorokunanahachikyu")

func BuildToken(id uint, username string) (string, error) {
	claims := &LocalClaims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2400).Unix(),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signKey)
}

type UserMeta struct {
	Id       uint
	Username string
}

func ValidToken(tokenStr string) (*UserMeta, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}
	cl := token.Claims.(jwt.MapClaims)
	return &UserMeta{
		Id:       uint(cl["id"].(float64)),
		Username: cl["username"].(string),
	}, nil
}
