package socket

import (
	"grid/pkg/model"
	"grid/pkg/repos/cache"
)

type Options[V model.Implementer] struct {
	Cache *cache.Repo[V]
}
