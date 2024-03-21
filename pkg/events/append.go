package events

//
// Consts
//

const (
	DefaultAppendReconcileEvents     = true
	DefaultAppendReconsilationWindow = 2
)

type AppendOptions struct {
	ReconcileEvents     bool
	ReconsilationWindow int
}

func AppendOptionsDefaults() *AppendOptions {
	return &AppendOptions{
		ReconcileEvents:     DefaultAppendReconcileEvents,
		ReconsilationWindow: DefaultAppendReconsilationWindow,
	}
}

func Append(payload *Event) error {
	return AppendWithOptions(payload, AppendOptionsDefaults())
}

func AppendWithOptions(payload *Event, options *AppendOptions) error {
	aircraftID := payload.AircraftID

	// sanity check

	if aircraftID == "" {
		return AppendAircraftIDEmptyError
	}

	if payload.StationID == "" {
		return AppendStationIDEmptyError
	}

	if payload.Timestamp == nil {
		return AppendTimestampEmptyError
	}

	// append

	data := append(Index[payload.AircraftID], payload)

	// check to see if we need to reconcile

	if options.ReconcileEvents == true && len(data) > options.ReconsilationWindow {
		events, err := Reconcile(aircraftID)

		if err != nil {
			return err
		}

		// flush

		data = []*Event{}

		err = Sink(events)

		if err != nil {
			return err
		}

	}

	// ensure index update

	Index[payload.AircraftID] = data

	// success

	return nil
}
