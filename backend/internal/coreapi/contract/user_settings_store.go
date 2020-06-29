package contract

import "github.com/awanku/awanku/pkg/core"

type UserSettingsStore interface {
    GetByID(id int64) (*core.Settings, error)
    GetByUserID(userID int64) (*core.Settings, error)
}
