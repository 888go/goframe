
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gring provides a concurrent-safe/unsafe ring(circular lists).
<原文结束>

# <翻译开始>
// 包gring提供了一个并发安全/不安全的环形列表（圆形队列）。 md5:bc78eee87d7b5c4b
# <翻译结束>


<原文开始>
// Ring is a struct of ring structure.
<原文结束>

# <翻译开始>
// Ring是一个环形结构的结构体。 md5:f371ac74ef187b03
# <翻译结束>


<原文开始>
// Length(already used size).
<原文结束>

# <翻译开始>
// 已使用长度。 md5:0093c138cefea3f0
# <翻译结束>


<原文开始>
// internalRingItem stores the ring element value.
<原文结束>

# <翻译开始>
// internalRingItem 用于存储环形元素的值。 md5:d1394bf8a3b910df
# <翻译结束>


<原文开始>
// New creates and returns a Ring structure of `cap` elements.
// The optional parameter `safe` specifies whether using this structure in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个具有`cap`个元素的Ring结构。
// 可选参数`safe`指定是否在并发安全环境下使用该结构，默认为false。 md5:70892e7ec9ed75d6
# <翻译结束>


<原文开始>
// Val returns the item's value of current position.
<原文结束>

# <翻译开始>
// Val 返回当前位置的项的值。 md5:b1027c8df14f08d2
# <翻译结束>


<原文开始>
// Len returns the size of ring.
<原文结束>

# <翻译开始>
// Len 返回环的大小。 md5:c4ff976cf0b72c58
# <翻译结束>


<原文开始>
// Cap returns the capacity of ring.
<原文结束>

# <翻译开始>
// Cap 返回环形缓冲区的容量。 md5:2ac015d8e20dce37
# <翻译结束>


<原文开始>
// Checks and updates the len and cap of ring when ring is dirty.
<原文结束>

# <翻译开始>
// 检查并更新ring的长度和容量，当ring被修改时。 md5:264d6fdc8ef33d31
# <翻译结束>


<原文开始>
// Set sets value to the item of current position.
<原文结束>

# <翻译开始>
// Set 将值设置为当前位置的项目。 md5:7140c77dfa3aa5dc
# <翻译结束>


<原文开始>
// Put sets `value` to current item of ring and moves position to next item.
<原文结束>

# <翻译开始>
// Put 将`value`设置为环形列表的当前项，并将位置移动到下一项。 md5:737e9a607801eee9
# <翻译结束>


<原文开始>
// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Move 函数根据给定的 n 值，将环(ring)中的元素向后（n < 0）或向前（n >= 0）移动 n % r.Len() 个位置，并返回移动后所在位置的元素。环(r)不能为空。 md5:92f786c9a5a8b8cd
# <翻译结束>


<原文开始>
// Prev returns the previous ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Prev 返回前一个环形元素。r 不能为空。 md5:574e755d8883dc1f
# <翻译结束>


<原文开始>
// Next returns the next ring element. r must not be empty.
<原文结束>

# <翻译开始>
// Next 返回下一个环元素。r必须不为空。 md5:f16b811ce23ee06b
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
// Link 将环 r 与环 s 连接起来，使得 r.Next() 变为 s，并返回 r.Next() 的原始值。
// r 不应为空。
//
// 如果 r 和 s 指向相同的环，链接它们会移除 r 和 s 之间的元素。移除的元素形成一个子环，结果是对这个子环的引用（如果没有移除任何元素，结果仍然是 r.Next() 的原始值，而不是 nil）。
//
// 如果 r 和 s 指向不同的环，链接它们会在 r 后插入 s 的元素，创建一个单一的环。结果指向插入后 s 的最后一个元素之后的元素。 md5:faa73e3f5f43468a
# <翻译结束>


<原文开始>
// Unlink removes n % r.Len() elements from the ring r, starting
// at r.Next(). If n % r.Len() == 0, r remains unchanged.
// The result is the removed sub-ring. r must not be empty.
<原文结束>

# <翻译开始>
// Unlink 从环 r 中移除 n % r.Len() 个元素，开始于 r.Next()。
// 如果 n % r.Len() == 0，那么 r 保持不变。
// 结果是被移除的子环。r 必须非空。 md5:00909914d8e87d32
# <翻译结束>


<原文开始>
// RLockIteratorNext iterates and locks reading forward
// with given callback function `f` within RWMutex.RLock.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// RLockIteratorNext 在RWMutex的RLock范围内向前迭代并加锁读取。
// 如果提供的回调函数`f`返回true，那么将继续迭代；如果返回false，则停止迭代。 md5:8cb144956023168f
# <翻译结束>


<原文开始>
// RLockIteratorPrev iterates and locks writing backward
// with given callback function `f` within RWMutex.RLock.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// RLockIteratorPrev 逆序迭代并锁定（写入）RWMutex.RLock。
// 使用给定的回调函数 `f`。如果 `f` 返回 true，那么继续迭代；否则停止。 md5:31f0f0af041e234a
# <翻译结束>


<原文开始>
// SliceNext returns a copy of all item values as slice forward from current position.
<原文结束>

# <翻译开始>
// SliceNext 返回一个从当前位置开始向前的所有项值的切片副本。 md5:54ba7b6ac01a38f8
# <翻译结束>


<原文开始>
// SlicePrev returns a copy of all item values as slice backward from current position.
<原文结束>

# <翻译开始>
// SlicePrev 从当前位置向前返回所有项目值的副本作为切片。 md5:632f85c2939f2e91
# <翻译结束>

