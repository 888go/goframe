// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcache

import (
	"context"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/os/gtimer"
)

// LRU 缓存对象。
// 它使用 stdlib 中的 list.List 作为其底层的双链表。
// md5:0865da04bb1ff4bb
type adapterMemoryLru struct {
	cache   *AdapterMemory // Parent cache object.
	data    *gmap.Map      // 键映射到列表中的项目。 md5:1783218fcc5a7851
	list    *glist.List    // Key list.
	rawList *glist.List    // 关于添加键的历史记录。 md5:73aaa8a4c7c9ca97
	closed  *gtype.Bool    // Closed or not.
}

// newMemCacheLru 创建并返回一个新的LRU对象。 md5:e52ac4e697ac0070
func newMemCacheLru(cache *AdapterMemory) *adapterMemoryLru {
	lru := &adapterMemoryLru{
		cache:   cache,
		data:    gmap.New(true),
		list:    glist.New(true),
		rawList: glist.New(true),
		closed:  gtype.NewBool(),
	}
	return lru
}

// Close 关闭 LRU 对象。 md5:5fbab2bd7f830bd3
func (lru *adapterMemoryLru) Close() {
	lru.closed.Set(true)
}

// Remove 从 LRU 缓存中删除 `key`。 md5:1b31a149f111557e
func (lru *adapterMemoryLru) Remove(key interface{}) {
	if v := lru.data.Get(key); v != nil {
		lru.data.Remove(key)
		lru.list.Remove(v.(*glist.Element))
	}
}

// Size 返回 lru 的大小。 md5:e6b8b41e660eeabd
func (lru *adapterMemoryLru) Size() int {
	return lru.data.Size()
}

// Push 将`key`推送到`lru`的尾部。 md5:d0793b82031a3f0e
func (lru *adapterMemoryLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop 从`lru`的尾部删除并返回键。 md5:e9a281592f5ec82e
func (lru *adapterMemoryLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.Remove(v)
		return v
	}
	return nil
}

// SyncAndClear 使用最近最少使用（LRU）算法，将键从`rawList`同步到`list`和`data`中，并清除不再需要的数据。
// md5:1da6cde3bc8d63d6
func (lru *adapterMemoryLru) SyncAndClear(ctx context.Context) {
	if lru.closed.Val() {
		gtimer.Exit()
		return
	}
	// Data synchronization.
	var alreadyExistItem interface{}
	for {
		if rawListItem := lru.rawList.PopFront(); rawListItem != nil {
			// 从列表中删除键。 md5:9044ea33db98a37a
			if alreadyExistItem = lru.data.Get(rawListItem); alreadyExistItem != nil {
				lru.list.Remove(alreadyExistItem.(*glist.Element))
			}
			// 将键推送到列表的头部
			// 并将其项目设置到哈希表中，以便快速索引。
			// md5:c4ec4de48ddb7b0c
			lru.data.Set(rawListItem, lru.list.PushFront(rawListItem))
		} else {
			break
		}
	}
	// Data cleaning up.
	for clearLength := lru.Size() - lru.cache.cap; clearLength > 0; clearLength-- {
		if topKey := lru.Pop(); topKey != nil {
			lru.cache.clearByKey(topKey, true)
		}
	}
}
