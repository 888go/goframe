// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcache

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

// Cache struct.
type Cache struct {
	localAdapter
}

// localAdapter is alias of Adapter, for embedded attribute purpose only.
type localAdapter = Adapter

// New creates and returns a new cache object using default memory adapter.
// Note that the LRU feature is only available using memory adapter.
// ff:创建
// lruCap:淘汰数量
func New(lruCap ...int) *Cache {
	memAdapter := NewAdapterMemory(lruCap...)
	c := &Cache{
		localAdapter: memAdapter,
	}
	return c
}

// NewWithAdapter creates and returns a Cache object with given Adapter implements.
// ff:创建并按适配器
// adapter:适配器
func NewWithAdapter(adapter Adapter) *Cache {
	return &Cache{
		localAdapter: adapter,
	}
}

// SetAdapter changes the adapter for this cache.
// Be very note that, this setting function is not concurrent-safe, which means you should not call
// this setting function concurrently in multiple goroutines.
// ff:设置适配器
// c:
// adapter:适配器
func (c *Cache) SetAdapter(adapter Adapter) {
	c.localAdapter = adapter
}

// GetAdapter returns the adapter that is set in current Cache.
// ff:取适配器
// c:
func (c *Cache) GetAdapter() Adapter {
	return c.localAdapter
}

// Removes deletes `keys` in the cache.
// ff:删除
// c:
// ctx:上下文
// keys:名称s
func (c *Cache) Removes(ctx context.Context, keys []interface{}) error {
	_, err := c.Remove(ctx, keys...)
	return err
}

// KeyStrings returns all keys in the cache as string slice.
// ff:取所有键文本
// c:
// ctx:上下文
func (c *Cache) KeyStrings(ctx context.Context) ([]string, error) {
	keys, err := c.Keys(ctx)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(keys), nil
}
