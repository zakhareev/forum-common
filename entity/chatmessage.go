package entity

import (
	"errors"
	"strings"
)

type ChatMessage struct {
	ID        int64  `json:"id"`
	AuthorID  int64  `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

func (m *ChatMessage) Validate() error {
	if len(strings.TrimSpace(m.Content)) == 0 {
		return errors.New("message content required")
	}
	return nil
}
