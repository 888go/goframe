// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package gmap

import (
	"bytes"
	"fmt"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
)

// ListMap 是一个保留插入顺序的映射。
//
// 它底层通过哈希表存储值，并通过双向链表来保存元素的顺序。
//
// 该结构不保证线程安全。
//
// 参考文献：http://en.wikipedia.org/wiki/关联数组
type ListMap struct {
	mu   rwmutex.RWMutex
	data map[interface{}]*glist.Element
	list *glist.List
}

type gListMapNode struct {
	key   interface{}
	value interface{}
}

// NewListMap 返回一个空的链式映射。
// ListMap 由哈希表（用于存储值）和双向链表（用于存储顺序）作为底层支持。
// 参数 `safe` 用于指定是否在并发环境中安全地使用映射，默认情况下为 false。
// 以下是更详细的翻译：
// ```go
// NewListMap 函数用于创建并返回一个新的、空的链式映射结构体实例。
// 这个 ListMap 结构体内部结合了哈希表和双向链表两种数据结构：
// 哈希表用于高效地存储和查找键值对，而双向链表则用于记录键值对的插入顺序。
// 
// 参数 `safe` 表示是否需要保证该链式映射在并发环境中的安全性（即线程安全），
// 如果设置为 true，则在多 goroutine 并发访问时会进行相应的同步控制；
// 若不特别指定，其默认值为 false，此时不提供并发安全保证。
func NewListMap(safe ...bool) *ListMap {
	return &ListMap{
		mu:   rwmutex.Create(safe...),
		data: make(map[interface{}]*glist.Element),
		list: glist.New(),
	}
}

// NewListMapFrom 从给定的 `data` map 中创建一个链接映射。
// 注意，参数 `data` 中的映射将被设置为底层数据映射（非深度复制），
// 因此在外部修改该映射时可能会存在一些并发安全问题。
func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *ListMap {
	m := NewListMap(safe...)
	m.Sets(data)
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (m *ListMap) Iterator(f func(key, value interface{}) bool) {
	m.IteratorAsc(f)
}

// IteratorAsc 以升序遍历给定回调函数 `f` 的只读映射。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
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

// IteratorDesc 以降序方式遍历给定的只读映射，并使用回调函数 `f` 进行处理。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
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

// Clone 返回一个新的链接映射，其中包含当前映射数据的副本。
func (m *ListMap) Clone(safe ...bool) *ListMap {
	return NewListMapFrom(m.Map(), safe...)
}

// 清空删除映射中的所有数据，它会重新创建一个新的底层数据映射。
func (m *ListMap) Clear() {
	m.mu.Lock()
	m.data = make(map[interface{}]*glist.Element)
	m.list = glist.New()
	m.mu.Unlock()
}

// 用给定的`data`替换map中的数据。
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

// Map 返回映射底层数据的一个副本。
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

// MapStrAny 返回该映射底层数据的一个副本，类型为 map[string]interface{}。
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

// FilterEmpty 删除所有值为空的键值对。
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

// Set 将键值对设置到映射中。
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

// 设置批量数据：将键值对设置到映射（map）中。
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

// Search 通过给定的 `key` 在映射中搜索。
// 第二个返回参数 `found` 如果找到了 key，则为 true，否则为 false。
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

// Get 通过给定的 `key` 返回对应的值。
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

// Pop 从映射中检索并删除一个项目。
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

// Pops 从映射中获取并删除 `size` 个元素。
// 当 size == -1 时，它返回所有元素。
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

// doSetWithLockCheck 在对mutex.Lock进行检查后，确认键对应的值是否存在，
// 如果不存在，则使用给定的`key`将值设置到映射中；
// 否则仅返回已存在的值。
//
// 在设置值时，如果`value`的类型为`func() interface {}`，
// 将在映射的mutex.Lock保护下执行该函数，
// 并将其返回值以`key`为键设置到映射中。
//
// 最终返回带有给定`key`的值。
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

// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
func (m *ListMap) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
func (m *ListMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键返回值，如果不存在该键，则使用回调函数`f`返回的值设置该值，
// 并随后返回这个新设置的值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在对 map 进行 mutex.Lock 锁定后执行函数 `f`。
// 这段代码注释翻译成中文后的意思是：
// ```markdown
// GetOrSetFuncLock 方法通过给定的键获取值，
// 若键对应的值不存在，则会使用回调函数 `f` 返回的值进行设置，
// 并最终返回这个已设置的值。
//
// GetOrSetFuncLock 方法与 GetOrSetFunc 方法的区别在于，
// 在对映射（map）执行操作前，它会先调用 mutex.Lock 进行锁定。
func (m *ListMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 通过给定的 `key` 返回一个具有相应值的 Var。
// 返回的 Var 不是线程安全的。
func (m *ListMap) GetVar(key interface{}) *gvar.Var {
	return gvar.New(m.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 Var。
// 返回的 Var 对象不保证线程安全。
func (m *ListMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(m.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 Var，其结果来自 GetOrSetFunc 的调用。
// 返回的 Var 不具备并发安全特性。
func (m *ListMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个 Var，其结果来自 GetOrSetFuncLock。
// 返回的 Var 不是并发安全的。
func (m *ListMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(m.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (m *ListMap) SetIfNotExist(key interface{}, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (m *ListMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 函数设置键值对，其值为回调函数 `f` 的返回值，并在设置成功时返回 true。
// 若 `key` 已存在，则返回 false，同时将忽略 `value`。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在操作 map 时使用了 mutex.Lock 进行加锁，确保在执行函数 `f` 期间数据同步安全。
func (m *ListMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Remove通过给定的`key`从map中删除值，并返回这个被删除的值。
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

// 删除map中通过keys指定的所有值，进行批量删除。
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

// Keys 返回映射的所有键，以升序排列的切片形式。
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

// Values 返回该映射的所有值作为一个切片。
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
// 如果 `key` 存在，则返回 true，否则返回 false。
func (m *ListMap) Contains(key interface{}) (ok bool) {
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Size 返回映射的大小。
func (m *ListMap) Size() (size int) {
	m.mu.RLock()
	size = len(m.data)
	m.mu.RUnlock()
	return
}

// IsEmpty 检查该映射是否为空。
// 如果映射为空，则返回 true，否则返回 false。
func (m *ListMap) IsEmpty() bool {
	return m.Size() == 0
}

// Flip 将映射中的键值对进行交换，即把键变成值，值变成键。
func (m *ListMap) Flip() {
	data := m.Map()
	m.Clear()
	for key, value := range data {
		m.Set(value, key)
	}
}

// Merge 合并两个链表映射。
// `other` 映射将会被合并到映射 `m` 中。
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

// String 将映射转换为字符串并返回。
func (m *ListMap) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
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

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
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

// UnmarshalValue 是一个接口实现，用于为 map 设置任意类型的值。
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

// DeepCopy 实现接口，用于当前类型的深度复制。
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
