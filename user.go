package main

import (
	"github.com/miguelpragier/pgkebab"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Crn   string `json:"crn"`
}

func (u *user) save() error {
	userPairs := pgkebab.Pairs(
		`name`, u.Name,
		`email`, u.Email,
		`password`, u.Pass,
		`crn`, u.Crn)
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

func (u *user) checkPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(password)); err != nil {
		return err
	}

	return nil
}
