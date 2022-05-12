package models

import (
	"errors"
	"strings"
	"time"
)

// Represent a user of one account
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserName  string    `json:"userName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Validates and formats a user
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("parameter 'name' is required and cannot be empty")
	}

	if user.UserName == "" {
		return errors.New("parameter 'username' is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("parameter 'email' is required and cannot be empty")
	}

	if user.Password == "" {
		return errors.New("parameter 'password' is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
}
