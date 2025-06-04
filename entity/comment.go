package entity

import (
	"errors"
	"strings"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	AuthorID  int64  `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

func (c *Comment) Validate() error {
	if len(strings.TrimSpace(c.Content)) < 1 {
		return errors.New("comment content required")
	}
	if c.PostID <= 0 {
		return errors.New("invalid post id")
	}
	return nil
}
