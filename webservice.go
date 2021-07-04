package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func corsConfig() middleware.CORSConfig {
	cc := middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderOrigin},
		AllowCredentials: true,
		ExposeHeaders:    []string{echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderOrigin},
	}

	return cc
}

func webserviceStart() {
	log.Println("starting webservice ...")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(corsConfig()))

	r := e.Group("/auth")

	r.POST("/create/", createToken)

	//r.POST("/check/", isValid)
	//
	//r.POST("/revoke/", revokeToken)
	//
	//r.POST("/refresh/", refreshToken)

	e.Logger.Fatal(e.Start(":4000"))
}