package cache_wrapper

import (
	"context"

	"golang.org/x/sync/singleflight"
)

func NewCacheWrapper(cache KVCacheI, request RequestFormatter, expired int) *cacheWrapper {
	return &cacheWrapper{
		cacheEngine:      cache,
		RequestFormatter: request,
		expired:          expired,
		group:            &singleflight.Group{},
	}
}

func (r *cacheWrapper) SetHandle(fn func(...interface{}) []interface{}) {
	r.handler = fn
}

func (r *cacheWrapper) SetUniqueKey(uniqueKey string) {
	r.uniqKey = uniqueKey
}

func (r *cacheWrapper) Request(ctx context.Context,
	reqs []interface{}, // 原函数的所有参数
	resp []interface{}, // UnMarshalWrapper 需要resp的类型
) ([]interface{}, error) {
	if r.handler == nil {
		panic("handler not set yet")
	}
	key := r.GetUniqKey(r.uniqKey, reqs...)
	var fn = func() (interface{}, error) {
		cache, err := r.cacheEngine.Get(ctx, key)
		if err == nil {
			resp, err := r.UnMarshalWrapper(cache, resp)
			if err != nil {
				return nil, err
			}
			return resp, nil
		}

		if r.cacheEngine.IsNotFoundErr(err) {
			curResp := r.handler(reqs...)
			respMarshal, err := r.MarshalWrapper(curResp)
			if err != nil {
				return nil, err
			}
			if err := r.cacheEngine.Set(ctx, key, respMarshal, int64(r.expired)); err != nil {
				return nil, err
			}
			return curResp, err
		}

		// 获取cache失败，直接访问原函数
		return r.handler(reqs...), nil
	}
	i, err, _ := r.group.Do(string(key), fn)
	if err != nil {
		return nil, err
	}

	ret := i.([]interface{})
	return ret, err
}

type cacheWrapper struct {
	cacheEngine KVCacheI
	group       *singleflight.Group
	RequestFormatter
	expired int // 缓存多久

	uniqKey string // 区分接口
	handler func(...interface{}) []interface{}
}
