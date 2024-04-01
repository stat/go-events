package events

import "grid/pkg/models"

func Empty(aircraftID string) error {
	delete(Index, aircraftID)
	Index[aircraftID] = []*models.ADSB{}
	return nil
}
