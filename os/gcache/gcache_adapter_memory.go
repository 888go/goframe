// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache
import (
	"context"
	"math"
	"time"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gtimer"
	)
// AdapterMemory 是一个使用内存实现的适配器。
type AdapterMemory struct {
// cap 限制缓存池的大小。
// 如果缓存的大小超过 cap，
// 缓存过期过程将根据 LRU 算法执行。
// 默认值为 0，表示无限制。
	cap         int
	data        *adapterMemoryData        // data 是底层缓存数据，存储在一个哈希表中。
	expireTimes *adapterMemoryExpireTimes // expireTimes 是一个过期键与其时间戳的映射，用于快速索引和删除。
	expireSets  *adapterMemoryExpireSets  // expireSets 是一个映射表，用于存储即将过期的时间戳及其对应的键集合。这个映射表用于快速索引和删除操作。
	lru         *adapterMemoryLru         // lru 是 LRU（最近最少使用）管理器，当属性 cap 大于 0 时启用。
	lruGetList  *glist.List               // lruGetList 是根据 Get 函数实现的 LRU（最近最少使用）历史记录列表。
	eventList   *glist.List               // eventList 是用于内部数据同步的异步事件列表。
	closed      *gtype.Bool               // closed 控制缓存是否关闭
}

// 内部缓存项。
type adapterMemoryItem struct {
	v interface{} // Value.
	e int64       // 过期时间戳（毫秒）
}

// 内部事件项
type adapterMemoryEvent struct {
	k interface{} // Key.
	e int64       // 过期时间（以毫秒为单位）
}

const (
// defaultMaxExpire 是未设置过期时间项目的默认过期时间。
// 它等于 math.MaxInt64/1000000。
	defaultMaxExpire = 9223372036854
)

// NewAdapterMemory 创建并返回一个新的内存缓存对象。
func NewAdapterMemory(lruCap ...int) Adapter {
	c := &AdapterMemory{
		data:        newAdapterMemoryData(),
		lruGetList:  glist.New(true),
		expireTimes: newAdapterMemoryExpireTimes(),
		expireSets:  newAdapterMemoryExpireSets(),
		eventList:   glist.New(true),
		closed:      gtype.NewBool(),
	}
	if len(lruCap) > 0 {
		c.cap = lruCap[0]
		c.lru = newMemCacheLru(c)
	}
	return c
}

// Set 通过 `key`-`value` 对设置缓存，该对在 `duration` 后过期。
//
// 如果 `duration` == 0，则永不过期。
// 如果 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 的键。
func (c *AdapterMemory) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	expireTime := c.getInternalExpire(duration)
	c.data.Set(key, adapterMemoryItem{
		v: value,
		e: expireTime,
	})
	c.eventList.PushBack(&adapterMemoryEvent{
		k: key,
		e: expireTime,
	})
	return nil
}

// SetMap 批量设置缓存，通过 `data` 参数中的键值对进行设置，并在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则会删除 `data` 中的键。
func (c *AdapterMemory) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	var (
		expireTime = c.getInternalExpire(duration)
		err        = c.data.SetMap(data, expireTime)
	)
	if err != nil {
		return err
	}
	for k := range data {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: k,
			e: expireTime,
		})
	}
	return nil
}

