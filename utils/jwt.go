package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// SecreteKey TODO: より秘匿性の髙い秘密鍵を生成しenvに切り出す
const SecreteKey = "secret"

// GenerateJWT IDからtokenを作成する
func GenerateJWT(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1DAY
	})

	return claims.SignedString([]byte(SecreteKey))
}

// ParseJWT tokenを確認し、IDを返す
func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecreteKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
