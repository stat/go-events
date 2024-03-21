package events

func Empty(aircraftID string) error {
	delete(Index, aircraftID)
	return nil
}
