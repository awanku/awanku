package model

import "time"

type User struct {
	ID                  int64
	Name                string
	Email               string
	GoogleLoginEmail    *string
	GithubLoginUsername *string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	DeletedAt           *time.Time
}

func (u *User) SetOauth2Identifier(provider string, identifier *string) {
	switch provider {
	case "github":
		u.GithubLoginUsername = identifier
	case "google":
		u.GoogleLoginEmail = identifier
	}
}
