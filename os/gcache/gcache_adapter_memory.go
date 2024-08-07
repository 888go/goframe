// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"context"
	"math"
	"time"

	glist "github.com/888go/goframe/container/glist"
	gset "github.com/888go/goframe/container/gset"
	gtype "github.com/888go/goframe/container/gtype"
	gvar "github.com/888go/goframe/container/gvar"
	gtime "github.com/888go/goframe/os/gtime"
	gtimer "github.com/888go/goframe/os/gtimer"
)

// AdapterMemory是一个适配器，它实现了使用内存。 md5:1058c2331fc6bbaa
type AdapterMemory struct {
	// cap 限制了缓存池的大小。
	// 如果缓存的大小超过了 cap，
	// 则按照 LRU（最近最少使用）算法进行缓存淘汰过程。
	// 默认值为 0，表示没有大小限制。
	// md5:70436dcd07b73070
	cap         int
	data        *adapterMemoryData        // data 是底层的缓存数据，它存储在一个哈希表中。 md5:7cfaf636328aa0e7
	expireTimes *adapterMemoryExpireTimes // expireTimes是过期键到其时间戳的映射，用于快速索引和删除。 md5:5e7fa0cd3e17ed6c
	expireSets  *adapterMemoryExpireSets  // expireSets 是过期时间戳到其键集合的映射，用于快速索引和删除。 md5:d2c25eb345e1ea19
	lru         *adapterMemoryLru         // lru 是 LRU（Least Recently Used）管理器，当属性 cap 大于 0 时启用。 md5:182c6471c0b4b317
	lruGetList  *glist.List               // lruGetList是根据Get函数的LRU历史记录。 md5:0ad54aeec8e8c762
	eventList   *glist.List               // eventList 是内部数据同步的异步事件列表。 md5:48cbe56e8d02ee7f
	closed      *gtype.Bool               // closed 控制缓存是否已关闭。 md5:8ebf4858be3c0e42
}

// Internal cache item.
type adapterMemoryItem struct {
	v interface{} // Value.
	e int64       // 过期时间戳，单位为毫秒。 md5:d7096ed51593fa59
}

// Internal event item.
type adapterMemoryEvent struct {
	k interface{} // Key.
	e int64       // 过期时间，以毫秒为单位。 md5:baebc3abd37be203
}

const (
	// defaultMaxExpire是不设置过期时间的默认过期时间。
	// 它等于math.MaxInt64除以1000000。
	// md5:75ccaa3b4b490a54
	defaultMaxExpire = 9223372036854
)

// X创建内存适配器 创建并返回一个新的内存缓存对象。 md5:188f107c550c0b2e
func X创建内存适配器(淘汰数量 ...int) Adapter {
	c := &AdapterMemory{
		data:        newAdapterMemoryData(),
		lruGetList:  glist.New(true),
		expireTimes: newAdapterMemoryExpireTimes(),
		expireSets:  newAdapterMemoryExpireSets(),
		eventList:   glist.New(true),
		closed:      gtype.NewBool(),
	}
	if len(淘汰数量) > 0 {
		c.cap = 淘汰数量[0]
		c.lru = newMemCacheLru(c)
	}
	// 如果适配器手动从内存适配器更改，这里可能存在“计时器泄露”。
	// 但不必担心这个问题，因为适配器的变更较少，并且如果未被使用，它也不会做什么。
	// md5:0d85b615ef8507fb
	gtimer.X加入单例循环任务(context.Background(), time.Second, c.syncEventAndClearExpired)
	return c
}

// X设置值 使用键值对 `key`-`value` 设置缓存，该缓存在 `duration` 时间后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，它将删除 `data` 中的键。
// md5:7faea7b643bffd7c
func (c *AdapterMemory) X设置值(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	expireTime := c.getInternalExpire(duration)
	c.data.X设置值(key, adapterMemoryItem{
		v: value,
		e: expireTime,
	})
	c.eventList.PushBack(&adapterMemoryEvent{
		k: key,
		e: expireTime,
	})
	return nil
}

