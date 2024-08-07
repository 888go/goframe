// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package map类

import (
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gconv "github.com/888go/goframe/util/gconv"
)

// IntStrMap 实现了一个带有开关的读写锁保护的 int 到 string 的映射。 md5:6cf655e78fde5a30
type IntStrMap struct {
	mu   rwmutex.RWMutex
	data map[int]string
}

// X创建IntStr 返回一个空的 IntStrMap 对象。
// 参数 `safe` 用于指定是否使用并发安全的 map，默认为 false。
// md5:2be815573a3b558d
func X创建IntStr(并发安全 ...bool) *IntStrMap {
	return &IntStrMap{
		mu:   rwmutex.Create(并发安全...),
		data: make(map[int]string),
	}
}

// X创建IntStr并从Map 从给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（不进行深拷贝），
// 在外部修改映射时可能会出现并发安全问题。
// md5:ad2e556b3e927079
func X创建IntStr并从Map(map值 map[int]string, 并发安全 ...bool) *IntStrMap {
	return &IntStrMap{
		mu:   rwmutex.Create(并发安全...),
		data: map值,
	}
}

// X遍历 使用自定义回调函数 `f` 读取只读哈希映射。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:52d024b320a69c3b
func (m *IntStrMap) X遍历(f func(k int, v string) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// X取副本 返回一个新的哈希映射，其中包含当前映射数据的副本。 md5:b9264f3636ead08a
func (m *IntStrMap) X取副本() *IntStrMap {
	return X创建IntStr并从Map(m.X浅拷贝(), m.mu.IsSafe())
}

// X取Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
func (m *IntStrMap) X取Map() map[int]string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[int]string, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// X取MapStrAny将映射的底层数据复制为map[string]interface{}。 md5:46db5a1110397522
func (m *IntStrMap) X取MapStrAny() map[string]interface{} {
	m.mu.RLock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[gconv.String(k)] = v
	}
	m.mu.RUnlock()
	return data
}

// X浅拷贝 返回哈希映射底层数据的一个副本。 md5:46f762167d5821b1
func (m *IntStrMap) X浅拷贝() map[int]string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[int]string, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// X删除所有空值 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
func (m *IntStrMap) X删除所有空值() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// X设置值 将键值对设置到哈希映射中。 md5:07ea2dd1ea28820a
func (m *IntStrMap) X设置值(key int, val string) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[int]string)
	}
	m.data[key] = val
	m.mu.Unlock()
}

// 将键值对设置到哈希映射中。 md5:e3f3f8a1b69eb832
func (m *IntStrMap) X设置值Map(map值 map[int]string) {
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
func (m *IntStrMap) X查找(名称 int) (value string, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// X取值 根据给定的 `key` 获取值。 md5:2b744a3e455aadfb
func (m *IntStrMap) X取值(名称 int) (value string) {
	m.mu.RLock()
	if m.data != nil {
		value = m.data[名称]
	}
	m.mu.RUnlock()
	return
}

// X出栈 从映射中获取并删除一个元素。 md5:2d364ca2b6054111
func (m *IntStrMap) X出栈() (名称 int, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for 名称, value = range m.data {
		delete(m.data, 名称)
		return
	}
	return
}

// X出栈多个 从映射中检索并删除 `size` 个项目。
// 如果 size 等于 -1，则返回所有项目。
// md5:0f2cdbc0238fdc37
func (m *IntStrMap) X出栈多个(数量 int) map[int]string {
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
		newMap = make(map[int]string, 数量)
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

// doSetWithLockCheck 使用互斥锁(mutex.Lock)检查键的值是否存在，
// 如果不存在，则将给定的`value`设置到映射中指定的`key`处，
// 否则，直接返回已存在的值。
//
// 它返回与给定`key`关联的值。
// md5:3a2d1537d3fe7230
func (m *IntStrMap) doSetWithLockCheck(key int, value string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]string)
	}
	if v, ok := m.data[key]; ok {
		return v
	}
	m.data[key] = value
	return value
}

// X取值或设置值 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (m *IntStrMap) X取值或设置值(名称 int, value string) string {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, value)
	} else {
		return v
	}
}

// X取值或设置值_函数通过键获取值，
// 如果不存在，它将使用回调函数`f`返回的值设置该值，并返回这个值。
// md5:c4de9d0fac2a8916
func (m *IntStrMap) X取值或设置值_函数(名称 int, f func() string) string {
	if v, ok := m.X查找(名称); !ok {
		return m.doSetWithLockCheck(名称, f())
	} else {
		return v
	}
}

// X取值或设置值_函数带锁 通过键获取值，
// 如果不存在，则使用回调函数 `f` 的返回值设置该键的值，并返回这个值。
//
// X取值或设置值_函数带锁 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的mutex。
// md5:ac8ad0e9416578ba
func (m *IntStrMap) X取值或设置值_函数带锁(名称 int, f func() string) string {
	if v, ok := m.X查找(名称); !ok {
		m.mu.Lock()
		defer m.mu.Unlock()
		if m.data == nil {
			m.data = make(map[int]string)
		}
		if v, ok = m.data[名称]; ok {
			return v
		}
		v = f()
		m.data[名称] = v
		return v
	} else {
		return v
	}
}

