package events

import (
	"fmt"
	"grid/pkg/models"
	"sort"
)

func Reconcile(aircraftID string) ([]*models.ADSB, error) {
	// get events by aircraftID

	data := Index[aircraftID]

	// ensure data

	if data == nil {
		return nil, fmt.Errorf("no data available for %s", aircraftID)
	}

	// sort by timestamp

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Timestamp.Before(*data[j].Timestamp)
	})

	// initialize reconciled events

	reconciledEvents := []*models.ADSB{}

	var event *models.ADSB = nil
	for _, item := range data {
		// determine if equal, then resume
		if isEventEqual(event, item) {
			continue
		}

		// append

		reconciledEvents = append(reconciledEvents, item)

		// rotate

		event = item
	}

	// success

	return reconciledEvents, nil
}