// SetIfNotExist 若`key`不存在于缓存中，则设置带有`key`-`value`对的缓存，该对在`duration`后过期。
// 如果`key`在缓存中不存在，它将返回true，并成功将`value`设置到缓存中，否则返回false。
//
// 如果`duration` == 0，则不会设置过期时间。
// 如果`duration` < 0 或给定的`value`为nil，则删除`key`。
func (c *AdapterMemory) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	isContained, err := c.Contains(ctx, key)
	if err != nil {
		return false, err
	}
	if !isContained {
		if _, err = c.doSetWithLockCheck(ctx, key, value, duration); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// SetIfNotExistFunc 函数用于设置 `key` 为函数 `f` 的计算结果，并在 `key` 不存在于缓存中时返回 true，
// 否则如果 `key` 已存在，则不做任何操作并返回 false。
//
// 参数 `value` 可以是类型 `func() interface{}`，但如果其结果为 nil，则该函数不会执行任何操作。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或给定的 `value` 为 nil，则会删除 `key`。
func (c *AdapterMemory) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	isContained, err := c.Contains(ctx, key)
	if err != nil {
		return false, err
	}
	if !isContained {
		value, err := f(ctx)
		if err != nil {
			return false, err
		}
		if _, err = c.doSetWithLockCheck(ctx, key, value, duration); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// SetIfNotExistFuncLock 将通过函数 `f` 计算的结果设置为 `key` 的值，并在以下情况下返回 true：
// 1. 如果 `key` 不存在于缓存中，则设置并返回 true。
// 2. 否则，如果 `key` 已经存在，则不做任何操作并返回 false。
// 若 `duration` 等于 0，则不设置过期时间。
// 若 `duration` 小于 0 或提供的 `value` 为 nil，则删除 `key`。
// 注意，此方法与函数 `SetIfNotExistFunc` 的不同之处在于，
// 函数 `f` 在写入互斥锁保护下执行，以确保并发安全。
func (c *AdapterMemory) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {
	isContained, err := c.Contains(ctx, key)
	if err != nil {
		return false, err
	}
	if !isContained {
		if _, err = c.doSetWithLockCheck(ctx, key, f, duration); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// Get 方法通过给定的 `key` 获取并返回关联的值。
// 若该 `key` 对应的值不存在，或者其值为 nil，或已过期，则返回 nil。
// 如果你想检查 `key` 是否存在于缓存中，最好使用 Contains 函数。
func (c *AdapterMemory) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	item, ok := c.data.Get(key)
	if ok && !item.IsExpired() {
		// 如果启用了LRU功能，则添加到LRU历史记录中。
		if c.cap > 0 {
			c.lruGetList.PushBack(key)
		}
		return gvar.New(item.v), nil
	}
	return nil, nil
}

// GetOrSet 获取并返回键`key`的值，如果`key`在缓存中不存在，则设置`key`-`value`对并返回`value`。
// 键值对在`duration`时间后过期。
//
// 如果`duration` == 0，则不会过期。
// 如果`duration` < 0 或者给定的`value`为nil，则删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
func (c *AdapterMemory) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return c.doSetWithLockCheck(ctx, key, value, duration)
	}
	return v, nil
}

// GetOrSetFunc 函数用于获取并返回 `key` 对应的值，如果 `key` 不存在于缓存中，则使用函数 `f` 的结果设置 `key` 并返回其结果。
// 这对键值在 `duration` 时间后将自动过期。
//
// 如果 `duration` 等于 0，则表示该键值对永不过期。
// 如果 `duration` 小于 0 或者给定的 `value` 为 nil，则会删除 `key`，但如果 `value` 是一个函数且函数结果为 nil，则不做任何操作。
func (c *AdapterMemory) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		value, err := f(ctx)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return c.doSetWithLockCheck(ctx, key, value, duration)
	}
	return v, nil
}

// GetOrSetFuncLock 从缓存中获取并返回`key`的值，如果`key`不存在，则使用函数`f`的结果设置`key`并返回其结果。键值对在`duration`时间后过期。
// 如果`duration`为0，则它不会过期。
// 如果`duration`小于0或给定的`value`为nil，它将删除`key`，但如果`value`是一个函数且函数结果为nil，则不做任何操作。
// 注意，该方法与函数`GetOrSetFunc`的不同之处在于，为了保证并发安全，函数`f`在写入互斥锁内执行。
func (c *AdapterMemory) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return c.doSetWithLockCheck(ctx, key, f, duration)
	}
	return v, nil
}

// Contains 检查并返回 true，如果 `key` 存在于缓存中；否则返回 false。
func (c *AdapterMemory) Contains(ctx context.Context, key interface{}) (bool, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return false, err
	}
	return v != nil, nil
}

// GetExpire 从缓存中检索并返回`key`的过期时间。
//
// 注意：
// 如果`key`永不过期，则返回0。
// 如果`key`在缓存中不存在，则返回-1。
func (c *AdapterMemory) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	if item, ok := c.data.Get(key); ok {
		return time.Duration(item.e-gtime.TimestampMilli()) * time.Millisecond, nil
	}
	return -1, nil
}

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回最后一个被删除项的值。
func (c *AdapterMemory) Remove(ctx context.Context, keys ...interface{}) (*gvar.Var, error) {
	var removedKeys []interface{}
	removedKeys, value, err := c.data.Remove(keys...)
	if err != nil {
		return nil, err
	}
	for _, key := range removedKeys {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: key,
			e: gtime.TimestampMilli() - 1000000,
		})
	}
	return gvar.New(value), nil
}

// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
func (c *AdapterMemory) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	v, exist, err := c.data.Update(key, value)
	return gvar.New(v), exist, err
}

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
func (c *AdapterMemory) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	newExpireTime := c.getInternalExpire(duration)
	oldDuration, err = c.data.UpdateExpire(key, newExpireTime)
	if err != nil {
		return
	}
	if oldDuration != -1 {
		c.eventList.PushBack(&adapterMemoryEvent{
			k: key,
			e: newExpireTime,
		})
	}
	return
}

// Size 返回缓存的大小。
func (c *AdapterMemory) Size(ctx context.Context) (size int, err error) {
	return c.data.Size()
}

