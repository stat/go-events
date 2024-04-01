package events

import "grid/pkg/models"

type TapOptions interface{}

func Tap(options ...*TapOptions) ([]*models.ADSB, error) {
	return nil, nil
}
