package contracts

import "github.com/awanku/awanku/backend/pkg/model"

type AuthProvider interface {
	LoginURL(state string) string
	ExchangeCode(code string) (*model.AuthUserData, error)
}
