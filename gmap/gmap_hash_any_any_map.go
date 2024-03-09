// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/888go/goframe/gmap/internal/deepcopy"
	"github.com/888go/goframe/gmap/internal/empty"
	"github.com/888go/goframe/gmap/internal/json"
	"github.com/888go/goframe/gmap/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

// AnyAnyMap 包装了 map 类型 `map[interface{}]interface{}`，并提供了更多的映射功能。
type AnyAnyMap struct {
	mu   rwmutex.RWMutex
	data map[interface{}]interface{}
}

// NewAnyAnyMap 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否使用线程安全的 map， 默认为 false。
func X创建AnyAny(并发安全 ...bool) *AnyAnyMap {
	return &AnyAnyMap{
		mu:   rwmutex.Create(并发安全...),
		data: make(map[interface{}]interface{}),
	}
}

// NewAnyAnyMapFrom 通过给定的 `data` 地图创建并返回一个哈希映射。
// 注意，参数 `data` 中的地图将被设置为底层数据地图（无深度复制），
// 因此在外部修改该映射时可能存在一些并发安全问题。
func X创建AnyAny并从Map(map值 map[interface{}]interface{}, 并发安全 ...bool) *AnyAnyMap {
	return &AnyAnyMap{
		mu:   rwmutex.Create(并发安全...),
		data: map值,
	}
}

// Iterator 使用自定义回调函数 `f` 以只读方式迭代哈希表。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止迭代。
func (m *AnyAnyMap) X遍历(回调函数 func(k interface{}, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !回调函数(k, v) {
			break
		}
	}
}

// Clone 返回一个新的哈希映射，其中包含当前映射数据的副本。
func (m *AnyAnyMap) X取副本(并发安全 ...bool) *AnyAnyMap {
	return X创建并从Map(m.X浅拷贝(), 并发安全...)
}

// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景下，将会返回底层数据的一个副本，
// 否则将返回指向底层数据的指针。
func (m *AnyAnyMap) X取Map() map[interface{}]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapCopy 返回哈希映射底层数据的浅复制副本。
func (m *AnyAnyMap) X浅拷贝() map[interface{}]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny 返回该映射底层数据的一个副本，类型为 map[string]interface{}。
func (m *AnyAnyMap) X取MapStrAny() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[gconv.String(k)] = v
	}
	return data
}

// FilterEmpty 删除所有值为空的键值对。
// 以下类型的值被视为空：0, nil, false, "", 切片/映射/通道长度为0。
func (m *AnyAnyMap) X删除所有空值() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
}

// FilterNil 删除所有值为nil的键值对。
func (m *AnyAnyMap) X删除所有nil值() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Set 将键值对设置到哈希映射中。
func (m *AnyAnyMap) X设置值(名称 interface{}, 值 interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	m.data[名称] = 值
	m.mu.Unlock()
}

// Sets批量设置键值对到哈希映射中。
func (m *AnyAnyMap) X设置值Map(map值 map[interface{}]interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = map值
	} else {
		for k, v := range map值 {
			m.data[k] = v
		}
	}
	m.mu.Unlock()
}

