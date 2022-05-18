package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	PosterID  uint64    `json:"poster_id,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

//
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" || post.Content == "" {
		return errors.New("title and content cannot be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
