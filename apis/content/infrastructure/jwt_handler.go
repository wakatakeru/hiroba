package infrastructure

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
)

type JWTHandler struct {
	PubKey *rsa.PublicKey
}

type JSONWebToken struct {
	Token jwt.Token
}

type CustomClaims struct {
	Name string `json:name`
	jwt.StandardClaims
}

func NewJWTHandler() *JWTHandler {
	pubBytes, err := ioutil.ReadFile(os.Getenv("HIROBA_CONTENT_JWT_PUB_KEY_PATH"))
	if err != nil {
		glog.Fatal(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		glog.Fatal(err)
	}

	jwtHandler := new(JWTHandler)
	jwtHandler.PubKey = pubKey

	return jwtHandler
}

func (handler *JWTHandler) Verify(token string) (userID int, err error) {
	jwtHandler := NewJWTHandler()
	jsonWebToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtHandler.PubKey, nil
	})

	if claims, ok := jsonWebToken.Claims.(*CustomClaims); ok && jsonWebToken.Valid {
		userID, _ = strconv.Atoi(claims.Id)
		return userID, nil
	}

	return userID, errors.New("Failed Parse JWT")
}
