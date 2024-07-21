// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随本文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:036a875c2d7cd8b1

// 包glist提供了最常见的双链表容器，同时也支持并发安全/不安全切换功能。 md5:0b7229b4fa0fbb49
package glist//bm:链表类

import (
	"bytes"
	"container/list"

	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	// List是一个包含并发安全/不安全切换的双向链表。初始化时应设置该开关，并且之后不能更改。
	// md5:54c981e147e0a03a
	List struct {
		mu   rwmutex.RWMutex
		list *list.List
	}
	// 列表中元素的类型。 md5:f22a0215543484b0
	Element = list.Element
)

// New 创建并返回一个新的空双向链表。 md5:d0d0b0225c460030
// ff:
// safe:
func New(safe ...bool) *List {
	return &List{
		mu:   rwmutex.Create(safe...),
		list: list.New(),
	}
}

// NewFrom 根据给定的切片 `array` 创建并返回一个新的列表。
// 参数 `safe` 用于指定是否在并发安全环境下使用列表，默认为 false。
// md5:bee3fb299025c2d8
// ff:
// array:
// safe:
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

// PushFront 在列表 `l` 的开头插入新元素 `e`，值为 `v`，并返回 `e`。 md5:efe14f0fd31ff77b
// ff:
// l:
// v:
// e:
func (l *List) PushFront(v interface{}) (e *Element) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.PushFront(v)
	l.mu.Unlock()
	return
}

// PushBack 在列表 `l` 的末尾插入一个新元素 `e`，值为 `v`，并返回 `e`。 md5:7f490aef9df259d7
// ff:
// l:
// v:
// e:
func (l *List) PushBack(v interface{}) (e *Element) {
	l.mu.Lock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.PushBack(v)
	l.mu.Unlock()
	return
}

// PushFronts 在列表 `l` 的前端插入多个具有值 `values` 的新元素。 md5:bd169f62b7c48e7d
// ff:
// l:
// values:
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

// PushBacks 将多个值为 `values` 的新元素插入到列表 `l` 的末尾。 md5:8760e724a5eb555e
// ff:
// l:
// values:
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

// PopBack 从 `l` 的尾部移除一个元素，并返回该元素的值。 md5:71aef7d06e374d4c
// ff:
// l:
// value:
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

// PopFront 从 `l` 的前端移除元素，并返回该元素的值。 md5:18dd24504d7e0084
// ff:
// l:
// value:
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

// PopBacks 从 `l` 的末尾移除 `max` 个元素，
// 并将移除元素的值作为切片返回。
// md5:100add87dc541cc8
// ff:
// l:
// max:
// values:
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

// PopFronts 从列表 `l` 的前面移除 `max` 个元素，
// 并将移除的元素值作为切片返回。
// md5:cced2abc2e709a67
// ff:
// l:
// max:
// values:
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

// PopBackAll 从 `l` 的尾部移除所有元素，并将移除的元素值作为切片返回。
// md5:6fd64ee47034d8b6
// ff:
// l:
func (l *List) PopBackAll() []interface{} {
	return l.PopBacks(-1)
}

// PopFrontAll 从 `l` 的前端移除所有元素，并将移除的元素值作为切片返回。
// md5:b1d251b985eb6a51
// ff:
// l:
func (l *List) PopFrontAll() []interface{} {
	return l.PopFronts(-1)
}

// FrontAll 复制并返回列表 `l` 前端所有元素的值作为一个切片。 md5:93c8d4452c927952
// ff:
// l:
// values:
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

// BackAll 复制并返回 `l` 后面所有元素的值，以切片形式返回。 md5:2dd8e946eed83cc0
// ff:
// l:
// values:
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

// FrontValue 返回列表 `l` 的第一个元素的值，如果列表为空，则返回 nil。 md5:c70a9c11634f5a74
// ff:
// l:
// value:
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

// BackValue 返回列表`l`的最后一个元素的值，如果列表为空，则返回nil。 md5:67d80721db31a403
// ff:
// l:
// value:
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

// Front 返回列表 `l` 的第一个元素，如果列表为空则返回 nil。 md5:24d42ffa6d3fd791
// ff:
// l:
// e:
func (l *List) Front() (e *Element) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	e = l.list.Front()
	return
}

// Back 返回列表 `l` 的最后一个元素，如果列表为空则返回 nil。 md5:655654a2cad68be9
// ff:
// l:
// e:
func (l *List) Back() (e *Element) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	e = l.list.Back()
	return
}

// Len 返回列表 `l` 的元素数量。
// 复杂度为 O(1)。
// md5:d2de4a4e990d787d
// ff:
// l:
// length:
func (l *List) Len() (length int) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list == nil {
		return
	}
	length = l.list.Len()
	return
}

// Size is alias of Len.
// ff:
// l:
func (l *List) Size() int {
	return l.Len()
}

// MoveBefore 将元素 `e` 移动到其新的位置，位于 `p` 之前。如果 `e` 或 `p` 不是 `l` 的元素，或者 `e` 等于 `p`，则列表不会被修改。元素 `e` 和 `p` 都不能为 nil。
// md5:b58644e1e9174539
// ff:
// l:
// e:
// p:
func (l *List) MoveBefore(e, p *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveBefore(e, p)
}

