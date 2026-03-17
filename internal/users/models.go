package users

import (
	"errors"
	"net/mail"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (req *CreateUserRequest) Validate() error {
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("invalid first name")
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}

func (req *UpdateUserRequest) Validate() error {
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("invalid first name")
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}
