
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with l file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随本文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:036a875c2d7cd8b1
# <翻译结束>


<原文开始>
// Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.
<原文结束>

# <翻译开始>
// 包glist提供了最常见的双链表容器，同时也支持并发安全/不安全切换功能。. md5:0b7229b4fa0fbb49
# <翻译结束>


<原文开始>
	// List is a doubly linked list containing a concurrent-safe/unsafe switch.
	// The switch should be set when its initialization and cannot be changed then.
<原文结束>

# <翻译开始>
// List是一个包含并发安全/不安全切换的双向链表。初始化时应设置该开关，并且之后不能更改。
// md5:54c981e147e0a03a
# <翻译结束>


<原文开始>
// Element the item type of the list.
<原文结束>

# <翻译开始>
// 列表中元素的类型。. md5:f22a0215543484b0
# <翻译结束>


<原文开始>
// New creates and returns a new empty doubly linked list.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的空双向链表。. md5:d0d0b0225c460030
# <翻译结束>


<原文开始>
// NewFrom creates and returns a list from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using list in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewFrom 根据给定的切片 `array` 创建并返回一个新的列表。
// 参数 `safe` 用于指定是否在并发安全环境下使用列表，默认为 false。
// md5:bee3fb299025c2d8
# <翻译结束>


<原文开始>
// PushFront inserts a new element `e` with value `v` at the front of list `l` and returns `e`.
<原文结束>

# <翻译开始>
// PushFront 在列表 `l` 的开头插入新元素 `e`，值为 `v`，并返回 `e`。. md5:efe14f0fd31ff77b
# <翻译结束>


<原文开始>
// PushBack inserts a new element `e` with value `v` at the back of list `l` and returns `e`.
<原文结束>

# <翻译开始>
// PushBack 在列表 `l` 的末尾插入一个新元素 `e`，值为 `v`，并返回 `e`。. md5:7f490aef9df259d7
# <翻译结束>


<原文开始>
// PushFronts inserts multiple new elements with values `values` at the front of list `l`.
<原文结束>

# <翻译开始>
// PushFronts 在列表 `l` 的前端插入多个具有值 `values` 的新元素。. md5:bd169f62b7c48e7d
# <翻译结束>


<原文开始>
// PushBacks inserts multiple new elements with values `values` at the back of list `l`.
<原文结束>

# <翻译开始>
// PushBacks 将多个值为 `values` 的新元素插入到列表 `l` 的末尾。. md5:8760e724a5eb555e
# <翻译结束>


<原文开始>
// PopBack removes the element from back of `l` and returns the value of the element.
<原文结束>

# <翻译开始>
// PopBack 从 `l` 的尾部移除一个元素，并返回该元素的值。. md5:71aef7d06e374d4c
# <翻译结束>


<原文开始>
// PopFront removes the element from front of `l` and returns the value of the element.
<原文结束>

# <翻译开始>
// PopFront 从 `l` 的前端移除元素，并返回该元素的值。. md5:18dd24504d7e0084
# <翻译结束>


<原文开始>
// PopBacks removes `max` elements from back of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopBacks 从 `l` 的末尾移除 `max` 个元素，
// 并将移除元素的值作为切片返回。
// md5:100add87dc541cc8
# <翻译结束>


<原文开始>
// PopFronts removes `max` elements from front of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopFronts 从列表 `l` 的前面移除 `max` 个元素，
// 并将移除的元素值作为切片返回。
// md5:cced2abc2e709a67
# <翻译结束>


<原文开始>
// PopBackAll removes all elements from back of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopBackAll 从 `l` 的尾部移除所有元素，并将移除的元素值作为切片返回。
// md5:6fd64ee47034d8b6
# <翻译结束>


<原文开始>
// PopFrontAll removes all elements from front of `l`
// and returns values of the removed elements as slice.
<原文结束>

# <翻译开始>
// PopFrontAll 从 `l` 的前端移除所有元素，并将移除的元素值作为切片返回。
// md5:b1d251b985eb6a51
# <翻译结束>


<原文开始>
// FrontAll copies and returns values of all elements from front of `l` as slice.
<原文结束>

# <翻译开始>
// FrontAll 复制并返回列表 `l` 前端所有元素的值作为一个切片。. md5:93c8d4452c927952
# <翻译结束>


<原文开始>
// BackAll copies and returns values of all elements from back of `l` as slice.
<原文结束>

# <翻译开始>
// BackAll 复制并返回 `l` 后面所有元素的值，以切片形式返回。. md5:2dd8e946eed83cc0
# <翻译结束>


<原文开始>
// FrontValue returns value of the first element of `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// FrontValue 返回列表 `l` 的第一个元素的值，如果列表为空，则返回 nil。. md5:c70a9c11634f5a74
# <翻译结束>


<原文开始>
// BackValue returns value of the last element of `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// BackValue 返回列表`l`的最后一个元素的值，如果列表为空，则返回nil。. md5:67d80721db31a403
# <翻译结束>


<原文开始>
// Front returns the first element of list `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// Front 返回列表 `l` 的第一个元素，如果列表为空则返回 nil。. md5:24d42ffa6d3fd791
# <翻译结束>


<原文开始>
// Back returns the last element of list `l` or nil if the list is empty.
<原文结束>

# <翻译开始>
// Back 返回列表 `l` 的最后一个元素，如果列表为空则返回 nil。. md5:655654a2cad68be9
# <翻译结束>


<原文开始>
// Len returns the number of elements of list `l`.
// The complexity is O(1).
<原文结束>

# <翻译开始>
// Len 返回列表 `l` 的元素数量。
// 复杂度为 O(1)。
// md5:d2de4a4e990d787d
# <翻译结束>


