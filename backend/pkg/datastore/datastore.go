package datastore

import hansip "github.com/asasmoyo/pq-hansip"

type DataStore struct {
	DB        *hansip.Cluster
	userStore *UserStore
}

func (d *DataStore) Init() {
	d.userStore = &UserStore{
		db: d.DB,
	}
}

func (d *DataStore) UserStore() *UserStore {
	return d.userStore
}
