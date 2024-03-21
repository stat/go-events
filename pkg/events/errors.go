package events

import (
	"errors"
)

var (
	// errors
	AppendEventAircraftIDEmptyError = errors.New("must include aircraft id!")
	AppendEventStationIDEmptyError  = errors.New("must include station id!")
	AppendEventTimestampEmptyError  = errors.New("must include timestamp!")
)
