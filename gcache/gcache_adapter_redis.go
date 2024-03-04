// 版权声明 2020 gf Author(https://github.com/gogf/gf)。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

package gcache

import (
	"context"
	"time"
	
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/util/gconv"
)

// AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。
type AdapterRedis struct {
	redis *gredis.Redis
}

// NewAdapterRedis 创建并返回一个新的内存缓存对象。
func NewAdapterRedis(redis *gredis.Redis) Adapter {
	return &AdapterRedis{
		redis: redis,
	}
}

// Set 通过 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 的键。
func (c *AdapterRedis) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error) {
	redisKey := gconv.String(key)
	if value == nil || duration < 0 {
		_, err = c.redis.Del(ctx, redisKey)
	} else {
		if duration == 0 {
			_, err = c.redis.Set(ctx, redisKey, value)
		} else {
			_, err = c.redis.Set(ctx, redisKey, value, gredis.SetOption{TTLOption: gredis.TTLOption{PX: gconv.PtrInt64(duration.Milliseconds())}})
		}
	}
	return err
}

// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则会删除 `data` 中的键。
func (c *AdapterRedis) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	// DEL.
	if duration < 0 {
		var (
			index = 0
			keys  = make([]string, len(data))
		)
		for k := range data {
			keys[index] = gconv.String(k)
			index += 1
		}
		_, err := c.redis.Del(ctx, keys...)
		if err != nil {
			return err
		}
	}
	if duration == 0 {
		err := c.redis.MSet(ctx, gconv.Map(data))
		if err != nil {
			return err
		}
	}
	if duration > 0 {
		var err error
		for k, v := range data {
			if err = c.Set(ctx, k, v, duration); err != nil {
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
func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	var (
		err      error
		redisKey = gconv.String(key)
	)
	// 执行函数并获取结果。
	f, ok := value.(Func)
	if !ok {
		// 与原始函数值兼容。
		f, ok = value.(func(ctx context.Context) (value interface{}, err error))
	}
	if ok {
		if value, err = f(ctx); err != nil {
			return false, err
		}
	}
	// DEL.
	if duration < 0 || value == nil {
		var delResult int64
		delResult, err = c.redis.Del(ctx, redisKey)
		if err != nil {
			return false, err
		}
		if delResult == 1 {
			return true, err
		}
		return false, err
	}
	ok, err = c.redis.SetNX(ctx, redisKey, value)
	if err != nil {
		return ok, err
	}
	if ok && duration > 0 {
		// 设置过期时间。
		_, err = c.redis.PExpire(ctx, redisKey, duration.Milliseconds())
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
func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return c.SetIfNotExist(ctx, key, value, duration)
}

// SetIfNotExistFuncLock 将通过函数 `f` 计算的结果设置为 `key` 的值，并在以下情况下返回 true：
// 1. 如果 `key` 不存在于缓存中，则设置并返回 true。
// 2. 否则，如果 `key` 已经存在，则不做任何操作并返回 false。
// 若 `duration` 等于 0，则不设置过期时间。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则删除 `key`。
// 注意，此方法与函数 `SetIfNotExistFunc` 的不同之处在于，
// 函数 `f` 在写入互斥锁保护下执行，以确保并发安全。
func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return c.SetIfNotExist(ctx, key, value, duration)
}

// Get 函数根据给定的 <key> 获取并返回其关联的值。
// 如果该键不存在，或者其对应的值为 nil，则返回 nil。
func (c *AdapterRedis) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return c.redis.Get(ctx, gconv.String(key))
}

// GetOrSet 获取并返回键`key`的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 键值对在`duration`时间后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或者给定的`value`为nil，则删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
func (c *AdapterRedis) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	result, err = c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if result.IsNil() {
		return gvar.New(value), c.Set(ctx, key, value, duration)
	}
	return
}

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值在 `duration` 时间后将自动过期。
//
// 如果 `duration` 等于 0，则表示该键值对永不过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
func (c *AdapterRedis) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v.IsNil() {
		value, err := f(ctx)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return gvar.New(value), c.Set(ctx, key, value, duration)
	} else {
		return v, nil
	}
}

// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`时间后过期。
// 如果`duration`为0，则它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
// 注意，该方法与函数`GetOrSetFunc`的不同之处在于，为了保证并发安全，函数`f`在写入互斥锁内执行。
func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error) {
	return c.GetOrSetFunc(ctx, key, f, duration)
}

