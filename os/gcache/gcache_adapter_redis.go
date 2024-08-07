// 版权所有 (c) 2020 gogf 作者(https://github.com/gogf/gf)。保留所有权利。
//
// 本源代码形式遵循MIT许可协议。若未随此文件一同分发MIT许可协议副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:7063305469ff40c2

package 缓存类

import (
	"context"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
	gredis "github.com/888go/goframe/database/gredis"
	gconv "github.com/888go/goframe/util/gconv"
)

// AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。 md5:7ac226ec6d59930e
type AdapterRedis struct {
	redis *gredis.Redis
}

// NewAdapterRedis 创建并返回一个新的内存缓存对象。 md5:ac9ad598fcd2adbb
func NewAdapterRedis(redis *gredis.Redis) Adapter {
	return &AdapterRedis{
		redis: redis,
	}
}

// X设置值 使用键值对 `key`-`value` 设置缓存，该缓存在 `duration` 时间后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，它将删除 `data` 中的键。
// md5:7faea7b643bffd7c
func (c *AdapterRedis) X设置值(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error) {
	redisKey := gconv.String(key)
	if value == nil || duration < 0 {
		_, err = c.redis.Del(ctx, redisKey)
	} else {
		if duration == 0 {
			_, err = c.redis.X设置值(ctx, redisKey, value)
		} else {
			_, err = c.redis.X设置值(ctx, redisKey, value, gredis.SetOption{TTLOption: gredis.TTLOption{PX: gconv.X取整数64位指针(duration.Milliseconds())}})
		}
	}
	return err
}

