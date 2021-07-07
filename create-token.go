package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func createToken(c echo.Context) error {
	var u user

	if err := c.Bind(&u); err != nil {
		log.Println(time.Now(), " erro no JSON fornecido => ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"msg": "invalid JSON"})
	}

	if err := u.hashPassword(u.Pass); err != nil {
		log.Println(time.Now(), " erro ao tentar realizar o hash da senha")
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar salvar a senha"})
	}

	if err := u.save(); err != nil {
		log.Println(time.Now(), " erro ao tentar gravar usuário no banco")
		c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar criar usuário"})
	}

	return c.JSON(http.StatusOK, &u)
}