// X设置Map 批量设置缓存，使用 `data` 映射（键值对）的方式，其在 `duration` 后过期。
//
// 如果 `duration` 等于 0，则不会过期。
// 如果 `duration` 小于 0 或给定的 `value` 为 `nil`，则会删除 `data` 中的键。
// md5:a09a11cd5d9d21e6
func (c *AdapterMemory) X设置Map(上下文 context.Context, 值 map[interface{}]interface{}, 时长 time.Duration) error {
	var (
		expireTime = c.getInternalExpire(时长)
		err        = c.data.X设置Map(值, expireTime)
	)
	if err != nil {
		return err
	}
	for k := range 值 {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: k,
			e: expireTime,
		})
	}
	return nil
}

// X设置值并跳过已存在 如果缓存中不存在`key`，则设置过期时间为`duration`的`key`-`value`对。如果成功将`value`设置到缓存中，它会返回`true`，表示`key`在缓存中不存在；否则返回`false`。
// 
// 如果`duration`为0，缓存不会过期。
// 如果`duration`小于0或给定的`value`为`nil`，它会删除`key`。
// md5:38aa90beb53ed441
func (c *AdapterMemory) X设置值并跳过已存在(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (bool, error) {
	isContained, err := c.X是否存在(上下文, 名称)
	if err != nil {
		return false, err
	}
	if !isContained {
		if _, err = c.doSetWithLockCheck(上下文, 名称, 值, 时长); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
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
func (c *AdapterMemory) X设置值并跳过已存在_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (bool, error) {
	isContained, err := c.X是否存在(上下文, 名称)
	if err != nil {
		return false, err
	}
	if !isContained {
		value, err := 回调函数(上下文)
		if err != nil {
			return false, err
		}
		if _, err = c.doSetWithLockCheck(上下文, 名称, value, 时长); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// X设置值并跳过已存在_并发安全函数 当`key`在缓存中不存在时，使用函数`f`的结果设置`key`，并返回true。
// 如果`key`已经存在，则不执行任何操作并返回false。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，将删除`key`。
//
// 注意，它与函数`SetIfNotExistFunc`的区别在于，函数`f`在写入互斥锁内部执行，以保证并发安全性。
// md5:629e13ace9eaf720
func (c *AdapterMemory) X设置值并跳过已存在_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (bool, error) {
	isContained, err := c.X是否存在(上下文, 名称)
	if err != nil {
		return false, err
	}
	if !isContained {
		if _, err = c.doSetWithLockCheck(上下文, 名称, 回调函数, 时长); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// X取值 从缓存中检索并返回给定 `key` 的关联值。如果不存在、值为nil或已过期，它将返回nil。如果你想检查`key`是否存在于缓存中，建议使用Contains函数。
// md5:f78c30f8338ce106
func (c *AdapterMemory) X取值(上下文 context.Context, 名称 interface{}) (*gvar.Var, error) {
	item, ok := c.data.X取值(名称)
	if ok && !item.IsExpired() {
				// 如果启用了LRU功能，则将其添加到LRU历史记录中。 md5:01c169ae5b2999b0
		if c.cap > 0 {
			c.lruGetList.PushBack(名称)
		}
		return gvar.X创建(item.v), nil
	}
	return nil, nil
}

// X取值或设置值 获取并返回`key`对应的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 这对键值将在指定的`duration`后过期。
//
// 如果`duration`为0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且函数结果为nil，它则不做任何操作。
// md5:b8646fcb99c81de9
func (c *AdapterMemory) X取值或设置值(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) (*gvar.Var, error) {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return c.doSetWithLockCheck(上下文, 名称, 值, 时长)
	}
	return v, nil
}

// X取值或设置值_函数 获取并返回`key`的值，如果缓存中不存在`key`，则使用函数`f`的结果设置`key`并返回该结果。键值对在`duration`时间后过期。
//
// 如果`duration`等于0，则不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但若`value`是一个函数且其结果为nil，则不执行任何操作。
// md5:822486c86baa87d1
func (c *AdapterMemory) X取值或设置值_函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (*gvar.Var, error) {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		return nil, err
	}
	if v == nil {
		value, err := 回调函数(上下文)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return c.doSetWithLockCheck(上下文, 名称, value, 时长)
	}
	return v, nil
}

// X取值或设置值_并发安全函数 获取并返回键`key`的值，或者如果`key`在缓存中不存在，则使用函数`f`的结果设置`key`，并返回其结果。键值对在`duration`后过期。
// 
// 如果`duration`为0，它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它会删除`key`；但如果`value`是一个函数并且函数结果为nil，它将不执行任何操作。
// 
// 注意，它与`GetOrSetFunc`函数不同，函数`f`是在写入互斥锁保护下执行的，以确保并发安全。
// md5:3e49c54e5e0c2857
func (c *AdapterMemory) X取值或设置值_并发安全函数(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) (*gvar.Var, error) {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return c.doSetWithLockCheck(上下文, 名称, 回调函数, 时长)
	}
	return v, nil
}

// X是否存在 检查并返回如果 `key` 在缓存中存在则为真，否则为假。 md5:4ff234995709b9ab
func (c *AdapterMemory) X是否存在(上下文 context.Context, 名称 interface{}) (bool, error) {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		return false, err
	}
	return v != nil, nil
}

// X取过期时间 从缓存中检索并返回 `key` 的过期时间。
// 
// 注意，
// 如果 `key` 没有过期，它将返回 0。
// 如果 `key` 不在缓存中，它将返回 -1。
// md5:d80ce12df8668b97
func (c *AdapterMemory) X取过期时间(上下文 context.Context, 名称 interface{}) (time.Duration, error) {
	if item, ok := c.data.X取值(名称); ok {
		return time.Duration(item.e-gtime.X取时间戳毫秒()) * time.Millisecond, nil
	}
	return -1, nil
}

// X删除并带返回值 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:d3b1c8af168b0ebf
func (c *AdapterMemory) X删除并带返回值(上下文 context.Context, 名称s ...interface{}) (*gvar.Var, error) {
	var removedKeys []interface{}
	removedKeys, value, err := c.data.X删除并带返回值(名称s...)
	if err != nil {
		return nil, err
	}
	for _, key := range removedKeys {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: key,
			e: gtime.X取时间戳毫秒() - 1000000,
		})
	}
	return gvar.X创建(value), nil
}

// X更新值 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。
// md5:6d92816db5b1d3bd
func (c *AdapterMemory) X更新值(上下文 context.Context, 名称 interface{}, 值 interface{}) (旧值 *gvar.Var, exist bool, err error) {
	v, exist, err := c.data.X更新值(名称, 值)
	return gvar.X创建(v), exist, err
}

// X更新过期时间 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
// md5:b974907dd46b44be
func (c *AdapterMemory) X更新过期时间(上下文 context.Context, 名称 interface{}, 时长 time.Duration) (旧过期时长 time.Duration, 错误 error) {
	newExpireTime := c.getInternalExpire(时长)
	旧过期时长, 错误 = c.data.X更新过期时间(名称, newExpireTime)
	if 错误 != nil {
		return
	}
	if 旧过期时长 != -1 {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: 名称,
			e: newExpireTime,
		})
	}
	return
}

