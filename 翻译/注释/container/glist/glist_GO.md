
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with l file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT协议条款。如果随同本文件未分发一份MIT协议副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.
<原文结束>

# <翻译开始>
// 包glist提供了最常用的双向链表容器，同时支持并发安全/非安全模式切换功能。
# <翻译结束>


<原文开始>
	// List is a doubly linked list containing a concurrent-safe/unsafe switch.
	// The switch should be set when its initialization and cannot be changed then.
<原文结束>

# <翻译开始>
// List 是一个包含并发安全/不安全切换功能的双向链表。
// 这个切换开关应在初始化时设定，并且之后不能更改。
// 在这里，“并发安全”意味着这个链表在多线程或协程环境下可以安全地进行读写操作，而“并发不安全”则表示在未采取额外同步措施的情况下，同时访问可能会导致数据竞争问题。初始化后不允许改变这个安全属性设置。
# <翻译结束>


<原文开始>
// Element the item type of the list.
<原文结束>

# <翻译开始>
// Element 表示列表中元素的类型。
# <翻译结束>


<原文开始>
// New creates and returns a new empty doubly linked list.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的空双向链表。
# <翻译结束>


<原文开始>
// NewFrom creates and returns a list from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using list in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewFrom 函数通过复制给定的切片 `array` 创建并返回一个列表。
// 参数 `safe` 用于指定是否在并发安全模式下使用该列表，默认情况下为 false。
# <翻译结束>


<原文开始>
// PushFront inserts a new element `e` with value `v` at the front of list `l` and returns `e`.
<原文结束>

# <翻译开始>
// PushFront在列表`l`的前端插入一个具有值`v`的新元素`e`，并返回`e`。
# <翻译结束>


<原文开始>
// PushBack inserts a new element `e` with value `v` at the back of list `l` and returns `e`.
<原文结束>

# <翻译开始>
// PushBack在列表`l`的尾部插入一个新元素，并将该元素的值设为`v`，然后返回这个新插入的元素`e`。
# <翻译结束>


<原文开始>
// PushFronts inserts multiple new elements with values `values` at the front of list `l`.
<原文结束>

# <翻译开始>
// PushFronts 在列表 `l` 的前端插入多个新元素，这些元素的值为 `values`。
# <翻译结束>


<原文开始>
// PushBacks inserts multiple new elements with values `values` at the back of list `l`.
<原文结束>

# <翻译开始>
// PushBacks在列表`l`的尾部插入多个新元素，其值为`values`。
# <翻译结束>


<原文开始>
// PopBack removes the element from back of `l` and returns the value of the element.
<原文结束>

# <翻译开始>
// PopBack从`l`的末尾移除元素，并返回该元素的值。
# <翻译结束>


<原文开始>
// PopFront removes the element from front of `l` and returns the value of the element.
<原文结束>

# <翻译开始>
// PopFront从`l`的前端移除元素，并返回该元素的值。
# <翻译结束>


<原文开始>
// PopBacks removes `max` elements from back of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopBacks 从 `l` 的尾部移除最多 `max` 个元素，
// 并将移除元素的值以切片形式返回。
# <翻译结束>


<原文开始>
// PopFronts removes `max` elements from front of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopFronts 从 `l`（列表）的前端移除最多`max`个元素，
// 并将已移除元素的值以切片形式返回。
# <翻译结束>


<原文开始>
// PopBackAll removes all elements from back of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopBackAll 从 `l` 的后部移除所有元素，并将移除元素的值以切片形式返回。
# <翻译结束>


<原文开始>
// PopFrontAll removes all elements from front of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopFrontAll从`l`的前端移除所有元素，并将已移除元素的值以切片形式返回。
# <翻译结束>


<原文开始>
// FrontAll copies and returns values of all elements from front of `l` as slice.
<原文结束>

# <翻译开始>
// FrontAll 从切片 `l` 的开头复制并返回所有元素的值作为新的切片。
# <翻译结束>


<原文开始>
// BackAll copies and returns values of all elements from back of `l` as slice.
<原文结束>

# <翻译开始>
// BackAll 从切片 `l` 的末尾开始复制所有元素的值，并以一个新的切片形式返回。
# <翻译结束>


<原文开始>
// FrontValue returns value of the first element of `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// FrontValue 返回 `l`（链表）的第一个元素的值，如果链表为空则返回 nil。
# <翻译结束>


<原文开始>
// BackValue returns value of the last element of `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// BackValue 返回 `l` 列表最后一个元素的值，如果列表为空则返回 nil。
# <翻译结束>


<原文开始>
// Front returns the first element of list `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// Front 返回列表 `l` 的第一个元素，如果列表为空，则返回 nil。
# <翻译结束>


<原文开始>
// Back returns the last element of list `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// Back 返回列表 `l` 的最后一个元素，如果列表为空则返回 nil。
# <翻译结束>


<原文开始>
// Len returns the number of elements of list `l`.
// The complexity is O(1).
<原文结束>

