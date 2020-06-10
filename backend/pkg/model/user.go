package model

import "time"

type User struct {
	ID                  int64      `json:"id"`
	Name                string     `json:"name"`
	Email               string     `json:"email"`
	GoogleLoginEmail    *string    `json:"google_login_email"`
	GithubLoginUsername *string    `json:"github_login_username"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
	DeletedAt           *time.Time `json:"-"`
}

func (u *User) SetOauth2Identifier(provider string, identifier *string) {
	switch provider {
	case "github":
		u.GithubLoginUsername = identifier
	case "google":
		u.GoogleLoginEmail = identifier
	}
}
