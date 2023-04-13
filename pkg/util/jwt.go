package util

import (
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt"
	"time"
)

var privatekey *ecdsa.PrivateKey

func GenerateToken(key string, obj interface{}) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": obj,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})
	return t.SignedString([]byte(key))
}

func SetPrivateKey(p *ecdsa.PrivateKey) {
	privatekey = p
}
