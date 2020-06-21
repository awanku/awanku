package core_test

import (
	"testing"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/pkg/core"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T, db *hansip.Cluster) *core.User {
	googleEmail := faker.Username() + "." + faker.Email()
	githubUsername := faker.Username() + "_" + faker.Username()
	user := &core.User{
		Name:                faker.Username() + " " + faker.Name(),
		Email:               googleEmail,
		GoogleLoginEmail:    &googleEmail,
		GithubLoginUsername: &githubUsername,
	}

	userStore := core.UserStore{DB: db}
	err := userStore.GetOrCreateByEmail(user)
	assert.NoError(t, err)
	return user
}
