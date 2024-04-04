package events_backends

import (
	"errors"
	"fmt"
	"sync"

	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
	//"github.com/hashicorp/go-set/v2"
)

type Local struct {
	// Data map[string][]*models.LocationEvent
	Data *sync.Map
}

const (
	LocalDLQSuffix = "-dlq"
)

var (
	LocalEventBackendKeyNotFound      = errors.New("Local event backend could not find key")
	LocalEventBackendIndexOutOfBounds = errors.New("Local event backend index out of bounds")
)

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{
		// Data: map[string][]*models.LocationEvent{},
		Data: &sync.Map{},
	}

	backend.Data = &sync.Map{}

	return concrete, nil
}

func (backend Local) Append(key string, v *models.LocationEvent) error {

	// l, ok := backend.Data[key]
	// iface, ok := backend.Data.Load(key)

	// if ok {
	//   l, ok = iface.([]*models.LocationEvent)
	// }

	// if !ok {
	//   l = []*models.LocationEvent{}
	// }

	l, err := backend.Get(key)

	if err != nil {
		l = []*models.LocationEvent{}
	}

	l = append(l, v)

	// backend.Data[key] = l
	backend.Data.Store(key, l)

	return nil
}

func (backend Local) AppendDLQ(key string, v *models.LocationEvent) error {
	dlqKey := fmt.Sprintf("%s%s", key, RedisDLQSuffix)
	if err := backend.Append(dlqKey, v); err != nil {
		return err
	}

	return nil
}

func (backend Local) Del(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend Local) DelAtIndex(key string, index int64) error {
	// TODO: implement or remove me
	return nil
}

func (backend Local) DelHead(key string) error {
	// l, ok := backend.Data[key]

	// if !ok {
	//   return nil
	// }

	l, err := backend.Get(key)

	if err != nil {
		return err
	}

	l = l[:len(l)-1]

	// backend.Data[key] = l
	backend.Data.Store(key, l)

	return nil
}

func (backend Local) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}
func (backend Local) Get(key string) ([]*models.LocationEvent, error) {
	iface, ok := backend.Data.Load(key)

	if !ok {
		return nil, LocalEventBackendKeyNotFound
	}

	l, ok := iface.([]*models.LocationEvent)

	if !ok {
		return nil, fmt.Errorf("event local get cast error")
	}

	return l, nil
}

func (backend Local) GetAtIndex(key string, index int64) (*models.LocationEvent, error) {
	// l, ok := backend.Data[key]

	// if !ok {
	//   return nil, LocalEventBackendKeyNotFound
	// }

	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	if int64(len(l)) < index {
		return nil, LocalEventBackendIndexOutOfBounds
	}

	return l[index], nil
}

func (backend Local) GetHead(key string) (*models.LocationEvent, error) {
	// l, ok := backend.Data[key]

	// if !ok {
	//   return nil, LocalEventBackendKeyNotFound
	// }

	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	return backend.GetAtIndex(key, int64(len(l))-1)
}

func (backend Local) GetTail(key string) (*models.LocationEvent, error) {
	return backend.GetAtIndex(key, 0)
}

// func (backend Local) Set(key string, l []*models.LocationEvent) error {
//   backend.Data.Store(k, l)
//   return nil
// }

// var (
//   local = map[string][]interface{}{}
// )

// func addEvent(key string, v interface{}) error {
//   cmd := tasks.AircraftDB.LPush(context.Background(), key, v)

//   return cmd.Err()
// }

// func addEventLocal(key string, v interface{}) error {
//   l, ok := local[key]

//   if !ok {
//     l = []interface{}{}
//   }

//   l = append(l, v)

//   local[key] = l

//   return nil
// }

// func getEventAtIndexLocal(key string, index int64) (*models.LocationEvent, error) {
//   l, ok := local[key]

//   if !ok {
//     return nil, fmt.Errorf("hi")
//   }

//   if int64(len(local)) < index {
//     return nil, fmt.Errorf("out of bounds")
//   }

//   data := l[index]

//   bytes, ok := data.([]byte)

//   if !ok {
//     return nil, fmt.Errorf("cannot cast to []byte")
//   }

//   event := &models.LocationEvent{}

//   // unmarshal

//   err := json.Unmarshal(bytes, event)

//   // delete event if mangled

//   if err != nil {
//     err = delLatestEventLocal(key)
//   }

//   return event, err
// }

// func getHeadEventLocal(key string) (*models.LocationEvent, error) {
//   return getEventAtIndexLocal(key, 0)
// }

// func getTailEventLocal(key string) (*models.LocationEvent, error) {
//   return getEventAtIndexLocal(key, 0)
// }

// func delLatestEventLocal(key string) error {
//   l, ok := local[key]

//   if !ok {
//     return nil
//   }

//   l = l[:len(l)-1]

//   local[key] = l

//   return nil
// }

// func updateCacheLocal(key string, v interface{}) error {
//   cache[key] = v
//   return nil
// }
