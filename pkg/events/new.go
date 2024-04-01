package events

import "grid/pkg/models"

func New() (map[string][]*models.ADSB, error) {
	return make(map[string][]*models.ADSB), nil
}
