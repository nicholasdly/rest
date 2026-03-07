package models

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (user *User) Validate() error {
	if strings.TrimSpace(user.FirstName) == "" {
		return errors.New("invalid first name")
	}
	if utf8.RuneCountInString(user.FirstName) > 100 {
		return errors.New("first name must not exceed 100 characters")
	}
	if strings.TrimSpace(user.LastName) == "" {
		return errors.New("invalid last name")
	}
	if utf8.RuneCountInString(user.LastName) > 100 {
		return errors.New("last name must not exceed 100 characters")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}
