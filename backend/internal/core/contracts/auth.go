package contracts

import "github.com/awanku/awanku/backend/pkg/model"

type AuthProvider interface {
	LoginURL(state string) string
	ExchangeCode(code string) (*model.AuthUserData, error)
}

type AuthStore interface {
	CreateAuthorizationCode(userID int64, code string) (*model.OauthAuthorizationCode, error)
	GetAuthorizationCodeByCode(code string) (*model.OauthAuthorizationCode, error)
	CreateOauthToken(token *model.OauthToken) error
	GetOauthToken(id int64) (*model.OauthToken, error)
	DeleteOauthToken(id int64) error
}
