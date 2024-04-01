package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"grid/pkg/db/queue"
	"grid/pkg/models"
	"grid/pkg/repos/events"
	"grid/pkg/tasks/producer"

	"github.com/hibiken/asynq"
)

type Payload struct {
	*models.ADSB
}

const (
	Type = "consumer"
	// DLQ_SUFFIX = "-dlq"
)

var (
	ProcessEventTimestampBeforeError = errors.New("event timestamp occurs before the latest entry")
	ProcessEventTimestampAfterError  = errors.New("event timestamp occurs after the current time")
	ProcessEventTimestampEqualError  = errors.New("event timestamp is the same as the current entry")
)

func Process(event *models.ADSB) error {
	// TODO: func sanity check

	// sanity check

	err := event.Validate()

	if err != nil {
		return err
	}

	key := event.AircraftID

	// get last

	lastEvent, err := events.Backend.GetHead(key)

	// compare

	if err == nil {
		// check if the event timestamp is before the latest
		if event.Timestamp.Before(*lastEvent.Timestamp) {
			return ProcessEventTimestampBeforeError
		}

		// check if the event timestamp is equal to the latest
		if event.Timestamp.Equal(*lastEvent.Timestamp) {
			return ProcessEventTimestampEqualError
		}

		// TODO: implement lat/long comp with stddev
	}

	// check if the event timestamp is in the future
	if event.Timestamp.After(time.Now()) {
		return ProcessEventTimestampAfterError
	}

	// success

	return nil
}

func writeDLQ[T any](key string, v *T, err error) error {
	// event := &events.DLQ[T]{
	//   Error: err,
	//   Event: v,
	// }

	// eventJSON, err := json.Marshal(event)

	// if err != nil {
	//   return err
	// }

	// dlq := fmt.Sprintf("%s%s", key, DLQ_SUFFIX)

	// addEvent(dlq, eventJSON)

	return nil
}

func (processor *Payload) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// unmarshal

	payload := &Payload{}

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	event := payload.ADSB

	// process

	err := Process(event)

	if err != nil {
		// TODO: write to DLQ
		return nil
	}

	// coreograph

	err = processor.Choreograph(event)

	if err != nil {
		return err
	}

	// success

	return nil
}

func (processor *Payload) Choreograph(event *models.ADSB) error {
	// enqueue producer

	next, err := producer.NewTask(event)

	if err != nil {
		return err
	}

	queue, err := queue.Instance()

	if err != nil {
		return err
	}

	_, err = queue.Enqueue(next)

	if err != nil {
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
