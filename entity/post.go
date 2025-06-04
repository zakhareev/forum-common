package entity

import (
	"errors"
	"strings"
)

type Post struct {
	ID        int64  `json:"id"`
	TopicID   int64  `json:"topic_id"`
	AuthorID  int64  `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

func (p *Post) Validate() error {
	if len(strings.TrimSpace(p.Title)) < 3 {
		return errors.New("post title too short")
	}
	if len(strings.TrimSpace(p.Content)) < 1 {
		return errors.New("post content required")
	}
	if p.TopicID <= 0 {
		return errors.New("invalid topic id")
	}
	return nil
}
