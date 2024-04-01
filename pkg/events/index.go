package events

import "grid/pkg/models"

var (
	// event buffer indexed by aircraftID
	Index map[string][]*models.ADSB = make(map[string][]*models.ADSB)
)
