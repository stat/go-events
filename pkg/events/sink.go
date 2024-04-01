package events

import "grid/pkg/models"

type SinkOptions interface{}

func Sink(events []*models.ADSB, options ...*SinkOptions) error {
	return nil
}
