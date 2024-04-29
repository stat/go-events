package consumer

import (
	"context"
	"encoding/json"
	"errors"

	// "time"

	"events/pkg/db/queue"
	"events/pkg/models"
	"events/pkg/tasks/producer"

	// "events/pkg/tasks/producer"

	"github.com/hibiken/asynq"
)

type Payload struct {
	*models.LocationEvent
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

func Process(event *models.LocationEvent) error {
	// // TODO: func sanity check

	// // sanity check

	// err := event.Validate()

	// if err != nil {
	//   return err
	// }

	// key := event.AircraftID

	// // get last

	// lastEvent, err := events.Backend.GetHead(key)

	// // compare

	// if err == nil {
	//   // check if the event timestamp is before the latest
	//   if event.Timestamp.Before(*lastEvent.Timestamp) {
	//     return ProcessEventTimestampBeforeError
	//   }

	//   // check if the event timestamp is equal to the latest
	//   if event.Timestamp.Equal(*lastEvent.Timestamp) {
	//     return ProcessEventTimestampEqualError
	//   }

	//   // TODO: implement lat/long comp with stddev
	// }

	// // check if the event timestamp is in the future
	// if event.Timestamp.After(time.Now()) {
	//   return ProcessEventTimestampAfterError
	// }

	// // success

	return nil
}

func (processor *Payload) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// unmarshal

	payload := &Payload{}

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	event := payload.LocationEvent

	// process

	err := Process(event)

	if err != nil {
		return nil
		// return events.Backend.AppendDLQ(event.AircraftID, event)
	}

	// coreograph

	err = processor.Choreograph(event)

	if err != nil {
		return err
	}

	// success

	return nil
}

func (processor *Payload) Choreograph(event *models.LocationEvent) error {
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

func NewTask(event *models.LocationEvent) (*asynq.Task, error) {
	payload, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(Type, payload)

	return task, nil
}
