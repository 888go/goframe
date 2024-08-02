// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"context"

	gconv "github.com/888go/goframe/util/gconv"
)

// Cache struct.
type Cache struct {
	localAdapter
}

// localAdapter 是 Adapter 的别名，仅用于嵌入属性的目的。 md5:be3fc375883fa166
type localAdapter = Adapter

// New 使用默认的内存适配器创建并返回一个新的缓存对象。
// 请注意，LRU（最近最少使用）功能仅在使用内存适配器时可用。
// md5:658995a71d08fbbe
func New(lruCap ...int) *Cache {
	memAdapter := NewAdapterMemory(lruCap...)
	c := &Cache{
		localAdapter: memAdapter,
	}
	return c
}

// NewWithAdapter 使用给定的实现了Adapter接口的适配器创建并返回一个Cache对象。 md5:0c92c6f9af030ccb
func NewWithAdapter(adapter Adapter) *Cache {
	return &Cache{
		localAdapter: adapter,
	}
}

// SetAdapter 更改此缓存的适配器。
// 非常注意，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用此设置函数。
// md5:5f950a554baddc2c
func (c *Cache) SetAdapter(adapter Adapter) {
	c.localAdapter = adapter
}

// GetAdapter 返回当前缓存中设置的适配器。 md5:e93da9e47a8b0c21
func (c *Cache) GetAdapter() Adapter {
	return c.localAdapter
}

// 从缓存中删除`keys`。 md5:370028bf9f2e1d24
func (c *Cache) Removes(ctx context.Context, keys []interface{}) error {
	_, err := c.Remove(ctx, keys...)
	return err
}

// KeyStrings返回缓存中的所有键作为字符串切片。 md5:3b0126221389825e
func (c *Cache) KeyStrings(ctx context.Context) ([]string, error) {
	keys, err := c.Keys(ctx)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(keys), nil
}
