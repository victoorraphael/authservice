package main

import (
	"github.com/miguelpragier/pgkebab"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Crn   string `json:"crn"`
}

func (u *user) save(exp time.Time) error {
	userPairs := pgkebab.Pairs(
		`user_name`, u.Name,
		`user_email`, u.Email,
		`user_pass`, u.Pass,
		`crn`, u.Crn,
		`expiration`, exp)
	if err := db.Insert(`credentials`, userPairs); err != nil {
		return err
	}

	return nil
}

func (u *user) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.Pass = string(bytes)

	return nil
}
