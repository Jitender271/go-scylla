package intializedb

import (
	"fmt"
	"time"

	"github.com/go-scylla/infrastructure/db/scylladb"
	"github.com/gocql/gocql"
)


type scyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	hosts       []string
}

func (conn *scyllaDBConnection) createCluster() *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}

	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	return cluster
}

func (conn *scyllaDBConnection) createSession(cluster *gocql.ClusterConfig) (scylladb.ISessionx, error) {
	session, err := scylladb.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Println("An error occurred while creating DB session", err.Error())
		return nil, err
	}
	return session, nil
}