// X设置值并跳过已存在 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (m *IntStrMap) X设置值并跳过已存在(名称 int, value string) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, value)
		return true
	}
	return false
}

// X设置值并跳过已存在_函数 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (m *IntStrMap) X设置值并跳过已存在_函数(名称 int, f func() string) bool {
	if !m.X是否存在(名称) {
		m.doSetWithLockCheck(名称, f())
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
func (m *IntStrMap) X设置值并跳过已存在_函数带锁(名称 int, f func() string) bool {
	if !m.X是否存在(名称) {
		m.mu.Lock()
		defer m.mu.Unlock()
		if m.data == nil {
			m.data = make(map[int]string)
		}
		if _, ok := m.data[名称]; !ok {
			m.data[名称] = f()
		}
		return true
	}
	return false
}

// 通过键删除map中的批删除值。 md5:57081208d84ca7e8
func (m *IntStrMap) X删除多个值(名称 []int) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range 名称 {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// X删除 通过给定的`key`从map中删除值，并返回被删除的值。 md5:5ee6dc9be17b4ab8
func (m *IntStrMap) X删除(名称 int) (value string) {
	m.mu.Lock()
	if m.data != nil {
		var ok bool
		if value, ok = m.data[名称]; ok {
			delete(m.data, 名称)
		}
	}
	m.mu.Unlock()
	return
}

// X取所有名称 返回映射中所有键的切片。 md5:425640fff4178659
func (m *IntStrMap) X取所有名称() []int {
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

// X取所有值 将地图中的所有值返回为一个切片。 md5:a89b5b485c966abd
func (m *IntStrMap) X取所有值() []string {
	m.mu.RLock()
	var (
		values = make([]string, len(m.data))
		index  = 0
	)
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

// X是否存在 检查键是否存在。
// 如果键存在，它返回 true，否则返回 false。
// md5:d8fb22313aadd65f
func (m *IntStrMap) X是否存在(名称 int) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[名称]
	}
	m.mu.RUnlock()
	return ok
}

// X取数量返回映射的大小。 md5:da42fb3955847483
func (m *IntStrMap) X取数量() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// X是否为空 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
func (m *IntStrMap) X是否为空() bool {
	return m.X取数量() == 0
}

// X清空 删除映射中的所有数据，它将重新创建一个新的底层数据映射。 md5:0553a5cd54a22f3c
func (m *IntStrMap) X清空() {
	m.mu.Lock()
	m.data = make(map[int]string)
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。 md5:a84ecf2839212d81
func (m *IntStrMap) X替换(map值 map[int]string) {
	m.mu.Lock()
	m.data = map值
	m.mu.Unlock()
}

// X遍历写锁定 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
func (m *IntStrMap) X遍历写锁定(回调函数 func(m map[int]string)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	回调函数(m.data)
}

// X遍历读锁定 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
func (m *IntStrMap) X遍历读锁定(回调函数 func(m map[int]string)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	回调函数(m.data)
}

// X名称值交换 将映射的键值对交换为值键。 md5:dbcb578f1b30fa01
func (m *IntStrMap) X名称值交换() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[int]string, len(m.data))
	for k, v := range m.data {
		n[gconv.X取整数(v)] = gconv.String(k)
	}
	m.data = n
}

// X合并 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
func (m *IntStrMap) X合并(map值 *IntStrMap) {
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
func (m *IntStrMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (m IntStrMap) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return json.Marshal(m.data)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (m *IntStrMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]string)
	}
	if err := json.UnmarshalUseNumber(b, &m.data); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。 md5:6f3087a6f7df5477
func (m *IntStrMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[int]string)
	}
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.X取字节集(value), &m.data)
	default:
		for k, v := range gconv.X取Map(value) {
			m.data[gconv.X取整数(k)] = gconv.String(v)
		}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (m *IntStrMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[int]string, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return X创建IntStr并从Map(data, m.mu.IsSafe())
}

// X是否为子集 检查当前映射是否是`other`的子映射。 md5:9a6c60859c5a0fbc
func (m *IntStrMap) X是否为子集(父集Map *IntStrMap) bool {
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
func (m *IntStrMap) X比较(map值 *IntStrMap) (addedKeys, removedKeys, updatedKeys []int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	map值.mu.RLock()
	defer map值.mu.RUnlock()

	for key := range m.data {
		if _, ok := map值.data[key]; !ok {
			removedKeys = append(removedKeys, key)
		} else if m.data[key] != map值.data[key] {
			updatedKeys = append(updatedKeys, key)
		}
	}
	for key := range map值.data {
		if _, ok := m.data[key]; !ok {
			addedKeys = append(addedKeys, key)
		}
	}
	return
}
