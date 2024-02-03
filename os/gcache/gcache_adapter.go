// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gvar"
)

// Adapter 是缓存功能的核心适配器实现。
//
// 注意：实现者自身应确保这些函数的并发安全性。
type Adapter interface {
// Set 用 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或给定的 `value` 为 nil，则会删除 `data` 的键。
	Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error

// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则表示删除 `data` 中的相应键。
	SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error

// SetIfNotExist 如果`key`不存在于缓存中，则设置带有`key`-`value`对的缓存，并在`duration`后过期。
// 如果`key`在缓存中不存在，它将返回true，并成功将`value`设置到缓存中，否则返回false。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或给定的`value`为nil，则删除`key`。
	SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error)

// SetIfNotExistFunc 设置键`key`为函数`f`的结果，并在`key`不存在于缓存中时返回true，
// 否则如果`key`已存在，则不做任何操作并返回false。
//
// 参数`value`可以是`func() interface{}`类型，但如果其结果为nil，则不会执行任何操作。
//
// 如果`duration` == 0，则不设置过期时间。
// 如果`duration` < 0 或给定的`value`为nil，则会删除`key`。
	SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

// SetIfNotExistFuncLock 将通过函数 `f` 计算得到的结果设置为 `key` 的值，并在以下情况下返回 true：
// 如果 `key` 不存在于缓存中。如果 `key` 已经存在，则不做任何操作并返回 false。
//
// 当 `duration` 等于 0 时，它不会设置过期时间。
// 如果 `duration` 小于 0 或者给定的 `value` 是 nil，则会删除 `key`。
//
// 注意，该函数与 `SetIfNotExistFunc` 函数的不同之处在于，
// 为了保证并发安全性，函数 `f` 在写入互斥锁内执行。
	SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

// Get 方法用于获取并返回给定`key`关联的值。
// 若该键不存在，或者其对应的值为nil，或者已过期，则返回nil。
// 如果你想检查`key`是否存在于缓存中，最好使用Contains函数。
	Get(ctx context.Context, key interface{}) (*gvar.Var, error)

// GetOrSet 函数用于检索并返回键 `key` 对应的值，如果 `key` 不存在于缓存中，则设置 `key`-`value` 键值对，并返回 `value`。该键值对在 `duration` 时间后过期。
// 如果 `duration` 等于 0，则表示永不过期。
// 如果 `duration` 小于 0 或提供的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
	GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值对将在 `duration` 时间后过期。
//
// 若 `duration` 等于 0，则表示永不过期。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
	GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或给定的`value`为 nil，则删除`key`，但如果`value`是一个函数且函数结果为 nil，则不做任何操作。
//
// 注意，它与函数`GetOrSetFunc`的不同之处在于，为了并发安全的目的，函数`f`在写入互斥锁内执行。
	GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

	// Contains 检查并返回 true 如果 `key` 存在于缓存中，否则返回 false。
	Contains(ctx context.Context, key interface{}) (bool, error)

	// Size 返回缓存中的项目数量。
	Size(ctx context.Context) (size int, err error)

// Data 返回缓存中所有键值对的副本，类型为 map 类型。
// 注意，此函数可能会导致大量内存使用，如有必要请自行实现该函数。
	Data(ctx context.Context) (data map[interface{}]interface{}, err error)

	// Keys 返回缓存中的所有键作为切片。
	Keys(ctx context.Context) (keys []interface{}, err error)

	// Values 返回缓存中的所有值作为一个切片。
	Values(ctx context.Context) (values []interface{}, err error)

// Update 更新`key`的值而不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，若`key`在缓存中不存在则为false。
//
// 若给出的`value`为nil，则会删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
	Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时间值。
//
// 如果 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则会删除 `key`。
	UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)

// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
	GetExpire(ctx context.Context, key interface{}) (time.Duration, error)

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后被删除项的值。
	Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)

// Clear 清除缓存中的所有数据。
// 注意：该函数较为敏感，应谨慎使用。
	Clear(ctx context.Context) error

	// Close 在必要时关闭缓存。
	Close(ctx context.Context) error
}
