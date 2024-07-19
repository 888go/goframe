// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with l file,
// You can obtain one at https://github.com/gogf/gf.
//

// Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.
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
	// List is a doubly linked list containing a concurrent-safe/unsafe switch.
	// The switch should be set when its initialization and cannot be changed then.
	List struct {
		mu   rwmutex.RWMutex
		list *list.List
	}
	// Element the item type of the list.
	Element = list.Element
)

// New creates and returns a new empty doubly linked list.
// ff:
// safe:
func New(safe ...bool) *List {
	return &List{
		mu:   rwmutex.Create(safe...),
		list: list.New(),
	}
}

// NewFrom creates and returns a list from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using list in concurrent-safety,
// which is false in default.
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

// PushFront inserts a new element `e` with value `v` at the front of list `l` and returns `e`.
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

// PushBack inserts a new element `e` with value `v` at the back of list `l` and returns `e`.
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

// PushFronts inserts multiple new elements with values `values` at the front of list `l`.
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

// PushBacks inserts multiple new elements with values `values` at the back of list `l`.
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

// PopBack removes the element from back of `l` and returns the value of the element.
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

// PopFront removes the element from front of `l` and returns the value of the element.
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

// PopBacks removes `max` elements from back of `l`
// and returns values of the removed elements as slice.
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

// PopFronts removes `max` elements from front of `l`
// and returns values of the removed elements as slice.
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

// PopBackAll removes all elements from back of `l`
// and returns values of the removed elements as slice.
// ff:
// l:
func (l *List) PopBackAll() []interface{} {
	return l.PopBacks(-1)
}

// PopFrontAll removes all elements from front of `l`
// and returns values of the removed elements as slice.
// ff:
// l:
func (l *List) PopFrontAll() []interface{} {
	return l.PopFronts(-1)
}

// FrontAll copies and returns values of all elements from front of `l` as slice.
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

// BackAll copies and returns values of all elements from back of `l` as slice.
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

// FrontValue returns value of the first element of `l` or nil if the list is empty.
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

// BackValue returns value of the last element of `l` or nil if the list is empty.
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

// Front returns the first element of list `l` or nil if the list is empty.
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

// Back returns the last element of list `l` or nil if the list is empty.
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

// Len returns the number of elements of list `l`.
// The complexity is O(1).
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

// MoveBefore moves element `e` to its new position before `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
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

// MoveAfter moves element `e` to its new position after `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
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

// MoveToFront moves element `e` to the front of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
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

// MoveToBack moves element `e` to the back of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
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

// PushBackList inserts a copy of an other list at the back of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
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

// PushFrontList inserts a copy of an other list at the front of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
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

// InsertAfter inserts a new element `e` with value `v` immediately after `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
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

// InsertBefore inserts a new element `e` with value `v` immediately before `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
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

// Remove removes `e` from `l` if `e` is an element of list `l`.
// It returns the element value e.Value.
// The element must not be nil.
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

// Removes removes multiple elements `es` from `l` if `es` are elements of list `l`.
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

// RemoveAll removes all elements from list `l`.
// ff:
// l:
func (l *List) RemoveAll() {
	l.mu.Lock()
	l.list = list.New()
	l.mu.Unlock()
}

// Clear is alias of RemoveAll.
// ff:
// l:
func (l *List) Clear() {
	l.RemoveAll()
}

// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
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

// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
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

// Iterator is alias of IteratorAsc.
// yx:true
// ff:X遍历
// l:
// f:
// e:
func (l *List) Iterator(f func(e *Element) bool) {
	l.IteratorAsc(f)
}

// IteratorAsc iterates the list readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
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

// IteratorDesc iterates the list readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
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

// Join joins list elements with a string `glue`.
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

// String returns current list as a string.
// ff:
// l:
func (l *List) String() string {
	if l == nil {
		return ""
	}
	return "[" + l.Join(",") + "]"
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// ff:
// l:
func (l List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.FrontAll())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
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

// UnmarshalValue is an interface implement which sets any type of value for list.
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

// DeepCopy implements interface for deep copy of current type.
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
