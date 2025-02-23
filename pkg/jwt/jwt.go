package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateAccessToken(userID int, accessSecret string) (string, time.Time, error) {
	accessExpireTime := time.Now().Add(7 * 24 * time.Hour)

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["exp"] = accessExpireTime.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessToken, err := at.SignedString([]byte(accessSecret))
	if err != nil {
		return "", accessExpireTime, fmt.Errorf("could not create access token: %v", err)
	}

	return accessToken, accessExpireTime, nil
}

func ValidateAccessToken(tokenString, accessSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(accessSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("unable to parse claims")
}
