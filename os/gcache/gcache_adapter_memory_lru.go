// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 缓存类

import (
	"context"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/os/gtimer"
)

// LRU 缓存对象.
// 它使用来自标准库的 list.List 作为其底层的双向链表实现。
type adapterMemoryLru struct {
	cache   *AdapterMemory // 父级缓存对象。
	data    *map类.Map      // Key 映射到列表中的项目。
	list    *链表类.List    // Key list.
	rawList *链表类.List    // key的添加历史记录
	closed  *安全变量类.Bool    // Closed or not.
}

// newMemCacheLru 创建并返回一个新的 LRU 对象。
func newMemCacheLru(cache *AdapterMemory) *adapterMemoryLru {
	lru := &adapterMemoryLru{
		cache:   cache,
		data:    map类.X创建(true),
		list:    链表类.New(true),
		rawList: 链表类.New(true),
		closed:  安全变量类.NewBool(),
	}
	return lru
}

// Close 关闭 LRU 对象。
func (lru *adapterMemoryLru) Close() {
	lru.closed.X设置值(true)
}

// Remove 从 `lru` 中删除键 `key`。
func (lru *adapterMemoryLru) X删除并带返回值(key interface{}) {
	if v := lru.data.X取值(key); v != nil {
		lru.data.X删除(key)
		lru.list.Remove(v.(*链表类.Element))
	}
}

// Size 返回 `lru` 的大小。
func (lru *adapterMemoryLru) X取数量() int {
	return lru.data.X取数量()
}

// Push将`key`推送到`lru`的尾部。
func (lru *adapterMemoryLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop从lru的尾部删除并返回键。
func (lru *adapterMemoryLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.X删除(v)
		return v
	}
	return nil
}

// SyncAndClear 使用“最近最少使用”(Least Recently Used, LRU)算法，将`rawList`中的键同步到`list`和`data`中，并进行清除操作。
func (lru *adapterMemoryLru) SyncAndClear(ctx context.Context) {
	if lru.closed.X取值() {
		定时类.X退出()
		return
	}
	// 数据同步。
	var alreadyExistItem interface{}
	for {
		if rawListItem := lru.rawList.PopFront(); rawListItem != nil {
			// 从列表中删除键。
			if alreadyExistItem = lru.data.X取值(rawListItem); alreadyExistItem != nil {
				lru.list.Remove(alreadyExistItem.(*链表类.Element))
			}
// 将键推送到列表的头部
// 并将其列表项设置到哈希表中以便快速索引。
			lru.data.X设置值(rawListItem, lru.list.PushFront(rawListItem))
		} else {
			break
		}
	}
	// Data cleaning up.
	for clearLength := lru.X取数量() - lru.cache.cap; clearLength > 0; clearLength-- {
		if topKey := lru.Pop(); topKey != nil {
			lru.cache.clearByKey(topKey, true)
		}
	}
}
