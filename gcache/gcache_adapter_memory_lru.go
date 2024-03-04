// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"context"
	
	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/os/gtimer"
)

// LRU 缓存对象.
// 它使用来自标准库的 list.List 作为其底层的双向链表实现。
type adapterMemoryLru struct {
	cache   *AdapterMemory // 父级缓存对象。
	data    *gmap.Map      // Key 映射到列表中的项目。
	list    *glist.List    // Key list.
	rawList *glist.List    // key的添加历史记录
	closed  *gtype.Bool    // Closed or not.
}

// newMemCacheLru 创建并返回一个新的 LRU 对象。
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

// Close 关闭 LRU 对象。
func (lru *adapterMemoryLru) Close() {
	lru.closed.Set(true)
}

// Remove 从 `lru` 中删除键 `key`。
func (lru *adapterMemoryLru) Remove(key interface{}) {
	if v := lru.data.Get(key); v != nil {
		lru.data.Remove(key)
		lru.list.Remove(v.(*glist.Element))
	}
}

// Size 返回 `lru` 的大小。
func (lru *adapterMemoryLru) Size() int {
	return lru.data.Size()
}

// Push将`key`推送到`lru`的尾部。
func (lru *adapterMemoryLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop从lru的尾部删除并返回键。
func (lru *adapterMemoryLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.Remove(v)
		return v
	}
	return nil
}

// SyncAndClear 使用“最近最少使用”(Least Recently Used, LRU)算法，将`rawList`中的键同步到`list`和`data`中，并进行清除操作。
func (lru *adapterMemoryLru) SyncAndClear(ctx context.Context) {
	if lru.closed.Val() {
		gtimer.Exit()
		return
	}
	// 数据同步。
	var alreadyExistItem interface{}
	for {
		if rawListItem := lru.rawList.PopFront(); rawListItem != nil {
			// 从列表中删除键。
			if alreadyExistItem = lru.data.Get(rawListItem); alreadyExistItem != nil {
				lru.list.Remove(alreadyExistItem.(*glist.Element))
			}
// 将键推送到列表的头部
// 并将其列表项设置到哈希表中以便快速索引。
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
