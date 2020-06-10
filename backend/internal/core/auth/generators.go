package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/awanku/awanku/backend/pkg/model"
)

func createRandomHash(length int) ([]byte, error) {
	buff := make([]byte, length)
	_, err := rand.Read(buff)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func buildAuthorizationCode(length int) (string, error) {
	buff, err := createRandomHash(length)
	if err != nil {
		return "", nil
	}
	code := base64.URLEncoding.EncodeToString(buff)
	return code, nil
}

func buildOauthToken(key []byte, tokensLength int) (*model.OauthToken, error) {
	accessToken, err := createRandomHash(tokensLength)
	if err != nil {
		return nil, err
	}
	refreshToken, err := createRandomHash(tokensLength)
	if err != nil {
		return nil, err
	}

	accessTokenHash, err := hmacHash(key, []byte(accessToken))
	if err != nil {
		return nil, err
	}

	refreshTokenHash, err := hmacHash(key, []byte(refreshToken))
	if err != nil {
		return nil, err
	}

	token := &model.OauthToken{
		AccessToken:      accessToken,
		AccessTokenHash:  accessTokenHash,
		RefreshToken:     refreshToken,
		RefreshTokenHash: refreshTokenHash,
		ExpiresAt:        time.Now().Add(60 * time.Minute),
	}
	return token, nil
}

func hmacHash(key, b []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(b)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}
