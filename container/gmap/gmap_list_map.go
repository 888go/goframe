// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package gmap

import (
	"bytes"
	"fmt"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

// ListMap 是一个保持插入顺序的映射。
//
// 它使用哈希表存储值，使用双向链表存储顺序。
//
// 该结构不支持多线程安全。
//
// 参考：http://en.wikipedia.org/wiki/Associative_array
// md5:a164418fa5f6798e
type ListMap struct {
	mu   rwmutex.RWMutex
	data map[interface{}]*glist.Element
	list *glist.List
}

type gListMapNode struct {
	key   interface{}
	value interface{}
}

// NewListMap 返回一个空的链接映射。
// ListMap 使用哈希表存储值，使用双向链表存储顺序。
// 参数 `safe` 用于指定是否在并发安全模式下使用映射，默认为 false。
// md5:515d74cfd6f50ee5
func NewListMap(safe ...bool) *ListMap {
	return &ListMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[interface{}]*glist.Element),
		list: glist.New(),
	}
}

// NewListMapFrom 从给定的映射（map）`data`创建一个链接映射。
// 注意，参数`data`映射将被设置为底层数据映射（不进行深拷贝），如果在外部修改该映射可能会导致并发安全问题。
// md5:d15c506b7dc77488
func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *ListMap {
	m := NewListMap(safe...)
	m.Sets(data)
	return m
}

// Iterator 是 IteratorAsc 的别名。. md5:1bfdea306db62845
func (m *ListMap) Iterator(f func(key, value interface{}) bool) {
	m.IteratorAsc(f)
}

// IteratorAsc 使用给定的回调函数 `f` 以升序遍历地图，并且是只读遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:da20ab57c78da7c6
func (m *ListMap) IteratorAsc(f func(key interface{}, value interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.list != nil {
		var node *gListMapNode
		m.list.IteratorAsc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			return f(node.key, node.value)
		})
	}
}

// IteratorDesc 使用给定的回调函数 `f` 以降序遍历只读映射。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止。
// md5:1745f0b396846901
func (m *ListMap) IteratorDesc(f func(key interface{}, value interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.list != nil {
		var node *gListMapNode
		m.list.IteratorDesc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			return f(node.key, node.value)
		})
	}
}

// Clone 返回一个新的链接映射，包含当前映射数据的副本。. md5:c24afe920899b3a6
func (m *ListMap) Clone(safe ...bool) *ListMap {
	return NewListMapFrom(m.Map(), safe...)
}

// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。. md5:0553a5cd54a22f3c
func (m *ListMap) Clear() {
	m.mu.Lock()
	m.data = make(map[interface{}]*glist.Element)
	m.list = glist.New()
	m.mu.Unlock()
}

// 用给定的 `data` 替换映射的数据。. md5:a84ecf2839212d81
func (m *ListMap) Replace(data map[interface{}]interface{}) {
	m.mu.Lock()
	m.data = make(map[interface{}]*glist.Element)
	m.list = glist.New()
	for key, value := range data {
		if e, ok := m.data[key]; !ok {
			m.data[key] = m.list.PushBack(&gListMapNode{key, value})
		} else {
			e.Value = &gListMapNode{key, value}
		}
	}
	m.mu.Unlock()
}

// Map返回map底层数据的副本。. md5:1be60fb31c8cf7e9
func (m *ListMap) Map() map[interface{}]interface{} {
	m.mu.RLock()
	var node *gListMapNode
	var data map[interface{}]interface{}
	if m.list != nil {
		data = make(map[interface{}]interface{}, len(m.data))
		m.list.IteratorAsc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			data[node.key] = node.value
			return true
		})
	}
	m.mu.RUnlock()
	return data
}

// MapStrAny将映射的底层数据复制为map[string]interface{}。. md5:46db5a1110397522
func (m *ListMap) MapStrAny() map[string]interface{} {
	m.mu.RLock()
	var node *gListMapNode
	var data map[string]interface{}
	if m.list != nil {
		data = make(map[string]interface{}, len(m.data))
		m.list.IteratorAsc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			data[gconv.String(node.key)] = node.value
			return true
		})
	}
	m.mu.RUnlock()
	return data
}

// FilterEmpty 删除所有值为空的键值对。. md5:77ba324f6e82e0c4
func (m *ListMap) FilterEmpty() {
	m.mu.Lock()
	if m.list != nil {
		var (
			keys = make([]interface{}, 0)
			node *gListMapNode
		)
		m.list.IteratorAsc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			if empty.IsEmpty(node.value) {
				keys = append(keys, node.key)
			}
			return true
		})
		if len(keys) > 0 {
			for _, key := range keys {
				if e, ok := m.data[key]; ok {
					delete(m.data, key)
					m.list.Remove(e)
				}
			}
		}
	}
	m.mu.Unlock()
}

