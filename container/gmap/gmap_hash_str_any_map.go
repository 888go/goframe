// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
//

package gmap

import (
	"reflect"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

// StrAnyMap implements map[string]interface{} with RWMutex that has switch.
type StrAnyMap struct {
	mu   rwmutex.RWMutex
	data map[string]interface{}
}

// NewStrAnyMap returns an empty StrAnyMap object.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
// ff:创建StrAny
// safe:并发安全
func NewStrAnyMap(safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[string]interface{}),
	}
}

// NewStrAnyMapFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// ff:创建AnyStr并从Map
// data:map值
// safe:并发安全
func NewStrAnyMapFrom(data map[string]interface{}, safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		mu:   rwmutex.Create(safe...),
		data: data,
	}
}

// Iterator iterates the hash map readonly with custom callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
// yx:true
// ff:X遍历
// m:
// f:
// k:
// v:
func (m *StrAnyMap) Iterator(f func(k string, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Clone returns a new hash map with copy of current map data.
// ff:取副本
// m:
func (m *StrAnyMap) Clone() *StrAnyMap {
	return NewStrAnyMapFrom(m.MapCopy(), m.mu.IsSafe())
}

// Map returns the underlying data map.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
// ff:取Map
// m:
func (m *StrAnyMap) Map() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
// yx:true
// ff:取MapStrAny
// m:
func (m *StrAnyMap) MapStrAny() map[string]interface{} {
	return m.Map()
}

// MapCopy returns a copy of the underlying data of the hash map.
// ff:浅拷贝
// m:
func (m *StrAnyMap) MapCopy() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// FilterEmpty deletes all key-value pair of which the value is empty.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
// ff:删除所有空值
// m:
func (m *StrAnyMap) FilterEmpty() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// FilterNil deletes all key-value pair of which the value is nil.
// ff:删除所有nil值
// m:
func (m *StrAnyMap) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Set sets key-value to the hash map.
// yx:true
// ff:设置值
// m:
// key:
// val:
func (m *StrAnyMap) Set(key string, val interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	m.data[key] = val
	m.mu.Unlock()
}

// Sets batch sets key-values to the hash map.
// ff:设置值Map
// m:
// data:map值
func (m *StrAnyMap) Sets(data map[string]interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = data
	} else {
		for k, v := range data {
			m.data[k] = v
		}
	}
	m.mu.Unlock()
}

// Search searches the map with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
// ff:查找
// m:
// key:名称
// value:
// found:
func (m *StrAnyMap) Search(key string) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Get returns the value by given `key`.
// ff:取值
// m:
// key:名称
// value:
func (m *StrAnyMap) Get(key string) (value interface{}) {
	m.mu.RLock()
	if m.data != nil {
		value = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Pop retrieves and deletes an item from the map.
// ff:出栈
// m:
// key:名称
// value:
func (m *StrAnyMap) Pop() (key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, value = range m.data {
		delete(m.data, key)
		return
	}
	return
}

// Pops retrieves and deletes `size` items from the map.
// It returns all items if size == -1.
// ff:出栈多个
// m:
// size:数量
func (m *StrAnyMap) Pops(size int) map[string]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if size > len(m.data) || size == -1 {
		size = len(m.data)
	}
	if size == 0 {
		return nil
	}
	var (
		index  = 0
		newMap = make(map[string]interface{}, size)
	)
	for k, v := range m.data {
		delete(m.data, k)
		newMap[k] = v
		index++
		if index == size {
			break
		}
	}
	return newMap
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given `key`,
// or else just return the existing value.
//
// When setting value, if `value` is type of `func() interface {}`,
// it will be executed with mutex.Lock of the hash map,
// and its return value will be set to the map with `key`.
//
// It returns value with given `key`.
func (m *StrAnyMap) doSetWithLockCheck(key string, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		m.data[key] = value
	}
	return value
}

// GetOrSet returns the value by key,
// or sets value with given `value` if it does not exist and then returns this value.
// ff:取值或设置值
// m:
// key:名称
// value:
func (m *StrAnyMap) GetOrSet(key string, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
// ff:取值或设置值_函数
// m:
// key:名称
// f:
func (m *StrAnyMap) GetOrSetFunc(key string, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.
// ff:取值或设置值_函数带锁
// m:
// key:名称
// f:
func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar returns a Var with the value by given `key`.
// The returned Var is un-concurrent safe.
// ff:取值泛型类
// m:
// key:名称
func (m *StrAnyMap) GetVar(key string) *gvar.Var {
	return gvar.New(m.Get(key))
}

// GetVarOrSet returns a Var with result from GetVarOrSet.
// The returned Var is un-concurrent safe.
// ff:取值或设置值泛型类
// m:
// key:名称
// value:
func (m *StrAnyMap) GetVarOrSet(key string, value interface{}) *gvar.Var {
	return gvar.New(m.GetOrSet(key, value))
}

// GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
// The returned Var is un-concurrent safe.
// ff:取值或设置值泛型类_函数
// m:
// key:名称
// f:
func (m *StrAnyMap) GetVarOrSetFunc(key string, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
// The returned Var is un-concurrent safe.
// ff:取值或设置值泛型类_函数带锁
// m:
// key:名称
// f:
func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFuncLock(key, f))
}

// SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
// ff:设置值并跳过已存在
// m:
// key:名称
// value:
func (m *StrAnyMap) SetIfNotExist(key string, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
// ff:设置值并跳过已存在_函数
// m:
// key:名称
// f:
func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
//
// SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that
// it executes function `f` with mutex.Lock of the hash map.
// ff:设置值并跳过已存在_函数带锁
// m:
// key:名称
// f:
func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Removes batch deletes values of the map by keys.
// ff:删除多个值
// m:
// keys:名称
func (m *StrAnyMap) Removes(keys []string) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Remove deletes value from map by given `key`, and return this deleted value.
// ff:删除
// m:
// key:名称
// value:
func (m *StrAnyMap) Remove(key string) (value interface{}) {
	m.mu.Lock()
	if m.data != nil {
		var ok bool
		if value, ok = m.data[key]; ok {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
	return
}

// Keys returns all keys of the map as a slice.
// ff:取所有名称
// m:
func (m *StrAnyMap) Keys() []string {
	m.mu.RLock()
	var (
		keys  = make([]string, len(m.data))
		index = 0
	)
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

// Values returns all values of the map as a slice.
// ff:取所有值
// m:
func (m *StrAnyMap) Values() []interface{} {
	m.mu.RLock()
	var (
		values = make([]interface{}, len(m.data))
		index  = 0
	)
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

// Contains checks whether a key exists.
// It returns true if the `key` exists, or else false.
// ff:是否存在
// m:
// key:名称
func (m *StrAnyMap) Contains(key string) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return ok
}

// Size returns the size of the map.
// ff:取数量
// m:
func (m *StrAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
// ff:是否为空
// m:
func (m *StrAnyMap) IsEmpty() bool {
	return m.Size() == 0
}

// Clear deletes all data of the map, it will remake a new underlying data map.
// ff:清空
// m:
func (m *StrAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]interface{})
	m.mu.Unlock()
}

// Replace the data of the map with given `data`.
// ff:替换
// m:
// data:map值
func (m *StrAnyMap) Replace(data map[string]interface{}) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
// ff:遍历写锁定
// m:
// f:回调函数
// m:
func (m *StrAnyMap) LockFunc(f func(m map[string]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
// ff:遍历读锁定
// m:
// f:回调函数
// m:
func (m *StrAnyMap) RLockFunc(f func(m map[string]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Flip exchanges key-value of the map to value-key.
// ff:名称值交换
// m:
func (m *StrAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		n[gconv.String(v)] = k
	}
	m.data = n
}

// Merge merges two hash maps.
// The `other` map will be merged into the map `m`.
// ff:合并
// m:
// other:map值
func (m *StrAnyMap) Merge(other *StrAnyMap) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = other.MapCopy()
		return
	}
	if other != m {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	for k, v := range other.data {
		m.data[k] = v
	}
}

// String returns the map as a string.
// ff:
// m:
func (m *StrAnyMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// ff:
// m:
func (m StrAnyMap) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
// ff:
// m:
// b:
func (m *StrAnyMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	if err := json.UnmarshalUseNumber(b, &m.data); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for map.
// ff:
// m:
// value:
// err:
func (m *StrAnyMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = gconv.Map(value)
	return
}

// DeepCopy implements interface for deep copy of current type.
// ff:
// m:
func (m *StrAnyMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = deepcopy.Copy(v)
	}
	return NewStrAnyMapFrom(data, m.mu.IsSafe())
}

// IsSubOf checks whether the current map is a sub-map of `other`.
// ff:是否为子集
// m:
// other:父集Map
func (m *StrAnyMap) IsSubOf(other *StrAnyMap) bool {
	if m == other {
		return true
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key, value := range m.data {
		otherValue, ok := other.data[key]
		if !ok {
			return false
		}
		if otherValue != value {
			return false
		}
	}
	return true
}

// Diff compares current map `m` with map `other` and returns their different keys.
// The returned `addedKeys` are the keys that are in map `m` but not in map `other`.
// The returned `removedKeys` are the keys that are in map `other` but not in map `m`.
// The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).
// ff:比较
// m:
// other:map值
// addedKeys:
// removedKeys:
// updatedKeys:
func (m *StrAnyMap) Diff(other *StrAnyMap) (addedKeys, removedKeys, updatedKeys []string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	for key := range m.data {
		if _, ok := other.data[key]; !ok {
			removedKeys = append(removedKeys, key)
		} else if !reflect.DeepEqual(m.data[key], other.data[key]) {
			updatedKeys = append(updatedKeys, key)
		}
	}
	for key := range other.data {
		if _, ok := m.data[key]; !ok {
			addedKeys = append(addedKeys, key)
		}
	}
	return
}
