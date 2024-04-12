package testing

import (
	"grid/app"
	"grid/pkg/env"
)

var (
	Vars *env.Vars
)

func init() {
	// vars

	vars, err := env.Load()

	if err != nil {
		panic(err)
	}

	// initialize

	err = app.Initialize(vars)

	if err != nil {
		panic(err)
	}

	Vars = vars

	// vars, err := env.LoadWithFile(".env-test", "env")

	// fmt.Println(vars)

	// if err != nil {
	//   panic(err)
	// }

	// err = app.InitializeWithVars(vars)

	// if err != nil {
	//   panic(err)
	// }
}
