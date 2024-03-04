
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// LRU cache object.
// It uses list.List from stdlib for its underlying doubly linked list.
<原文结束>

# <翻译开始>
// LRU 缓存对象.
// 它使用来自标准库的 list.List 作为其底层的双向链表实现。
# <翻译结束>







<原文开始>
// Key mapping to the item of the list.
<原文结束>

# <翻译开始>
// Key 映射到列表中的项目。
# <翻译结束>







<原文开始>
// newMemCacheLru creates and returns a new LRU object.
<原文结束>

# <翻译开始>
// newMemCacheLru 创建并返回一个新的 LRU 对象。
# <翻译结束>


<原文开始>
// Close closes the LRU object.
<原文结束>

# <翻译开始>
// Close 关闭 LRU 对象。
# <翻译结束>


<原文开始>
// Remove deletes the `key` FROM `lru`.
<原文结束>

# <翻译开始>
// Remove 从 `lru` 中删除键 `key`。
# <翻译结束>


<原文开始>
// Size returns the size of `lru`.
<原文结束>

# <翻译开始>
// Size 返回 `lru` 的大小。
# <翻译结束>


<原文开始>
// Push pushes `key` to the tail of `lru`.
<原文结束>

# <翻译开始>
// Push将`key`推送到`lru`的尾部。
# <翻译结束>


<原文开始>
// Pop deletes and returns the key from tail of `lru`.
<原文结束>

# <翻译开始>
// Pop从lru的尾部删除并返回键。
# <翻译结束>


<原文开始>
// SyncAndClear synchronizes the keys from `rawList` to `list` and `data`
// using Least Recently Used algorithm.
<原文结束>

# <翻译开始>
// SyncAndClear 使用“最近最少使用”(Least Recently Used, LRU)算法，将`rawList`中的键同步到`list`和`data`中，并进行清除操作。
# <翻译结束>







<原文开始>
// Deleting the key from list.
<原文结束>

# <翻译开始>
// 从列表中删除键。
# <翻译结束>


<原文开始>
			// Pushing key to the head of the list
			// and setting its list item to hash table for quick indexing.
<原文结束>

# <翻译开始>
// 将键推送到列表的头部
// 并将其列表项设置到哈希表中以便快速索引。
# <翻译结束>







<原文开始>
// Parent cache object.
<原文结束>

# <翻译开始>
// 父级缓存对象。
# <翻译结束>


<原文开始>
// History for key adding.
<原文结束>

# <翻译开始>
// key的添加历史记录
# <翻译结束>


<原文开始>
// Data synchronization.
<原文结束>

# <翻译开始>
// 数据同步。
# <翻译结束>