// X设置Map 批量设置缓存，使用 `data` 映射（键值对）的方式，其在 `duration` 后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或给定的 `value` 为 `nil`，则会删除 `data` 中的键。
// md5:a09a11cd5d9d21e6
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
			keys[index] = gconv.String(k)
			index += 1
		}
		_, err := c.redis.Del(上下文, keys...)
		if err != nil {
			return err
		}
	}
	if 时长 == 0 {
		err := c.redis.MSet(上下文, gconv.X取Map(值))
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

// X设置值并跳过已存在 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
// 
// 如果`duration`为0，缓存不会过期。
// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。
// md5:38aa90beb53ed441
func (c *AdapterRedis) X设置值并跳过已存在(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (bool, error) {
	var (
		err      error
		redisKey = gconv.String(名称)
	)
		// 执行函数并获取结果。 md5:1443cd3171693ec8
	f, ok := 值.(Func)
	if !ok {
				// 与原始函数值兼容。 md5:b6980bd817389e7f
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
		// Set the expiration.
		_, err = c.redis.PExpire(上下文, redisKey, 时长.Milliseconds())
		if err != nil {
			return ok, err
		}
		return ok, err
	}
	return ok, err
}

// X设置值并跳过已存在_函数 如果`key`不存在于缓存中，则使用函数`f`的结果设置`key`并返回true。
// 否则，如果`key`已存在，则不做任何操作并返回false。
//
// 参数`value`可以是类型为`func() interface{}`的函数，
// 但如果其结果为nil，则不会执行任何操作。
//
// 如果`duration`等于0，则不设置过期时间。
// 如果`duration`小于0或给定的`value`为nil，则删除该`key`。
// md5:8300c80b9bab735d
func (c *AdapterRedis) X设置值并跳过已存在_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (成功 bool, 错误 error) {
	value, 错误 := 回调函数(上下文)
	if 错误 != nil {
		return false, 错误
	}
	return c.X设置值并跳过已存在(上下文, 名称, value, 时长)
}

// X设置值并跳过已存在_并发安全函数 当`key`在缓存中不存在时，使用函数`f`的结果设置`key`，并返回true。
// 如果`key`已经存在，则不执行任何操作并返回false。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，将删除`key`。
//
// 注意，它与函数`SetIfNotExistFunc`的区别在于，函数`f`在写入互斥锁内部执行，以保证并发安全性。
// md5:629e13ace9eaf720
func (c *AdapterRedis) X设置值并跳过已存在_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (成功 bool, 错误 error) {
	value, 错误 := 回调函数(上下文)
	if 错误 != nil {
		return false, 错误
	}
	return c.X设置值并跳过已存在(上下文, 名称, value, 时长)
}

// X取值 通过给定的 <key> 获取并返回关联的值。如果不存在或其值为 nil，则返回 nil。
// md5:ecb61eca16fb4324
func (c *AdapterRedis) X取值(上下文 context.Context, 名称 interface{}) (*gvar.Var, error) {
	return c.redis.Get(上下文, gconv.String(名称))
}

// X取值或设置值 获取并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 这对键值将在指定的`duration`后过期。
//
// 如果`duration`为0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil，它则不做任何操作。
// md5:b8646fcb99c81de9
func (c *AdapterRedis) X取值或设置值(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (结果 *gvar.Var, err error) {
	结果, err = c.X取值(上下文, 名称)
	if err != nil {
		return nil, err
	}
	if 结果.X是否为Nil() {
		return gvar.X创建(值), c.X设置值(上下文, 名称, 值, 时长)
	}
	return
}

// X取值或设置值_函数 获取并返回`key`的值，如果缓存中不存在`key`，则使用函数`f`的结果设置`key`并返回该结果。键值对在`duration`时间后过期。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且其结果为nil，则不执行任何操作。
// md5:822486c86baa87d1
func (c *AdapterRedis) X取值或设置值_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (结果 *gvar.Var, err error) {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		return nil, err
	}
	if v.X是否为Nil() {
		value, err := 回调函数(上下文)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return gvar.X创建(value), c.X设置值(上下文, 名称, value, 时长)
	} else {
		return v, nil
	}
}

// X取值或设置值_并发安全函数 获取并返回键`key`的值，或者如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`，并返回其结果。键值对在`duration`后过期。
// 
// 如果`duration`为0，它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数并且函数结果为nil，它将不执行任何操作。
// 
// 注意，它与`GetOrSetFunc`函数不同，函数`f`是在写入互斥锁保护下执行的，以确保并发安全。
// md5:3e49c54e5e0c2857
func (c *AdapterRedis) X取值或设置值_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (结果 *gvar.Var, err error) {
	return c.X取值或设置值_函数(上下文, 名称, 回调函数, 时长)
}

// X是否存在 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
func (c *AdapterRedis) X是否存在(上下文 context.Context, 名称 interface{}) (bool, error) {
	n, err := c.redis.Exists(上下文, gconv.String(名称))
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// X取数量 返回缓存中的项目数量。 md5:2122f80de9340261
func (c *AdapterRedis) X取数量(上下文 context.Context) (数量 int, 错误 error) {
	n, 错误 := c.redis.DBSize(上下文)
	if 错误 != nil {
		return 0, 错误
	}
	return int(n), nil
}

// X取所有键值Map副本 返回缓存中所有键值对的副本，以映射类型形式呈现。
// 注意：此函数可能会占用大量内存，请根据需要决定是否实现该功能。
// md5:c44cdbd9b10ab98f
func (c *AdapterRedis) X取所有键值Map副本(上下文 context.Context) (map[interface{}]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*gvar.Var
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

// X取所有键 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
func (c *AdapterRedis) X取所有键(上下文 context.Context) ([]interface{}, error) {
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	return gconv.X取any切片(keys), nil
}

// X取所有值 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
func (c *AdapterRedis) X取所有值(上下文 context.Context) ([]interface{}, error) {
	// Keys.
	keys, err := c.redis.Keys(上下文, "*")
	if err != nil {
		return nil, err
	}
	// Key-Value pairs.
	var m map[string]*gvar.Var
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

// X更新值 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。
// md5:6d92816db5b1d3bd
func (c *AdapterRedis) X更新值(上下文 context.Context, 名称 interface{}, 值 interface{}) (旧值 *gvar.Var, exist bool, err error) {
	var (
		v        *gvar.Var
		oldPTTL  int64
		redisKey = gconv.String(名称)
	)
	// TTL.
	oldPTTL, err = c.redis.PTTL(上下文, redisKey) // update ttl -> 更新时间戳到毫秒级的pttl. md5:a9616c495a46fa50
	if err != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。 md5:a51ac96e5909ca59
		return
	}
	// Check existence.
	v, err = c.redis.Get(上下文, redisKey)
	if err != nil {
		return
	}
	旧值 = v
	// DEL.
	if 值 == nil {
		_, err = c.redis.Del(上下文, redisKey)
		if err != nil {
			return
		}
		return
	}
	// Update the value.
	if oldPTTL == -1 {
		_, err = c.redis.X设置值(上下文, redisKey, 值)
	} else {
		// 更新 SetEX -> 设置PX选项（毫秒）
		// 从Redis版本2.6.12开始：添加了EX、PX、NX和XX选项。
		// md5:490be86df7cc2df5
		_, err = c.redis.X设置值(上下文, redisKey, 值, gredis.SetOption{TTLOption: gredis.TTLOption{PX: gconv.X取整数64位指针(oldPTTL)}})
	}
	return 旧值, true, err
}

// X更新过期时间 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
// md5:b974907dd46b44be
func (c *AdapterRedis) X更新过期时间(上下文 context.Context, 名称 interface{}, 时长 time.Duration) (旧过期时长 time.Duration, 错误 error) {
	var (
		v        *gvar.Var
		oldPTTL  int64
		redisKey = gconv.String(名称)
	)
	// TTL.
	oldPTTL, 错误 = c.redis.PTTL(上下文, redisKey)
	if 错误 != nil {
		return
	}
	if oldPTTL == -2 || oldPTTL == 0 {
		// 它不存在或已过期。 md5:a51ac96e5909ca59
		oldPTTL = -1
		return
	}
	旧过期时长 = time.Duration(oldPTTL) * time.Millisecond
	// DEL.
	if 时长 < 0 {
		_, 错误 = c.redis.Del(上下文, redisKey)
		return
	}
	// Update the expiration.
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

// X取过期时间 从缓存中检索并返回 `key` 的过期时间。
// 
// 注意，
// 如果 `key` 没有过期，它将返回 0。
// 如果 `key` 不在缓存中，它将返回 -1。
// md5:d80ce12df8668b97
func (c *AdapterRedis) X取过期时间(上下文 context.Context, 名称 interface{}) (time.Duration, error) {
	pttl, err := c.redis.PTTL(上下文, gconv.String(名称))
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

// X删除并带返回值 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:b3f23906b769df08
func (c *AdapterRedis) X删除并带返回值(上下文 context.Context, 名称s ...interface{}) (最后一个删除值 *gvar.Var, err error) {
	if len(名称s) == 0 {
		return nil, nil
	}
		// 获取最后一个键值。 md5:c348d395d5ea0c9f
	if 最后一个删除值, err = c.redis.Get(上下文, gconv.String(名称s[len(名称s)-1])); err != nil {
		return nil, err
	}
		// 删除所有给定的键。 md5:5c8528683a62a6e5
	_, err = c.redis.Del(上下文, gconv.X取文本切片(名称s)...)
	return
}

// X清空 清空缓存中的所有数据。
// 注意，此函数具有敏感性，应谨慎使用。
// 它使用了 Redis 服务器中的 `FLUSHDB` 命令，但该命令可能在服务器中被禁用。
// md5:e9b895cf3a7760c0
func (c *AdapterRedis) X清空(上下文 context.Context) (错误 error) {
		// "FLUSHDB"可能不可用。 md5:95fb09eb47c6baab
	错误 = c.redis.FlushDB(上下文)
	return
}

// X关闭 关闭缓存。 md5:c1a9d7a347be93a8
func (c *AdapterRedis) X关闭(上下文 context.Context) error {
	// It does nothing.
	return nil
}