// Data 返回缓存中所有键值对的副本，类型为 map。
func (c *AdapterMemory) Data(ctx context.Context) (map[interface{}]interface{}, error) {
	return c.data.Data()
}

// Keys 返回缓存中的所有键作为切片。
func (c *AdapterMemory) Keys(ctx context.Context) ([]interface{}, error) {
	return c.data.Keys()
}

// Values 返回缓存中的所有值作为一个切片。
func (c *AdapterMemory) Values(ctx context.Context) ([]interface{}, error) {
	return c.data.Values()
}

// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
func (c *AdapterMemory) Clear(ctx context.Context) error {
	return c.data.Clear()
}

// Close 关闭缓存。
func (c *AdapterMemory) Close(ctx context.Context) error {
	if c.cap > 0 {
		c.lru.Close()
	}
	c.closed.Set(true)
	return nil
}

// doSetWithLockCheck 函数用于在缓存中设置键值对 `key`-`value`，如果 `key` 不存在于缓存中且设置了过期时间 `duration`。
//
// 如果 `duration` 等于0，则不会设置过期时间。
// 参数 `value` 可以是 <func() interface{}> 类型，但如果函数结果为 nil，则不做任何操作。
//
// 在设置缓存前，它会通过互斥写锁对 `key` 是否存在于缓存进行双重检查。
func (c *AdapterMemory) doSetWithLockCheck(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	expireTimestamp := c.getInternalExpire(duration)
	v, err := c.data.SetWithLock(ctx, key, value, expireTimestamp)
	c.eventList.PushBack(&adapterMemoryEvent{k: key, e: expireTimestamp})
	return gvar.New(v), err
}

// getInternalExpire 将给定的以毫秒为单位的过期时长转换并返回其对应的过期时间。
func (c *AdapterMemory) getInternalExpire(duration time.Duration) int64 {
	if duration == 0 {
		return defaultMaxExpire
	}
	return gtime.TimestampMilli() + duration.Nanoseconds()/1000000
}

// makeExpireKey 将以毫秒为单位的 `expire` 分组到相应的秒数。
func (c *AdapterMemory) makeExpireKey(expire int64) int64 {
	return int64(math.Ceil(float64(expire/1000)+1) * 1000)
}

// syncEventAndClearExpired 执行异步任务循环:
// 1. 异步处理事件列表中的数据，
// 并将处理结果同步到 `expireTimes` 和 `expireSets` 属性上。
// 2. 清理已过期的键值对数据。
func (c *AdapterMemory) syncEventAndClearExpired(ctx context.Context) {
	if c.closed.Val() {
		gtimer.Exit()
		return
	}
	var (
		event         *adapterMemoryEvent
		oldExpireTime int64
		newExpireTime int64
	)
// ========================================
// 数据同步.
// ========================================
	for {
		v := c.eventList.PopFront()
		if v == nil {
			break
		}
		event = v.(*adapterMemoryEvent)
		// 获取旧的过期集合。
		oldExpireTime = c.expireTimes.Get(event.k)
		// 计算新设置的过期时间
		newExpireTime = c.makeExpireKey(event.e)
		if newExpireTime != oldExpireTime {
			c.expireSets.GetOrNew(newExpireTime).Add(event.k)
			if oldExpireTime != 0 {
				c.expireSets.GetOrNew(oldExpireTime).Remove(event.k)
			}
			// 更新<event.k>的过期时间。
			c.expireTimes.Set(event.k, newExpireTime)
		}
		// 通过写操作将键添加到LRU历史记录中。
		if c.cap > 0 {
			c.lru.Push(event.k)
		}
	}
	// 处理LRU中已过期的键。
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
// ==================================
// 数据清理
// ==================================
	var (
		expireSet *gset.Set
		ek        = c.makeExpireKey(gtime.TimestampMilli())
		eks       = []int64{ek - 1000, ek - 2000, ek - 3000, ek - 4000, ek - 5000}
	)
	for _, expireTime := range eks {
		if expireSet = c.expireSets.Get(expireTime); expireSet != nil {
			// 遍历集合以删除其中的所有键。
			expireSet.Iterator(func(key interface{}) bool {
				c.clearByKey(key)
				return true
			})
			// 在其所有键都被删除后，删除该集合。
			c.expireSets.Delete(expireTime)
		}
	}
}

// clearByKey 通过给定的 `key` 删除键值对。
// 参数 `force` 指定是否强制执行此删除操作。
func (c *AdapterMemory) clearByKey(key interface{}, force ...bool) {
	// 在真正从缓存中删除之前进行双重检查。
	c.data.DeleteWithDoubleCheck(key, force...)

	// 从`expireTimes`中删除其过期时间。
	c.expireTimes.Delete(key)

	// 从LRU中删除它
	if c.cap > 0 {
		c.lru.Remove(key)
	}
}
