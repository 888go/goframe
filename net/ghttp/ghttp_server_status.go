// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"fmt"
)

// getStatusHandler根据给定的状态码获取并返回处理程序。
func (s *X服务) getStatusHandler(status int, r *X请求) []HandlerFunc {
	domains := []string{r.X取主机名(), DefaultDomainName}
	for _, domain := range domains {
		if f, ok := s.statusHandlerMap[s.statusHandlerKey(status, domain)]; ok {
			return f
		}
	}
	return nil
}

// addStatusHandler 为给定的状态码设置处理器。
// 参数 `pattern` 形如：domain#status
func (s *X服务) addStatusHandler(pattern string, handler HandlerFunc) {
	if s.statusHandlerMap[pattern] == nil {
		s.statusHandlerMap[pattern] = make([]HandlerFunc, 0)
	}
	s.statusHandlerMap[pattern] = append(s.statusHandlerMap[pattern], handler)
}

// statusHandlerKey 为给定的状态和域名创建并返回一个键。
func (s *X服务) statusHandlerKey(status int, domain string) string {
	return fmt.Sprintf("%s#%d", domain, status)
}

// BindStatusHandler 为给定的状态码注册处理器。
func (s *X服务) X绑定状态码中间件(状态码 int, 处理函数 HandlerFunc) {
	s.addStatusHandler(s.statusHandlerKey(状态码, DefaultDomainName), 处理函数)
}

// BindStatusHandlerByMap 通过映射注册给定状态码的处理器。
func (s *X服务) X绑定状态码中间件Map(中间件Map map[int]HandlerFunc) {
	for k, v := range 中间件Map {
		s.X绑定状态码中间件(k, v)
	}
}
