package events_backends

import (
	"context"
	"encoding/json"
	"fmt"

	"grid/pkg/db/redis"
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
	"grid/pkg/utils"
)

type Redis struct {
	*redis.Client
}

const (
	RedisDLQSuffix = "-dlq"
	RedisDLQIndex  = "dlq-index"
)

func (backend Redis) Initialize(vars *env.Vars) (provider.Provider, error) {
	args := *vars
	args.RedisDB = vars.RedisDBEvents

	concrete := Redis{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return Redis{}, err
	}

	concrete.Client = client

	return concrete, nil
}

func (backend Redis) Append(key string, v *models.LocationEvent) error {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return err
	}

	cmd := backend.LPush(context.Background(), key, marshalled)

	return cmd.Err()
}

func (backend Redis) AppendDLQ(key string, v *models.LocationEvent) error {
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

func (backend Redis) Del(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend Redis) DelAtIndex(key string, index int64) error {
	// TODO: implement or remove me
	return nil
}

func (backend Redis) DelHead(key string) error {
	cmd := backend.LPop(context.Background(), key)

	return cmd.Err()
}

func (backend Redis) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend Redis) Get(key string) ([]*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Redis) GetAtIndex(key string, index int64) (*models.LocationEvent, error) {
	cmd := backend.LIndex(context.Background(), key, index)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	// unmarshal

	event := &models.LocationEvent{}
	err := json.Unmarshal([]byte(cmd.Val()), event)

	// delete event if mangled

	if err != nil {
		err = backend.DelHead(key)
	}

	// success

	return event, err
}

func (backend Redis) GetHead(key string) (*models.LocationEvent, error) {
	return backend.GetAtIndex(key, 0)
}

func (backend Redis) GetTail(key string) (*models.LocationEvent, error) {
	return backend.GetAtIndex(key, -1)
}

// func Set(key, []*models.LocationEvent) error {
//   // TODO: implement or remove me
//   return nil
// }
