package environments

import (
	"os"

	"github.com/gocql/gocql"
)

func initializeDevEnv() {
	os.Setenv("SCYLLA_CONSISTENCY", gocql.Quorum.String())
	os.Setenv("SCYLLA_HOSTS", "127.0.0.1:9042")
	os.Setenv("HOST", "0.0.0.0")
	os.Setenv("PORT", "4444")
}

func init() {
	ENV := os.Getenv("GO_ENV")

	switch ENV {
	case "development":
		initializeDevEnv()
	default:
		initializeDevEnv()
	}
}
