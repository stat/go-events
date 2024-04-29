package backends

import (
	"context"
	"encoding/json"
	"fmt"

	"events/pkg/db/redis"
	"events/pkg/env"
	"events/pkg/model"
	"events/pkg/utils"
)

type Redis[V model.Implementer] struct {
	*redis.Client
}

const (
	RedisDLQSuffix = "-dlq"
	RedisDLQIndex  = "dlq-index"
)

func (backend *Redis[V]) Initialize(vars *env.Vars) error {
	args := *vars
	args.RedisDB = vars.RedisDBEvents

	// concrete := *Redis[V]{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return err
	}

	backend.Client = client

	return nil
}

func (backend *Redis[V]) Append(key string, v *V) error {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return err
	}

	cmd := backend.LPush(context.Background(), key, marshalled)

	return cmd.Err()
}

func (backend *Redis[V]) AppendDLQ(key string, v *V) error {
	// write to DLQ

	dlqKey := fmt.Sprintf("%s%s", key, RedisDLQSuffix)
	if err := backend.Append(dlqKey, v); err != nil {
		return err
	}

	// add DLQ key to cassandra write buffer

	cmd := backend.SAdd(context.Background(), RedisDLQIndex, dlqKey)
	if err := cmd.Err(); err != nil {
		return err
	}

	// success

	return nil
}

func (backend *Redis[V]) Del(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Redis[V]) DelAtIndex(key string, index int64) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Redis[V]) DelHead(key string) error {
	cmd := backend.LPop(context.Background(), key)

	return cmd.Err()
}

func (backend *Redis[V]) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Redis[V]) Get(key string) ([]*V, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend *Redis[V]) GetAtIndex(key string, index int64) (*V, error) {
	cmd := backend.LIndex(context.Background(), key, index)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	// unmarshal

	var event V

	err := json.Unmarshal([]byte(cmd.Val()), &event)

	// delete event if mangled

	if err != nil {
		err = backend.DelHead(key)
	}

	// success

	return &event, err
}

func (backend *Redis[V]) GetHead(key string) (*V, error) {
	return backend.GetAtIndex(key, 0)
}

func (backend *Redis[V]) GetTail(key string) (*V, error) {
	return backend.GetAtIndex(key, -1)
}

// func Set(key, []*models.LocationEvent) error {
//   // TODO: implement or remove me
//   return nil
// }
