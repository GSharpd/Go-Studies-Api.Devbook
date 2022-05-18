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
		insert into posts (title, content, posterId)
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

// Gets a specific post from the database
func (repo postsRepo) GetPostByID(postID uint64) (models.Post, error) {
	row, err := repo.db.Query(`
		select p.*, u.userName from posts p
		inner join users u on u.id = p.posterId
		where p.id = ?
	`, postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PosterID,
			&post.Likes,
			&post.CreatedAt,
			&post.UserName,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
