package core_test

import (
	"testing"

	"github.com/awanku/awanku/pkg/core"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestUserStore_GetOrCreateByEmail(t *testing.T) {
	userStore := &core.UserStore{DB: db}

	t.Run("no existing user with google", func(t *testing.T) {
		googleEmail := faker.Email()
		user := &core.User{
			Name:             faker.Name(),
			Email:            faker.Email(),
			GoogleLoginEmail: &googleEmail,
		}
		err := userStore.GetOrCreateByEmail(user)
		assert.NoError(t, err)
		assert.Greater(t, user.ID, int64(0))
	})

	t.Run("there is existing user with google", func(t *testing.T) {
		googleEmail := faker.Email()
		user1 := core.User{
			Name:             faker.Name(),
			Email:            faker.Email(),
			GoogleLoginEmail: &googleEmail,
		}
		user2 := user1

		err := userStore.GetOrCreateByEmail(&user1)
		assert.NoError(t, err)
		assert.Greater(t, user1.ID, int64(0))
		assert.Nil(t, user1.UpdatedAt)

		err = userStore.GetOrCreateByEmail(&user2)
		assert.NoError(t, err)
		assert.Greater(t, user2.ID, int64(0))
		assert.Equal(t, user2.ID, user1.ID)

		// must update updatedAt
		assert.NotNil(t, user2.UpdatedAt)
	})

	t.Run("no existing user with github", func(t *testing.T) {
		githubUsername := faker.Username()
		user := &core.User{
			Name:                faker.Name(),
			Email:               faker.Email(),
			GithubLoginUsername: &githubUsername,
		}
		err := userStore.GetOrCreateByEmail(user)
		assert.NoError(t, err)
		assert.Greater(t, user.ID, int64(0))
	})

	t.Run("there is existing user with github", func(t *testing.T) {
		githubUsername := faker.Username()
		user1 := core.User{
			Name:                faker.Name(),
			Email:               faker.Email(),
			GithubLoginUsername: &githubUsername,
		}
		user2 := user1

		err := userStore.GetOrCreateByEmail(&user1)
		assert.NoError(t, err)
		assert.Greater(t, user1.ID, int64(0))
		assert.Nil(t, user1.UpdatedAt)

		err = userStore.GetOrCreateByEmail(&user2)
		assert.NoError(t, err)
		assert.Greater(t, user2.ID, int64(0))
		assert.Equal(t, user2.ID, user1.ID)

		// must update updatedAt
		assert.NotNil(t, user2.UpdatedAt)
	})
}

func TestUserStore_GetByID(t *testing.T) {
	userStore := &core.UserStore{DB: db}

	t.Run("user does not exist", func(t *testing.T) {
		user, err := userStore.GetByID(-1)
		assert.NoError(t, err)
		assert.Zero(t, user.ID)
	})

	t.Run("user exists", func(t *testing.T) {
		githubUsername := faker.Username()
		user := &core.User{
			Name:                faker.Name(),
			Email:               faker.Email(),
			GithubLoginUsername: &githubUsername,
		}
		err := userStore.GetOrCreateByEmail(user)

		retrievedUser, err := userStore.GetByID(user.ID)
		assert.NoError(t, err)
		assert.Greater(t, retrievedUser.ID, int64(0))
	})
}

func TestUserStore_Save(t *testing.T) {
	userStore := &core.UserStore{DB: db}

	githubUsername := faker.Username()
	user := &core.User{
		Name:                faker.Name(),
		Email:               faker.Email(),
		GithubLoginUsername: &githubUsername,
	}
	err := userStore.GetOrCreateByEmail(user)
	assert.NoError(t, err)

	assert.Nil(t, user.GoogleLoginEmail)

	email := faker.Email()
	user.GoogleLoginEmail = &email
	err = userStore.Save(user)
	assert.NoError(t, err)
	assert.NotNil(t, user.UpdatedAt)
}