// Search 通过给定的 `key` 在映射中搜索。
// 第二个返回参数 `found` 如果找到了 key，则为 true，否则为 false。
func (m *AnyAnyMap) X查找(名称 interface{}) (值 interface{}, 成功 bool) {
	m.mu.RLock()
	if m.data != nil {
		值, 成功 = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// Get 通过给定的 `key` 返回对应的值。
func (m *AnyAnyMap) X取值(名称 interface{}) (值 interface{}) {
	m.mu.RLock()
	if m.data != nil {
		值 = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// Pop 从映射中检索并删除一个项目。
func (m *AnyAnyMap) X出栈() (名称, 值 interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for 名称, 值 = range m.data {
		delete(m.data, 名称)
		return
	}
	return
}

// Pops 从映射中获取并删除 `size` 个元素。
// 当 size == -1 时，它返回所有元素。
func (m *AnyAnyMap) X出栈多个(数量 int) map[interface{}]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if 数量 > len(m.data) || 数量 == -1 {
		数量 = len(m.data)
	}
	if 数量 == 0 {
		return nil
	}
	var (
		index  = 0
		newMap = make(map[interface{}]interface{}, 数量)
	)
	for k, v := range m.data {
		delete(m.data, k)
		newMap[k] = v
		index++
		if index == 数量 {
			break
		}
	}
	return newMap
}

// doSetWithLockCheck 检查在对 mutex 锁定后，给定 key 的值是否存在，
// 如果不存在，则使用给定的 `key` 将 value 设置到 map 中，
// 否则仅返回已存在的值。
//
// 在设置值时，如果 `value` 是 `func() interface {}` 类型，
// 它将在哈希映射的 mutex 锁定下执行，
// 并将其返回值以 `key` 为键设置到 map 中。
//
// 它最后返回给定 `key` 对应的值。
func (m *AnyAnyMap) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
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

// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
func (m *AnyAnyMap) X取值或设置值(名称 interface{}, 值 interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 值)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
func (m *AnyAnyMap) X取值或设置值_函数(名称 interface{}, 回调函数 func() interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 回调函数())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键返回值，如果该键不存在，则使用回调函数 `f` 返回的值设置并返回这个新值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
func (m *AnyAnyMap) X取值或设置值_函数带锁(名称 interface{}, 回调函数 func() interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 回调函数)
	} else {
		return v
	}
}

// GetVar 通过给定的 `key` 返回一个具有相应值的 Var。
// 返回的 Var 不是线程安全的。
func (m *AnyAnyMap) X取值泛型类(名称 interface{}) *gvar.Var {
	return gvar.New(m.X取值(名称))
}

// GetVarOrSet 返回一个从 GetOrSet 获取结果的 Var。
// 返回的 Var 不是线程安全的。
func (m *AnyAnyMap) X取值或设置值泛型类(名称 interface{}, 值 interface{}) *gvar.Var {
	return gvar.New(m.X取值或设置值(名称, 值))
}

// GetVarOrSetFunc 返回一个 Var，其结果来自 GetOrSetFunc 的调用。
// 返回的 Var 不具备并发安全特性。
func (m *AnyAnyMap) X取值或设置值泛型类_函数(名称 interface{}, 回调函 func() interface{}) *gvar.Var {
	return gvar.New(m.X取值或设置值_函数(名称, 回调函))
}

// GetVarOrSetFuncLock 返回一个 Var，其结果来自 GetOrSetFuncLock。
// 返回的 Var 不是并发安全的。
func (m *AnyAnyMap) X取值或设置值泛型类_函数带锁(名称 interface{}, 回调函数 func() interface{}) *gvar.Var {
	return gvar.New(m.X取值或设置值_函数带锁(名称, 回调函数))
}

// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (m *AnyAnyMap) X设置值并跳过已存在(名称 interface{}, 值 interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 值)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (m *AnyAnyMap) X设置值并跳过已存在_函数(名称 interface{}, 回调函数 func() interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 回调函数())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 函数用于设置键值对，其值为回调函数 `f` 的返回值，并在设置成功时返回 true。
// 若 `key` 已存在，则返回 false，并且将忽略 `value` 参数。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在执行回调函数 `f` 时会锁定哈希表的 mutex 锁。
func (m *AnyAnyMap) X设置值并跳过已存在_函数带锁(名称 interface{}, 回调函数 func() interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 回调函数)
		return true
	}
	return false
}

// Remove通过给定的`key`从map中删除值，并返回这个被删除的值。
func (m *AnyAnyMap) X删除(名称 interface{}) (值 interface{}) {
	m.mu.Lock()
	if m.data != nil {
		var ok bool
		if 值, ok = m.data[名称]; ok {
			delete(m.data, 名称)
		}
	}
	m.mu.Unlock()
	return
}

// 删除map中通过keys指定的所有值，进行批量删除。
func (m *AnyAnyMap) X删除多个值(名称 []interface{}) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range 名称 {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Keys 返回该映射的所有键作为一个切片。
func (m *AnyAnyMap) X取所有名称() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var (
		keys  = make([]interface{}, len(m.data))
		index = 0
	)
	for key := range m.data {
		keys[index] = key
		index++
	}
	return keys
}

// Values 返回该映射的所有值作为一个切片。
func (m *AnyAnyMap) X取所有值() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var (
		values = make([]interface{}, len(m.data))
		index  = 0
	)
	for _, value := range m.data {
		values[index] = value
		index++
	}
	return values
}

