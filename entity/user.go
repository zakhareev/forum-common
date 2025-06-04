package entity

import (
	"errors"
	"regexp"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	CreatedAt    int64  `json:"created_at"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return errors.New("username too short")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(u.Username) {
		return errors.New("invalid username: only letters, numbers, underscore allowed")
	}
	// Role can be "guest", "user", "admin"
	switch u.Role {
	case "guest", "user", "admin":
	default:
		return errors.New("invalid role")
	}
	return nil
}
