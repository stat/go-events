package cassandra

import (
	"events/pkg/env"
)

type Client struct {
	Options *Options
}

type Options struct {
	CassandraKeyspace string
}

func Initialize(vars *env.Vars) error {
	return nil
}

//
