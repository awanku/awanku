package hansip

import (
	"errors"
	"time"
)

var (
	ErrNoReplicaAvailable    = errors.New("no replica connection available")
	ErrNoPrimaryAvailable    = errors.New("no primary connection available")
	ErrClusterNotInitialized = errors.New("cluster was not initialized")
)

type Config struct {
	PrependQueryWithCaller bool
	MaxConnAttempt         int
	ConnRetryDelay         time.Duration
	ConnCheckDelay         time.Duration
	ConnPingTimeout        time.Duration
}

type Cluster struct {
	Primary  SQL
	Replicas []SQL

	PingTimeout    time.Duration
	ConnCheckDelay time.Duration

	manager     *connectionManager
	initialized bool
}

func (c *Cluster) Init() {
	if c.initialized {
		return
	}
	c.initialized = true

	primary := newConnection(c.Primary, c.PingTimeout, c.ConnCheckDelay)
	replicas := make([]*connection, 0)
	for _, replica := range c.Replicas {
		replicas = append(replicas, newConnection(replica, c.PingTimeout, c.ConnCheckDelay))
	}

	c.manager = newConnectionManager(
		primary,
		replicas,
		c.ConnCheckDelay/2,
	)
	go c.manager.updateActiveReplicas()
	go c.manager.loop()
}

func (c *Cluster) Query(dest interface{}, query string, args ...interface{}) error {
	if !c.initialized {
		return ErrClusterNotInitialized
	}

	conn := c.manager.getReplica()
	if conn == nil {
		return ErrNoReplicaAvailable
	}
	return conn.Query(dest, query, args...)
}

func (c *Cluster) WriterExec(query string, args ...interface{}) error {
	if !c.initialized {
		return ErrClusterNotInitialized
	}

	conn := c.manager.getPrimary()
	if conn == nil {
		return ErrNoPrimaryAvailable
	}
	return conn.Exec(query, args...)
}

func (c *Cluster) WriterQuery(dest interface{}, query string, args ...interface{}) error {
	if !c.initialized {
		return ErrClusterNotInitialized
	}

	conn := c.manager.getPrimary()
	if conn == nil {
		return ErrNoPrimaryAvailable
	}
	return conn.Query(dest, query, args...)
}

func (c *Cluster) NewTransaction() (Transaction, error) {
	if !c.initialized {
		return nil, ErrClusterNotInitialized
	}

	conn := c.manager.getPrimary()
	if conn == nil {
		return nil, ErrNoPrimaryAvailable
	}
	return conn.NewTransaction()
}

func (c *Cluster) Shutdown() {
	c.manager.quit()
}
