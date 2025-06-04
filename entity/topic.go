package entity

import (
	"errors"
	"strings"
)

type Topic struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"category_id"`
	Title      string `json:"title"`
	AuthorID   int64  `json:"author_id"`
	CreatedAt  int64  `json:"created_at"`
}

func (t *Topic) Validate() error {
	if len(strings.TrimSpace(t.Title)) < 5 {
		return errors.New("topic title too short")
	}
	if t.CategoryID <= 0 {
		return errors.New("invalid category id")
	}
	return nil
}
