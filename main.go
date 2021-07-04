package main

import (
	"github.com/golangsugar/env"
	"github.com/miguelpragier/pgkebab"
	"log"
	"os"
)

var (
	db        *pgkebab.DBLink
	secretKey []byte
)

func dbConnection() error {
	log.Println("connecting database ...")

	cn := pgkebab.ConnStringDirect("postgres://postgres:postgres123@localhost:5432/authservice?sslmode=disable")

	const (
		connTimeout                        = 10
		execTimeout                        = 10
		connAttemptsMax                    = 10
		connAttemptsMaxMinutes             = 10
		secondsBetweenReconnectionAttempts = 10
		debugPrint                         = true
	)

	opts := pgkebab.Options(cn, connTimeout, execTimeout, connAttemptsMax, connAttemptsMaxMinutes, secondsBetweenReconnectionAttempts, debugPrint)

	_db, err := pgkebab.NewConnected(opts)

	if err != nil {
		log.Println("failed to connect database")
		return err
	}

	db = _db

	return nil
}

func loadSettings() error {
	return env.Check(`SECRET_KEY`, `@xente#nutri`, true, true)
}

func main() {
	if err := dbConnection(); err != nil {
		log.Fatal(err)
	}

	if err := loadSettings(); err != nil {
		log.Fatal(err)
	}

	secretKey = []byte(os.Getenv("SECRET_KEY"))

	webserviceStart()
}
