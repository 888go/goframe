// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package map类

import (
	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gconv "github.com/888go/goframe/util/gconv"
	"reflect"
)

// AnyAnyMap 包装了 map 类型 `map[interface{}]interface{}` 并提供了更多的 map 功能。 md5:fdeb4ee0e757ccc7
type AnyAnyMap struct {
	mu   rwmutex.RWMutex
	data map[interface{}]interface{}
}

// X创建AnyAny 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否使用并发安全的映射，默认为false。
// md5:f5ea72ba91a61ee2
func X创建AnyAny(并发安全 ...bool) *AnyAnyMap {
	return &AnyAnyMap{
		mu:   rwmutex.Create(并发安全...),
		data: make(map[interface{}]interface{}),
	}
}

// X创建AnyAny并从Map 根据给定的映射 `data` 创建并返回一个哈希映射。
// 请注意，参数 `data` 映射将被设置为底层数据映射（无深拷贝），
// 当在外部修改此映射时，可能存在并发安全问题。
// md5:e7327f6b619e71b1
func X创建AnyAny并从Map(map值 map[interface{}]interface{}, 并发安全 ...bool) *AnyAnyMap {
	return &AnyAnyMap{
		mu:   rwmutex.Create(并发安全...),
		data: map值,
	}
}

// X遍历 使用自定义回调函数 `f` 读取只读哈希映射。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:52d024b320a69c3b
func (m *AnyAnyMap) X遍历(f func(k interface{}, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// X取副本 返回一个新的哈希映射，其中包含当前映射数据的副本。 md5:b9264f3636ead08a
func (m *AnyAnyMap) X取副本(并发安全 ...bool) *AnyAnyMap {
	return X创建并从Map(m.X浅拷贝(), 并发安全...)
}

// X取Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
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

// X浅拷贝 返回哈希映射底层数据的浅拷贝。 md5:7c14284a426ba2a8
func (m *AnyAnyMap) X浅拷贝() map[interface{}]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// X取MapStrAny将映射的底层数据复制为map[string]interface{}。 md5:46db5a1110397522
func (m *AnyAnyMap) X取MapStrAny() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[gconv.String(k)] = v
	}
	return data
}

// X删除所有空值 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
func (m *AnyAnyMap) X删除所有空值() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
}

// X删除所有nil值 删除所有值为 nil 的键值对。 md5:3c964818401771a4
func (m *AnyAnyMap) X删除所有nil值() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// X设置值 将键值对设置到哈希映射中。 md5:07ea2dd1ea28820a
func (m *AnyAnyMap) X设置值(key interface{}, value interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	m.data[key] = value
	m.mu.Unlock()
}

// 将键值对设置到哈希映射中。 md5:e3f3f8a1b69eb832
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

// X查找 在给定的`key`下搜索映射。
// 第二个返回参数`found`如果找到键，则为true，否则为false。
// md5:99336de9941a3b02
func (m *AnyAnyMap) X查找(名称 interface{}) (值 interface{}, 成功 bool) {
	m.mu.RLock()
	if m.data != nil {
		值, 成功 = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// X取值 根据给定的 `key` 获取值。 md5:2b744a3e455aadfb
func (m *AnyAnyMap) X取值(名称 interface{}) (值 interface{}) {
	m.mu.RLock()
	if m.data != nil {
		值 = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// X出栈 从映射中获取并删除一个元素。 md5:2d364ca2b6054111
func (m *AnyAnyMap) X出栈() (名称, 值 interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for 名称, 值 = range m.data {
		delete(m.data, 名称)
		return
	}
	return
}

// X出栈多个 从映射中检索并删除 `size` 个项目。
// 如果 size 等于 -1，则返回所有项目。
// md5:0f2cdbc0238fdc37
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

// doSetWithLockCheck 会使用 mutex.Lock 检查给定键的值是否存在。
// 如果不存在，将使用给定的 `key` 将值设置到映射中；否则，直接返回已存在的值。
// 
// 当设置值时，如果 `value` 类型为 `func() interface{}`，它将在映射的 mutex.Lock 保护下执行，
// 并将返回值设置到映射中，键为 `key`。
// 
// 它返回给定 `key` 的值。
// md5:60f1f50efa66e173
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

// X取值或设置值 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (m *AnyAnyMap) X取值或设置值(名称 interface{}, 值 interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 值)
	} else {
		return v
	}
}

// X取值或设置值_函数 通过键获取值，
// 如果键不存在，则使用回调函数`f`的返回值设置值，
// 并返回这个设置的值。
// md5:f584dd7547dfbcc0
func (m *AnyAnyMap) X取值或设置值_函数(名称 interface{}, 回调函数 func() interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 回调函数())
	} else {
		return v
	}
}

// X取值或设置值_函数带锁 通过键获取值，
// 如果不存在，它将使用回调函数 `f` 的返回值设置该值，然后返回这个值。
//
// X取值或设置值_函数带锁 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的 mutex。
// md5:d32fdee586d84dde
func (m *AnyAnyMap) X取值或设置值_函数带锁(名称 interface{}, 回调函数 func() interface{}) interface{} {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, 回调函数)
	} else {
		return v
	}
}

