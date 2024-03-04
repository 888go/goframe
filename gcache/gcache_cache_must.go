// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcache

import (
	"context"
	"time"
	
	"github.com/gogf/gf/v2/container/gvar"
)

// MustGet 行为类似于 Get，但当发生任何错误时，它会触发panic。
func (c *Cache) MustGet(ctx context.Context, key interface{}) *gvar.Var {
	v, err := c.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSet 行为类似于 GetOrSet，但是当发生任何错误时，它会触发panic（异常）。
func (c *Cache) MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSet(ctx, key, value, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当发生任何错误时它会触发panic。
func (c *Cache) MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSetFunc(ctx, key, f, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFuncLock 类似于 GetOrSetFuncLock，但如果发生任何错误，它会触发 panic。
func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {
	v, err := c.GetOrSetFuncLock(ctx, key, f, duration)
	if err != nil {
		panic(err)
	}
	return v
}

// MustContains 行为类似于 Contains，但当发生任何错误时，它会触发panic（异常退出）。
func (c *Cache) MustContains(ctx context.Context, key interface{}) bool {
	v, err := c.Contains(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetExpire 的行为类似于 GetExpire，但是当发生任何错误时，它会触发panic。
func (c *Cache) MustGetExpire(ctx context.Context, key interface{}) time.Duration {
	v, err := c.GetExpire(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustSize 的行为类似于 Size，但如果发生任何错误，它会触发 panic。
func (c *Cache) MustSize(ctx context.Context) int {
	v, err := c.Size(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData 的行为类似于 Data，但是当发生任何错误时它会触发panic（异常）。
func (c *Cache) MustData(ctx context.Context) map[interface{}]interface{} {
	v, err := c.Data(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeys 行为类似 Keys，但当发生任何错误时会触发 panic。
func (c *Cache) MustKeys(ctx context.Context) []interface{} {
	v, err := c.Keys(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeyStrings 行为类似 KeyStrings，但当发生任何错误时，它会引发 panic。
func (c *Cache) MustKeyStrings(ctx context.Context) []string {
	v, err := c.KeyStrings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// MustValues 行为类似 Values，但是当发生任何错误时它会触发panic（异常）。
func (c *Cache) MustValues(ctx context.Context) []interface{} {
	v, err := c.Values(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
