package intializedb

import (
	"github.com/go-scylla/infrastructure/db/scylladb"
	"github.com/gocql/gocql"
)

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency,
		keyspace,
		hosts,
	}
}

func GetConnection(connection *scyllaDBConnection) (scylladb.ISessionx, error) {
	cluster := connection.createCluster()
	return connection.createSession(cluster)
}