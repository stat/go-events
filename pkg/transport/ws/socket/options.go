package socket

import (
	"events/pkg/model"
	"events/pkg/repos/cache"
)

type Options[V model.Implementer] struct {
	Cache *cache.Repo[V]
}
