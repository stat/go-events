package workers

import (
	"fmt"
	"grid/pkg/env"
	"grid/pkg/tasks/consumer"
	"grid/pkg/tasks/producer"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
)

var (
	instance *Server
)

type Server struct {
	*asynq.Server

	Options *Options
}

type Options struct {
	ASYNQConcurrency int
	RedisDB          int `copier:"RedisBDAsynq"`
	RedisHost        string
	RedisPort        string
}

func Initialize(vars *env.Vars) error {
	options := &Options{}

	err := copier.Copy(options, vars)

	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%s", vars.RedisHost, vars.RedisPort)

	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: address,
			DB:   options.RedisDB,
		},
		asynq.Config{
			Concurrency: vars.ASYNQConcurrency,
		},
	)

	mux := asynq.NewServeMux()
	mux.Handle(consumer.Type, &consumer.Payload{})
	mux.Handle(producer.Type, &producer.Payload{})

	// err = server.Run(mux)

	go func() {
		err := server.Run(mux)

		if err != nil {
			panic(err)
		}
	}()

	return nil
}
