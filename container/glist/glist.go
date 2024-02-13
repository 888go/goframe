// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT协议条款。如果随同本文件未分发一份MIT协议副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

// 包glist提供了最常用的双向链表容器，同时支持并发安全/非安全模式切换功能。
package 链表类

import (
	"bytes"
	"container/list"
	
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
)

type (
// List 是一个包含并发安全/不安全切换功能的双向链表。
// 这个切换开关应在初始化时设定，并且之后不能更改。
// 在这里，“并发安全”意味着这个链表在多线程或协程环境下可以安全地进行读写操作，而“并发不安全”则表示在未采取额外同步措施的情况下，同时访问可能会导致数据竞争问题。初始化后不允许改变这个安全属性设置。
	List struct {
		mu   rwmutex.RWMutex
		list *list.List
	}
	// Element 表示列表中元素的类型。
	Element = list.Element
)

// New 创建并返回一个新的空双向链表。
func New(safe ...bool) *List {
	return &List{
		mu:   rwmutex.Create(safe...),
		list: list.New(),
	}
}

// NewFrom 函数通过复制给定的切片 `array` 创建并返回一个列表。
// 参数 `safe` 用于指定是否在并发安全模式下使用该列表，默认情况下为 false。
func NewFrom(array []interface{}, safe ...bool) *List {
	l := list.New()
	for _, v := range array {
		l.PushBack(v)
	}
	return &List{
		mu:   rwmutex.Create(safe...),
		list: l,
	}
}

// PushFront在列表`l`的前端插入一个具有值`v`的新元素`e`，并返回`e`。
func (l *List) PushFront(v interface{}) (e *Element) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.PushFront(v)
	l.mu.Unlock()
	return
}

// PushBack在列表`l`的尾部插入一个新元素，并将该元素的值设为`v`，然后返回这个新插入的元素`e`。
func (l *List) PushBack(v interface{}) (e *Element) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.PushBack(v)
	l.mu.Unlock()
	return
}

// PushFronts 在列表 `l` 的前端插入多个新元素，这些元素的值为 `values`。
func (l *List) PushFronts(values []interface{}) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	for _, v := range values {
		l.list.PushFront(v)
	}
	l.mu.Unlock()
}

// PushBacks在列表`l`的尾部插入多个新元素，其值为`values`。
func (l *List) PushBacks(values []interface{}) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	for _, v := range values {
		l.list.PushBack(v)
	}
	l.mu.Unlock()
}

// PopBack从`l`的末尾移除元素，并返回该元素的值。
func (l *List) PopBack() (value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
		return
	}
	if e := l.list.Back(); e != nil {
		value = l.list.Remove(e)
	}
	return
}

// PopFront从`l`的前端移除元素，并返回该元素的值。
func (l *List) PopFront() (value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
		return
	}
	if e := l.list.Front(); e != nil {
		value = l.list.Remove(e)
	}
	return
}

// PopBacks 从 `l` 的尾部移除最多 `max` 个元素，
// 并将移除元素的值以切片形式返回。
func (l *List) PopBacks(max int) (values []interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
		return
	}
	length := l.list.Len()
	if length > 0 {
		if max > 0 && max < length {
			length = max
		}
		values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			values[i] = l.list.Remove(l.list.Back())
		}
	}
	return
}

// PopFronts 从 `l`（列表）的前端移除最多`max`个元素，
// 并将已移除元素的值以切片形式返回。
func (l *List) PopFronts(max int) (values []interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
		return
	}
	length := l.list.Len()
	if length > 0 {
		if max > 0 && max < length {
			length = max
		}
		values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			values[i] = l.list.Remove(l.list.Front())
		}
	}
	return
}

// PopBackAll 从 `l` 的后部移除所有元素，并将移除元素的值以切片形式返回。
func (l *List) PopBackAll() []interface{} {
	return l.PopBacks(-1)
}

// PopFrontAll从`l`的前端移除所有元素，并将已移除元素的值以切片形式返回。
func (l *List) PopFrontAll() []interface{} {
	return l.PopFronts(-1)
}

// FrontAll 从切片 `l` 的开头复制并返回所有元素的值作为新的切片。
func (l *List) FrontAll() (values []interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length := l.list.Len()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			values[i] = e.Value
		}
	}
	return
}

// BackAll 从切片 `l` 的末尾开始复制所有元素的值，并以一个新的切片形式返回。
func (l *List) BackAll() (values []interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length := l.list.Len()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, l.list.Back(); i < length; i, e = i+1, e.Prev() {
			values[i] = e.Value
		}
	}
	return
}

// FrontValue 返回 `l`（链表）的第一个元素的值，如果链表为空则返回 nil。
func (l *List) FrontValue() (value interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	if e := l.list.Front(); e != nil {
		value = e.Value
	}
	return
}

// BackValue 返回 `l` 列表最后一个元素的值，如果列表为空则返回 nil。
func (l *List) BackValue() (value interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	if e := l.list.Back(); e != nil {
		value = e.Value
	}
	return
}

// Front 返回列表 `l` 的第一个元素，如果列表为空，则返回 nil。
func (l *List) Front() (e *Element) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	e = l.list.Front()
	return
}

// Back 返回列表 `l` 的最后一个元素，如果列表为空则返回 nil。
func (l *List) Back() (e *Element) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	e = l.list.Back()
	return
}

// Len 返回列表 `l` 的元素个数。
// 时间复杂度为 O(1)。
func (l *List) Len() (length int) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length = l.list.Len()
	return
}

// Size 是 Len 的别名。
func (l *List) Size() int {
	return l.Len()
}

