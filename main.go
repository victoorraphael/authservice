package main

import (
	"github.com/miguelpragier/pgkebab"
	"log"
)

var db *pgkebab.DBLink

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

func main()  {
	if err := dbConnection(); err != nil {
		log.Fatal(err)
	}
}