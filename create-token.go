package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func createToken(c echo.Context) error {
	var u user

	if err := c.Bind(&u); err != nil {
		log.Println("ERROR: invalid JSON")
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "invalid JSON"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = u.Id
	claims["pass"] = u.Pass
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Println("Somenthing wrong while generating token", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "erro ao tentar gerar token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
