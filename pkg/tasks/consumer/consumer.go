package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"grid/pkg/db/queue"
	"grid/pkg/models"
	"grid/pkg/repos/events"
	"grid/pkg/tasks"
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
		return err
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

var (
	local = map[string][]interface{}{}
	cache = map[string]interface{}{}
)

func addEvent(key string, v interface{}) error {
	cmd := tasks.AircraftDB.LPush(context.Background(), key, v)

	return cmd.Err()
}

func addEventLocal(key string, v interface{}) error {
	l, ok := local[key]

	if !ok {
		l = []interface{}{}
	}

	l = append(l, v)

	local[key] = l

	return nil
}

func getEventAtIndexLocal(key string, index int64) (*models.ADSB, error) {
	l, ok := local[key]

	if !ok {
		return nil, fmt.Errorf("hi")
	}

	if int64(len(local)) < index {
		return nil, fmt.Errorf("out of bounds")
	}

	data := l[index]

	bytes, ok := data.([]byte)

	if !ok {
		return nil, fmt.Errorf("cannot cast to []byte")
	}

	event := &models.ADSB{}

	// unmarshal

	err := json.Unmarshal(bytes, event)

	// delete event if mangled

	if err != nil {
		err = delLatestEventLocal(key)
	}

	return event, err
}

func getHeadEventLocal(key string) (*models.ADSB, error) {
	return getEventAtIndexLocal(key, 0)
}

func getTailEventLocal(key string) (*models.ADSB, error) {
	return getEventAtIndexLocal(key, 0)
}

func delLatestEventLocal(key string) error {
	l, ok := local[key]

	if !ok {
		return nil
	}

	l = l[:len(l)-1]

	local[key] = l

	return nil
}

func updateCacheLocal(key string, v interface{}) error {
	cache[key] = v
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
