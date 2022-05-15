package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type user struct {
	db *sql.DB
}

// Creates a new user repository using the given db object
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

// Gets the specified users in the database with the specified Name or UserName
func (repo user) Get(nameOrUserName string) ([]models.User, error) {
	nameOrUserName = fmt.Sprintf("%%%s%%", nameOrUserName) // First % escapes the string, second one is wildcard for the SQL statement, resulting in "%nameOrUserName%"

	rows, err := repo.db.Query(
		"select id, name, userName, email, createdAt from users where name like ? or userName like ? ",
		nameOrUserName,
		nameOrUserName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.UserName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Gets a specific user by its ID from the database
func (repo user) GetUserByID(id uint64) (models.User, error) {
	rows, err := repo.db.Query(
		"select id, name, userName, email, createdAt from users where id = ?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.UserName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Updates a user by its ID on the database
func (repo user) UpdateUser(ID uint64, user models.User) error {
	statement, err := repo.db.Prepare("update users set name = ?, userName = ?, email = ? where ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.UserName, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Removes a user by its ID from the database
func (repo user) DeleteUser(ID uint64) error {
	statement, err := repo.db.Prepare("delete from users where ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// Gets an user by is email and returns its ID and hashed password if it exists
func (repo user) GetUserByEmail(email string) (models.User, error) {
	row, err := repo.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Inserts a follow record for the specified users in the database
func (repo user) Follow(userID, followerID uint64) error {
	statement, err := repo.db.Prepare(
		"insert ignore into followers (userId, followerId) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}
