package events

func Empty(aircraftID string) error {
	delete(Index, aircraftID)
	Index[aircraftID] = []*Event{}
	return nil
}
