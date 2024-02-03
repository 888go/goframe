// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"context"
	"sync"
	"time"
	
	"github.com/888go/goframe/os/gtime"
)

type adapterMemoryData struct {
	mu   sync.RWMutex                      // dataMu 用于确保底层数据映射的并发安全性。
	data map[interface{}]adapterMemoryItem // data 是底层缓存数据，存储在一个哈希表中。
}

func newAdapterMemoryData() *adapterMemoryData {
	return &adapterMemoryData{
		data: make(map[interface{}]adapterMemoryItem),
	}
}

// Update 更新`key`的值，但不改变其过期时间，并返回旧值。
// 返回的布尔值`exist`，如果`key`在缓存中不存在，则为false。
//
// 如果给定的`value`为nil，则删除`key`。
// 若`key`在缓存中不存在，则不做任何操作。
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

// UpdateExpire 更新键 `key` 的过期时间，并返回旧的过期持续时长值。
//
// 若 `key` 不存在于缓存中，则返回 -1 并不做任何操作。
// 若 `duration` 小于 0，则删除 `key`。
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

// Remove 从缓存中删除一个或多个键，并返回其对应的值。
// 如果提供了多个键，它将返回被删除的最后一个项目的值。
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

// Data 返回缓存中所有键值对的副本，类型为 map。
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

// Keys 返回缓存中的所有键作为切片。
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

// Values 返回缓存中的所有值作为一个切片。
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

// Size 返回缓存的大小。
func (d *adapterMemoryData) Size() (size int, err error) {
	d.mu.RLock()
	size = len(d.data)
	d.mu.RUnlock()
	return size, nil
}

// Clear 清除缓存中的所有数据。
// 注意：此函数较为敏感，应谨慎使用。
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

// SetMap 通过 `data` 批量设置缓存键值对，缓存将在 `duration` 后过期。
//
// 如果 `duration` == 0，则表示永不过期。
// 若 `duration` < 0 或提供的 `value` 为 nil，则删除 `data` 中的键。
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
		// 与原始函数值兼容。
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
	// 在真正从缓存中删除之前进行双重检查。
	if item, ok := d.data[key]; (ok && item.IsExpired()) || (len(force) > 0 && force[0]) {
		delete(d.data, key)
	}
	d.mu.Unlock()
}