// X取数量 返回缓存的大小。 md5:c939a4ed87cd79ce
func (c *AdapterMemory) X取数量(上下文 context.Context) (数量 int, 错误 error) {
	return c.data.X取数量()
}

// X取所有键值Map副本 返回一个缓存中所有键值对的副本，以映射类型表示。 md5:d88afdf7cfc66604
func (c *AdapterMemory) X取所有键值Map副本(上下文 context.Context) (map[interface{}]interface{}, error) {
	return c.data.X取所有键值Map副本()
}

// X取所有键 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
func (c *AdapterMemory) X取所有键(上下文 context.Context) ([]interface{}, error) {
	return c.data.X取所有键()
}

// X取所有值 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
func (c *AdapterMemory) X取所有值(上下文 context.Context) ([]interface{}, error) {
	return c.data.X取所有值()
}

// X清空 清空缓存中的所有数据。
// 注意，此函数涉及敏感操作，应谨慎使用。
// md5:9212cab88870d3df
func (c *AdapterMemory) X清空(ctx context.Context) error {
	return c.data.X清空()
}

// X关闭 关闭缓存。 md5:c1a9d7a347be93a8
func (c *AdapterMemory) X关闭(ctx context.Context) error {
	if c.cap > 0 {
		c.lru.Close()
	}
	c.closed.X设置值(true)
	return nil
}

// doSetWithLockCheck 如果缓存中不存在键为`key`的项，将`key-value`对设置到缓存中，且该项的过期时间为`duration`。
//
// 如果`duration`为0，则不过期。参数`value`可以是类型为`func() interface{}`的函数，但如果函数结果为nil，则不执行任何操作。
//
// 在将`key-value`对设置到缓存之前，它会使用写入锁双重检查`key`是否已存在于缓存中。
// md5:17967ab63e2b200c
func (c *AdapterMemory) doSetWithLockCheck(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	expireTimestamp := c.getInternalExpire(duration)
	v, err := c.data.SetWithLock(ctx, key, value, expireTimestamp)
	c.eventList.PushBack(&adapterMemoryEvent{k: key, e: expireTimestamp})
	return gvar.X创建(v), err
}

