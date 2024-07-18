// 版权归GoFrame作者(https://goframe.org)所有。
//
// 本源代码遵循MIT许可证条款。
// 如果gm文件未随附MIT许可证的副本，
// 您可以在https://github.com/gogf/gf获取一个。
// md5:c99fd05f11d37c36

package gmap

import (
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

// StrIntMap 实现了具有开关的字符串到整数的 RWMutex 映射。 md5:9126e7c711d57c87
type StrIntMap struct {
	mu   rwmutex.RWMutex
	data map[string]int
}

// NewStrIntMap 返回一个空的 StrIntMap 对象。
// 参数 `safe` 用于指定是否使用并发安全的 map，默认为 false。
// md5:040c01e687e6526f
// ff:创建StrInt
// safe:并发安全
func NewStrIntMap(safe ...bool) *StrIntMap {
	return &StrIntMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[string]int),
	}
}

// NewStrIntMapFrom 根据给定的映射 `data` 创建并返回一个新的哈希映射。
// 请注意，参数 `data` 映射将被设置为底层数据映射（无深拷贝），
// 当在外部修改该映射时，可能存在一些并发安全问题。
// md5:def2dd5f3c6bbd06
// ff:创建StrInt并从Map
// data:map值
// safe:并发安全
func NewStrIntMapFrom(data map[string]int, safe ...bool) *StrIntMap {
	return &StrIntMap{
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
func (m *StrIntMap) Iterator(f func(k string, v int) bool) {
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
func (m *StrIntMap) Clone() *StrIntMap {
	return NewStrIntMapFrom(m.MapCopy(), m.mu.IsSafe())
}

// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
// ff:取Map
// m:
func (m *StrIntMap) Map() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[string]int, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny将映射的底层数据复制为map[string]interface{}。 md5:46db5a1110397522
// yx:true
// ff:取MapStrAny
// m:
func (m *StrIntMap) MapStrAny() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapCopy 返回哈希映射底层数据的一个副本。 md5:46f762167d5821b1
// ff:浅拷贝
// m:
func (m *StrIntMap) MapCopy() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]int, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// FilterEmpty 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
// ff:删除所有空值
// m:
func (m *StrIntMap) FilterEmpty() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// Set 将键值对设置到哈希映射中。 md5:07ea2dd1ea28820a
// yx:true
// ff:设置值
// m:
// key:
// val:
func (m *StrIntMap) Set(key string, val int) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = val
	m.mu.Unlock()
}

// 将键值对设置到哈希映射中。 md5:e3f3f8a1b69eb832
// ff:设置值Map
// m:
// data:map值
func (m *StrIntMap) Sets(data map[string]int) {
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
// value:值
// found:成功
func (m *StrIntMap) Search(key string) (value int, found bool) {
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
func (m *StrIntMap) Get(key string) (value int) {
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
func (m *StrIntMap) Pop() (key string, value int) {
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
func (m *StrIntMap) Pops(size int) map[string]int {
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
		newMap = make(map[string]int, size)
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

// doSetWithLockCheck 使用互斥锁(mutex.Lock)检查键的值是否存在，
// 如果不存在，则将给定的`value`设置到映射中指定的`key`处，
// 否则，直接返回已存在的值。
//
// 它返回与给定`key`关联的值。
// md5:3a2d1537d3fe7230
func (m *StrIntMap) doSetWithLockCheck(key string, value int) int {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[string]int)
	}
	if v, ok := m.data[key]; ok {
		m.mu.Unlock()
		return v
	}
	m.data[key] = value
	m.mu.Unlock()
	return value
}

// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
// ff:取值或设置值
// m:
// key:名称
// value:值
func (m *StrIntMap) GetOrSet(key string, value int) int {
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
// ff:取值或设置值_函数
// m:
// key:名称
// f:
func (m *StrIntMap) GetOrSetFunc(key string, f func() int) int {
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
// ff:取值或设置值_函数带锁
// m:
// key:名称
// f:
func (m *StrIntMap) GetOrSetFuncLock(key string, f func() int) int {
	if v, ok := m.Search(key); !ok {
		m.mu.Lock()
		defer m.mu.Unlock()
		if m.data == nil {
			m.data = make(map[string]int)
		}
		if v, ok = m.data[key]; ok {
			return v
		}
		v = f()
		m.data[key] = v
		return v
	} else {
		return v
	}
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
// ff:设置值并跳过已存在
// m:
// key:名称
// value:值
func (m *StrIntMap) SetIfNotExist(key string, value int) bool {
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
func (m *StrIntMap) SetIfNotExistFunc(key string, f func() int) bool {
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
func (m *StrIntMap) SetIfNotExistFuncLock(key string, f func() int) bool {
	if !m.Contains(key) {
		m.mu.Lock()
		defer m.mu.Unlock()
		if m.data == nil {
			m.data = make(map[string]int)
		}
		if _, ok := m.data[key]; !ok {
			m.data[key] = f()
		}
		return true
	}
	return false
}

// 通过键删除map中的批删除值。 md5:57081208d84ca7e8
// ff:删除多个值
// m:
// keys:名称
func (m *StrIntMap) Removes(keys []string) {
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
// value:被删除值
func (m *StrIntMap) Remove(key string) (value int) {
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
func (m *StrIntMap) Keys() []string {
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

// Values 将地图中的所有值返回为一个切片。 md5:a89b5b485c966abd
// ff:取所有值
// m:
func (m *StrIntMap) Values() []int {
	m.mu.RLock()
	var (
		values = make([]int, len(m.data))
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
func (m *StrIntMap) Contains(key string) bool {
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
func (m *StrIntMap) Size() int {
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
func (m *StrIntMap) IsEmpty() bool {
	return m.Size() == 0
}

// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。 md5:0553a5cd54a22f3c
// ff:清空
// m:
func (m *StrIntMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]int)
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。 md5:a84ecf2839212d81
// ff:替换
// m:
// data:map值
func (m *StrIntMap) Replace(data map[string]int) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
// ff:遍历写锁定
// m:
// f:回调函数
// m:
func (m *StrIntMap) LockFunc(f func(m map[string]int)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
// ff:遍历读锁定
// m:
// f:回调函数
// m:
func (m *StrIntMap) RLockFunc(f func(m map[string]int)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Flip 将映射的键值对交换为值键。 md5:dbcb578f1b30fa01
// ff:名称值交换
// m:
func (m *StrIntMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]int, len(m.data))
	for k, v := range m.data {
		n[gconv.String(v)] = gconv.Int(k)
	}
	m.data = n
}

// Merge 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
// ff:合并
// m:
// other:map值
func (m *StrIntMap) Merge(other *StrIntMap) {
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
func (m *StrIntMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
// ff:
// m:
func (m StrIntMap) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return json.Marshal(m.data)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
// ff:
// m:
// b:
func (m *StrIntMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[string]int)
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
func (m *StrIntMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[string]int)
	}
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.Bytes(value), &m.data)
	default:
		for k, v := range gconv.Map(value) {
			m.data[k] = gconv.Int(v)
		}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
// ff:
// m:
func (m *StrIntMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]int, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return NewStrIntMapFrom(data, m.mu.IsSafe())
}

// IsSubOf 检查当前映射是否是`other`的子映射。 md5:9a6c60859c5a0fbc
// ff:是否为子集
// m:
// other:父集Map
func (m *StrIntMap) IsSubOf(other *StrIntMap) bool {
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
// addedKeys:增加的名称
// removedKeys:删除的名称
// updatedKeys:更新数据的名称
func (m *StrIntMap) Diff(other *StrIntMap) (addedKeys, removedKeys, updatedKeys []string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	for key := range m.data {
		if _, ok := other.data[key]; !ok {
			removedKeys = append(removedKeys, key)
		} else if m.data[key] != other.data[key] {
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
