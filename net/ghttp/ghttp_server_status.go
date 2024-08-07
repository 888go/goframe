// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"fmt"
)

// getStatusHandler 获取并返回给定状态码的处理器。 md5:d572ff16b68347d0
func (s *X服务) getStatusHandler(status int, r *Request) []HandlerFunc {
	domains := []string{r.X取主机名(), DefaultDomainName}
	for _, domain := range domains {
		if f, ok := s.statusHandlerMap[s.statusHandlerKey(status, domain)]; ok {
			return f
		}
	}
	return nil
}

// addStatusHandler 为给定的状态码设置处理器。
// 参数 `pattern` 的格式形如：domain#status
// md5:cd0f0b9754ee5b43
func (s *X服务) addStatusHandler(pattern string, handler HandlerFunc) {
	if s.statusHandlerMap[pattern] == nil {
		s.statusHandlerMap[pattern] = make([]HandlerFunc, 0)
	}
	s.statusHandlerMap[pattern] = append(s.statusHandlerMap[pattern], handler)
}

// statusHandlerKey 根据给定的状态和域名创建并返回一个键。 md5:1a4aa99d2a1f13c7
func (s *X服务) statusHandlerKey(status int, domain string) string {
	return fmt.Sprintf("%s#%d", domain, status)
}

// X绑定状态码中间件 为给定的状态码注册处理器。 md5:c94c3dd2e5b4197e
func (s *X服务) X绑定状态码中间件(状态码 int, 处理函数 HandlerFunc) {
	s.addStatusHandler(s.statusHandlerKey(状态码, DefaultDomainName), 处理函数)
}

// X绑定状态码中间件Map 使用映射为给定的状态码注册处理器。 md5:a9ee1be3cd8089de
func (s *X服务) X绑定状态码中间件Map(中间件Map map[int]HandlerFunc) {
	for k, v := range 中间件Map {
		s.X绑定状态码中间件(k, v)
	}
}
