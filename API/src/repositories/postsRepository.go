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

// Gets all the posts for a specific user id
func (repo postsRepo) GetPostsForUser(userID uint64) ([]models.Post, error) {
	rows, err := repo.db.Query(`
		SELECT DISTINCT
		p.id, p.posterId , u.userName , p.title , p.content , p.likes , p.createdAt 
		FROM posts p
		INNER JOIN users u ON u.id = p.posterId 
		INNER JOIN followers f ON f.followerId  = ?
		WHERE p.posterId  = ? OR p.posterId  = f.userId 
		ORDER BY p.id DESC;
	`, userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.PosterID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Updates the specified post by its ID
func (repo postsRepo) UpdatePost(postID uint64, post models.Post) error {
	statement, err := repo.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// Deletes the specified post by its ID
func (repo postsRepo) DeletePost(postID uint64) error {
	statement, err := repo.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Gets the specified user posts from the database
func (repo postsRepo) GetPostsByUserID(userID uint64) ([]models.Post, error) {
	rows, err := repo.db.Query(`
		select p.*, u.userName from posts p
		join users u on u.id = p.posterId 
		where p.posterId = ?`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PosterID,
			&post.Likes,
			&post.CreatedAt,
			&post.UserName,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
