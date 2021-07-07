package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtWrapper struct {
	SecretKey  string
	Issuer     string
	Expiration int64
}

type jwtClaims struct {
	Name  string
	Email string
	jwt.StandardClaims
}

func (j *jwtWrapper) generateToken(name, email string) (signedToken string, err error) {
	claims := &jwtClaims{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.Expiration)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

func (j *jwtWrapper) validateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		err = errors.New("Não foi possível realizar o parse dos claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expirado")
		return
	}

	return
}
