// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"context"
	"sync"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
)

type adapterMemoryData struct {
	mu   sync.RWMutex                      // dataMu 确保底层数据映射的并发安全性。 md5:ddcd414a151f3cf2
	data map[interface{}]adapterMemoryItem // data 是底层的缓存数据，它存储在一个哈希表中。 md5:7cfaf636328aa0e7
}

func newAdapterMemoryData() *adapterMemoryData {
	return &adapterMemoryData{
		data: make(map[interface{}]adapterMemoryItem),
	}
}

// Update 更新`key`的值，不改变其过期时间，并返回旧的值。
// 如果`key`在缓存中不存在，返回的值`exist`为false。
//
// 如果给定的`value`为nil，它会删除`key`。
// 如果`key`不在缓存中，它不会做任何操作。
// md5:6d92816db5b1d3bd
func (d *adapterMemoryData) Update(key interface{}, value interface{}) (oldValue interface{}, exist bool, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if item, ok := d.data[key]; ok {
		d.data[key] = adapterMemoryItem{
			v: value,
			e: item.e,
		}
		return item.v, true, nil
	}
	return nil, false, nil
}

// UpdateExpire 更新键`key`的过期时间，并返回旧的过期持续时间值。
//
// 如果`key`在缓存中不存在，它将返回-1并什么都不做。如果`duration`小于0，它会删除`key`。
// md5:b974907dd46b44be
func (d *adapterMemoryData) UpdateExpire(key interface{}, expireTime int64) (oldDuration time.Duration, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if item, ok := d.data[key]; ok {
		d.data[key] = adapterMemoryItem{
			v: item.v,
			e: expireTime,
		}
		return time.Duration(item.e-gtime.TimestampMilli()) * time.Millisecond, nil
	}
	return -1, nil
}

// Remove 从缓存中删除一个或多个键，并返回其值。
// 如果给出了多个键，它将返回最后删除项的值。
// md5:b3f23906b769df08
func (d *adapterMemoryData) Remove(keys ...interface{}) (removedKeys []interface{}, value interface{}, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	removedKeys = make([]interface{}, 0)
	for _, key := range keys {
		item, ok := d.data[key]
		if ok {
			value = item.v
			delete(d.data, key)
			removedKeys = append(removedKeys, key)
		}
	}
	return removedKeys, value, nil
}

// Data 返回一个缓存中所有键值对的副本，以映射类型表示。 md5:d88afdf7cfc66604
func (d *adapterMemoryData) Data() (map[interface{}]interface{}, error) {
	d.mu.RLock()
	m := make(map[interface{}]interface{}, len(d.data))
	for k, v := range d.data {
		if !v.IsExpired() {
			m[k] = v.v
		}
	}
	d.mu.RUnlock()
	return m, nil
}

// Keys 返回缓存中所有键的切片。 md5:7ebd9dba01282dc2
func (d *adapterMemoryData) Keys() ([]interface{}, error) {
	d.mu.RLock()
	var (
		index = 0
		keys  = make([]interface{}, len(d.data))
	)
	for k, v := range d.data {
		if !v.IsExpired() {
			keys[index] = k
			index++
		}
	}
	d.mu.RUnlock()
	return keys, nil
}

// Values 返回缓存中所有的值作为切片。 md5:dc00b32eb8913e9b
func (d *adapterMemoryData) Values() ([]interface{}, error) {
	d.mu.RLock()
	var (
		index  = 0
		values = make([]interface{}, len(d.data))
	)
	for _, v := range d.data {
		if !v.IsExpired() {
			values[index] = v.v
			index++
		}
	}
	d.mu.RUnlock()
	return values, nil
}

// Size 返回缓存的大小。 md5:c939a4ed87cd79ce
func (d *adapterMemoryData) Size() (size int, err error) {
	d.mu.RLock()
	size = len(d.data)
	d.mu.RUnlock()
	return size, nil
}

// Clear 清空缓存中的所有数据。
// 注意，此函数涉及敏感操作，应谨慎使用。
// md5:9212cab88870d3df
func (d *adapterMemoryData) Clear() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data = make(map[interface{}]adapterMemoryItem)
	return nil
}

func (d *adapterMemoryData) Get(key interface{}) (item adapterMemoryItem, ok bool) {
	d.mu.RLock()
	item, ok = d.data[key]
	d.mu.RUnlock()
	return
}

func (d *adapterMemoryData) Set(key interface{}, value adapterMemoryItem) {
	d.mu.Lock()
	d.data[key] = value
	d.mu.Unlock()
}

// SetMap 通过 `data` 批量设置缓存键值对，这些缓存在 `duration` 时间后过期。
//
// 如果 `duration` == 0，则不设置过期时间。
// 如果 `duration` < 0 或者给定的 `value` 为 nil，则删除 `data` 中的键。
// md5:cc6156a6df071b21
func (d *adapterMemoryData) SetMap(data map[interface{}]interface{}, expireTime int64) error {
	d.mu.Lock()
	for k, v := range data {
		d.data[k] = adapterMemoryItem{
			v: v,
			e: expireTime,
		}
	}
	d.mu.Unlock()
	return nil
}

func (d *adapterMemoryData) SetWithLock(ctx context.Context, key interface{}, value interface{}, expireTimestamp int64) (interface{}, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var (
		err error
	)
	if v, ok := d.data[key]; ok && !v.IsExpired() {
		return v.v, nil
	}
	f, ok := value.(Func)
	if !ok {
				// 与原始函数值兼容。 md5:b6980bd817389e7f
		f, ok = value.(func(ctx context.Context) (value interface{}, err error))
	}
	if ok {
		if value, err = f(ctx); err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
	}
	d.data[key] = adapterMemoryItem{v: value, e: expireTimestamp}
	return value, nil
}

func (d *adapterMemoryData) DeleteWithDoubleCheck(key interface{}, force ...bool) {
	d.mu.Lock()
		// 在从缓存中真正删除之前，再双检查一次。 md5:53767fc86cbfbf5e
	if item, ok := d.data[key]; (ok && item.IsExpired()) || (len(force) > 0 && force[0]) {
		delete(d.data, key)
	}
	d.mu.Unlock()
}
