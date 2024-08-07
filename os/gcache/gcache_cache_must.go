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

// X取值PANI 的行为就像 Get 一样，但如果发生任何错误，它会引发 panic。 md5:9004545d221e9637
func (c *Cache) X取值PANI(上下文 context.Context, 名称 interface{}) *gvar.Var {
	v, err := c.X取值(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// X取值或设置值PANI 的行为类似于 GetOrSet，但是如果发生任何错误，它会直接 panic。 md5:684c6b06451a2f6f
func (c *Cache) X取值或设置值PANI(上下文 context.Context, 名称 interface{}, 值 interface{}, 时长 time.Duration) *gvar.Var {
	v, err := c.X取值或设置值(上下文, 名称, 值, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// X取值或设置值_函数PANI 行为类似于 GetOrSetFunc，但如果发生任何错误，则会引发 panic。 md5:07fd1ef2dbfce0b4
func (c *Cache) X取值或设置值_函数PANI(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) *gvar.Var {
	v, err := c.X取值或设置值_函数(上下文, 名称, 回调函数, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// X取值或设置值_并发安全函数PANI 行为与 GetOrSetFuncLock 类似，但如果发生任何错误，它将引发恐慌。 md5:7f84f54a71da5305
func (c *Cache) X取值或设置值_并发安全函数PANI(上下文 context.Context, 名称 interface{}, 回调函数 Func, 时长 time.Duration) *gvar.Var {
	v, err := c.X取值或设置值_并发安全函数(上下文, 名称, 回调函数, 时长)
	if err != nil {
		panic(err)
	}
	return v
}

// X是否存在PANI 的行为就像 Contains，但如果发生任何错误，它将引发恐慌。 md5:63cc1bbb0025d8b1
func (c *Cache) X是否存在PANI(上下文 context.Context, 名称 interface{}) bool {
	v, err := c.X是否存在(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// X取过期时间PANI 的行为类似于 GetExpire，但如果发生任何错误，它会直接 panic。 md5:c97fa5941bbc47a3
func (c *Cache) X取过期时间PANI(上下文 context.Context, 名称 interface{}) time.Duration {
	v, err := c.X取过期时间(上下文, 名称)
	if err != nil {
		panic(err)
	}
	return v
}

// X取数量PANI 行为类似于 Size，但在发生错误时会引发 panic。 md5:cee955b74cc42d5c
func (c *Cache) X取数量PANI(上下文 context.Context) int {
	v, err := c.X取数量(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// X取所有键值Map副本PANI的行为就像Data一样，但如果发生任何错误，它会引发恐慌。 md5:b53b751e2003cd20
func (c *Cache) X取所有键值Map副本PANI(上下文 context.Context) map[interface{}]interface{} {
	v, err := c.X取所有键值Map副本(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// X取所有键PANI 行为与 Keys 类似，但如果发生任何错误，它将引发 panic。 md5:7f7801d0cd170166
func (c *Cache) X取所有键PANI(上下文 context.Context) []interface{} {
	v, err := c.X取所有键(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// X取所有键文本PANI 的行为类似于 KeyStrings，但如果发生任何错误，它会直接 panic。 md5:3efe93008da2eb0f
func (c *Cache) X取所有键文本PANI(上下文 context.Context) []string {
	v, err := c.X取所有键文本(上下文)
	if err != nil {
		panic(err)
	}
	return v
}

// X取所有值PANI 行为类似于 Values，但如果发生任何错误则会引发 panic。 md5:859aff610512a748
func (c *Cache) X取所有值PANI(上下文 context.Context) []interface{} {
	v, err := c.X取所有值(上下文)
	if err != nil {
		panic(err)
	}
	return v
}
