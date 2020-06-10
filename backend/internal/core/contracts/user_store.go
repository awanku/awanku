package contracts

import "github.com/awanku/awanku/backend/pkg/model"

type UserStore interface {
	FindOrCreateByEmail(user *model.User) error
	FindByID(id int64) (*model.User, error)
	Save(user *model.User) error
}
