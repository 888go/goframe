// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 缓存类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gvar"
)

// MustGet 行为类似于 Get，但当发生任何错误时，它会触发panic。
func (c *Cache) X取值PANI(上下文 context.Context, 名称 interface{}) *泛型类.Var {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSet 行为类似于 GetOrSet，但是当发生任何错误时，它会触发panic（异常）。
func (c *Cache) X取值或设置值PANI(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) *泛型类.Var {
	v, err := c.X取值或设置值(上下文, 名称, 值, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFunc 行为类似于 GetOrSetFunc，但当发生任何错误时它会触发panic。
func (c *Cache) X取值或设置值_函数PANI(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) *泛型类.Var {
	v, err := c.X取值或设置值_函数(上下文, 名称, 回调函数, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetOrSetFuncLock 类似于 GetOrSetFuncLock，但如果发生任何错误，它会触发 panic。
func (c *Cache) X取值或设置值_并发安全函数PANI(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) *泛型类.Var {
	v, err := c.X取值或设置值_并发安全函数(上下文, 名称, 回调函数, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// MustContains 行为类似于 Contains，但当发生任何错误时，它会触发panic（异常退出）。
func (c *Cache) X是否存在PANI(上下文 context.Context, 名称 interface{}) bool {
	v, err := c.X是否存在(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetExpire 的行为类似于 GetExpire，但是当发生任何错误时，它会触发panic。
func (c *Cache) X取过期时间PANI(上下文 context.Context, 名称 interface{}) time.Duration {
	v, err := c.X取过期时间(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// MustSize 的行为类似于 Size，但如果发生任何错误，它会触发 panic。
func (c *Cache) X取数量PANI(上下文 context.Context) int {
	v, err := c.X取数量(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData 的行为类似于 Data，但是当发生任何错误时它会触发panic（异常）。
func (c *Cache) X取所有键值Map副本PANI(上下文 context.Context) map[interface{}]interface{} {
	v, err := c.X取所有键值Map副本(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeys 行为类似 Keys，但当发生任何错误时会触发 panic。
func (c *Cache) X取所有键PANI(上下文 context.Context) []interface{} {
	v, err := c.X取所有键(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// MustKeyStrings 行为类似 KeyStrings，但当发生任何错误时，它会引发 panic。
func (c *Cache) X取所有键文本PANI(上下文 context.Context) []string {
	v, err := c.X取所有键文本(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// MustValues 行为类似 Values，但是当发生任何错误时它会触发panic（异常）。
func (c *Cache) X取所有值PANI(上下文 context.Context) []interface{} {
	v, err := c.X取所有值(上下文)
	if err != nil {
		panic(err)
	}
	return v
}
