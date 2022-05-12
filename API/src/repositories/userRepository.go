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
func (repo user) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, userName, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.UserName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
