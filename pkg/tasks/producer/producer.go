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
	*models.ADSB
}

const (
	Type = "producer"
)

func Process(event *models.ADSB) error {
	// key

	key := event.AircraftID

	// marshal

	eventJSON, err := json.Marshal(event)

	if err != nil {
		return err
	}

	// add event to airplaneID

	err = events.Backend.Append(key, eventJSON)

	if err != nil {
		return err
	}

	// update cache

	err = cache.Backend.UpsertAircraftLocation(key, eventJSON)

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

	if err := Process(payload.ADSB); err != nil {
		return err
	}

	return nil
}

func NewTask(event *models.ADSB) (*asynq.Task, error) {
	payload, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(Type, payload)

	return task, nil
}