// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsvc 提供了服务注册与发现的定义。
package gsvc

import (
	"github.com/888go/goframe/text/gstr"
)

// NewEndpoints 创建并返回 Endpoints，它可以从多个地址构建，如：
// "192.168.1.100:80,192.168.1.101:80"。
func NewEndpoints(addresses string) Endpoints {
	endpoints := make([]Endpoint, 0)
	for _, address := range gstr.SplitAndTrim(addresses, EndpointsDelimiter) {
		endpoints = append(endpoints, NewEndpoint(address))
	}
	return endpoints
}

// String函数将Endpoints格式化并以如下形式返回字符串：
// "192.168.1.100:80,192.168.1.101:80"
func (es Endpoints) String() string {
	var s string
	for _, endpoint := range es {
		if s != "" {
			s += EndpointsDelimiter
		}
		s += endpoint.String()
	}
	return s
}
