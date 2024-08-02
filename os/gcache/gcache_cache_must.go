// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	"context"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
)

// MustGet 的行为就像 Get 一样，但如果发生任何错误，它会引发 panic。 md5:9004545d221e9637
func (c *Cache) MustGet(ctx context.Context, key interface{}) *gvar.Var {
	v, err := c.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSet 的行为类似于 GetOrSet，但是如果发生任何错误，它会直接 panic。 md5:684c6b06451a2f6f
func (c *Cache) MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSet(ctx, key, value, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但如果发生任何错误，则会引发 panic。 md5:07fd1ef2dbfce0b4
func (c *Cache) MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSetFunc(ctx, key, f, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFuncLock 行为与 GetOrSetFuncLock 类似，但如果发生任何错误，它将引发恐慌。 md5:7f84f54a71da5305
func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSetFuncLock(ctx, key, f, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustContains 的行为就像 Contains，但如果发生任何错误，它将引发恐慌。 md5:63cc1bbb0025d8b1
func (c *Cache) MustContains(ctx context.Context, key interface{}) bool {
	v, err := c.Contains(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会直接 panic。 md5:c97fa5941bbc47a3
func (c *Cache) MustGetExpire(ctx context.Context, key interface{}) time.Duration {
	v, err := c.GetExpire(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustSize 行为类似于 Size，但在发生错误时会引发 panic。 md5:cee955b74cc42d5c
func (c *Cache) MustSize(ctx context.Context) int {
	v, err := c.Size(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData的行为就像Data一样，但如果发生任何错误，它会引发恐慌。 md5:b53b751e2003cd20
func (c *Cache) MustData(ctx context.Context) map[interface{}]interface{} {
	v, err := c.Data(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeys 行为与 Keys 类似，但如果发生任何错误，它将引发 panic。 md5:7f7801d0cd170166
func (c *Cache) MustKeys(ctx context.Context) []interface{} {
	v, err := c.Keys(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它会直接 panic。 md5:3efe93008da2eb0f
func (c *Cache) MustKeyStrings(ctx context.Context) []string {
	v, err := c.KeyStrings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustValues 行为类似于 Values，但如果发生任何错误则会引发 panic。 md5:859aff610512a748
func (c *Cache) MustValues(ctx context.Context) []interface{} {
	v, err := c.Values(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
