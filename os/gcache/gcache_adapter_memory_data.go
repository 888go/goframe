// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcache

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type adapterMemoryData struct {
	mu   sync.RWMutex                      // dataMu ensures the concurrent safety of underlying data map.
	data map[interface{}]adapterMemoryItem // data is the underlying cache data which is stored in a hash table.
}

func newAdapterMemoryData() *adapterMemoryData {
	return &adapterMemoryData{
		data: make(map[interface{}]adapterMemoryItem),
	}
}

// Update updates the value of `key` without changing its expiration and returns the old value.
// The returned value `exist` is false if the `key` does not exist in the cache.
//
// It deletes the `key` if given `value` is nil.
// It does nothing if `key` does not exist in the cache.

// ff:更新值
// err:错误
// exist:是否已存在
// oldValue:旧值
// value:值
// key:名称
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

// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
//
// It returns -1 and does nothing if the `key` does not exist in the cache.
// It deletes the `key` if `duration` < 0.

// ff:更新过期时间
// err:错误
// oldDuration:旧过期时长
// expireTime:时长
// key:名称
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

// Remove deletes the one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the deleted last item.

// ff:删除并带返回值
// err:错误
// value:值
// removedKeys:被删除名称
// keys:名称
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

// Data returns a copy of all key-value pairs in the cache as map type.

// ff:取所有键值Map副本
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

// Keys returns all keys in the cache as slice.

// ff:取所有键
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

// Values returns all values in the cache as slice.

// ff:取所有值
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

// Size returns the size of the cache.

// ff:取数量
// err:错误
// size:数量
func (d *adapterMemoryData) Size() (size int, err error) {
	d.mu.RLock()
	size = len(d.data)
	d.mu.RUnlock()
	return size, nil
}

// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.

// ff:清空
func (d *adapterMemoryData) Clear() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data = make(map[interface{}]adapterMemoryItem)
	return nil
}


// ff:取值
// ok:成功
// item:
// key:名称
func (d *adapterMemoryData) Get(key interface{}) (item adapterMemoryItem, ok bool) {
	d.mu.RLock()
	item, ok = d.data[key]
	d.mu.RUnlock()
	return
}


// ff:设置值
// yx:true
// value:
// key:
func (d *adapterMemoryData) Set(key interface{}, value adapterMemoryItem) {
	d.mu.Lock()
	d.data[key] = value
	d.mu.Unlock()
}

// SetMap batch sets cache with key-value pairs by `data`, which is expired after `duration`.
//
// It does not expire if `duration` == 0.
// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

// ff:设置Map
// expireTime:
// data:
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


// ff:
// expireTimestamp:
// value:
// key:
// ctx:
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
		// Compatible with raw function value.
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


// ff:
// force:
// key:
func (d *adapterMemoryData) DeleteWithDoubleCheck(key interface{}, force ...bool) {
	d.mu.Lock()
	// Doubly check before really deleting it from cache.
	if item, ok := d.data[key]; (ok && item.IsExpired()) || (len(force) > 0 && force[0]) {
		delete(d.data, key)
	}
	d.mu.Unlock()
}
