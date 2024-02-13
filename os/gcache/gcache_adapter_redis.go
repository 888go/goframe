// 版权声明 2020 gf Author(https://github.com/gogf/gf)。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

package 缓存类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/util/gconv"
)

// AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。
type AdapterRedis struct {
	redis *redis类.Redis
}

// NewAdapterRedis 创建并返回一个新的内存缓存对象。
func NewAdapterRedis(redis *redis类.Redis) Adapter {
	return &AdapterRedis{
		redis: redis,
	}
}

// Set 通过 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 的键。
func (c *AdapterRedis) X设置值(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (错误 error) {
	redisKey := 转换类.String(名称)
	if 值 == nil || 时长 < 0 {
		_, 错误 = c.redis.Del(上下文, redisKey)
	} else {
		if 时长 == 0 {
			_, 错误 = c.redis.X设置值(上下文, redisKey, 值)
		} else {
			_, 错误 = c.redis.X设置值(上下文, redisKey, 值, redis类.SetOption{TTLOption: redis类.TTLOption{PX: 转换类.X取整数64位指针(时长.Milliseconds())}})
		}
	}
	return 错误
}

// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则会删除 `data` 中的键。
func (c *AdapterRedis) X设置Map(上下文 context.Context, 值 map[interface{}]interface{}, 时长 time.Duration) error {
	if len(值) == 0 {
		return nil
	}
	// DEL.
	if 时长 < 0 {
		var (
			index = 0
			keys  = make([]string, len(值))
		)
		for k := range 值 {
			keys[index] = 转换类.String(k)
			index += 1
		}
		_, err := c.redis.Del(上下文, keys...)
		if err != nil {
			return err
		}
	}
	if 时长 == 0 {
		err := c.redis.MSet(上下文, 转换类.X取Map(值))
		if err != nil {
			return err
		}
	}
	if 时长 > 0 {
		var err error
		for k, v := range 值 {
			if err = c.X设置值(上下文, k, v, 时长); err != nil {
				return err
			}
		}
	}
	return nil
}

// SetIfNotExist 若`key`不存在于缓存中，则设置带有`key`-`value`对的缓存，该对在`duration`后过期。
// 如果`key`在缓存中不存在，它将返回true，并成功将`value`设置到缓存中，否则返回false。
//
// 如果`duration` == 0，则不会设置过期时间。
// 如果`duration` < 0 或给定的`value`为nil，则删除`key`。
func (c *AdapterRedis) X设置值并跳过已存在(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (bool, error) {
	var (
		err      error
		redisKey = 转换类.String(名称)
	)
	// 执行函数并获取结果。
	f, ok := 值.(Func)
	if !ok {
		// 与原始函数值兼容。
		f, ok = 值.(func(ctx context.Context) (value interface{}, err error))
	}
	if ok {
		if 值, err = f(上下文); err != nil {
			return false, err
		}
	}
	// DEL.
	if 时长 < 0 || 值 == nil {
		var delResult int64
		delResult, err = c.redis.Del(上下文, redisKey)
		if err != nil {
			return false, err
		}
		if delResult == 1 {
			return true, err
		}
		return false, err
	}
	ok, err = c.redis.SetNX(上下文, redisKey, 值)
	if err != nil {
		return ok, err
	}
	if ok && 时长 > 0 {
		// 设置过期时间。
		_, err = c.redis.PExpire(上下文, redisKey, 时长.Milliseconds())
		if err != nil {
			return ok, err
		}
		return ok, err
	}
	return ok, err
}

// SetIfNotExistFunc 函数用于设置 `key` 为函数 `f` 的计算结果，并在 `key` 不存在于缓存中时返回 true，
// 否则如果 `key` 已存在，则不做任何操作并返回 false。
//
// 参数 `value` 可以是类型 `func() interface{}`，但如果其结果为 nil，则该函数不会执行任何操作。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或给定的 `value` 为 nil，则会删除 `key`。
func (c *AdapterRedis) X设置值并跳过已存在_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (成功 bool, 错误 error) {
	value, 错误 := 回调函数(上下文)
	if 错误 != nil {
		return false, 错误
	}
	return c.X设置值并跳过已存在(上下文, 名称, value, 时长)
}

// SetIfNotExistFuncLock 将通过函数 `f` 计算的结果设置为 `key` 的值，并在以下情况下返回 true：
// 1. 如果 `key` 不存在于缓存中，则设置并返回 true。
// 2. 否则，如果 `key` 已经存在，则不做任何操作并返回 false。
// 若 `duration` 等于 0，则不设置过期时间。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则删除 `key`。
// 注意，此方法与函数 `SetIfNotExistFunc` 的不同之处在于，
// 函数 `f` 在写入互斥锁保护下执行，以确保并发安全。
func (c *AdapterRedis) X设置值并跳过已存在_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (成功 bool, 错误 error) {
	value, 错误 := 回调函数(上下文)
	if 错误 != nil {
		return false, 错误
	}
	return c.X设置值并跳过已存在(上下文, 名称, value, 时长)
}

// Get 函数根据给定的 <key> 获取并返回其关联的值。
// 如果该键不存在，或者其对应的值为 nil，则返回 nil。
func (c *AdapterRedis) X取值(上下文 context.Context, 名称 interface{}) (*泛型类.Var, error) {
	return c.redis.Get(上下文, 转换类.String(名称))
}

// GetOrSet 获取并返回键`key`的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 键值对在`duration`时间后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或者给定的`value`为nil，则删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
func (c *AdapterRedis) X取值或设置值(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (结果 *泛型类.Var, 错误 error) {
	结果, 错误 = c.X取值(上下文, 名称)
	if 错误 != nil {
		return nil, 错误
	}
	if 结果.X是否为Nil() {
		return 泛型类.X创建(值), c.X设置值(上下文, 名称, 值, 时长)
	}
	return
}

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值在 `duration` 时间后将自动过期。
//
// 如果 `duration` 等于 0，则表示该键值对永不过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
func (c *AdapterRedis) X取值或设置值_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (结果 *泛型类.Var, 错误 error) {
	v, 错误 := c.X取值(上下文, 名称)
	if 错误 != nil {
		return nil, 错误
	}
	if v.X是否为Nil() {
		value, err := 回调函数(上下文)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return 泛型类.X创建(value), c.X设置值(上下文, 名称, value, 时长)
	} else {
		return v, nil
	}
}

// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`时间后过期。
// 如果`duration`为0，则它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
// 注意，该方法与函数`GetOrSetFunc`的不同之处在于，为了保证并发安全，函数`f`在写入互斥锁内执行。
func (c *AdapterRedis) X取值或设置值_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (结果 *泛型类.Var, 错误 error) {
	return c.X取值或设置值_函数(上下文, 名称, 回调函数, 时长)
}

// Contains 检查并返回 true，如果 `key` 存在于缓存中；否则返回 false。
func (c *AdapterRedis) X是否存在(上下文 context.Context, 名称 interface{}) (bool, error) {
	n, err := c.redis.Exists(上下文, 转换类.String(名称))
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// Size 返回缓存中的项目数量。
func (c *AdapterRedis) X取数量(上下文 context.Context) (数量 int, 错误 error) {
	n, 错误 := c.redis.DBSize(上下文)
	if 错误 != nil {
		return 0, 错误
	}
	return int(n), nil
}

// Data 返回缓存中所有键值对的副本，类型为 map。注意，此函数可能会导致大量内存使用，
// 因此请按需实现该函数。
func (c *AdapterRedis) X取所有键值Map副本(上下文 context.Context) (map[interface{}]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*泛型类.Var
	m, err = c.redis.MGet(上下文, keys...)
	if err != nil {
		return nil, err
	}
	// Type converting.
	data := make(map[interface{}]interface{})
	for k, v := range m {
		data[k] = v.X取值()
	}
	return data, nil
}

// Keys 返回缓存中的所有键作为切片。
func (c *AdapterRedis) X取所有键(上下文 context.Context) ([]interface{}, error) {
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	return 转换类.X取any数组(keys), nil
}

// Values 返回缓存中的所有值作为一个切片。
func (c *AdapterRedis) X取所有值(上下文 context.Context) ([]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*泛型类.Var
	m, err = c.redis.MGet(上下文, keys...)
	if err != nil {
		return nil, err
	}
	// Values.
	var values []interface{}
	for _, key := range keys {
		if v := m[key]; !v.X是否为Nil() {
			values = append(values, v.X取值())
		}
	}
	return values, nil
}

// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
func (c *AdapterRedis) X更新值(上下文 context.Context, 名称 interface{}, 值 interface{}) (旧值 *泛型类.Var, 是否存在 bool, 错误 error) {
	var (
		v        *泛型类.Var
		oldPTTL  int64
		redisKey = 转换类.String(名称)
	)
	// TTL.
	oldPTTL, 错误 = c.redis.PTTL(上下文, redisKey) // 更新ttl -> pttl（毫秒）
	if 错误 != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。
		return
	}
	// Check existence.
	v, 错误 = c.redis.Get(上下文, redisKey)
	if 错误 != nil {
		return
	}
	旧值 = v
	// DEL.
	if 值 == nil {
		_, 错误 = c.redis.Del(上下文, redisKey)
		if 错误 != nil {
			return
		}
		return
	}
	// Update the value.
	if oldPTTL == -1 {
		_, 错误 = c.redis.X设置值(上下文, redisKey, 值)
	} else {
// 更新 SetEX -> SET PX 选项（毫秒）
// 自 Redis 版本 2.6.12 开始：添加了 EX、PX、NX 和 XX 选项。
// 这段 Go 语言代码注释的中文翻译如下：
// ```go
// 将 SetEX 更新为使用 PX 选项（以毫秒为单位）
// 从 Redis 版本 2.6.12 开始：新增了 EX、PX、NX 和 XX 等选项功能。
		_, 错误 = c.redis.X设置值(上下文, redisKey, 值, redis类.SetOption{TTLOption: redis类.TTLOption{PX: 转换类.X取整数64位指针(oldPTTL)}})
	}
	return 旧值, true, 错误
}

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
func (c *AdapterRedis) X更新过期时间(上下文 context.Context, 名称 interface{}, 时长 time.Duration) (旧过期时长 time.Duration, 错误 error) {
	var (
		v        *泛型类.Var
		oldPTTL  int64
		redisKey = 转换类.String(名称)
	)
	// TTL.
	oldPTTL, 错误 = c.redis.PTTL(上下文, redisKey)
	if 错误 != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。
		oldPTTL = -1
		return
	}
	旧过期时长 = time.Duration(oldPTTL) * time.Millisecond
	// DEL.
	if 时长 < 0 {
		_, 错误 = c.redis.Del(上下文, redisKey)
		return
	}
	// 更新过期时间
	if 时长 > 0 {
		_, 错误 = c.redis.PExpire(上下文, redisKey, 时长.Milliseconds())
	}
	// No expire.
	if 时长 == 0 {
		v, 错误 = c.redis.Get(上下文, redisKey)
		if 错误 != nil {
			return
		}
		_, 错误 = c.redis.X设置值(上下文, redisKey, v.X取值())
	}
	return
}

// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
func (c *AdapterRedis) X取过期时间(上下文 context.Context, 名称 interface{}) (time.Duration, error) {
	pttl, err := c.redis.PTTL(上下文, 转换类.String(名称))
	if err != nil {
		return 0, err
	}
	switch pttl {
	case -1:
		return 0, nil
	case -2, 0: // 它不存在或已过期。
		return -1, nil
	default:
		return time.Duration(pttl) * time.Millisecond, nil
	}
}

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回被删除的最后一个项目的值。
func (c *AdapterRedis) X删除并带返回值(上下文 context.Context, 名称s ...interface{}) (最后一个删除值 *泛型类.Var, 错误 error) {
	if len(名称s) == 0 {
		return nil, nil
	}
	// 获取最后一个键值对。
	if 最后一个删除值, 错误 = c.redis.Get(上下文, 转换类.String(名称s[len(名称s)-1])); 错误 != nil {
		return nil, 错误
	}
	// 删除所有给定的键。
	_, 错误 = c.redis.Del(上下文, 转换类.X取文本数组(名称s)...)
	return
}

// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
// 它在 Redis 服务器中使用了 `FLUSHDB` 命令，该命令可能在服务器中被禁用。
func (c *AdapterRedis) X清空(上下文 context.Context) (错误 error) {
	// "FLUSHDB" 可能不可用。
	错误 = c.redis.FlushDB(上下文)
	return
}

// Close 关闭缓存。
func (c *AdapterRedis) X关闭(上下文 context.Context) error {
	// It does nothing.
	return nil
}
