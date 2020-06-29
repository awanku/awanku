package core

import hansip "github.com/asasmoyo/pq-hansip"

type Config struct {
	DB                  *hansip.Cluster
	OauthTokenSecretKey []byte
}

type CoreService struct {
	db *hansip.Cluster

	userStore *UserStore
	authStore *AuthStore
    userSettingsStore *UserSettingsStore
}

func NewCoreService(conf *Config) (*CoreService, error) {
	return nil, nil
}

func (s *CoreService) UserStore() *UserStore {
	return s.userStore
}

func (s *CoreService) AuthStore() *AuthStore {
	return s.authStore
}

func (s *CoreService) UserSettingsStore() *UserSettingsStore {
    return s.userSettingsStore
}
