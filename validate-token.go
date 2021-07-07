package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func isValid(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		log.Println(time.Now().Local(), " token ausente")
		return c.JSON(http.StatusForbidden, map[string]string{"msg": "formato incorreto de authorization header"})
	}

	jw := jwtWrapper{
		SecretKey: os.Getenv("SECRET_KEY"),
		Issuer:    "authservice",
	}

	claims, err := jw.validateToken(token)
	if err != nil {
		log.Println(time.Now().Local(), " token expirado ou inválido")
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "token inválido"})
	}

	log.Println(time.Now().Local(), " usuário autorizado")
	return c.JSON(http.StatusOK, claims)
}
