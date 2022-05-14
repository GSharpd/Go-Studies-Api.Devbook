package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("parameter 'name' is required and cannot be empty")
	}

	if user.UserName == "" {
		return errors.New("parameter 'username' is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("parameter 'email' is required and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if stage == "signup" && user.Password == "" {
		return errors.New("parameter 'password' is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
}
