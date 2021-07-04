package main

import (
	"github.com/miguelpragier/pgkebab"
	"time"
)

type user struct {
	Id   int    `json:"user_id"`
	Pass string `json:"user_pass"`
}

func (u *user) save(exp time.Time) error {
	userPairs := pgkebab.Pairs(
		`user_id`, u.Id,
		`user_pass`, u.Pass,
		`expiration`, exp)
	if err := db.Insert(`credentials`, userPairs); err != nil {
		return err
	}

	return nil
}
