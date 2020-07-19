package contract

import "github.com/awanku/awanku/pkg/core"

type UserSettingsStore interface {
    GetByID(id int64) (*core.Settings, error)
    GetOrCreateByUserID(userID int64) (*core.UserSettings, error)
    Update(userSettings *core.UserSettings) (*core.Settings, error)
}
