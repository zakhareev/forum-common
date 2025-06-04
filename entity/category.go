package entity

import (
	"errors"
	"strings"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Validate() error {
	if len(strings.TrimSpace(c.Name)) < 2 {
		return errors.New("category name too short")
	}
	return nil
}