<原文开始>
// MoveBefore moves element `e` to its new position before `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
<原文结束>

# <翻译开始>
// MoveBefore 将元素 `e` 移动到其新的位置，位于 `p` 之前。如果 `e` 或 `p` 不是 `l` 的元素，或者 `e` 等于 `p`，则列表不会被修改。元素 `e` 和 `p` 都不能为 nil。
// md5:b58644e1e9174539
# <翻译结束>


<原文开始>
// MoveAfter moves element `e` to its new position after `p`.
// If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified.
// The element and `p` must not be nil.
<原文结束>

# <翻译开始>
// MoveAfter 将元素 `e` 移动到 `p` 之后的新位置。
// 如果 `e` 或 `p` 不是 `l` 的元素，或者 `e` 等于 `p`，则列表不作任何修改。
// 元素 `e` 和 `p` 都不能为 nil。
// md5:18e13c9c5720547c
# <翻译结束>


<原文开始>
// MoveToFront moves element `e` to the front of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
<原文结束>

# <翻译开始>
// MoveToFront 将元素 `e` 移动到列表 `l` 的前面。
// 如果 `e` 不是 `l` 中的元素，列表将不会被修改。
// 元素必须不为 nil。
// md5:8b3809d7912952aa
# <翻译结束>


<原文开始>
// MoveToBack moves element `e` to the back of list `l`.
// If `e` is not an element of `l`, the list is not modified.
// The element must not be nil.
<原文结束>

# <翻译开始>
// MoveToBack 将元素 `e` 移动到列表 `l` 的末尾。
// 如果 `e` 不是 `l` 的元素，列表不会被修改。
// 元素不能为空。
// md5:97cb0a61b230357a
# <翻译结束>


<原文开始>
// PushBackList inserts a copy of an other list at the back of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
<原文结束>

# <翻译开始>
// PushBackList 在列表 `l` 的末尾插入另一个列表的副本。
// 列表 `l` 和 `other` 可以是相同的，但它们不能为 nil。
// md5:9bb4d2888e02946d
# <翻译结束>


<原文开始>
// PushFrontList inserts a copy of an other list at the front of list `l`.
// The lists `l` and `other` may be the same, but they must not be nil.
<原文结束>

# <翻译开始>
// PushFrontList 将另一个列表 `other` 的副本插入到列表 `l` 的前端。
// 列表 `l` 和 `other` 可以是相同的列表，但它们都不能为空。
// md5:0b7e24dd279b0ec0
# <翻译结束>


<原文开始>
// InsertAfter inserts a new element `e` with value `v` immediately after `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
<原文结束>

# <翻译开始>
// InsertAfter 在元素 `p` 之后立即插入一个新元素 `e`，其值为 `v`，并返回 `e`。
// 如果 `p` 不是 `l` 的元素，列表不会被修改。
// `p` 不能为 nil。
// md5:18fa91d04a81c29d
# <翻译结束>


<原文开始>
// InsertBefore inserts a new element `e` with value `v` immediately before `p` and returns `e`.
// If `p` is not an element of `l`, the list is not modified.
// The `p` must not be nil.
<原文结束>

# <翻译开始>
// InsertBefore 在`p`元素之前插入新元素`e`，值为`v`，然后返回`e`。
// 如果`p`不是`l`中的元素，则不修改列表。
// `p`不能为nil。
// md5:b4054a0ba93bd780
# <翻译结束>


<原文开始>
// Remove removes `e` from `l` if `e` is an element of list `l`.
// It returns the element value e.Value.
// The element must not be nil.
<原文结束>

# <翻译开始>
// Remove 从列表 `l` 中移除元素 `e`，如果 `e` 是 `l` 的元素。它返回元素的值 `e.Value`。元素必须不为 nil。
// md5:49dd42047b93518c
# <翻译结束>


<原文开始>
// Removes removes multiple elements `es` from `l` if `es` are elements of list `l`.
<原文结束>

# <翻译开始>
// Removes 从列表 `l` 中移除多个元素 `es`，前提是 `es` 是列表 `l` 的成员。. md5:19a1f18ca5d0cf06
# <翻译结束>


<原文开始>
// RemoveAll removes all elements from list `l`.
<原文结束>

# <翻译开始>
// RemoveAll 从列表 `l` 中移除所有元素。. md5:183c16a2ab7fbbfa
# <翻译结束>


<原文开始>
// Clear is alias of RemoveAll.
<原文结束>

# <翻译开始>
// Clear是RemoveAll的别名。. md5:a37765a4c78aba68
# <翻译结束>


<原文开始>
// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
<原文结束>

# <翻译开始>
// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。. md5:4ae51d9b7445f043
# <翻译结束>


<原文开始>
// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。. md5:e73dbc0381ebb3dc
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。. md5:1bfdea306db62845
# <翻译结束>


<原文开始>
// IteratorAsc iterates the list readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 按升序遍历列表，只读方式，使用给定的回调函数 `f`。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止。
// md5:0a077491be342096
# <翻译结束>


<原文开始>
// IteratorDesc iterates the list readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历列表，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:b9a7d34f2e3426a7
# <翻译结束>


<原文开始>
// Join joins list elements with a string `glue`.
<原文结束>

# <翻译开始>
// Join使用字符串`glue`将list元素连接起来。. md5:daf9e3877e4dd942
# <翻译结束>


<原文开始>
// String returns current list as a string.
<原文结束>

# <翻译开始>
// String 将当前列表作为字符串返回。. md5:e5f56499b5c2f331
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for list.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置到列表中。. md5:a6e906ab9decb788
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
# <翻译结束>

