package core

import hansip "github.com/asasmoyo/pq-hansip"

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

func (u *UserSettingsStore) GetByUserID(userID int64) (*Settings, error) {
	var query = `
		SELECT *
		FROM user_settings
		WHERE user_id = ?
	`

	var userSettings UserSettings
	err := u.db.Query(&userSettings, query, userID)
	if err != nil {
		return nil, err
	}

	return userSettings.Settings, nil
}
