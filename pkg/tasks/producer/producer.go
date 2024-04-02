package producer

import (
	"context"
	"encoding/json"
	"grid/pkg/models"
	"grid/pkg/repos/cache"
	"grid/pkg/repos/events"

	"github.com/hibiken/asynq"
)

type Payload struct {
	*models.LocationEvent
}

const (
	Type = "producer"
)

func Process(event *models.LocationEvent) error {
	// key

	key := event.AircraftID

	// marshal

	// eventJSON, err := json.Marshal(event)

	// if err != nil {
	//   return err
	// }

	// add event to airplaneID

	err := events.Backend.Append(key, event)

	if err != nil {
		return err
	}

	// update cache

	err = cache.Backend.UpsertAircraftLocation(key, event)

	if err != nil {
		return err
	}

	return nil
}

func (processor *Payload) ProcessTask(ctx context.Context, t *asynq.Task) error {
	payload := &Payload{}

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	if err := Process(payload.LocationEvent); err != nil {
		return err
	}

	return nil
}

func NewTask(event *models.LocationEvent) (*asynq.Task, error) {
	payload, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(Type, payload)

	return task, nil
}
