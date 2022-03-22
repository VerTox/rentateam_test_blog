package model

import (
	"errors"
	"time"
)

type Post struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	Tags      []string   `json:"tags"`
	CreatedAt *time.Time `json:"created_at"`
}

type PostRepository interface {
	GetList() ([]*Post, error)
	Store(p *Post) (*Post, error)
}

func (p *Post) Valid() error {
	if p.Title == "" {
		return errors.New("empty title")
	}
	if p.Text == "" {
		return errors.New("empty text")
	}
	return nil
}
