package repositories

import (
	"api/src/models"
	"database/sql"
)

type postsRepo struct {
	db *sql.DB
}

// Creates a new instance of the posts repository struct
func NewPostsRepository(db *sql.DB) *postsRepo {
	return &postsRepo{db}
}

// Inserts a new post into the database
func (repo postsRepo) CreateNewPost(post models.Post) (uint64, error) {
	statement, err := repo.db.Prepare(`
		inser into posts (title, content, posterId)
		values (?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.PosterID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
