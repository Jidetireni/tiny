package cassandra

import (
	"github.com/Jidetireni/tiny/config"
	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type Cassandra struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

func New(config *config.CassandraConfig) (*Cassandra, error) {
	cluster := gocql.NewCluster(config.Host)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = config.KeySpace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &Cassandra{
		Cluster: cluster,
		Session: session,
	}, nil
}
