package events

import (
	"grid/pkg/utils"
)

func Clear() error {
	Index = utils.Must(New())
	return nil
}
