package datastore

import hansip "github.com/asasmoyo/pq-hansip"

type DataStore struct {
	DB        *hansip.Cluster
	userStore *UserStore
	authStore *AuthStore
}

func (d *DataStore) Init() {
	d.userStore = &UserStore{
		db: d.DB,
	}
	d.authStore = &AuthStore{
		db: *d.DB,
	}
}

func (d *DataStore) UserStore() *UserStore {
	return d.userStore
}

func (d *DataStore) AuthStore() *AuthStore {
	return d.authStore
}
