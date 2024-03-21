package events

import (
	"errors"
)

var (
	// errors
	AppendAircraftIDEmptyError = errors.New("must include aircraft id!")
	AppendStationIDEmptyError  = errors.New("must include station id!")
	AppendTimestampEmptyError  = errors.New("must include timestamp!")
)
