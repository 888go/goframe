// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包gsvc提供了服务注册和发现的定义。 md5:d3c854663f57d96a
package gsvc

import (
	"github.com/gogf/gf/v2/text/gstr"
)

// NewEndpoints 从多个地址创建并返回 Endpoints，例如：
// "192.168.1.100:80,192.168.1.101:80"。 md5:a9ff1a4a1317ab38
func NewEndpoints(addresses string) Endpoints {
	endpoints := make([]Endpoint, 0)
	for _, address := range gstr.SplitAndTrim(addresses, EndpointsDelimiter) {
		endpoints = append(endpoints, NewEndpoint(address))
	}
	return endpoints
}

// String 方法将 Endpoints 格式化为字符串，类似于：
// "192.168.1.100:80,192.168.1.101:80" md5:f254f1ef4b38d633
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
