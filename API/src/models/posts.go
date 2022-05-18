package models

import "time"

type Post struct {
	ID         uint64    `json:"id, omitempty"`
	Title      string    `json:"title, omitempty"`
	Content    string    `json:"content, omitempty"`
	PosterID   uint64    `json:"poster_id, omitempty"`
	PosterName string    `json:"poster_name, omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at, omitempty"`
}
