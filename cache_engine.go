package cache_wrapper

import (
	"context"
)

//go:generate mockery --name=KVCacheI --outpkg=mocks
type KVCacheI interface {
	Get(ctx context.Context, key []byte) ([]byte, error)
	Set(ctx context.Context, key []byte, value []byte, expired int64) error
	IsNotFoundErr(err error) bool
}
