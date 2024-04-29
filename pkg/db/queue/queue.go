package queue

import (
	"errors"
	"events/pkg/env"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
)

type Client struct {
	*asynq.Client

	Options *Options
}

type Options struct {
	RedisDB   int `copier:"RedisBDAsynq"`
	RedisHost string
	RedisPort string
}

var (
	instance *Client

	QueueNotInitializedError = errors.New("HTTP queue has not been initialized")
)

func Initialize(vars *env.Vars) error {
	options := &Options{}

	err := copier.Copy(options, vars)

	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%s", vars.RedisHost, vars.RedisPort)

	clientDriver := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr: address,
			DB:   options.RedisDB,
		},
	)

	instance = &Client{
		Client:  clientDriver,
		Options: options,
	}

	return nil
}

func Instance() (*Client, error) {
	if instance == nil {
		return nil, QueueNotInitializedError
	}

	return instance, nil
}

func Terminate() error {
	return nil
}
