package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var defaultKey = "secret"

func CreateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "TPAWebBack",
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		})

	key := GetDotENVVariable("JWT_KEY", defaultKey)

	jwtString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return jwtString, nil
}

func ParseJWT(jwtToken string) (string, error) {

	key := GetDotENVVariable("JWT_KEY", defaultKey)

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	} else {
		return "", fmt.Errorf("Invalid Token")
	}
}
