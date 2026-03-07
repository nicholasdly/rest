package models

import (
	"fmt"
	"strings"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

func (user *User) Validate() error {
	if strings.TrimSpace(user.FirstName) == "" {
		return fmt.Errorf("invalid first name \"%s\"", user.FirstName)
	}
	if strings.TrimSpace(user.LastName) == "" {
		return fmt.Errorf("invalid last name \"%s\"", user.LastName)
	}
	if strings.TrimSpace(user.Email) == "" {
		return fmt.Errorf("invalid email \"%s\"", user.Email)
	}
	return nil
}
