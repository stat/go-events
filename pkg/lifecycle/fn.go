package lifecycle

type Fn[T any] func(*T) error