// Contains 检查键是否存在。
// 如果 `key` 存在，则返回 true，否则返回 false。
func (m *AnyAnyMap) X是否存在(名称 interface{}) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[名称]
	}
	m.mu.RUnlock()
	return ok
}

// Size 返回映射的大小。
func (m *AnyAnyMap) X取数量() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty 检查该映射是否为空。
// 如果映射为空，则返回 true，否则返回 false。
func (m *AnyAnyMap) X是否为空() bool {
	return m.X取数量() == 0
}

// 清空删除映射中的所有数据，它会重新创建一个新的底层数据映射。
func (m *AnyAnyMap) X清空() {
	m.mu.Lock()
	m.data = make(map[interface{}]interface{})
	m.mu.Unlock()
}

// 用给定的`data`替换map中的数据。
func (m *AnyAnyMap) X替换(map值 map[interface{}]interface{}) {
	m.mu.Lock()
	m.data = map值
	m.mu.Unlock()
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 内锁定写入操作。
func (m *AnyAnyMap) X遍历写锁定(回调函数 func(m map[interface{}]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	回调函数(m.data)
}

// RLockFunc 在 RWMutex.RLock 内使用给定的回调函数 `f` 进行读取锁定。
func (m *AnyAnyMap) X遍历读锁定(回调函数 func(m map[interface{}]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	回调函数(m.data)
}

// Flip 将映射中的键值对进行交换，即把键变成值，值变成键。
func (m *AnyAnyMap) X名称值交换() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		n[v] = k
	}
	m.data = n
}

// Merge 合并两个哈希映射。
// `other` 映射将会被合并到映射 `m` 中。
func (m *AnyAnyMap) X合并(map值 *AnyAnyMap) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = map值.X浅拷贝()
		return
	}
	if map值 != m {
		map值.mu.RLock()
		defer map值.mu.RUnlock()
	}
	for k, v := range map值.data {
		m.data[k] = v
	}
}

// String 将映射转换为字符串并返回。
func (m *AnyAnyMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (m AnyAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(gconv.Map(m.X取Map()))
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (m *AnyAnyMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	var data map[string]interface{}
	if err := json.UnmarshalUseNumber(b, &data); err != nil {
		return err
	}
	for k, v := range data {
		m.data[k] = v
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 map 设置任意类型的值。
func (m *AnyAnyMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	for k, v := range gconv.Map(value) {
		m.data[k] = v
	}
	return
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (m *AnyAnyMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}

	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = deepcopy.Copy(v)
	}
	return X创建并从Map(data, m.mu.IsSafe())
}

// IsSubOf 检查当前 map 是否为 `other` 的子集。
func (m *AnyAnyMap) X是否为子集(父集Map *AnyAnyMap) bool {
	if m == 父集Map {
		return true
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	父集Map.mu.RLock()
	defer 父集Map.mu.RUnlock()
	for key, value := range m.data {
		otherValue, ok := 父集Map.data[key]
		if !ok {
			return false
		}
		if otherValue != value {
			return false
		}
	}
	return true
}

// Diff 函数用于比较当前映射 `m` 与映射 `other`，并返回它们不同的键。
// 返回的 `addedKeys` 是存在于映射 `m` 中但不在映射 `other` 中的键。
// 返回的 `removedKeys` 是存在于映射 `other` 中但不在映射 `m` 中的键。
// 返回的 `updatedKeys` 是同时存在于映射 `m` 和 `other` 中，但其对应值不相等（`!=`）的键。
func (m *AnyAnyMap) X比较(map值 *AnyAnyMap) (增加的名称, 删除的名称, 更新数据的名称 []interface{}) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	map值.mu.RLock()
	defer map值.mu.RUnlock()

	for key := range m.data {
		if _, ok := map值.data[key]; !ok {
			删除的名称 = append(删除的名称, key)
		} else if !reflect.DeepEqual(m.data[key], map值.data[key]) {
			更新数据的名称 = append(更新数据的名称, key)
		}
	}
	for key := range map值.data {
		if _, ok := m.data[key]; !ok {
			增加的名称 = append(增加的名称, key)
		}
	}
	return
}
