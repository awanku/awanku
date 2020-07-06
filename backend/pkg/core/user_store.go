package core

import (
	"errors"
	"time"

	hansip "github.com/asasmoyo/pq-hansip"
)

type UserStore struct {
	DB *hansip.Cluster
}

func (s *UserStore) GetOrCreateByEmail(user *User) error {
	var query = `
        insert into users (name, email, google_login_email, github_login_username)
        values (?, ?, ?, ?)
        on conflict (email) do update set updated_at = now()
        returning id, created_at, updated_at
    `

	returned := struct {
		ID        int64
		CreatedAt time.Time
		UpdatedAt *time.Time
	}{}
	err := s.DB.WriterQuery(&returned, query, user.Name, user.Email, user.GoogleLoginEmail, user.GithubLoginUsername)
	if err != nil {
		return err
	}

	user.ID = returned.ID
	user.CreatedAt = returned.CreatedAt
	user.UpdatedAt = returned.UpdatedAt
	return nil
}

func (s *UserStore) GetByID(id int64) (*User, error) {
	var query = `
        select *
        from users
        where id = ? and deleted_at is null
    `
	var user User
	err := s.DB.Query(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) Save(user *User) error {
	if user.ID <= 0 {
		return errors.New("model does not have id set")
	}

	var query = `
        update users
        set name = ?, email = ?, google_login_email = ?, github_login_username = ?, updated_at = now()
        where id = ?
        returning updated_at
    `
	var returned struct {
		UpdatedAt time.Time
	}
	err := s.DB.WriterQuery(&returned, query, user.Name, user.Email, user.GoogleLoginEmail, user.GithubLoginUsername, user.ID)
	if err != nil {
		return err
	}
	user.UpdatedAt = &returned.UpdatedAt
	return nil
}
