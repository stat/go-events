package socket

import (
	"fmt"
	"net/http"

	"grid/pkg/env"
	"grid/pkg/repos/cache"
	"grid/pkg/transport/ws/socket/pump"

	"github.com/olahol/melody"
)

var (
	instance     *melody.Melody
	instancePump *pump.Pump
)

func Initialize(vars *env.Vars) error {
	m := melody.New()

	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	// run

	go func() {
		http.ListenAndServe(fmt.Sprintf(":%s", vars.WebSocketPort), nil)
	}()

	instance = m

	// pump

	p, err := pump.New(m, cache.Backend)

	if err != nil {
		return err
	}

	err = p.Start()

	if err != nil {
		return err
	}

	p.Start()

	instancePump = p

	// success

	return nil
}
