package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func loginUser(c echo.Context) error {
	var userResponse user
	var u user

	if err := c.Bind(&userResponse); err != nil {
		log.Println(" erro no JSON fornecido => ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"msg": "invalid JSON"})
	}

	row, err := db.GetOne("SELECT name, email, password, crn FROM users WHERE email=$1", userResponse.Email)
	if err != nil {
		if db.IsEmptyErr(err) {
			log.Println(" usuário não existe no banco")
			return c.JSON(http.StatusBadRequest, map[string]string{"msg": "usuário não existe"})
		}
		log.Println(" erro ao tentar buscar usuário no banco")
		c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar buscar usuário"})
	}

	u.Name = row.String("name")
	u.Email = row.String("email")
	u.Pass = row.String("password")
	u.Crn = row.String("crn")

	if err := u.checkPassword(userResponse.Pass); err != nil {
		log.Println(" senha inválida")
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "credenciais inválidas"})
	}

	jw := jwtWrapper{
		SecretKey:  os.Getenv("SECRET_KEY"),
		Issuer:     "authservice",
		Expiration: 24,
	}

	signedToken, err := jw.generateToken(u.Name, u.Email)

	if err != nil {
		log.Println("erro ao tentar gerar token")
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "erro ao tentar gerar token"})
	}

	log.Println("token gerado com sucesso")
	return c.JSON(http.StatusOK, map[string]string{"token": signedToken})
}