// getInternalExpire 将给定的过期毫秒数转换并返回过期时间。 md5:176ebdcfb2a89f78
func (c *AdapterMemory) getInternalExpire(duration time.Duration) int64 {
	if duration == 0 {
		return defaultMaxExpire
	}
	return gtime.X取时间戳毫秒() + duration.Nanoseconds()/1000000
}

// makeExpireKey 将毫秒级的 `expire` 值归类到其对应的秒级单位。 md5:40d29c22e827fc9e
func (c *AdapterMemory) makeExpireKey(expire int64) int64 {
	return int64(math.Ceil(float64(expire/1000)+1) * 1000)
}

// syncEventAndClearExpired 执行异步任务循环：
// 1. 异步处理事件列表中的数据，
// 并将结果同步到 `expireTimes` 和 `expireSets` 属性。
// 2. 清理过期的键值对数据。
// md5:ce52abd32c5f232e
func (c *AdapterMemory) syncEventAndClearExpired(ctx context.Context) {
	if c.closed.X取值() {
		gtimer.X退出()
		return
	}
	var (
		event         *adapterMemoryEvent
		oldExpireTime int64
		newExpireTime int64
	)
	// ========================
	// 数据同步。
	// ========================
	// md5:a7203ea428e10983
	for {
		v := c.eventList.PopFront()
		if v == nil {
			break
		}
		event = v.(*adapterMemoryEvent)
				// 获取旧的过期集合。 md5:e6633f31f39e1499
		oldExpireTime = c.expireTimes.X取值(event.k)
				// 计算新的过期时间设置。 md5:57b48d53f5270f91
		newExpireTime = c.makeExpireKey(event.e)
		if newExpireTime != oldExpireTime {
			c.expireSets.GetOrNew(newExpireTime).X加入(event.k)
			if oldExpireTime != 0 {
				c.expireSets.GetOrNew(oldExpireTime).X删除(event.k)
			}
						// 更新<event.k>的过期时间。 md5:f04ccde84655d99f
			c.expireTimes.X设置值(event.k, newExpireTime)
		}
				// 通过写操作将键添加到LRU历史中。 md5:ca17e775d3b31310
		if c.cap > 0 {
			c.lru.Push(event.k)
		}
	}
		// 从最近最少使用（Least Recently Used，LRU）缓存中处理过期的键。 md5:c555319093b1296e
	if c.cap > 0 {
		if c.lruGetList.Len() > 0 {
			for {
				if v := c.lruGetList.PopFront(); v != nil {
					c.lru.Push(v)
				} else {
					break
				}
			}
		}
		c.lru.SyncAndClear(ctx)
	}
	// ========================
	// 数据清理。
	// ========================
	// md5:c845ec8cb41f31ac
	var (
		expireSet *gset.Set
		ek        = c.makeExpireKey(gtime.X取时间戳毫秒())
		eks       = []int64{ek - 1000, ek - 2000, ek - 3000, ek - 4000, ek - 5000}
	)
	for _, expireTime := range eks {
		if expireSet = c.expireSets.X取值(expireTime); expireSet != nil {
						// 遍历集合以删除其中的所有键。 md5:de77c90f243260c0
			expireSet.X遍历(func(key interface{}) bool {
				c.clearByKey(key)
				return true
			})
						// 在删除所有键之后，删除集合。 md5:d34b6cd2767c7800
			c.expireSets.Delete(expireTime)
		}
	}
}

// clearByKey 删除给定`key`的键值对。参数`force`指定是否强制执行删除操作。
// md5:5b26398959f735ad
func (c *AdapterMemory) clearByKey(key interface{}, force ...bool) {
		// 在从缓存中真正删除之前，再双检查一次。 md5:53767fc86cbfbf5e
	c.data.DeleteWithDoubleCheck(key, force...)

		// 从`expireTimes`中删除其过期时间。 md5:d2320f7b4a5f1c26
	c.expireTimes.Delete(key)

	// Deleting it from LRU.
	if c.cap > 0 {
		c.lru.X删除并带返回值(key)
	}
}
