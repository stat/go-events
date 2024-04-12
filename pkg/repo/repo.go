package repo

// func InitializeFn[V model.Implementer](options *Options[V]) (func(vars *env.Vars) error, error) {
//   if options == nil {
//     options = &Options[V]{}
//   }

//   if options.Backend == nil {
//     options.Backend = &backends.Local[V]{}
//   }

//   backend := options.Backend

//   fn := func(vars *env.Vars) error {
//     return backend.Initialize(vars)
//   }

//   implementation.Backend = backend

//   return fn, nil
// }
