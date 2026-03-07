package models

import (
	"strings"
	"testing"
)

func TestValidateFirstName(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"valid first name": {
			input:  "John",
			output: true,
		},
		"empty first name": {
			input:  "  ",
			output: false,
		},
		"too long of first name": {
			input:  strings.Repeat("foobar", 25),
			output: false,
		},
	}

	for name, test := range tests {
		user := User{
			Id:        0,
			FirstName: test.input,
			LastName:  "Doe",
			Email:     "hello@email.com",
		}

		t.Run(name, func(t *testing.T) {
			result := user.Validate() == nil
			if result != test.output {
				t.Errorf("got %t, want %t", result, test.output)
			}
		})
	}
}

func TestValidateLastName(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"valid last name": {
			input:  "Doe",
			output: true,
		},
		"empty last name": {
			input:  "  ",
			output: false,
		},
		"too long of last name": {
			input:  strings.Repeat("foobar", 25),
			output: false,
		},
	}

	for name, test := range tests {
		user := User{
			Id:        0,
			FirstName: "John",
			LastName:  test.input,
			Email:     "hello@email.com",
		}

		t.Run(name, func(t *testing.T) {
			result := user.Validate() == nil
			if result != test.output {
				t.Errorf("got %t, want %t", result, test.output)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"valid email": {
			input:  "hello@email.com",
			output: true,
		},
		"empty email": {
			input:  "  ",
			output: false,
		},
		"invalid email": {
			input:  "foobar",
			output: false,
		},
	}

	for name, test := range tests {
		user := User{
			Id:        0,
			FirstName: "John",
			LastName:  "Doe",
			Email:     test.input,
		}

		t.Run(name, func(t *testing.T) {
			result := user.Validate() == nil
			if result != test.output {
				t.Errorf("got %t, want %t", result, test.output)
			}
		})
	}
}
