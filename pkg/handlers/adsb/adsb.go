package adsb

import (
	"grid/pkg/models"
)

//
// Consts
//

// const (
//   DefaultAppendReconcileEvents     = true
//   DefaultAppendReconsilationWindow = 2
// )

// type AppendOptions struct {
//   ReconcileEvents     bool
//   ReconsilationWindow int
// }

// func AppendOptionsDefaults() *AppendOptions {
//   return &AppendOptions{
//     ReconcileEvents:     DefaultAppendReconcileEvents,
//     ReconsilationWindow: DefaultAppendReconsilationWindow,
//   }
// }

// func Append(payload *LocationEvent) error {
//   return AppendWithOptions(payload, AppendOptionsDefaults())
// }

// func AppendWithOptions(payload *LocationEvent, options *AppendOptions) error {
func Process(event *models.ADSB) error {

	// validate

	if err := Validate(event); err != nil {
		return err
	}

	// append

	data := append(Index[event.AircraftID], event)

	// check to see if we need to reconcile

	// if options.ReconcileEvents == true && len(data) > options.ReconsilationWindow {
	//   events, err := Reconcile(aircraftID)

	//   if err != nil {
	//     return err
	//   }

	//   // flush

	//   data = []*LocationEvent{}

	//   err = Sink(events)

	//   if err != nil {
	//     return err
	//   }

	// }

	// ensure index update

	Index[event.AircraftID] = data

	// success

	return nil
}

func Validate(event *models.ADSB) error {

	return nil
}
