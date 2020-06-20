package core

import (
	hansip "github.com/asasmoyo/pq-hansip"
)

type AuthStore struct {
	db *hansip.Cluster
}

func (s *AuthStore) CreateOauthAuthorizationCode(userID int64, code string) (*OauthAuthorizationCode, error) {
	var query = `
        insert into oauth_authorization_codes (user_id, code, expires_at)
        values (?, ?, now() + interval '5 minutes')
        returning *
    `
	var codeObj OauthAuthorizationCode
	err := s.db.WriterQuery(&codeObj, query, userID, code)
	if err != nil {
		return nil, err
	}
	return &codeObj, nil
}

func (s *AuthStore) GetOauthAuthorizationCodeByCode(code string) (*OauthAuthorizationCode, error) {
	var query = `
        delete from oauth_authorization_codes
        where code = ? and expires_at > now()
        returning *
    `
	var codeObj OauthAuthorizationCode
	err := s.db.Query(&codeObj, query, code)
	if err != nil {
		return nil, err
	}
	return &codeObj, nil
}

func (s *AuthStore) CreateOauthToken(token *OauthToken) error {
	var query = `
        insert into oauth_tokens (user_id, access_token_hash, refresh_token_hash, expires_at, requester_ip, requester_user_agent)
        values (?, ?, ?, ?, ?, ?)
        returning id
    `

	var id int64
	err := s.db.WriterQuery(&id, query, token.UserID, token.AccessTokenHash, token.RefreshTokenHash, token.ExpiresAt, token.RequesterIP, token.RequesterUserAgent)
	if err != nil {
		return err
	}

	token.ID = id
	return nil
}

func (s *AuthStore) GetOauthTokenByID(id int64) (*OauthToken, error) {
	var query = `
        select *
        from oauth_tokens
        where id = ? and expires_at > now() and deleted_at is null
    `
	var token OauthToken
	err := s.db.Query(&token, query, id)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *AuthStore) DeleteOauthToken(id int64) error {
	var query = `
        update oauth_tokens
        set deleted_at = now()
        where id = ?
    `
	return s.db.WriterExec(query, id)
}
