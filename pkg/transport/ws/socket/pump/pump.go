package pump

import (
	"encoding/json"
	"time"

	"grid/pkg/repos/cache/provider"

	"github.com/olahol/melody"
)

type Pump struct {
	Cache  provider.Provider
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

func New(socket *melody.Melody, cache provider.Provider) (*Pump, error) {
	return NewWithOptions(socket, cache, DefaultOptions)
}

func NewWithOptions(socket *melody.Melody, cache provider.Provider, options *PumpOptions) (*Pump, error) {
	ticker := time.NewTicker(options.Interval)

	pump := &Pump{
		Cache:  cache,
		Socket: socket,
		Ticker: ticker,
		Done:   make(chan bool),
	}

	return pump, nil
}

func (publisher *Pump) Publish() error {
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

func (publisher *Pump) Start() error {
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

func (publisher *Pump) Stop() error {
	publisher.Done <- true
	return nil
}