// X取值泛型类通过给定的`key`返回一个Var。返回的Var是非并发安全的。
// md5:debfb1b2bd13312b
func (m *AnyAnyMap) X取值泛型类(名称 interface{}) *gvar.Var {
	return gvar.X创建(m.X取值(名称))
}

// X取值或设置值泛型类 返回一个 Var，其结果来自 GetOrSet。
// 返回的 Var 是非并发安全的。
// md5:5d4b8a2f15c827e0
func (m *AnyAnyMap) X取值或设置值泛型类(名称 interface{}, 值 interface{}) *gvar.Var {
	return gvar.X创建(m.X取值或设置值(名称, 值))
}

// X取值或设置值泛型类_函数 返回一个Var，其结果来自GetOrSetFunc。
// 返回的Var不具备并发安全性。
// md5:7d7674129b73ead1
func (m *AnyAnyMap) X取值或设置值泛型类_函数(名称 interface{}, 回调函 func() interface{}) *gvar.Var {
	return gvar.X创建(m.X取值或设置值_函数(名称, 回调函))
}

// X取值或设置值泛型类_函数带锁 返回一个从 GetOrSetFuncLock 获得结果的 Var。返回的 Var 不是线程安全的。
// md5:bdab644d14c89234
func (m *AnyAnyMap) X取值或设置值泛型类_函数带锁(名称 interface{}, 回调函数 func() interface{}) *gvar.Var {
	return gvar.X创建(m.X取值或设置值_函数带锁(名称, 回调函数))
}

// X设置值并跳过已存在 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (m *AnyAnyMap) X设置值并跳过已存在(名称 interface{}, 值 interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 值)
		return true
	}
	return false
}

// X设置值并跳过已存在_函数 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (m *AnyAnyMap) X设置值并跳过已存在_函数(名称 interface{}, 回调函数 func() interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 回调函数())
		return true
	}
	return false
}

// X设置值并跳过已存在_函数带锁 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，则返回 false，`value` 将被忽略。
//
// X设置值并跳过已存在_函数带锁 与 SetIfNotExistFunc 函数的区别在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
// md5:a6ee84b157328f61
func (m *AnyAnyMap) X设置值并跳过已存在_函数带锁(名称 interface{}, 回调函数 func() interface{}) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, 回调函数)
		return true
	}
	return false
}

// X删除 通过给定的`key`从map中删除值，并返回被删除的值。 md5:5ee6dc9be17b4ab8
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

// 通过键删除map中的批删除值。 md5:57081208d84ca7e8
func (m *AnyAnyMap) X删除多个值(名称 []interface{}) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range 名称 {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// X取所有名称 返回映射中所有键的切片。 md5:425640fff4178659
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

// X取所有值 将地图中的所有值返回为一个切片。 md5:a89b5b485c966abd
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

// X是否存在 检查键是否存在。
// 如果键存在，它返回 true，否则返回 false。
// md5:d8fb22313aadd65f
func (m *AnyAnyMap) X是否存在(名称 interface{}) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[名称]
	}
	m.mu.RUnlock()
	return ok
}

// X取数量返回映射的大小。 md5:da42fb3955847483
func (m *AnyAnyMap) X取数量() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// X是否为空 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
func (m *AnyAnyMap) X是否为空() bool {
	return m.X取数量() == 0
}

// X清空 删除映射中的所有数据，它将重新创建一个新的底层数据映射。 md5:0553a5cd54a22f3c
func (m *AnyAnyMap) X清空() {
	m.mu.Lock()
	m.data = make(map[interface{}]interface{})
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。 md5:a84ecf2839212d81
func (m *AnyAnyMap) X替换(map值 map[interface{}]interface{}) {
	m.mu.Lock()
	m.data = map值
	m.mu.Unlock()
}

// X遍历写锁定 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
func (m *AnyAnyMap) X遍历写锁定(回调函数 func(m map[interface{}]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	回调函数(m.data)
}

// X遍历读锁定 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
func (m *AnyAnyMap) X遍历读锁定(回调函数 func(m map[interface{}]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	回调函数(m.data)
}

// X名称值交换 将映射的键值对交换为值键。 md5:dbcb578f1b30fa01
func (m *AnyAnyMap) X名称值交换() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		n[v] = k
	}
	m.data = n
}

// X合并 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
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

// String 将地图转换为字符串形式并返回。 md5:6473318e71d3dfd0
func (m *AnyAnyMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (m AnyAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(gconv.X取Map(m.X取Map()))
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
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

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。 md5:6f3087a6f7df5477
func (m *AnyAnyMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	for k, v := range gconv.X取Map(value) {
		m.data[k] = v
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
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

// X是否为子集 检查当前映射是否是`other`的子映射。 md5:9a6c60859c5a0fbc
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

// X比较 函数比较当前地图 `m` 与地图 `other` 并返回它们不同的键。
// 返回的 `addedKeys` 是存在于地图 `m` 中但不在地图 `other` 中的键。
// 返回的 `removedKeys` 是存在于地图 `other` 中但不在地图 `m` 中的键。
// 返回的 `updatedKeys` 是同时存在于地图 `m` 和 `other` 中，但其值不相等（`!=`）的键。
// md5:d3bf0bf8c70e9093
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
