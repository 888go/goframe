
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// LRU cache object.
// It uses list.List from stdlib for its underlying doubly linked list.
<原文结束>

# <翻译开始>
// LRU 缓存对象。
// 它使用 stdlib 中的 list.List 作为其底层的双链表。
// md5:0865da04bb1ff4bb
# <翻译结束>


<原文开始>
// Key mapping to the item of the list.
<原文结束>

# <翻译开始>
// 键映射到列表中的项目。. md5:1783218fcc5a7851
# <翻译结束>


<原文开始>
// History for key adding.
<原文结束>

# <翻译开始>
// 关于添加键的历史记录。. md5:73aaa8a4c7c9ca97
# <翻译结束>


<原文开始>
// newMemCacheLru creates and returns a new LRU object.
<原文结束>

# <翻译开始>
// newMemCacheLru 创建并返回一个新的LRU对象。. md5:e52ac4e697ac0070
# <翻译结束>


<原文开始>
// Close closes the LRU object.
<原文结束>

# <翻译开始>
// Close 关闭 LRU 对象。. md5:5fbab2bd7f830bd3
# <翻译结束>


<原文开始>
// Remove deletes the `key` FROM `lru`.
<原文结束>

# <翻译开始>
// Remove 从 LRU 缓存中删除 `key`。. md5:1b31a149f111557e
# <翻译结束>


<原文开始>
// Size returns the size of `lru`.
<原文结束>

# <翻译开始>
// Size 返回 lru 的大小。. md5:e6b8b41e660eeabd
# <翻译结束>


<原文开始>
// Push pushes `key` to the tail of `lru`.
<原文结束>

# <翻译开始>
// Push 将`key`推送到`lru`的尾部。. md5:d0793b82031a3f0e
# <翻译结束>


<原文开始>
// Pop deletes and returns the key from tail of `lru`.
<原文结束>

# <翻译开始>
// Pop 从`lru`的尾部删除并返回键。. md5:e9a281592f5ec82e
# <翻译结束>


<原文开始>
// SyncAndClear synchronizes the keys from `rawList` to `list` and `data`
// using Least Recently Used algorithm.
<原文结束>

# <翻译开始>
// SyncAndClear 使用最近最少使用（LRU）算法，将键从`rawList`同步到`list`和`data`中，并清除不再需要的数据。
// md5:1da6cde3bc8d63d6
# <翻译结束>


<原文开始>
// Deleting the key from list.
<原文结束>

# <翻译开始>
// 从列表中删除键。. md5:9044ea33db98a37a
# <翻译结束>


<原文开始>
			// Pushing key to the head of the list
			// and setting its list item to hash table for quick indexing.
<原文结束>

# <翻译开始>
// 将键推送到列表的头部
// 并将其项目设置到哈希表中，以便快速索引。
// md5:c4ec4de48ddb7b0c
# <翻译结束>

