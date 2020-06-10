package model

import (
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

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
