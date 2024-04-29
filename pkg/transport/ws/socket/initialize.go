package socket

import (
	"fmt"
	"net/http"

	"events/pkg/env"
	"events/pkg/model"
	"events/pkg/transport/ws/socket/pump"

	"github.com/olahol/melody"
)

// var (
//   instance     *melody.Melody
//   instancePump *pump.Pump
// )

type Server[CacheModel model.Implementer] struct {
	*melody.Melody

	Pump *pump.Pump[CacheModel]
}

func (implementation *Server[V]) InitializeFn(options *Options[V]) (func(*env.Vars) error, error) {
	cache := options.Cache

	fn := func(vars *env.Vars) error {
		//   return nil, nil
		// }

		// func Initialize(vars *env.Vars) error {
		m := melody.New()

		http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
			m.HandleRequest(w, r)
		})

		// run

		go func() {
			http.ListenAndServe(fmt.Sprintf(":%s", vars.WebSocketPort), nil)
		}()

		// instance = m
		implementation.Melody = m

		// pump

		p, err := pump.New[V](m, cache)

		if err != nil {
			return err
		}

		err = p.Start()

		if err != nil {
			return err
		}

		p.Start()

		// instancePump = p
		implementation.Pump = p

		// success

		return nil
	}

	return fn, nil
}
