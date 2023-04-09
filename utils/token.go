package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type myCustomClaims struct {
	jwt.RegisteredClaims
	TokenData interface{} `json:"data"`
}

// Generate token
func GenerateToken(customData interface{}, issuer string, validDuration time.Duration, signingKey string) (string, error) {

	// Create the claims
	claims := myCustomClaims{
		jwt.RegisteredClaims{
			// 表示Token的过期时间，应该设置为当前时间加上一定时间间隔
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(validDuration)),
			// 表示Token的签发时间，也可以使用jwt.NewNumericDate()函数设置为当前时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 表示Token不能在此时间之前使用，也可以使用jwt.NewNumericDate()函数设置为当前时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 表示Token的签发者
			Issuer: issuer,
		},
		customData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	return ss, err
}

// Parse token
func ParseToken(tokenString string, signingKey string) (customeData interface{}, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*myCustomClaims); ok && token.Valid {
		return claims.TokenData, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
