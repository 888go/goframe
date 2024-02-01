// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"github.com/888go/goframe/container/gvar"
	)
// GetRouterMap 获取并返回路由器映射表的一个副本。
func (r *Request) GetRouterMap() map[string]string {
	if r.routerMap != nil {
		m := make(map[string]string, len(r.routerMap))
		for k, v := range r.routerMap {
			m[k] = v
		}
		return m
	}
	return nil
}

// GetRouter 根据给定的键名 `key` 获取并返回路由器值。
// 如果 `key` 不存在，则返回 `def`。
func (r *Request) GetRouter(key string, def ...interface{}) *gvar.Var {
	if r.routerMap != nil {
		if v, ok := r.routerMap[key]; ok {
			return gvar.New(v)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}
