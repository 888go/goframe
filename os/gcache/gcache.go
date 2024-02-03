// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gcache 提供了多种用于进程的缓存管理功能。
//
// 默认情况下，它提供了一个线程安全的内存缓存适配器，用于进程中的缓存管理。
package gcache

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gvar"
)

// Func 是缓存函数，用于计算并返回值。
type Func func(ctx context.Context) (value interface{}, err error)

const (
	DurationNoExpire = 0 // Expire 永不过期的持续时间。
)

// 默认缓存对象。
var defaultCache = New()

// Set 通过 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 的键。
func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return defaultCache.Set(ctx, key, value, duration)
}

// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则会删除 `data` 中的键。
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	return defaultCache.SetMap(ctx, data, duration)
}

// SetIfNotExist 若`key`不存在于缓存中，则设置带有`key`-`value`对的缓存，该对在`duration`后过期。
// 如果`key`在缓存中不存在，它将返回true，并成功将`value`设置到缓存中，否则返回false。
//
// 如果`duration` == 0，则不会设置过期时间。
// 如果`duration` < 0 或给定的`value`为nil，则删除`key`。
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExist(ctx, key, value, duration)
}

// SetIfNotExistFunc 函数用于设置 `key` 为函数 `f` 的计算结果，并在 `key` 不存在于缓存中时返回 true，
// 否则如果 `key` 已存在，则不做任何操作并返回 false。
//
// 参数 `value` 可以是类型 `func() interface{}`，但如果其结果为 nil，则该函数不会执行任何操作。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或给定的 `value` 为 nil，则会删除 `key`。
func SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExistFunc(ctx, key, f, duration)
}

// SetIfNotExistFuncLock 将通过函数 `f` 计算的结果设置为 `key` 的值，并在以下情况下返回 true：
// 1. 如果 `key` 不存在于缓存中，则设置并返回 true。
// 2. 否则，如果 `key` 已经存在，则不做任何操作并返回 false。
// 若 `duration` 等于 0，则不设置过期时间。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则删除 `key`。
// 注意，此方法与函数 `SetIfNotExistFunc` 的不同之处在于，
// 函数 `f` 在写入互斥锁保护下执行，以确保并发安全。
func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExistFuncLock(ctx, key, f, duration)
}

// Get 方法通过给定的 `key` 获取并返回关联的值。
// 若该 `key` 对应的值不存在，或者其值为 nil，或已过期，则返回 nil。
// 如果你想检查 `key` 是否存在于缓存中，最好使用 Contains 函数。
func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return defaultCache.Get(ctx, key)
}

// GetOrSet 获取并返回键`key`的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 键值对在`duration`时间后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或者给定的`value`为nil，则删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSet(ctx, key, value, duration)
}

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值在 `duration` 时间后将自动过期。
//
// 如果 `duration` 等于 0，则表示该键值对永不过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
func GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSetFunc(ctx, key, f, duration)
}

// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`时间后过期。
// 如果`duration`为0，则它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
// 注意，该方法与函数`GetOrSetFunc`的不同之处在于，为了保证并发安全，函数`f`在写入互斥锁内执行。
func GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSetFuncLock(ctx, key, f, duration)
}

// Contains 检查并返回 true，如果 `key` 存在于缓存中；否则返回 false。
func Contains(ctx context.Context, key interface{}) (bool, error) {
	return defaultCache.Contains(ctx, key)
}

// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return defaultCache.GetExpire(ctx, key)
}

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后一个被删除项的值。
func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error) {
	return defaultCache.Remove(ctx, keys...)
}

// 删除缓存中的`keys`。
func Removes(ctx context.Context, keys []interface{}) error {
	return defaultCache.Removes(ctx, keys)
}

// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	return defaultCache.Update(ctx, key, value)
}

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	return defaultCache.UpdateExpire(ctx, key, duration)
}

// Size 返回缓存中的项目数量。
func Size(ctx context.Context) (int, error) {
	return defaultCache.Size(ctx)
}

// Data 返回缓存中所有键值对的副本，类型为 map。注意，此函数可能会导致大量内存使用，
// 因此请按需实现该函数。
func Data(ctx context.Context) (map[interface{}]interface{}, error) {
	return defaultCache.Data(ctx)
}

// Keys 返回缓存中的所有键作为切片。
func Keys(ctx context.Context) ([]interface{}, error) {
	return defaultCache.Keys(ctx)
}

// KeyStrings 返回缓存中的所有键，以字符串切片的形式。
func KeyStrings(ctx context.Context) ([]string, error) {
	return defaultCache.KeyStrings(ctx)
}

// Values 返回缓存中的所有值作为一个切片。
func Values(ctx context.Context) ([]interface{}, error) {
	return defaultCache.Values(ctx)
}

// MustGet 行为类似于 Get，但当发生任何错误时，它会触发panic。
func MustGet(ctx context.Context, key interface{}) *gvar.Var {
	return defaultCache.MustGet(ctx, key)
}

// MustGetOrSet 行为类似于 GetOrSet，但是当发生任何错误时，它会触发panic（异常）。
func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSet(ctx, key, value, duration)
}

// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当发生任何错误时它会触发panic。
func MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSetFunc(ctx, key, f, duration)
}

// MustGetOrSetFuncLock 类似于 GetOrSetFuncLock，但如果发生任何错误，它会触发 panic。
func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSetFuncLock(ctx, key, f, duration)
}

// MustContains 行为类似于 Contains，但当发生任何错误时，它会触发panic（异常退出）。
func MustContains(ctx context.Context, key interface{}) bool {
	return defaultCache.MustContains(ctx, key)
}

// MustGetExpire 的行为类似于 GetExpire，但是当发生任何错误时，它会触发panic。
func MustGetExpire(ctx context.Context, key interface{}) time.Duration {
	return defaultCache.MustGetExpire(ctx, key)
}

// MustSize 的行为类似于 Size，但如果发生任何错误，它会触发 panic。
func MustSize(ctx context.Context) int {
	return defaultCache.MustSize(ctx)
}

// MustData 的行为类似于 Data，但是当发生任何错误时它会触发panic（异常）。
func MustData(ctx context.Context) map[interface{}]interface{} {
	return defaultCache.MustData(ctx)
}

// MustKeys 行为类似 Keys，但当发生任何错误时会触发 panic。
func MustKeys(ctx context.Context) []interface{} {
	return defaultCache.MustKeys(ctx)
}

// MustKeyStrings 行为类似 KeyStrings，但当发生任何错误时，它会引发 panic。
func MustKeyStrings(ctx context.Context) []string {
	return defaultCache.MustKeyStrings(ctx)
}

// MustValues 行为类似 Values，但是当发生任何错误时它会触发panic（异常）。
func MustValues(ctx context.Context) []interface{} {
	return defaultCache.MustValues(ctx)
}
