package events_backends

import (
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
)

type Local struct {
	Data map[string][]interface{}
}

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{}

	return concrete, nil
}

func (backend Local) Append(key string, v interface{}) error {
	l, ok := backend.Data[key]

	if !ok {
		l = []interface{}{}
	}

	l = append(l, v)

	backend.Data[key] = l

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
	l, ok := backend.Data[key]

	if !ok {
		return nil
	}

	l = l[:len(l)-1]

	backend.Data[key] = l

	return nil
}

func (backend Local) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}
func (backend Local) Get(key string) (*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Local) GetAtIndex(key string, index int64) (*models.LocationEvent, error) {
	return nil, nil
}

func (backend Local) GetHead(key string) (*models.LocationEvent, error) {
	return nil, nil
}

func (backend Local) GetTail(key string) (*models.LocationEvent, error) {
	return nil, nil
}

// var (
//   local = map[string][]interface{}{}
//   cache = map[string]interface{}{}
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
