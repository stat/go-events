package workers

import (
	"fmt"

	"grid/pkg/env"
	"grid/pkg/model"
	"grid/pkg/repos/cache"
	"grid/pkg/repos/events"
	"grid/pkg/tasks/consumer"
	"grid/pkg/tasks/producer"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
)

// var (
//   instance *Server
// )

type Server[CacheModel, EventModel model.Implementer] struct {
	*asynq.Server

	Mux     *asynq.ServeMux
	Options *Options[CacheModel, EventModel]
}

type Options[CacheModel, EventModel model.Implementer] struct {
	ASYNQConcurrency int
	RedisDB          int `copier:"RedisBDAsynq"`
	RedisHost        string
	RedisPort        string

	Cache  *cache.Repo[CacheModel]
	Events *events.Repo[EventModel]
}

// var (
//   Instance *Server[config.CacheModel, config.EventModel]
// )

// func InitializeFn[CacheModel, EventModel model.Implementer](vars *env.Vars) error {
func (implementation *Server[CacheModel, EventModel]) InitializeFn(options *Options[CacheModel, EventModel]) (func(*env.Vars) error, error) {
	// options := &Options{}

	fn := func(vars *env.Vars) error {
		err := copier.Copy(options, vars)

		if err != nil {
			return err
		}

		address := fmt.Sprintf("%s:%s", vars.RedisHost, vars.RedisPort)

		// server

		server := asynq.NewServer(
			asynq.RedisClientOpt{
				Addr: address,
				DB:   options.RedisDB,
			},
			asynq.Config{
				Concurrency: vars.ASYNQConcurrency,
			},
		)

		implementation.Server = server

		// mux

		mux := asynq.NewServeMux()
		mux.Handle(consumer.Type, &consumer.Payload{})
		mux.Handle(producer.Type, &producer.Payload{})

		implementation.Mux = mux

		// err = server.Run(mux)

		// TODO: handle this error

		go func() {
			err := server.Run(mux)

			if err != nil {
				panic(err)
			}
		}()

		return nil
	}

	return fn, nil
}
