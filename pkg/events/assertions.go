package events

import (
	"grid/pkg/models"
	"time"
)

func isEventEqual(e1, e2 *models.ADSB) bool {
	// sanity check

	if e1 == nil || e2 == nil {
		return false
	}

	// compute
	// TODO: implmenet additional comp algos

	result :=
		isTimeEqual(e1.Timestamp, e2.Timestamp)

	// success

	return result
}

func isTimeEqual(t1, t2 *time.Time) bool {
	// sanity check

	if t1 == nil || t2 == nil {
		return false
	}

	// compute
	// TODO: implmenet theshold algo

	result := t1.Sub(*t2) == 0

	// success

	return result
}
