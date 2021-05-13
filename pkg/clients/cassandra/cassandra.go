package cassandra

import (
	"log"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "auth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		log.Fatalf("Error occurred connecting to the database: %s", err.Error())
	}
}

func GetSession() *gocql.Session {
	return session
}