// MoveAfter 将元素 `e` 移动到 `p` 之后的新位置。
// 如果 `e` 或 `p` 不是 `l` 的元素，或者 `e` 等于 `p`，则列表不作任何修改。
// 元素 `e` 和 `p` 都不能为 nil。
// md5:18e13c9c5720547c
// ff:
// l:
// e:
// p:
func (l *List) MoveAfter(e, p *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveAfter(e, p)
}

// MoveToFront 将元素 `e` 移动到列表 `l` 的前面。
// 如果 `e` 不是 `l` 中的元素，列表将不会被修改。
// 元素必须不为 nil。
// md5:8b3809d7912952aa
// ff:
// l:
// e:
func (l *List) MoveToFront(e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveToFront(e)
}

// MoveToBack 将元素 `e` 移动到列表 `l` 的末尾。
// 如果 `e` 不是 `l` 的元素，列表不会被修改。
// 元素不能为空。
// md5:97cb0a61b230357a
// ff:
// l:
// e:
func (l *List) MoveToBack(e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	l.list.MoveToBack(e)
}

// PushBackList 在列表 `l` 的末尾插入另一个列表的副本。
// 列表 `l` 和 `other` 可以是相同的，但它们不能为 nil。
// md5:9bb4d2888e02946d
// ff:
// l:
// other:
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

// PushFrontList 将另一个列表 `other` 的副本插入到列表 `l` 的前端。
// 列表 `l` 和 `other` 可以是相同的列表，但它们都不能为空。
// md5:0b7e24dd279b0ec0
// ff:
// l:
// other:
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

// InsertAfter 在元素 `p` 之后立即插入一个新元素 `e`，其值为 `v`，并返回 `e`。
// 如果 `p` 不是 `l` 的元素，列表不会被修改。
// `p` 不能为 nil。
// md5:18fa91d04a81c29d
// ff:
// l:
// p:
// v:
// e:
func (l *List) InsertAfter(p *Element, v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.InsertAfter(v, p)
	return
}

// InsertBefore 在`p`元素之前插入新元素`e`，值为`v`，然后返回`e`。
// 如果`p`不是`l`中的元素，则不修改列表。
// `p`不能为nil。
// md5:b4054a0ba93bd780
// ff:
// l:
// p:
// v:
// e:
func (l *List) InsertBefore(p *Element, v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	e = l.list.InsertBefore(v, p)
	return
}

// Remove 从列表 `l` 中移除元素 `e`，如果 `e` 是 `l` 的元素。它返回元素的值 `e.Value`。元素必须不为 nil。
// md5:49dd42047b93518c
// ff:
// l:
// e:
// value:
func (l *List) Remove(e *Element) (value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	value = l.list.Remove(e)
	return
}

// Removes 从列表 `l` 中移除多个元素 `es`，前提是 `es` 是列表 `l` 的成员。 md5:19a1f18ca5d0cf06
// ff:
// l:
// es:
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

// RemoveAll 从列表 `l` 中移除所有元素。 md5:183c16a2ab7fbbfa
// ff:
// l:
func (l *List) RemoveAll() {
	l.mu.Lock()
	l.list = list.New()
	l.mu.Unlock()
}

// Clear是RemoveAll的别名。 md5:a37765a4c78aba68
// ff:
// l:
func (l *List) Clear() {
	l.RemoveAll()
}

// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
// ff:
// l:
// f:
// list:
func (l *List) RLockFunc(f func(list *list.List)) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.list != nil {
		f(l.list)
	}
}

// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
// ff:
// l:
// f:
// list:
func (l *List) LockFunc(f func(list *list.List)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	f(l.list)
}

// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
// yx:true
// ff:X遍历
// l:
// f:
// e:
func (l *List) Iterator(f func(e *Element) bool) {
	l.IteratorAsc(f)
}

// IteratorAsc 按升序遍历列表，只读方式，使用给定的回调函数 `f`。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止。
// md5:0a077491be342096
// ff:
// l:
// f:
// e:
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

// IteratorDesc 以降序方式遍历列表，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:b9a7d34f2e3426a7
// ff:
// l:
// f:
// e:
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

// Join使用字符串`glue`将list元素连接起来。 md5:daf9e3877e4dd942
// ff:
// l:
// glue:
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
			buffer.WriteString(gconv.String(e.Value))
			if i != length-1 {
				buffer.WriteString(glue)
			}
		}
	}
	return buffer.String()
}

// String 将当前列表作为字符串返回。 md5:e5f56499b5c2f331
// ff:
// l:
func (l *List) String() string {
	if l == nil {
		return ""
	}
	return "[" + l.Join(",") + "]"
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
// ff:
// l:
func (l List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.FrontAll())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
// ff:
// l:
// b:
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

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置到列表中。 md5:a6e906ab9decb788
// ff:
// l:
// value:
// err:
func (l *List) UnmarshalValue(value interface{}) (err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.list == nil {
		l.list = list.New()
	}
	var array []interface{}
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &array)
	default:
		array = gconv.SliceAny(value)
	}
	l.PushBacks(array)
	return err
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
// ff:
// l:
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
