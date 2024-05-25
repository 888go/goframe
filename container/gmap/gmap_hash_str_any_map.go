// 版权归GoFrame作者(https://goframe.org)所有。
//
// 本源代码遵循MIT许可证条款。
// 如果gm文件未随附MIT许可证的副本，
// 您可以在https://github.com/gogf/gf获取一个。
// md5:c99fd05f11d37c36

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

// StrAnyMap 实现了一个使用 RWMutex（读写互斥锁）并带有开关功能的 map[string]interface{}。. md5:1389fe7e3914ba43
type StrAnyMap struct {
	mu   rwmutex.RWMutex
	data map[string]interface{}
}

// NewStrAnyMap 返回一个空的 StrAnyMap 对象。
// 参数 `safe` 用于指定是否在并发安全环境下使用该映射，其默认值为 false。
// md5:cbe007353b41d97e
func NewStrAnyMap(safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[string]interface{}),
	}
}

// NewStrAnyMapFrom 从给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（不进行深拷贝），在外部修改映射时可能会存在并发安全问题。
// md5:3f15992f776bd362
func NewStrAnyMapFrom(data map[string]interface{}, safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		mu:   rwmutex.Create(safe...),
		data: data,
	}
}

// Iterator 使用自定义回调函数 `f` 读取只读哈希映射。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:52d024b320a69c3b
func (m *StrAnyMap) Iterator(f func(k string, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Clone 返回一个新的哈希映射，其中包含当前映射数据的副本。. md5:b9264f3636ead08a
func (m *StrAnyMap) Clone() *StrAnyMap {
	return NewStrAnyMapFrom(m.MapCopy(), m.mu.IsSafe())
}

// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
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

// MapStrAny将映射的底层数据复制为map[string]interface{}。. md5:46db5a1110397522
func (m *StrAnyMap) MapStrAny() map[string]interface{} {
	return m.Map()
}

// MapCopy 返回哈希映射底层数据的一个副本。. md5:46f762167d5821b1
func (m *StrAnyMap) MapCopy() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// FilterEmpty 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
func (m *StrAnyMap) FilterEmpty() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// FilterNil 删除所有值为 nil 的键值对。. md5:3c964818401771a4
func (m *StrAnyMap) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Set 将键值对设置到哈希映射中。. md5:07ea2dd1ea28820a
func (m *StrAnyMap) Set(key string, val interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	m.data[key] = val
	m.mu.Unlock()
}

// 将键值对设置到哈希映射中。. md5:e3f3f8a1b69eb832
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

// Search 在给定的`key`下搜索映射。
// 第二个返回参数`found`如果找到键，则为true，否则为false。
// md5:99336de9941a3b02
func (m *StrAnyMap) Search(key string) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Get 根据给定的 `key` 获取值。. md5:2b744a3e455aadfb
func (m *StrAnyMap) Get(key string) (value interface{}) {
	m.mu.RLock()
	if m.data != nil {
		value = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Pop 从映射中获取并删除一个元素。. md5:2d364ca2b6054111
func (m *StrAnyMap) Pop() (key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, value = range m.data {
		delete(m.data, key)
		return
	}
	return
}

// Pops 从映射中检索并删除 `size` 个项目。
// 如果 size 等于 -1，则返回所有项目。
// md5:0f2cdbc0238fdc37
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

// doSetWithLockCheck 会使用 mutex.Lock 检查给定键的值是否存在。
// 如果不存在，将使用给定的 `key` 将值设置到映射中；否则，直接返回已存在的值。
// 
// 当设置值时，如果 `value` 类型为 `func() interface{}`，它将在映射的 mutex.Lock 保护下执行，
// 并将返回值设置到映射中，键为 `key`。
// 
// 它返回给定 `key` 的值。
// md5:60f1f50efa66e173
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

// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (m *StrAnyMap) GetOrSet(key string, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键获取值，
// 如果键不存在，则使用回调函数`f`的返回值设置值，
// 并返回这个设置的值。
// md5:f584dd7547dfbcc0
func (m *StrAnyMap) GetOrSetFunc(key string, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键获取值，
// 如果不存在，它将使用回调函数 `f` 的返回值设置该值，然后返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的 mutex。
// md5:d32fdee586d84dde
func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar通过给定的`key`返回一个Var。返回的Var是非并发安全的。
// md5:debfb1b2bd13312b
func (m *StrAnyMap) GetVar(key string) *gvar.Var {
	return gvar.New(m.Get(key))
}

// GetVarOrSet 返回一个 Var，其中包含从 GetVarOrSet 获取的结果。
// 返回的 Var 是非并发安全的。
// md5:c3730f368b7f62b5
func (m *StrAnyMap) GetVarOrSet(key string, value interface{}) *gvar.Var {
	return gvar.New(m.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个Var，其结果来自GetOrSetFunc。
// 返回的Var不具备并发安全性。
// md5:7d7674129b73ead1
func (m *StrAnyMap) GetVarOrSetFunc(key string, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个从 GetOrSetFuncLock 获得结果的 Var。返回的 Var 不是线程安全的。
// md5:bdab644d14c89234
func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (m *StrAnyMap) SetIfNotExist(key string, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，则返回 false，`value` 将被忽略。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
// md5:a6ee84b157328f61
func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// 通过键删除map中的批删除值。. md5:57081208d84ca7e8
func (m *StrAnyMap) Removes(keys []string) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Remove 通过给定的`key`从map中删除值，并返回被删除的值。. md5:5ee6dc9be17b4ab8
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

// Keys 返回映射中所有键的切片。. md5:425640fff4178659
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

// Values 将地图中的所有值返回为一个切片。. md5:a89b5b485c966abd
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

// Contains 检查键是否存在。
// 如果键存在，它返回 true，否则返回 false。
// md5:d8fb22313aadd65f
func (m *StrAnyMap) Contains(key string) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return ok
}

// Size返回映射的大小。. md5:da42fb3955847483
func (m *StrAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
func (m *StrAnyMap) IsEmpty() bool {
	return m.Size() == 0
}

// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。. md5:0553a5cd54a22f3c
func (m *StrAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]interface{})
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。. md5:a84ecf2839212d81
func (m *StrAnyMap) Replace(data map[string]interface{}) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。. md5:e73dbc0381ebb3dc
func (m *StrAnyMap) LockFunc(f func(m map[string]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。. md5:4ae51d9b7445f043
func (m *StrAnyMap) RLockFunc(f func(m map[string]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Flip 将映射的键值对交换为值键。. md5:dbcb578f1b30fa01
func (m *StrAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		n[gconv.String(v)] = k
	}
	m.data = n
}

// Merge 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
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

// String 将地图转换为字符串形式并返回。. md5:6473318e71d3dfd0
func (m *StrAnyMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (m StrAnyMap) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return json.Marshal(m.data)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
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

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。. md5:6f3087a6f7df5477
func (m *StrAnyMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = gconv.Map(value)
	return
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
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

// IsSubOf 检查当前映射是否是`other`的子映射。. md5:9a6c60859c5a0fbc
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

// Diff 函数比较当前地图 `m` 与地图 `other` 并返回它们不同的键。
// 返回的 `addedKeys` 是存在于地图 `m` 中但不在地图 `other` 中的键。
// 返回的 `removedKeys` 是存在于地图 `other` 中但不在地图 `m` 中的键。
// 返回的 `updatedKeys` 是同时存在于地图 `m` 和 `other` 中，但其值不相等（`!=`）的键。
// md5:d3bf0bf8c70e9093
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
