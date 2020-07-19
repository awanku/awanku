package core

import (
	"errors"

	hansip "github.com/asasmoyo/pq-hansip"
)

type UserSettingsStore struct {
	db *hansip.Cluster
}

func (u *UserSettingsStore) GetByID(id int64) (*Settings, error) {
	var query = `
		SELECT *
		FROM user_settings
		WHERE id = ?
	`

	var userSettings UserSettings
	err := u.db.Query(&userSettings, query, id)
	if err != nil {
		return nil, err
	}

	return userSettings.Settings, nil
}

func (u *UserSettingsStore) GetOrCreateByUserID(userID int64) (*UserSettings, error) {
    var query = `
        insert into user_settings (user_id, settings)
        values (?, ?)
        on conflict (user_id) do update set user_id = ?
        returning user_id, settings
    `

    userSettings := &UserSettings{}
    settings := &Settings{}

    err := u.db.WriterQuery(&userSettings, query, userID, settings, userID)
    if err != nil {
        return nil, err
    }

    return userSettings, nil
}

func (u *UserSettingsStore) Update(userSettings *UserSettings) (*Settings, error) {
    if userSettings.ID <= 0 {
        return nil, errors.New("model does not have id set")
    }

    var query = `
        update user_settings
        set settings = settings | ?
        where id = ?
        returning settings
    `

    settings := &Settings{}

    err := u.db.WriterQuery(settings, query, userSettings.Settings, userSettings.ID)
    if err != nil {
        return nil, err
    }
    return settings, nil
}