// Contains 检查并返回 true，如果 `key` 存在于缓存中；否则返回 false。
func (c *AdapterRedis) Contains(ctx context.Context, key interface{}) (bool, error) {
	n, err := c.redis.Exists(ctx, gconv.String(key))
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// Size 返回缓存中的项目数量。
func (c *AdapterRedis) Size(ctx context.Context) (size int, err error) {
	n, err := c.redis.DBSize(ctx)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// Data 返回缓存中所有键值对的副本，类型为 map。注意，此函数可能会导致大量内存使用，
// 因此请按需实现该函数。
func (c *AdapterRedis) Data(ctx context.Context) (map[interface{}]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(ctx, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*gvar.Var
	m, err = c.redis.MGet(ctx, keys...)
	if err != nil {
		return nil, err
	}
	// Type converting.
	data := make(map[interface{}]interface{})
	for k, v := range m {
		data[k] = v.Val()
	}
	return data, nil
}

// Keys 返回缓存中的所有键作为切片。
func (c *AdapterRedis) Keys(ctx context.Context) ([]interface{}, error) {
	keys, err := c.redis.Keys(ctx, "*")
	if err != nil {
		return nil, err
	}
	return gconv.Interfaces(keys), nil
}

// Values 返回缓存中的所有值作为一个切片。
func (c *AdapterRedis) Values(ctx context.Context) ([]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(ctx, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*gvar.Var
	m, err = c.redis.MGet(ctx, keys...)
	if err != nil {
		return nil, err
	}
	// Values.
	var values []interface{}
	for _, key := range keys {
		if v := m[key]; !v.IsNil() {
			values = append(values, v.Val())
		}
	}
	return values, nil
}

// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
func (c *AdapterRedis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	var (
		v        *gvar.Var
		oldPTTL  int64
		redisKey = gconv.String(key)
	)
	// TTL.
	oldPTTL, err = c.redis.PTTL(ctx, redisKey) // 更新ttl -> pttl（毫秒）
	if err != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。
		return
	}
	// Check existence.
	v, err = c.redis.Get(ctx, redisKey)
	if err != nil {
		return
	}
	oldValue = v
	// DEL.
	if value == nil {
		_, err = c.redis.Del(ctx, redisKey)
		if err != nil {
			return
		}
		return
	}
	// Update the value.
	if oldPTTL == -1 {
		_, err = c.redis.Set(ctx, redisKey, value)
	} else {
// 更新 SetEX -> SET PX 选项（毫秒）
// 自 Redis 版本 2.6.12 开始：添加了 EX、PX、NX 和 XX 选项。
// 这段 Go 语言代码注释的中文翻译如下：
// ```go
// 将 SetEX 更新为使用 PX 选项（以毫秒为单位）
// 从 Redis 版本 2.6.12 开始：新增了 EX、PX、NX 和 XX 等选项功能。
		_, err = c.redis.Set(ctx, redisKey, value, gredis.SetOption{TTLOption: gredis.TTLOption{PX: gconv.PtrInt64(oldPTTL)}})
	}
	return oldValue, true, err
}

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
func (c *AdapterRedis) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	var (
		v        *gvar.Var
		oldPTTL  int64
		redisKey = gconv.String(key)
	)
	// TTL.
	oldPTTL, err = c.redis.PTTL(ctx, redisKey)
	if err != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。
		oldPTTL = -1
		return
	}
	oldDuration = time.Duration(oldPTTL) * time.Millisecond
	// DEL.
	if duration < 0 {
		_, err = c.redis.Del(ctx, redisKey)
		return
	}
	// 更新过期时间
	if duration > 0 {
		_, err = c.redis.PExpire(ctx, redisKey, duration.Milliseconds())
	}
	// No expire.
	if duration == 0 {
		v, err = c.redis.Get(ctx, redisKey)
		if err != nil {
			return
		}
		_, err = c.redis.Set(ctx, redisKey, v.Val())
	}
	return
}

// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
func (c *AdapterRedis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	pttl, err := c.redis.PTTL(ctx, gconv.String(key))
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
func (c *AdapterRedis) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// 获取最后一个键值对。
	if lastValue, err = c.redis.Get(ctx, gconv.String(keys[len(keys)-1])); err != nil {
		return nil, err
	}
	// 删除所有给定的键。
	_, err = c.redis.Del(ctx, gconv.Strings(keys)...)
	return
}

// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
// 它在 Redis 服务器中使用了 `FLUSHDB` 命令，该命令可能在服务器中被禁用。
func (c *AdapterRedis) Clear(ctx context.Context) (err error) {
	// "FLUSHDB" 可能不可用。
	err = c.redis.FlushDB(ctx)
	return
}

// Close 关闭缓存。
func (c *AdapterRedis) Close(ctx context.Context) error {
	// It does nothing.
	return nil
}
