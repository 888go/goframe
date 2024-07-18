// 版权所有 (c) 2020 gogf 作者(https://github.com/gogf/gf)。保留所有权利。
//
// 本源代码形式遵循MIT许可协议。若未随此文件一同分发MIT许可协议副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:7063305469ff40c2

package gcache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/util/gconv"
)

// AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。 md5:7ac226ec6d59930e
type AdapterRedis struct {
	redis *gredis.Redis
}

// NewAdapterRedis 创建并返回一个新的内存缓存对象。 md5:ac9ad598fcd2adbb
// ff:
// redis:
func NewAdapterRedis(redis *gredis.Redis) Adapter {
	return &AdapterRedis{
		redis: redis,
	}
}

// Set 使用键值对 `key`-`value` 设置缓存，该缓存在 `duration` 时间后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，它将删除 `data` 中的键。
// md5:7faea7b643bffd7c
// yx:true
// ff:设置值
// c:
// ctx:
// key:
// value:
// duration:
// err:
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

// SetMap 批量设置缓存，使用 `data` 映射（键值对）的方式，其在 `duration` 后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或给定的 `value` 为 `nil`，则会删除 `data` 中的键。
// md5:a09a11cd5d9d21e6
// ff:设置Map
// c:
// ctx:上下文
// data:值
// duration:时长
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

// SetIfNotExist 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
// 
// 如果`duration`为0，缓存不会过期。
// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。
// md5:38aa90beb53ed441
// ff:设置值并跳过已存在
// c:
// ctx:上下文
// key:名称
// value:值
// duration:时长
func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	var (
		err      error
		redisKey = gconv.String(key)
	)
	// 执行函数并获取结果。 md5:1443cd3171693ec8
	f, ok := value.(Func)
	if !ok {
		// 与原始函数值兼容。 md5:b6980bd817389e7f
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
		// Set the expiration.
		_, err = c.redis.PExpire(ctx, redisKey, duration.Milliseconds())
		if err != nil {
			return ok, err
		}
		return ok, err
	}
	return ok, err
}

// SetIfNotExistFunc 如果`key`不存在于缓存中，则使用函数`f`的结果设置`key`并返回true。
// 否则，如果`key`已存在，则不做任何操作并返回false。
//
// 参数`value`可以是类型为`func() interface{}`的函数，
// 但如果其结果为nil，则不会执行任何操作。
//
// 如果`duration`等于0，则不设置过期时间。
// 如果`duration`小于0或给定的`value`为nil，则删除该`key`。
// md5:8300c80b9bab735d
// ff:设置值并跳过已存在_函数
// c:
// ctx:上下文
// key:名称
// f:回调函数
// duration:时长
// ok:成功
// err:错误
func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return c.SetIfNotExist(ctx, key, value, duration)
}

// SetIfNotExistFuncLock 当`key`在缓存中不存在时，使用函数`f`的结果设置`key`，并返回true。
// 如果`key`已经存在，则不执行任何操作并返回false。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，将删除`key`。
//
// 注意，它与函数`SetIfNotExistFunc`的区别在于，函数`f`在写入互斥锁内部执行，以保证并发安全性。
// md5:629e13ace9eaf720
// ff:设置值并跳过已存在_并发安全函数
// c:
// ctx:上下文
// key:名称
// f:回调函数
// duration:时长
// ok:成功
// err:错误
func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return c.SetIfNotExist(ctx, key, value, duration)
}

// Get 通过给定的 <key> 获取并返回关联的值。如果不存在或其值为 nil，则返回 nil。
// md5:ecb61eca16fb4324
// ff:取值
// c:
// ctx:上下文
// key:名称
func (c *AdapterRedis) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return c.redis.Get(ctx, gconv.String(key))
}

// GetOrSet 获取并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 这对键值将在指定的`duration`后过期。
//
// 如果`duration`为0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil，它则不做任何操作。
// md5:b8646fcb99c81de9
// ff:取值或设置值
// c:
// ctx:上下文
// key:名称
// value:值
// duration:时长
// result:结果
// err:
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

// GetOrSetFunc 获取并返回`key`的值，如果缓存中不存在`key`，则使用函数`f`的结果设置`key`并返回该结果。键值对在`duration`时间后过期。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且其结果为nil，则不执行任何操作。
// md5:822486c86baa87d1
// ff:取值或设置值_函数
// c:
// ctx:上下文
// key:名称
// f:回调函数
// duration:时长
// result:结果
// err:
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

// GetOrSetFuncLock 获取并返回键`key`的值，或者如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`，并返回其结果。键值对在`duration`后过期。
// 
// 如果`duration`为0，它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数并且函数结果为nil，它将不执行任何操作。
// 
// 注意，它与`GetOrSetFunc`函数不同，函数`f`是在写入互斥锁保护下执行的，以确保并发安全。
// md5:3e49c54e5e0c2857
// ff:取值或设置值_并发安全函数
// c:
// ctx:上下文
// key:名称
// f:回调函数
// duration:时长
// result:结果
// err:
func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error) {
	return c.GetOrSetFunc(ctx, key, f, duration)
}

