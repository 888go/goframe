// 版权归GoFrame作者(https://goframe.org)所有。
//
// 本源代码遵循MIT许可证条款。
// 如果gm文件未随附MIT许可证的副本，
// 您可以在https://github.com/gogf/gf获取一个。
// md5:c99fd05f11d37c36

package gmap

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

// IntAnyMap 实现了具有开关的RWMutex接口的[int]interface{}映射。 md5:91c3a42c4b361517
type IntAnyMap struct {
	mu   rwmutex.RWMutex
	data map[int]interface{}
}

// NewIntAnyMap 返回一个空的 IntAnyMap 对象。
// 参数 `safe` 用于指定是否使用并发安全的 map，默认为 false。
// md5:1939fd81edf46e9b
// ff:创建IntAny
// safe:并发安全
func NewIntAnyMap(safe ...bool) *IntAnyMap {
	return &IntAnyMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[int]interface{}),
	}
}

// NewIntAnyMapFrom 根据给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（非深度复制），
// 因此，在外部修改该映射时可能会存在一些并发安全问题。
// md5:69d7f3d651b3336d
// ff:创建IntAny并从Map
// data:map值
// safe:并发安全
func NewIntAnyMapFrom(data map[int]interface{}, safe ...bool) *IntAnyMap {
	return &IntAnyMap{
		mu:   rwmutex.Create(safe...),
		data: data,
	}
}

// Iterator 使用自定义回调函数 `f` 读取只读哈希映射。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:52d024b320a69c3b
// yx:true
// ff:X遍历
// m:
// f:
// k:
// v:
func (m *IntAnyMap) Iterator(f func(k int, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Clone 返回一个新的哈希映射，其中包含当前映射数据的副本。 md5:b9264f3636ead08a
// ff:取副本
// m:
func (m *IntAnyMap) Clone() *IntAnyMap {
	return NewIntAnyMapFrom(m.MapCopy(), m.mu.IsSafe())
}

// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
// ff:取Map
// m:
func (m *IntAnyMap) Map() map[int]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[int]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny将映射的底层数据复制为map[string]interface{}。 md5:46db5a1110397522
// yx:true
// ff:取MapStrAny
// m:
func (m *IntAnyMap) MapStrAny() map[string]interface{} {
	m.mu.RLock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[gconv.String(k)] = v
	}
	m.mu.RUnlock()
	return data
}

// MapCopy 返回哈希映射底层数据的一个副本。 md5:46f762167d5821b1
// ff:浅拷贝
// m:
func (m *IntAnyMap) MapCopy() map[int]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[int]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// FilterEmpty 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
// ff:删除所有空值
// m:
func (m *IntAnyMap) FilterEmpty() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// FilterNil 删除所有值为 nil 的键值对。 md5:3c964818401771a4
// ff:删除所有nil值
// m:
func (m *IntAnyMap) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Set 将键值对设置到哈希映射中。 md5:07ea2dd1ea28820a
// yx:true
// ff:设置值
// m:
// key:
// val:
func (m *IntAnyMap) Set(key int, val interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[int]interface{})
	}
	m.data[key] = val
	m.mu.Unlock()
}

