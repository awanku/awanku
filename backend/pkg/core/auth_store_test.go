package core_test

import (
	"testing"
	"time"

	"github.com/awanku/awanku/pkg/core"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestAuthStore_CreateOauthAuthorizationCode(t *testing.T) {
	authStore := &core.AuthStore{DB: db}

	t.Run("user does not exist", func(t *testing.T) {
		_, err := authStore.CreateOauthAuthorizationCode(-1, "-1randomcode")
		assert.Error(t, err, "index violation")
	})

	t.Run("user exists", func(t *testing.T) {
		user := createUser(t, db)
		code := faker.Word() + faker.Word()
		codeObj, err := authStore.CreateOauthAuthorizationCode(user.ID, code)
		assert.NoError(t, err)
		assert.NotNil(t, codeObj)
		assert.True(t, codeObj.ExpiresAt.Before(time.Now().Add(5*time.Minute)), "max 5 minutes")
	})
}

func TestAuthStore_GetOauthAuthorizationCodeByCode(t *testing.T) {
	authStore := &core.AuthStore{DB: db}

	t.Run("code does not exist", func(t *testing.T) {
		code, err := authStore.GetOauthAuthorizationCodeByCode("doesnotexist1451")
		assert.NoError(t, err)
		assert.Nil(t, code)
	})

	t.Run("code exists", func(t *testing.T) {
		user := createUser(t, db)
		code := faker.Word() + faker.Word()
		codeObj1, err := authStore.CreateOauthAuthorizationCode(user.ID, code)
		assert.NoError(t, err)

		codeObj2, err := authStore.GetOauthAuthorizationCodeByCode(code)
		assert.NoError(t, err)

		assert.Equal(t, codeObj1.Code, codeObj2.Code)
	})
}
