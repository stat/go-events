package events_backends

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

// func getEventAtIndexLocal(key string, index int64) (*models.ADSB, error) {
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

//   event := &models.ADSB{}

//   // unmarshal

//   err := json.Unmarshal(bytes, event)

//   // delete event if mangled

//   if err != nil {
//     err = delLatestEventLocal(key)
//   }

//   return event, err
// }

// func getHeadEventLocal(key string) (*models.ADSB, error) {
//   return getEventAtIndexLocal(key, 0)
// }

// func getTailEventLocal(key string) (*models.ADSB, error) {
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
