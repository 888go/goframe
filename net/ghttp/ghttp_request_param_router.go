// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	gvar "github.com/888go/goframe/container/gvar"
)

// X取路由器Map副本 获取并返回路由器映射的副本。 md5:c47609cda93e62b3
func (r *Request) X取路由器Map副本() map[string]string {
	if r.routerMap != nil {
		m := make(map[string]string, len(r.routerMap))
		for k, v := range r.routerMap {
			m[k] = v
		}
		return m
	}
	return nil
}

// X取路由器值到泛型类 通过给定的键名 `key` 获取并返回路由器值。如果 `key` 不存在，它将返回 `def`。
// md5:25ec580295596f49
func (r *Request) X取路由器值到泛型类(名称 string, 默认值 ...interface{}) *gvar.Var {
	if r.routerMap != nil {
		if v, ok := r.routerMap[名称]; ok {
			return gvar.X创建(v)
		}
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}
