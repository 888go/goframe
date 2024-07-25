// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包 gcache 为进程提供各种缓存管理。
//
// 默认情况下，它提供了一个并发安全的内存缓存适配器给进程。 md5:83aa9516287cdc99
package gcache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

// Func是缓存函数，它计算并返回值。 md5:a8aeba091cce1386
type Func func(ctx context.Context) (value interface{}, err error)

const (
	DurationNoExpire = time.Duration(0) // 永不过期的持续时间。 md5:2536f018477cbf65
)

// Default cache object.
var defaultCache = New()

// Set 使用键值对 `key`-`value` 设置缓存，该缓存在 `duration` 时间后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，它将删除 `data` 中的键。 md5:7faea7b643bffd7c
func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return defaultCache.Set(ctx, key, value, duration)
}

// SetMap 批量设置缓存，使用 `data` 映射（键值对）的方式，其在 `duration` 后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或给定的 `value` 为 `nil`，则会删除 `data` 中的键。 md5:a09a11cd5d9d21e6
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	return defaultCache.SetMap(ctx, data, duration)
}

// SetIfNotExist 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
//
// 如果`duration`为0，缓存不会过期。
// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。 md5:38aa90beb53ed441
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExist(ctx, key, value, duration)
}

// SetIfNotExistFunc 如果`key`不存在于缓存中，则使用函数`f`的结果设置`key`并返回true。
// 否则，如果`key`已存在，则不做任何操作并返回false。
//
// 参数`value`可以是类型为`func() interface{}`的函数，
// 但如果其结果为nil，则不会执行任何操作。
//
// 如果`duration`等于0，则不设置过期时间。
// 如果`duration`小于0或给定的`value`为nil，则删除该`key`。 md5:8300c80b9bab735d
func SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExistFunc(ctx, key, f, duration)
}

// SetIfNotExistFuncLock 当`key`在缓存中不存在时，使用函数`f`的结果设置`key`，并返回true。
// 如果`key`已经存在，则不执行任何操作并返回false。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，将删除`key`。
//
// 注意，它与函数`SetIfNotExistFunc`的区别在于，函数`f`在写入互斥锁内部执行，以保证并发安全性。 md5:629e13ace9eaf720
func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	return defaultCache.SetIfNotExistFuncLock(ctx, key, f, duration)
}

// Get 从缓存中检索并返回给定 `key` 的关联值。如果不存在、值为nil或已过期，它将返回nil。如果你想检查`key`是否存在于缓存中，建议使用Contains函数。 md5:f78c30f8338ce106
func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return defaultCache.Get(ctx, key)
}

// GetOrSet 获取并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 这对键值将在指定的`duration`后过期。
//
// 如果`duration`为0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil，它则不做任何操作。 md5:b8646fcb99c81de9
func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSet(ctx, key, value, duration)
}

// GetOrSetFunc 获取并返回`key`的值，如果缓存中不存在`key`，则使用函数`f`的结果设置`key`并返回该结果。键值对在`duration`时间后过期。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且其结果为nil，则不执行任何操作。 md5:822486c86baa87d1
func GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSetFunc(ctx, key, f, duration)
}

// GetOrSetFuncLock 获取并返回键`key`的值，或者如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`，并返回其结果。键值对在`duration`后过期。
//
// 如果`duration`为0，它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数并且函数结果为nil，它将不执行任何操作。
//
// 注意，它与`GetOrSetFunc`函数不同，函数`f`是在写入互斥锁保护下执行的，以确保并发安全。 md5:3e49c54e5e0c2857
func GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	return defaultCache.GetOrSetFuncLock(ctx, key, f, duration)
}

// Contains 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
func Contains(ctx context.Context, key interface{}) (bool, error) {
	return defaultCache.Contains(ctx, key)
}

// GetExpire 从缓存中检索并返回 `key` 的过期时间。
//
// 注意，
// 如果 `key` 没有过期，它将返回 0。
// 如果 `key` 不在缓存中，它将返回 -1。 md5:d80ce12df8668b97
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return defaultCache.GetExpire(ctx, key)
}

// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。 md5:d3b1c8af168b0ebf
func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error) {
	return defaultCache.Remove(ctx, keys...)
}

// 从缓存中删除`keys`。 md5:370028bf9f2e1d24
func Removes(ctx context.Context, keys []interface{}) error {
	return defaultCache.Removes(ctx, keys)
}

// Update 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。 md5:6d92816db5b1d3bd
func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	return defaultCache.Update(ctx, key, value)
}

// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。 md5:b974907dd46b44be
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	return defaultCache.UpdateExpire(ctx, key, duration)
}

// Size 返回缓存中的项目数量。 md5:2122f80de9340261
func Size(ctx context.Context) (int, error) {
	return defaultCache.Size(ctx)
}

// Data 返回缓存中所有键值对的副本，以映射类型形式呈现。
// 注意：此函数可能会占用大量内存，请根据需要决定是否实现该功能。 md5:c44cdbd9b10ab98f
func Data(ctx context.Context) (map[interface{}]interface{}, error) {
	return defaultCache.Data(ctx)
}

// Keys 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
func Keys(ctx context.Context) ([]interface{}, error) {
	return defaultCache.Keys(ctx)
}

// KeyStrings返回缓存中的所有键作为字符串切片。 md5:3b0126221389825e
func KeyStrings(ctx context.Context) ([]string, error) {
	return defaultCache.KeyStrings(ctx)
}

// Values 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
func Values(ctx context.Context) ([]interface{}, error) {
	return defaultCache.Values(ctx)
}

// MustGet 的行为就像 Get 一样，但如果发生任何错误，它会引发 panic。 md5:9004545d221e9637
func MustGet(ctx context.Context, key interface{}) *gvar.Var {
	return defaultCache.MustGet(ctx, key)
}

// MustGetOrSet 的行为类似于 GetOrSet，但是如果发生任何错误，它会直接 panic。 md5:684c6b06451a2f6f
func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSet(ctx, key, value, duration)
}

// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但如果发生任何错误，则会引发 panic。 md5:07fd1ef2dbfce0b4
func MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSetFunc(ctx, key, f, duration)
}

// MustGetOrSetFuncLock 行为与 GetOrSetFuncLock 类似，但如果发生任何错误，它将引发恐慌。 md5:7f84f54a71da5305
func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	return defaultCache.MustGetOrSetFuncLock(ctx, key, f, duration)
}

// MustContains 的行为就像 Contains，但如果发生任何错误，它将引发恐慌。 md5:63cc1bbb0025d8b1
func MustContains(ctx context.Context, key interface{}) bool {
	return defaultCache.MustContains(ctx, key)
}

// MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会直接 panic。 md5:c97fa5941bbc47a3
func MustGetExpire(ctx context.Context, key interface{}) time.Duration {
	return defaultCache.MustGetExpire(ctx, key)
}

// MustSize 行为类似于 Size，但在发生错误时会引发 panic。 md5:cee955b74cc42d5c
func MustSize(ctx context.Context) int {
	return defaultCache.MustSize(ctx)
}

// MustData的行为就像Data一样，但如果发生任何错误，它会引发恐慌。 md5:b53b751e2003cd20
func MustData(ctx context.Context) map[interface{}]interface{} {
	return defaultCache.MustData(ctx)
}

// MustKeys 行为与 Keys 类似，但如果发生任何错误，它将引发 panic。 md5:7f7801d0cd170166
func MustKeys(ctx context.Context) []interface{} {
	return defaultCache.MustKeys(ctx)
}

// MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它会直接 panic。 md5:3efe93008da2eb0f
func MustKeyStrings(ctx context.Context) []string {
	return defaultCache.MustKeyStrings(ctx)
}

// MustValues 行为类似于 Values，但如果发生任何错误则会引发 panic。 md5:859aff610512a748
func MustValues(ctx context.Context) []interface{} {
	return defaultCache.MustValues(ctx)
}
