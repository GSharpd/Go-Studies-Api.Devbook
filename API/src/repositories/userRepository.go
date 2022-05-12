package repositories

import (
	"api/src/models"
	"database/sql"
)

type user struct {
	db *sql.DB
}

// Creates a new user repository using the given User object
func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

// Creates a new user and inserts it into the database
func (u user) Create(user models.User) (uint64, error) {
	return 0, nil
}
