
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gring provides a concurrent-safe/unsafe ring(circular lists).
<原文结束>

# <翻译开始>
// gring包提供了一种并发安全/不安全的循环链表（环形列表）。
# <翻译结束>


<原文开始>
// Ring is a struct of ring structure.
<原文结束>

# <翻译开始>
// Ring 是一个环形结构的结构体。
# <翻译结束>


<原文开始>
// internalRingItem stores the ring element value.
<原文结束>

# <翻译开始>
// internalRingItem 存储环形元素的值。
# <翻译结束>


<原文开始>
// New creates and returns a Ring structure of `cap` elements.
// The optional parameter `safe` specifies whether using this structure in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个容量为`cap`的Ring结构体。
// 可选参数`sage`用于指定该结构体是否在并发场景下安全使用，默认为false（不安全）。
// func New(cap int, safe ...bool) *Ring {
    // ...
// }
# <翻译结束>


<原文开始>
// Val returns the item's value of current position.
<原文结束>

# <翻译开始>
// Val 返回当前位置项的值。
# <翻译结束>


<原文开始>
// Len returns the size of ring.
<原文结束>

# <翻译开始>
// Len 返回环形结构的大小。
# <翻译结束>


<原文开始>
// Cap returns the capacity of ring.
<原文结束>

# <翻译开始>
// Cap 返回环形缓冲区的容量。
# <翻译结束>


<原文开始>
// Checks and updates the len and cap of ring when ring is dirty.
<原文结束>

# <翻译开始>
// 当ring脏时，检查并更新ring的长度(len)和容量(cap)。
# <翻译结束>


<原文开始>
// Set sets value to the item of current position.
<原文结束>

# <翻译开始>
// Set 将值设置为当前位置的项。
# <翻译结束>


<原文开始>
// Put sets `value` to current item of ring and moves position to next item.
<原文结束>

# <翻译开始>
// Put 将 `value` 设置为环形结构当前项的值，并将位置移动到下一个项。
# <翻译结束>


<原文开始>
// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Move 函数将循环队列中的元素向后（n < 0）或向前（n >= 0）移动 n % r.Len() 个位置，
// 并返回该移动后的位置上的元素。注意，r 不得为空。
# <翻译结束>


<原文开始>
// Prev returns the previous ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Prev 返回上一个环形元素。r不能为空。
# <翻译结束>


<原文开始>
// Next returns the next ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Next 返回下一个环形元素。r 必须非空。
# <翻译结束>


<原文开始>
// Link connects ring r with ring s such that r.Next()
// becomes s and returns the original value for r.Next().
// r must not be empty.
//
// If r and s point to the same ring, linking
// them removes the elements between r and s from the ring.
// The removed elements form a sub-ring and the result is a
// reference to that sub-ring (if no elements were removed,
// the result is still the original value for r.Next(),
// and not nil).
//
// If r and s point to different rings, linking
// them creates a single ring with the elements of s inserted
// after r. The result points to the element following the
// last element of s after insertion.
<原文结束>

# <翻译开始>
// Link 将环 r 与环 s 连接，使得 r.Next() 指向 s，并返回连接前 r.Next() 的原始值。
// r 必须非空。
//
// 如果 r 和 s 指向同一个环，将它们连接会从环中移除 r 和 s 之间的元素。被移除的元素形成一个子环，结果是对该子环的一个引用（如果未移除任何元素，则结果仍然是原始的 r.Next() 值，而非 nil）。
//
// 如果 r 和 s 指向不同的环，将它们连接会创建一个新的单个环，在 r 后面插入 s 中的所有元素。结果指向在插入后 s 的最后一个元素之后的那个元素。
# <翻译结束>


<原文开始>
// Unlink removes n % r.Len() elements from the ring r, starting
// at r.Next(). If n % r.Len() == 0, r remains unchanged.
// The result is the removed sub-ring. r must not be empty.
<原文结束>

# <翻译开始>
// Unlink 从环形链表 r 中移除 n % r.Len() 个元素，从 r.Next() 开始移除。
// 若 n % r.Len() 等于 0，则 r 保持不变。
// 返回值为被移除的子环。r 必须非空。
# <翻译结束>


<原文开始>
// RLockIteratorNext iterates and locks reading forward
// with given callback function `f` within RWMutex.RLock.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// RLockIteratorNext 在 RWMutex.RLock 的保护下，以给定的回调函数 `f` 进行正向迭代并进行读取锁定。
// 若 `f` 返回 true，则继续进行迭代；若返回 false，则停止迭代。
// 这段代码的中文注释如下：
// ```go
// RLockIteratorNext 函数在读写互斥锁（RWMutex）的读锁状态下，通过给定的回调函数 `f` 实现向前遍历并加读锁。
// 当 `f` 返回值为 true 时，将继续进行遍历操作；若返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// RLockIteratorPrev iterates and locks writing backward
// with given callback function `f` within RWMutex.RLock.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// RLockIteratorPrev 在RWMutex.RLock的保护下，以给定回调函数`f`向后遍历并加写锁。
// 如果`f`返回true，则继续迭代；若返回false，则停止遍历。
# <翻译结束>


<原文开始>
// SliceNext returns a copy of all item values as slice forward from current position.
<原文结束>

# <翻译开始>
// SliceNext 从当前位置开始向前复制所有项目值，并以切片形式返回。
# <翻译结束>


<原文开始>
// SlicePrev returns a copy of all item values as slice backward from current position.
<原文结束>

# <翻译开始>
// SlicePrev 从当前位置开始向后返回所有项值的切片副本。
# <翻译结束>


<原文开始>
// Length(already used size).
<原文结束>

# <翻译开始>
// 长度（已使用大小）
# <翻译结束>


<原文开始>
// Capability(>=len).
<原文结束>

# <翻译开始>
// Capability(>=len) 表示能力（或功能）至少为 len。
# <翻译结束>