// 将键值对设置到哈希映射中。 md5:e3f3f8a1b69eb832
// ff:设置值Map
// m:
// data:map值
func (m *IntAnyMap) Sets(data map[int]interface{}) {
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
// ff:查找
// m:
// key:名称
// value:
// found:
func (m *IntAnyMap) Search(key int) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Get 根据给定的 `key` 获取值。 md5:2b744a3e455aadfb
// ff:取值
// m:
// key:名称
// value:
func (m *IntAnyMap) Get(key int) (value interface{}) {
	m.mu.RLock()
	if m.data != nil {
		value = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Pop 从映射中获取并删除一个元素。 md5:2d364ca2b6054111
// ff:出栈
// m:
// key:名称
// value:
func (m *IntAnyMap) Pop() (key int, value interface{}) {
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
// ff:出栈多个
// m:
// size:数量
func (m *IntAnyMap) Pops(size int) map[int]interface{} {
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
		newMap = make(map[int]interface{}, size)
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
func (m *IntAnyMap) doSetWithLockCheck(key int, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]interface{})
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
// ff:取值或设置值
// m:
// key:名称
// value:
func (m *IntAnyMap) GetOrSet(key int, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc通过键获取值，
// 如果不存在，它将使用回调函数`f`返回的值设置该值，并返回这个值。
// md5:c4de9d0fac2a8916
// ff:取值或设置值_函数
// m:
// key:名称
// f:
func (m *IntAnyMap) GetOrSetFunc(key int, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键获取值，
// 如果不存在，则使用回调函数 `f` 的返回值设置该键的值，并返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的mutex。
// md5:ac8ad0e9416578ba
// ff:取值或设置值_函数带锁
// m:
// key:名称
// f:
func (m *IntAnyMap) GetOrSetFuncLock(key int, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar通过给定的`key`返回一个Var。返回的Var是非并发安全的。
// md5:debfb1b2bd13312b
// ff:取值泛型类
// m:
// key:名称
func (m *IntAnyMap) GetVar(key int) *gvar.Var {
	return gvar.New(m.Get(key))
}

// GetVarOrSet 返回一个 Var，其中包含从 GetVarOrSet 获取的结果。
// 返回的 Var 是非并发安全的。
// md5:c3730f368b7f62b5
// ff:取值或设置值泛型类
// m:
// key:名称
// value:
func (m *IntAnyMap) GetVarOrSet(key int, value interface{}) *gvar.Var {
	return gvar.New(m.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个Var，其结果来自GetOrSetFunc。
// 返回的Var不具备并发安全性。
// md5:7d7674129b73ead1
// ff:取值或设置值泛型类_函数
// m:
// key:名称
// f:
func (m *IntAnyMap) GetVarOrSetFunc(key int, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个从 GetOrSetFuncLock 获得结果的 Var。返回的 Var 不是线程安全的。
// md5:bdab644d14c89234
// ff:取值或设置值泛型类_函数带锁
// m:
// key:名称
// f:
func (m *IntAnyMap) GetVarOrSetFuncLock(key int, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
// ff:设置值并跳过已存在
// m:
// key:名称
// value:
func (m *IntAnyMap) SetIfNotExist(key int, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
// ff:设置值并跳过已存在_函数
// m:
// key:名称
// f:
func (m *IntAnyMap) SetIfNotExistFunc(key int, f func() interface{}) bool {
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
// ff:设置值并跳过已存在_函数带锁
// m:
// key:名称
// f:
func (m *IntAnyMap) SetIfNotExistFuncLock(key int, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// 通过键删除map中的批删除值。 md5:57081208d84ca7e8
// ff:删除多个值
// m:
// keys:名称
func (m *IntAnyMap) Removes(keys []int) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Remove 通过给定的`key`从map中删除值，并返回被删除的值。 md5:5ee6dc9be17b4ab8
// ff:删除
// m:
// key:名称
// value:
func (m *IntAnyMap) Remove(key int) (value interface{}) {
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

// Keys 返回映射中所有键的切片。 md5:425640fff4178659
// ff:取所有名称
// m:
func (m *IntAnyMap) Keys() []int {
	m.mu.RLock()
	var (
		keys  = make([]int, len(m.data))
		index = 0
	)
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

// Values 将地图中的所有值返回为一个切片。 md5:a89b5b485c966abd
// ff:取所有值
// m:
func (m *IntAnyMap) Values() []interface{} {
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
// ff:是否存在
// m:
// key:名称
func (m *IntAnyMap) Contains(key int) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return ok
}

// Size返回映射的大小。 md5:da42fb3955847483
// ff:取数量
// m:
func (m *IntAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
// ff:是否为空
// m:
func (m *IntAnyMap) IsEmpty() bool {
	return m.Size() == 0
}

// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。 md5:0553a5cd54a22f3c
// ff:清空
// m:
func (m *IntAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[int]interface{})
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。 md5:a84ecf2839212d81
// ff:替换
// m:
// data:map值
func (m *IntAnyMap) Replace(data map[int]interface{}) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
// ff:遍历写锁定
// m:
// f:回调函数
// m:
func (m *IntAnyMap) LockFunc(f func(m map[int]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
// ff:遍历读锁定
// m:
// f:回调函数
// m:
func (m *IntAnyMap) RLockFunc(f func(m map[int]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Flip 将映射的键值对交换为值键。 md5:dbcb578f1b30fa01
// ff:名称值交换
// m:
func (m *IntAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[int]interface{}, len(m.data))
	for k, v := range m.data {
		n[gconv.Int(v)] = k
	}
	m.data = n
}

// Merge 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
// ff:合并
// m:
// other:map值
func (m *IntAnyMap) Merge(other *IntAnyMap) {
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

// String 将地图转换为字符串形式并返回。 md5:6473318e71d3dfd0
// ff:
// m:
func (m *IntAnyMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
// ff:
// m:
func (m IntAnyMap) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return json.Marshal(m.data)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
// ff:
// m:
// b:
func (m *IntAnyMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]interface{})
	}
	if err := json.UnmarshalUseNumber(b, &m.data); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。 md5:6f3087a6f7df5477
// ff:
// m:
// value:
// err:
func (m *IntAnyMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]interface{})
	}
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.Bytes(value), &m.data)
	default:
		for k, v := range gconv.Map(value) {
			m.data[gconv.Int(k)] = v
		}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
// ff:
// m:
func (m *IntAnyMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[int]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = deepcopy.Copy(v)
	}
	return NewIntAnyMapFrom(data, m.mu.IsSafe())
}

// IsSubOf 检查当前映射是否是`other`的子映射。 md5:9a6c60859c5a0fbc
// ff:是否为子集
// m:
// other:父集Map
func (m *IntAnyMap) IsSubOf(other *IntAnyMap) bool {
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
// ff:比较
// m:
// other:map值
// addedKeys:
// removedKeys:
// updatedKeys:
func (m *IntAnyMap) Diff(other *IntAnyMap) (addedKeys, removedKeys, updatedKeys []int) {
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