# <翻译开始>
// Len 返回列表 `l` 的元素个数。
// 时间复杂度为 O(1)。
# <翻译结束>







<原文开始>
// MoveBefore moves element `e` to its new position before `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
<原文结束>

# <翻译开始>
// MoveBefore 将元素 `e` 移动到其在 `p` 之前的新位置。
// 如果 `e` 或 `p` 不是 `l` 列表中的元素，或者 `e` 和 `p` 相等，则列表不会被修改。
// 元素 `e` 和 `p` 都必须不为 nil。
# <翻译结束>


<原文开始>
// MoveAfter moves element `e` to its new position after `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
<原文结束>

# <翻译开始>
// MoveAfter 将元素 `e` 移动到新位置，即在元素 `p` 之后。
// 如果 `e` 或 `p` 不是列表 `l` 的成员，或者 `e` 和 `p` 相等，则列表不会被修改。
// 元素 `e` 和 `p` 都不能为空（nil）。
# <翻译结束>


<原文开始>
// MoveToFront moves element `e` to the front of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
<原文结束>

# <翻译开始>
// MoveToFront 将元素 `e` 移动到列表 `l` 的前端。
// 如果 `e` 不是列表 `l` 的一个元素，则列表不会被修改。
// 此元素必须不为空（nil）。
# <翻译结束>


<原文开始>
// MoveToBack moves element `e` to the back of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
<原文结束>

# <翻译开始>
// MoveToBack 将元素 `e` 移动到列表 `l` 的末尾。
// 如果 `e` 不是列表 `l` 中的元素，则列表不会被修改。
// 此元素必须不为 nil。
# <翻译结束>


<原文开始>
// PushBackList inserts a copy of an other list at the back of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
<原文结束>

# <翻译开始>
// PushBackList将另一个列表的副本插入到列表`l`的末尾。
// 列表`l`和`other`可以是同一个列表，但它们都必须不为nil。
# <翻译结束>


<原文开始>
// PushFrontList inserts a copy of an other list at the front of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
<原文结束>

# <翻译开始>
// PushFrontList 将另一个列表的副本插入到列表 `l` 的前端。
// 列表 `l` 和 `other` 可能相同，但它们都必须不为 nil。
# <翻译结束>


<原文开始>
// InsertAfter inserts a new element `e` with value `v` immediately after `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
<原文结束>

# <翻译开始>
// InsertAfter 在元素 `p` 后立即插入一个新元素 `e`，其值为 `v`，并返回 `e`。
// 若 `p` 不是列表 `l` 的成员，则列表不会被修改。
// 参数 `p` 必须不为空。
# <翻译结束>


<原文开始>
// InsertBefore inserts a new element `e` with value `v` immediately before `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
<原文结束>

# <翻译开始>
// InsertBefore在元素`p`之前立即插入一个新元素`e`，其值为`v`并返回`e`。
// 如果`p`不是列表`l`中的元素，则列表不会被修改。
// 参数`p`必须不为nil。
# <翻译结束>


<原文开始>
// Remove removes `e` from `l` if `e` is an element of list `l`.
// It returns the element value e.Value.
// The element must not be nil.
<原文结束>

# <翻译开始>
// Remove从列表`l`中删除元素`e`，如果`e`是列表`l`中的一个元素。
// 它返回元素的值e.Value。
// 该元素必须不为nil。
# <翻译结束>


<原文开始>
// Removes removes multiple elements `es` from `l` if `es` are elements of list `l`.
<原文结束>

# <翻译开始>
// Removes 从列表 `l` 中移除多个元素 `es`，条件是 `es` 是列表 `l` 中的元素。
# <翻译结束>


<原文开始>
// RemoveAll removes all elements from list `l`.
<原文结束>

# <翻译开始>
// RemoveAll 从列表 `l` 中移除所有元素。
# <翻译结束>


<原文开始>
// Clear is alias of RemoveAll.
<原文结束>

# <翻译开始>
// Clear 是 RemoveAll 的别名。
# <翻译结束>


<原文开始>
// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
<原文结束>

# <翻译开始>
// RLockFunc 在 RWMutex.RLock 内使用给定的回调函数 `f` 进行读取锁定。
# <翻译结束>


<原文开始>
// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 内锁定写入操作。
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。
# <翻译结束>


<原文开始>
// IteratorAsc iterates the list readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 以升序遍历列表，并使用给定的回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// IteratorDesc iterates the list readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历给定的只读列表，并使用回调函数 `f` 进行处理。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
# <翻译结束>


<原文开始>
// Join joins list elements with a string `glue`.
<原文结束>

# <翻译开始>
// Join通过字符串`glue`连接列表元素。
# <翻译结束>


<原文开始>
// String returns current list as a string.
<原文结束>

# <翻译开始>
// String 函数返回当前列表作为字符串表示。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for list.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为列表设置任何类型的值。
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>


<原文开始>
// Size is alias of Len.
<原文结束>

# <翻译开始>
// Size 是 Len 的别名。
# <翻译结束>

