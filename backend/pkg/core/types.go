package core

import (
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

// supported oauth providers
const (
	OauthProviderGithub               = "github"
	OauthProviderGoogle               = "google"
	OauthAuthorizationCodeMaxDuration = 5 * time.Minute
)

// OauthUserData represents user data provided by third party oauth services
type OauthUserData struct {
	Provider   string `json:"provider"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Identifier string `json:"identifier"`
}

// OauthAuthorizationCode represents oauth authorization code
type OauthAuthorizationCode struct {
	Code      string
	UserID    int64
	ExpiresAt time.Time
}

// OauthToken represents oauth token
type OauthToken struct {
	ID                 int64
	UserID             int64
	AccessToken        []byte
	AccessTokenHash    []byte
	RefreshToken       []byte
	RefreshTokenHash   []byte
	ExpiresAt          time.Time
	RequesterIP        string
	RequesterUserAgent string
	DeletedAt          *time.Time
}

// Token returns standar token representation
func (t *OauthToken) Token() *oauth2.Token {
	encodedAccessToken := base64.URLEncoding.EncodeToString(t.AccessToken)
	encodedRefreshToken := base64.URLEncoding.EncodeToString(t.RefreshToken)
	return &oauth2.Token{
		AccessToken:  fmt.Sprintf("%d:%s", t.ID, encodedAccessToken),
		RefreshToken: fmt.Sprintf("%d:%s", t.ID, encodedRefreshToken),
		Expiry:       t.ExpiresAt,
		TokenType:    "bearer",
	}
}

// User represents User
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

// SetOauth2Identifier sets identifier based on provider
func (u *User) SetOauth2Identifier(provider string, identifier *string) {
	switch provider {
	case OauthProviderGithub:
		u.GithubLoginUsername = identifier
	case OauthProviderGoogle:
		u.GoogleLoginEmail = identifier
	}
}
