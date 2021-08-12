package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(getSecret())

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

type JwtCustomClaim struct {
	UuID string `json:"uu_id"`
	jwt.StandardClaims
}

func JwtGenerate(uuID string) string {
	stdClaim := &JwtCustomClaim{
		UuID: uuID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "mylocalhost.com",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaim)

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return t
}

func JwtValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %+v", t.Header["alg"])
		}
		return jwtSecret, nil
	})
}
