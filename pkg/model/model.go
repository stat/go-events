package model

type Implementer interface {
	Key() (string, error)
}
