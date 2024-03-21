package events

//
// Consts
//

const (
	DefaultReconsileEvents     = true
	DefaultReconsilationWindow = 2
)

type AppendOptions struct {
	ReconsileEvents     bool
	ReconsilationWindow int
}

func AppendOptionsDefaults() *AppendOptions {
	return &AppendOptions{
		ReconsileEvents:     DefaultReconsileEvents,
		ReconsilationWindow: DefaultReconsilationWindow,
	}
}

func AppendEvent(payload *Event) error {
	return AppendEventWithOptions(payload, AppendOptionsDefaults())
}

func AppendEventWithOptions(payload *Event, options *AppendOptions) error {
	aircraftID := payload.AircraftID

	// sanity check

	if aircraftID == "" {
		return AppendEventAircraftIDEmptyError
	}

	if payload.StationID == "" {
		return AppendEventStationIDEmptyError
	}

	if payload.Timestamp == nil {
		return AppendEventTimestampEmptyError
	}

	// append

	data := append(Index[payload.AircraftID], payload)

	// check to see if we need to reconsile

	if options.ReconsileEvents == true && len(data) > options.ReconsilationWindow {
		events, err := ReconsileEvents(aircraftID)

		if err != nil {
			return err
		}

		// flush

		data = []*Event{}

		err = storeEvents(events)

		if err != nil {
			return err
		}

	}

	// ensure index update

	Index[payload.AircraftID] = data

	// success

	return nil
}