// Contains 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
// ff:是否存在
// c:
// ctx:上下文
// key:名称
func (c *AdapterRedis) Contains(ctx context.Context, key interface{}) (bool, error) {
	n, err := c.redis.Exists(ctx, gconv.String(key))
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// Size 返回缓存中的项目数量。 md5:2122f80de9340261
// ff:取数量
// c:
// ctx:上下文
// size:数量
// err:错误
func (c *AdapterRedis) Size(ctx context.Context) (size int, err error) {
	n, err := c.redis.DBSize(ctx)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// Data 返回缓存中所有键值对的副本，以映射类型形式呈现。
// 注意：此函数可能会占用大量内存，请根据需要决定是否实现该功能。
// md5:c44cdbd9b10ab98f
// ff:取所有键值Map副本
// c:
// ctx:上下文
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

// Keys 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
// ff:取所有键
// c:
// ctx:上下文
func (c *AdapterRedis) Keys(ctx context.Context) ([]interface{}, error) {
	keys, err := c.redis.Keys(ctx, "*")
	if err != nil {
		return nil, err
	}
	return gconv.Interfaces(keys), nil
}

// Values 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
// ff:取所有值
// c:
// ctx:上下文
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

// Update 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。
// md5:6d92816db5b1d3bd
// ff:更新值
// c:
// ctx:上下文
// key:名称
// value:值
// oldValue:旧值
// exist:
// err:
func (c *AdapterRedis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	var (
		v        *gvar.Var
		oldPTTL  int64
		redisKey = gconv.String(key)
	)
	// TTL.
	oldPTTL, err = c.redis.PTTL(ctx, redisKey) // update ttl -> 更新时间戳到毫秒级的pttl. md5:a9616c495a46fa50
	if err != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。 md5:a51ac96e5909ca59
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
// 更新 SetEX -> 设置PX选项（毫秒）
// 从Redis版本2.6.12开始：添加了EX、PX、NX和XX选项。
// md5:490be86df7cc2df5
		_, err = c.redis.Set(ctx, redisKey, value, gredis.SetOption{TTLOption: gredis.TTLOption{PX: gconv.PtrInt64(oldPTTL)}})
	}
	return oldValue, true, err
}

// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
// md5:b974907dd46b44be
// ff:更新过期时间
// c:
// ctx:上下文
// key:名称
// duration:时长
// oldDuration:旧过期时长
// err:错误
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
		// 它不存在或已过期。 md5:a51ac96e5909ca59
		oldPTTL = -1
		return
	}
	oldDuration = time.Duration(oldPTTL) * time.Millisecond
	// DEL.
	if duration < 0 {
		_, err = c.redis.Del(ctx, redisKey)
		return
	}
	// Update the expiration.
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

// GetExpire 从缓存中检索并返回 `key` 的过期时间。
// 
// 注意，
// 如果 `key` 没有过期，它将返回 0。
// 如果 `key` 不在缓存中，它将返回 -1。
// md5:d80ce12df8668b97
// ff:取过期时间
// c:
// ctx:上下文
// key:名称
func (c *AdapterRedis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	pttl, err := c.redis.PTTL(ctx, gconv.String(key))
	if err != nil {
		return 0, err
	}
	switch pttl {
	case -1:
		return 0, nil
	case -2, 0: // 它不存在或已过期。 md5:a51ac96e5909ca59
		return -1, nil
	default:
		return time.Duration(pttl) * time.Millisecond, nil
	}
}

// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:b3f23906b769df08
// ff:删除并带返回值
// c:
// ctx:上下文
// keys:名称s
// lastValue:最后一个删除值
// err:
func (c *AdapterRedis) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// 获取最后一个键值。 md5:c348d395d5ea0c9f
	if lastValue, err = c.redis.Get(ctx, gconv.String(keys[len(keys)-1])); err != nil {
		return nil, err
	}
	// 删除所有给定的键。 md5:5c8528683a62a6e5
	_, err = c.redis.Del(ctx, gconv.Strings(keys)...)
	return
}

// Clear 清空缓存中的所有数据。
// 注意，此函数具有敏感性，应谨慎使用。
// 它使用了 Redis 服务器中的 `FLUSHDB` 命令，但该命令可能在服务器中被禁用。
// md5:e9b895cf3a7760c0
// ff:清空
// c:
// ctx:上下文
// err:错误
func (c *AdapterRedis) Clear(ctx context.Context) (err error) {
	// "FLUSHDB"可能不可用。 md5:95fb09eb47c6baab
	err = c.redis.FlushDB(ctx)
	return
}

// Close 关闭缓存。 md5:c1a9d7a347be93a8
// ff:关闭
// c:
// ctx:上下文
func (c *AdapterRedis) Close(ctx context.Context) error {
	// It does nothing.
	return nil
}
