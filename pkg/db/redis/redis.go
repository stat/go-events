package redis

import (
	"context"
	"events/pkg/env"
	"fmt"

	"github.com/jinzhu/copier"
	v9 "github.com/redis/go-redis/v9"
)

type Client struct {
	*v9.Client

	Options *Options
}

type Options struct {
	RedisDB   int
	RedisHost string
	RedisPort string
}

func NewWithEnv(vars *env.Vars) (*Client, error) {
	options := &Options{}
	copier.Copy(options, vars)

	// create client

	client, err := NewWithOptions(options)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewWithOptions(options *Options) (*Client, error) {
	address := fmt.Sprintf("%s:%s", options.RedisHost, options.RedisPort)

	client := v9.NewClient(&v9.Options{
		Addr: address,
		DB:   options.RedisDB,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	instance := &Client{
		Client:  client,
		Options: options,
	}

	return instance, nil
}
