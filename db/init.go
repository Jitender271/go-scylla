package db

import (
	"github.com/gocql/gocql"
)

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency,
		keyspace,
		hosts,
	}
}

func GetConnection(connection *scyllaDBConnection) (ISessionx, error) {
	cluster := connection.createCluster()
	return connection.createSession(cluster)
}