// Set 将键值对设置到映射中。. md5:78bf4bf763bbf6e3
func (m *ListMap) Set(key interface{}, value interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	if e, ok := m.data[key]; !ok {
		m.data[key] = m.list.PushBack(&gListMapNode{key, value})
	} else {
		e.Value = &gListMapNode{key, value}
	}
	m.mu.Unlock()
}

// 将键值对设置到映射中。. md5:863aacdcc54fd6e1
func (m *ListMap) Sets(data map[interface{}]interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	for key, value := range data {
		if e, ok := m.data[key]; !ok {
			m.data[key] = m.list.PushBack(&gListMapNode{key, value})
		} else {
			e.Value = &gListMapNode{key, value}
		}
	}
	m.mu.Unlock()
}

// Search 在给定的`key`下搜索映射。
// 第二个返回参数`found`如果找到键，则为true，否则为false。
// md5:99336de9941a3b02
func (m *ListMap) Search(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		if e, ok := m.data[key]; ok {
			value = e.Value.(*gListMapNode).value
			found = ok
		}
	}
	m.mu.RUnlock()
	return
}

// Get 根据给定的 `key` 获取值。. md5:2b744a3e455aadfb
func (m *ListMap) Get(key interface{}) (value interface{}) {
	m.mu.RLock()
	if m.data != nil {
		if e, ok := m.data[key]; ok {
			value = e.Value.(*gListMapNode).value
		}
	}
	m.mu.RUnlock()
	return
}

// Pop 从映射中获取并删除一个元素。. md5:2d364ca2b6054111
func (m *ListMap) Pop() (key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, e := range m.data {
		value = e.Value.(*gListMapNode).value
		delete(m.data, k)
		m.list.Remove(e)
		return k, value
	}
	return
}

// Pops 从映射中检索并删除 `size` 个项目。
// 如果 size 等于 -1，则返回所有项目。
// md5:0f2cdbc0238fdc37
func (m *ListMap) Pops(size int) map[interface{}]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if size > len(m.data) || size == -1 {
		size = len(m.data)
	}
	if size == 0 {
		return nil
	}
	index := 0
	newMap := make(map[interface{}]interface{}, size)
	for k, e := range m.data {
		value := e.Value.(*gListMapNode).value
		delete(m.data, k)
		m.list.Remove(e)
		newMap[k] = value
		index++
		if index == size {
			break
		}
	}
	return newMap
}

// doSetWithLockCheck 使用 mutex.Lock 检查给定键的值是否存在。
// 如果不存在，使用给定的 `key` 将值设置到映射中；否则，直接返回现有的值。
// 
// 当设置值时，如果 `value` 是类型为 `func() interface{}`，它将在映射的 mutex.Lock 保护下执行，
// 并将返回值设置为映射中的 `key`。
// 
// 它返回给定 `key` 的值。
// md5:b667e8828a47a6d9
func (m *ListMap) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	if e, ok := m.data[key]; ok {
		return e.Value.(*gListMapNode).value
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		m.data[key] = m.list.PushBack(&gListMapNode{key, value})
	}
	return value
}

// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (m *ListMap) GetOrSet(key interface{}, value interface{}) interface{} {
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
func (m *ListMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键获取值，
// 如果该值不存在，则使用回调函数 `f` 的返回值进行设置，
// 然后返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它在执行映射的 mutex.Lock 保护下执行函数 `f`。
// md5:f5e408a3393171bc
func (m *ListMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar通过给定的`key`返回一个Var。返回的Var是非并发安全的。
// md5:debfb1b2bd13312b
func (m *ListMap) GetVar(key interface{}) *gvar.Var {
	return gvar.New(m.Get(key))
}

// GetVarOrSet 返回一个 Var，其中包含从 GetVarOrSet 获取的结果。
// 返回的 Var 是非并发安全的。
// md5:c3730f368b7f62b5
func (m *ListMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(m.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个Var，其结果来自GetOrSetFunc。
// 返回的Var不具备并发安全性。
// md5:7d7674129b73ead1
func (m *ListMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个从 GetOrSetFuncLock 获得结果的 Var。返回的 Var 不是线程安全的。
// md5:bdab644d14c89234
func (m *ListMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (m *ListMap) SetIfNotExist(key interface{}, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (m *ListMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，它将返回 false，`value` 将被忽略。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在执行函数 `f` 时会获取映射的 mutex.Lock。
// md5:12e78d7edb4c4c12
func (m *ListMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Remove 通过给定的`key`从map中删除值，并返回被删除的值。. md5:5ee6dc9be17b4ab8
func (m *ListMap) Remove(key interface{}) (value interface{}) {
	m.mu.Lock()
	if m.data != nil {
		if e, ok := m.data[key]; ok {
			value = e.Value.(*gListMapNode).value
			delete(m.data, key)
			m.list.Remove(e)
		}
	}
	m.mu.Unlock()
	return
}

// 通过键删除map中的批删除值。. md5:57081208d84ca7e8
func (m *ListMap) Removes(keys []interface{}) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			if e, ok := m.data[key]; ok {
				delete(m.data, key)
				m.list.Remove(e)
			}
		}
	}
	m.mu.Unlock()
}

// Keys返回映射中所有键作为升序排列的切片。. md5:140d43c5cccae9d9
func (m *ListMap) Keys() []interface{} {
	m.mu.RLock()
	var (
		keys  = make([]interface{}, m.list.Len())
		index = 0
	)
	if m.list != nil {
		m.list.IteratorAsc(func(e *glist.Element) bool {
			keys[index] = e.Value.(*gListMapNode).key
			index++
			return true
		})
	}
	m.mu.RUnlock()
	return keys
}

// Values 将地图中的所有值返回为一个切片。. md5:a89b5b485c966abd
func (m *ListMap) Values() []interface{} {
	m.mu.RLock()
	var (
		values = make([]interface{}, m.list.Len())
		index  = 0
	)
	if m.list != nil {
		m.list.IteratorAsc(func(e *glist.Element) bool {
			values[index] = e.Value.(*gListMapNode).value
			index++
			return true
		})
	}
	m.mu.RUnlock()
	return values
}

// Contains 检查键是否存在。
// 如果键存在，它返回 true，否则返回 false。
// md5:d8fb22313aadd65f
func (m *ListMap) Contains(key interface{}) (ok bool) {
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Size返回映射的大小。. md5:da42fb3955847483
func (m *ListMap) Size() (size int) {
	m.mu.RLock()
	size = len(m.data)
	m.mu.RUnlock()
	return
}

// IsEmpty 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
func (m *ListMap) IsEmpty() bool {
	return m.Size() == 0
}

// Flip 将映射的键值对交换为值键。. md5:dbcb578f1b30fa01
func (m *ListMap) Flip() {
	data := m.Map()
	m.Clear()
	for key, value := range data {
		m.Set(value, key)
	}
}

// Merge 合并两个链接映射。
// 将将`other`映射合并到`m`映射中。
// md5:2ec13ae7c16e16f8
func (m *ListMap) Merge(other *ListMap) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	if other != m {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	var node *gListMapNode
	other.list.IteratorAsc(func(e *glist.Element) bool {
		node = e.Value.(*gListMapNode)
		if e, ok := m.data[node.key]; !ok {
			m.data[node.key] = m.list.PushBack(&gListMapNode{node.key, node.value})
		} else {
			e.Value = &gListMapNode{node.key, node.value}
		}
		return true
	})
}

// String 将地图转换为字符串形式并返回。. md5:6473318e71d3dfd0
func (m *ListMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (m ListMap) MarshalJSON() (jsonBytes []byte, err error) {
	if m.data == nil {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('{')
	m.Iterator(func(key, value interface{}) bool {
		valueBytes, valueJsonErr := json.Marshal(value)
		if valueJsonErr != nil {
			err = valueJsonErr
			return false
		}
		if buffer.Len() > 1 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(fmt.Sprintf(`"%v":%s`, key, valueBytes))
		return true
	})
	buffer.WriteByte('}')
	return buffer.Bytes(), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
func (m *ListMap) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	var data map[string]interface{}
	if err := json.UnmarshalUseNumber(b, &data); err != nil {
		return err
	}
	for key, value := range data {
		if e, ok := m.data[key]; !ok {
			m.data[key] = m.list.PushBack(&gListMapNode{key, value})
		} else {
			e.Value = &gListMapNode{key, value}
		}
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。. md5:6f3087a6f7df5477
func (m *ListMap) UnmarshalValue(value interface{}) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]*glist.Element)
		m.list = glist.New()
	}
	for k, v := range gconv.Map(value) {
		if e, ok := m.data[k]; !ok {
			m.data[k] = m.list.PushBack(&gListMapNode{k, v})
		} else {
			e.Value = &gListMapNode{k, v}
		}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
func (m *ListMap) DeepCopy() interface{} {
	if m == nil {
		return nil
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[interface{}]interface{}, len(m.data))
	if m.list != nil {
		var node *gListMapNode
		m.list.IteratorAsc(func(e *glist.Element) bool {
			node = e.Value.(*gListMapNode)
			data[node.key] = deepcopy.Copy(node.value)
			return true
		})
	}
	return NewListMapFrom(data, m.mu.IsSafe())
}
