package pump

import (
	"encoding/json"
	"time"

	"grid/pkg/model"
	"grid/pkg/repos/cache/provider"

	"github.com/olahol/melody"
)

type Pump[V model.Implementer] struct {
	Cache  provider.Implementer[V]
	Socket *melody.Melody
	Ticker *time.Ticker

	Done chan bool

	Options *PumpOptions
}

type PumpOptions struct {
	Interval time.Duration
}

const (
	// TODO: implement with env vars
	DefaultInterval = 1 * time.Second
)

var (
	DefaultOptions = &PumpOptions{
		Interval: DefaultInterval,
	}
)

func New[V model.Implementer](socket *melody.Melody, cache provider.Implementer[V]) (*Pump[V], error) {
	return NewWithOptions(socket, cache, DefaultOptions)
}

func NewWithOptions[V model.Implementer](socket *melody.Melody, cache provider.Implementer[V], options *PumpOptions) (*Pump[V], error) {
	ticker := time.NewTicker(options.Interval)

	pump := &Pump[V]{
		Cache:  cache,
		Socket: socket,
		Ticker: ticker,
		Done:   make(chan bool),
	}

	return pump, nil
}

func (publisher *Pump[V]) Publish() error {
	data, err := publisher.Cache.GetAircraftsLocations()

	if err != nil {
		return err
	}

	bytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = publisher.Socket.Broadcast(bytes)

	if err != nil {
		return err
	}

	return nil
}

func (publisher *Pump[V]) Start() error {
	go func() {
		for {
			select {
			case <-publisher.Done:
				return
			case _ = <-publisher.Ticker.C:
				_ = publisher.Publish()
			}
		}
	}()

	return nil
}

func (publisher *Pump[V]) Stop() error {
	publisher.Done <- true
	return nil
}
