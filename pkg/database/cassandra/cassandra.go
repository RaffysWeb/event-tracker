package cassandra

import (
	"sync"

	"github.com/gocql/gocql"
)

var session *gocql.Session
var once sync.Once

// Init initializes the Cassandra session once as a singleton.
func Init() {
	once.Do(func() {
		cluster := gocql.NewCluster("localhost:9042") // need to update to connect to docker cassandra:9042
		cluster.Keyspace = "mykeyspace"
		var err error
		session, err = cluster.CreateSession()
		if err != nil {
			panic(err)
		}
	})
}

// Close closes the Cassandra session.
func Close() {
	session.Close()
}

// GetSession returns the Cassandra session.
func GetSession() *gocql.Session {
	return session
}
