package datastore

import (
	"errors"
	"time"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/backend/pkg/model"
)

type UserStore struct {
	db *hansip.Cluster
}

func (s *UserStore) FindOrCreateByEmail(user *model.User) error {
	var query = `
        insert into users (name, email, google_login_email, github_login_username)
        values (?, ?, ?, ?)
        on conflict (email) do update set updated_at = now()
        returning id, created_at
    `

	returned := struct {
		ID        int64
		CreatedAt time.Time
	}{}
	err := s.db.WriterQuery(&returned, query, user.Name, user.Email, user.GoogleLoginEmail, user.GithubLoginUsername)
	if err != nil {
		return err
	}

	user.ID = returned.ID
	user.CreatedAt = returned.CreatedAt
	return nil
}

func (s *UserStore) FindByID(id int64) (*model.User, error) {
	var query = `
        select *
        from users
        where id = ? and deleted_at is null
    `
	var user model.User
	err := s.db.Query(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) Save(user *model.User) error {
	if user.ID <= 0 {
		return errors.New("model does not have id set")
	}

	var query = `
        update users
        set name = ?, email = ?, google_login_email = ?, github_login_username = ?, updated_at = now()
        where id = ?
        returning updated_at
    `
	var updatedAt time.Time
	err := s.db.WriterQuery(&updatedAt, query, user.Name, user.Email, user.GoogleLoginEmail, user.GithubLoginUsername, user.ID)
	if err != nil {
		return err
	}
	return nil
}
