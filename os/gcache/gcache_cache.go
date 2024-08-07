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

// X创建 使用默认的内存适配器创建并返回一个新的缓存对象。
// 请注意，LRU（最近最少使用）功能仅在使用内存适配器时可用。
// md5:658995a71d08fbbe
func X创建(淘汰数量 ...int) *Cache {
	memAdapter := X创建内存适配器(淘汰数量...)
	c := &Cache{
		localAdapter: memAdapter,
	}
	return c
}

// X创建并按适配器 使用给定的实现了Adapter接口的适配器创建并返回一个Cache对象。 md5:0c92c6f9af030ccb
func X创建并按适配器(适配器 Adapter) *Cache {
	return &Cache{
		localAdapter: 适配器,
	}
}

// X设置适配器 更改此缓存的适配器。
// 非常注意，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用此设置函数。
// md5:5f950a554baddc2c
func (c *Cache) X设置适配器(适配器 Adapter) {
	c.localAdapter = 适配器
}

// X取适配器 返回当前缓存中设置的适配器。 md5:e93da9e47a8b0c21
func (c *Cache) X取适配器() Adapter {
	return c.localAdapter
}

// 从缓存中删除`keys`。 md5:370028bf9f2e1d24
func (c *Cache) X删除(上下文 context.Context, 名称s []interface{}) error {
	_, err := c.X删除并带返回值(上下文, 名称s...)
	return err
}

// X取所有键文本返回缓存中的所有键作为字符串切片。 md5:3b0126221389825e
func (c *Cache) X取所有键文本(上下文 context.Context) ([]string, error) {
	keys, err := c.X取所有键(上下文)
	if err != nil {
		return nil, err
	}
	return gconv.X取文本切片(keys), nil
}
