package model

import "time"

type OauthAuthorizationCode struct {
	Code      string
	UserID    int64
	ExpiresAt time.Time
}
