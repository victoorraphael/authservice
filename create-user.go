package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func createUser(c echo.Context) error {
	var u user

	if err := c.Bind(&u); err != nil {
		log.Println(" erro no JSON fornecido => ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"msg": "invalid JSON"})
	}

	if err := u.hashPassword(u.Pass); err != nil {
		log.Println(" erro ao tentar realizar o hash da senha")
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar salvar a senha"})
	}

	if err := u.save(); err != nil {
		log.Println(" erro ao tentar gravar usuário no banco")
		c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar criar usuário"})
	}

	log.Println("usuario criado com sucesso")
	return c.JSON(http.StatusOK, &u)
}
