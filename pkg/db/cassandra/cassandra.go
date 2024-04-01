package cassandra

import (
	"grid/pkg/env"
)

type Client struct {
	Options *Options
}

type Options struct {
	CassandraKeyspace string
	Cassandra
}

func Initialize(vars *env.Vars) error {
	return nil
}
