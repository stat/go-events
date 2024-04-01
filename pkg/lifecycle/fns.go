package lifecycle

type Fns[T any] []Fn[T]

func (fns Fns[T]) Execute(t *T) error {
	for _, fn := range fns {
		if err := fn(t); err != nil {
			return err
		}
	}

	return nil
}
