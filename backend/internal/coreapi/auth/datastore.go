package auth

import (
	"context"
	"time"

	"github.com/awanku/awanku/internal/coreapi/appctx"
	"github.com/awanku/awanku/pkg/core"
)

func getUserByID(ctx context.Context, id int64) (*core.User, error) {
	db := appctx.Database(ctx)

	var query = `
        select *
        from users
        where id = ? and deleted_at is null
    `
	var returned core.User
	err := db.Query(&returned, query, id)
	if err != nil {
		return nil, err
	}
	if returned.ID == 0 {
		return nil, nil
	}
	return &returned, nil
}

func getOrCreateUserByEmail(ctx context.Context, user *core.User) error {
	db := appctx.Database(ctx)

	var query = `
        insert into users (name, email, google_login_email, github_login_username)
        values (?, ?, ?, ?)
        on conflict (email) do update set updated_at = now()
        returning id, created_at, updated_at
    `
	var returned struct {
		ID        int64
		CreatedAt time.Time
		UpdatedAt *time.Time
	}
	err := db.WriterQuery(&returned, query, user.Name, user.Email, user.GoogleLoginEmail, user.GithubLoginUsername)
	if err != nil {
		return err
	}

	user.ID = returned.ID
	user.CreatedAt = returned.CreatedAt
	user.UpdatedAt = returned.UpdatedAt
	return nil
}

func saveOauthAuthorizationCode(ctx context.Context, userID int64, code string) (*core.OauthAuthorizationCode, error) {
	db := appctx.Database(ctx)

	var query = `
        insert into oauth_authorization_codes (user_id, code, expires_at)
        values (?, ?, now() + interval '5 minutes')
        returning *
    `
	var returned core.OauthAuthorizationCode
	err := db.WriterQuery(&returned, query, userID, code)
	if err != nil {
		return nil, err
	}
	return &returned, nil
}

func getOauthAuthorizationCodeBycode(ctx context.Context, code string) (*core.OauthAuthorizationCode, error) {
	db := appctx.Database(ctx)

	var query = `
        delete from oauth_authorization_codes
        where code = ? and expires_at > now()
        returning *
    `
	var returned core.OauthAuthorizationCode
	err := db.Query(&returned, query, code)
	if err != nil {
		return nil, err
	}
	if returned.UserID == 0 || returned.Code == "" {
		return nil, nil
	}
	return &returned, nil
}

func getOauthTokenByID(ctx context.Context, id int64) (*core.OauthToken, error) {
	db := appctx.Database(ctx)

	var query = `
        select *
        from oauth_tokens
        where id = ? and expires_at > now() and deleted_at is null
    `
	var returned core.OauthToken
	err := db.Query(&returned, query, id)
	if err != nil {
		return nil, err
	}
	if returned.ID == 0 {
		return nil, nil
	}
	return &returned, nil
}

func deleteOauthToken(ctx context.Context, id int64) error {
	db := appctx.Database(ctx)

	var query = `
        update oauth_tokens
        set deleted_at = now()
        where id = ?
    `
	return db.WriterExec(query, id)
}

func saveOauthToken(ctx context.Context, token *core.OauthToken) error {
	db := appctx.Database(ctx)

	var query = `
        insert into oauth_tokens (user_id, access_token_hash, refresh_token_hash, expires_at, requester_ip, requester_user_agent)
        values (?, ?, ?, ?, ?, ?)
        returning id
    `
	var returned struct {
		ID int64
	}
	err := db.WriterQuery(&returned, query, token.UserID, token.AccessTokenHash, token.RefreshTokenHash, token.ExpiresAt, token.RequesterIP, token.RequesterUserAgent)
	if err != nil {
		return err
	}
	token.ID = returned.ID
	return nil
}
