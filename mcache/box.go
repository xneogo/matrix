/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/7/3 -- 10:41
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: mcache mcache/box.go
*/

package mcache

import (
	"context"
	"time"
)

type ZBox interface {
	// Get from backend and store the value in dst.
	Get(ctx context.Context, key string, dst interface{}) error
	MustGet(ctx context.Context, key string, dst interface{})

	// GetMulti get from backend and store the value in dst,
	// dst must be a map or pointer-to-map in format of
	// map[string]interface{} where key is the passed key if
	// key exists.
	GetMulti(ctx context.Context, keys []string, dstMap interface{}) error
	MustGetMulti(ctx context.Context, keys []string, dstMap interface{})

	// Exists ask for backend whether specified item exists.
	Exists(ctx context.Context, key string) (bool, error)
	MustExists(ctx context.Context, key string) bool

	// ExistsMulti ask for backend whether specified items exists.
	ExistsMulti(ctx context.Context, keys ...string) ([]bool, error)
	MustExistsMulti(ctx context.Context, keys ...string) []bool

	// Set set key and value with timeout.
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	MustSet(ctx context.Context, key string, value interface{}, ttl time.Duration)

	// SetMulti set keys with values. values must be a slice
	SetMulti(ctx context.Context, keys []string, values interface{}, ttl time.Duration) error
	MustSetMulti(ctx context.Context, keys []string, values interface{}, ttl time.Duration)

	// Delete remove the specified item by key.
	Delete(ctx context.Context, keys ...string) error
	MustDelete(ctx context.Context, keys ...string)
}
