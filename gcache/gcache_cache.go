// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"context"
	"time"
	
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
)

// Cache struct.
type Cache struct {
	localAdapter
}

// localAdapter 是 Adapter 的别名，仅用于嵌入式属性的目的。
type localAdapter = Adapter

// New 函数创建并返回一个使用默认内存适配器的新缓存对象。
// 注意，LRU（最近最少使用）特性仅在使用内存适配器时可用。
func New(lruCap ...int) *Cache {
	memAdapter := NewAdapterMemory(lruCap...)
	c := &Cache{
		localAdapter: memAdapter,
	}
// 如果适配器手动从内存适配器更改，这里可能存在“计时器泄漏”的问题。
// 不必担心这一点，因为适配器很少更改，并且如果未使用则不会造成任何影响。
	gtimer.AddSingleton(context.Background(), time.Second, memAdapter.(*AdapterMemory).syncEventAndClearExpired)
	return c
}

// NewWithAdapter 使用给定的已实现Adapter接口的对象创建并返回一个Cache对象。
func NewWithAdapter(adapter Adapter) *Cache {
	return &Cache{
		localAdapter: adapter,
	}
}

// SetAdapter 更改此缓存的适配器。
// 非常需要注意的是，这个设置函数不是并发安全的，这意味着你不应该在多个goroutine中并发调用这个设置函数。
func (c *Cache) SetAdapter(adapter Adapter) {
	c.localAdapter = adapter
}

// GetAdapter 返回当前 Cache 中设置的适配器。
func (c *Cache) GetAdapter() Adapter {
	return c.localAdapter
}

// 删除缓存中的`keys`。
func (c *Cache) Removes(ctx context.Context, keys []interface{}) error {
	_, err := c.Remove(ctx, keys...)
	return err
}

// KeyStrings 返回缓存中的所有键，以字符串切片的形式。
func (c *Cache) KeyStrings(ctx context.Context) ([]string, error) {
	keys, err := c.Keys(ctx)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(keys), nil
}
