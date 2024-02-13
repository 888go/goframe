// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 缓存类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/util/gconv"
)

// Cache struct.
type Cache struct {
	localAdapter
}

// localAdapter 是 Adapter 的别名，仅用于嵌入式属性的目的。
type localAdapter = Adapter

// New 函数创建并返回一个使用默认内存适配器的新缓存对象。
// 注意，LRU（最近最少使用）特性仅在使用内存适配器时可用。
func X创建(淘汰数量 ...int) *Cache {
	memAdapter := X创建内存适配器(淘汰数量...)
	c := &Cache{
		localAdapter: memAdapter,
	}
// 如果适配器手动从内存适配器更改，这里可能存在“计时器泄漏”的问题。
// 不必担心这一点，因为适配器很少更改，并且如果未使用则不会造成任何影响。
	定时类.X加入单例循环任务(context.Background(), time.Second, memAdapter.(*AdapterMemory).syncEventAndClearExpired)
	return c
}

// NewWithAdapter 使用给定的已实现Adapter接口的对象创建并返回一个Cache对象。
func X创建并按适配器(适配器 Adapter) *Cache {
	return &Cache{
		localAdapter: 适配器,
	}
}

// SetAdapter 更改此缓存的适配器。
// 非常需要注意的是，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用这个设置函数。
func (c *Cache) X设置适配器(适配器 Adapter) {
	c.localAdapter = 适配器
}

// GetAdapter 返回当前 Cache 中设置的适配器。
func (c *Cache) X取适配器() Adapter {
	return c.localAdapter
}

// 删除缓存中的`keys`。
func (c *Cache) X删除(上下文 context.Context, 名称s []interface{}) error {
	_, err := c.X删除并带返回值(上下文, 名称s...)
	return err
}

// KeyStrings 返回缓存中的所有键，以字符串切片的形式。
func (c *Cache) X取所有键文本(上下文 context.Context) ([]string, error) {
	keys, err := c.X取所有键(上下文)
	if err != nil {
		return nil, err
	}
	return 转换类.X取文本数组(keys), nil
}