// MoveBefore 将元素 `e` 移动到其在 `p` 之前的新位置。
// 如果 `e` 或 `p` 不是 `l` 列表中的元素，或者 `e` 和 `p` 相等，则列表不会被修改。
// 元素 `e` 和 `p` 都必须不为 nil。
func (l *List) MoveBefore(e, p *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveBefore(e, p)
}

// MoveAfter 将元素 `e` 移动到新位置，即在元素 `p` 之后。
// 如果 `e` 或 `p` 不是列表 `l` 的成员，或者 `e` 和 `p` 相等，则列表不会被修改。
// 元素 `e` 和 `p` 都不能为空（nil）。
func (l *List) MoveAfter(e, p *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveAfter(e, p)
}

// MoveToFront 将元素 `e` 移动到列表 `l` 的前端。
// 如果 `e` 不是列表 `l` 的一个元素，则列表不会被修改。
// 此元素必须不为空（nil）。
func (l *List) MoveToFront(e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveToFront(e)
}

// MoveToBack 将元素 `e` 移动到列表 `l` 的末尾。
// 如果 `e` 不是列表 `l` 中的元素，则列表不会被修改。
// 此元素必须不为 nil。
func (l *List) MoveToBack(e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveToBack(e)
}

// PushBackList将另一个列表的副本插入到列表`l`的末尾。
// 列表`l`和`other`可以是同一个列表，但它们都必须不为nil。
func (l *List) PushBackList(other *List) {
	if l != other {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.PushBackList(other.list)
}

// PushFrontList 将另一个列表的副本插入到列表 `l` 的前端。
// 列表 `l` 和 `other` 可能相同，但它们都必须不为 nil。
func (l *List) PushFrontList(other *List) {
	if l != other {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.PushFrontList(other.list)
}

// InsertAfter 在元素 `p` 后立即插入一个新元素 `e`，其值为 `v`，并返回 `e`。
// 若 `p` 不是列表 `l` 的成员，则列表不会被修改。
// 参数 `p` 必须不为空。
func (l *List) InsertAfter(p *Element, v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.InsertAfter(v, p)
	return
}

// InsertBefore在元素`p`之前立即插入一个新元素`e`，其值为`v`并返回`e`。
// 如果`p`不是列表`l`中的元素，则列表不会被修改。
// 参数`p`必须不为nil。
func (l *List) InsertBefore(p *Element, v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.InsertBefore(v, p)
	return
}

// Remove从列表`l`中删除元素`e`，如果`e`是列表`l`中的一个元素。
// 它返回元素的值e.Value。
// 该元素必须不为nil。
func (l *List) Remove(e *Element) (value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	value = l.list.Remove(e)
	return
}

// Removes 从列表 `l` 中移除多个元素 `es`，条件是 `es` 是列表 `l` 中的元素。
func (l *List) Removes(es []*Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	for _, e := range es {
		l.list.Remove(e)
	}
}

// RemoveAll 从列表 `l` 中移除所有元素。
func (l *List) RemoveAll() {
	l.mu.Lock()
	l.list = list.New()
	l.mu.Unlock()
}

// Clear 是 RemoveAll 的别名。
func (l *List) Clear() {
	l.RemoveAll()
}

// RLockFunc 在 RWMutex.RLock 内使用给定的回调函数 `f` 进行读取锁定。
func (l *List) RLockFunc(f func(list *list.List)) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list != nil {
		f(l.list)
	}
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 内锁定写入操作。
func (l *List) LockFunc(f func(list *list.List)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	f(l.list)
}

// Iterator 是 IteratorAsc 的别名。
func (l *List) X遍历(f func(e *Element) bool) {
	l.IteratorAsc(f)
}

// IteratorAsc 以升序遍历列表，并使用给定的回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (l *List) IteratorAsc(f func(e *Element) bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length := l.list.Len()
	if length > 0 {
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			if !f(e) {
				break
			}
		}
	}
}

// IteratorDesc 以降序方式遍历给定的只读列表，并使用回调函数 `f` 进行处理。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
func (l *List) IteratorDesc(f func(e *Element) bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length := l.list.Len()
	if length > 0 {
		for i, e := 0, l.list.Back(); i < length; i, e = i+1, e.Prev() {
			if !f(e) {
				break
			}
		}
	}
}

// Join通过字符串`glue`连接列表元素。
func (l *List) Join(glue string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	length := l.list.Len()
	if length > 0 {
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			buffer.WriteString(转换类.String(e.Value))
			if i != length-1 {
				buffer.WriteString(glue)
			}
		}
	}
	return buffer.String()
}

// String 函数返回当前列表作为字符串表示。
func (l *List) String() string {
	if l == nil {
		return ""
	}
	return "[" + l.Join(",") + "]"
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (l List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.FrontAll())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (l *List) UnmarshalJSON(b []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	var array []interface{}
	if err := json.UnmarshalUseNumber(b, &array); err != nil {
		return err
	}
	l.PushBacks(array)
	return nil
}

// UnmarshalValue 是一个接口实现，用于为列表设置任何类型的值。
func (l *List) UnmarshalValue(value interface{}) (err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	var array []interface{}
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(转换类.X取字节集(value), &array)
	default:
		array = 转换类.SliceAny别名(value)
	}
	l.PushBacks(array)
	return err
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (l *List) DeepCopy() interface{} {
	if l == nil {
		return nil
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.list == nil {
		return nil
	}
	var (
		length = l.list.Len()
		values = make([]interface{}, length)
	)
	if length > 0 {
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			values[i] = deepcopy.Copy(e.Value)
		}
	}
	return NewFrom(values, l.mu.IsSafe())
}
