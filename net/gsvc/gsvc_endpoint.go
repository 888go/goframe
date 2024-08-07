// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gsvc提供了服务注册和发现的定义。 md5:d3c854663f57d96a
package gsvc

import (
	"fmt"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// LocalEndpoint 实现了接口 Endpoint。 md5:2c8da8dce28b09e7
type LocalEndpoint struct {
	host string // host可以是IPv4或IPv6地址。 md5:a35907a310997b41
	port int    // port 是通常所说的端口。 md5:543d477387d04665
}

// NewEndpoint 从地址字符串（格式为"host:port"，如："192.168.1.100:80"）创建并返回一个Endpoint。
// md5:837de544fe4ec26d
func NewEndpoint(address string) Endpoint {
	array := gstr.X分割并忽略空值(address, EndpointHostPortDelimiter)
	if len(array) != 2 {
		panic(gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`invalid address "%s" for creating endpoint, endpoint address is like "ip:port"`,
			address,
		))
	}
	return &LocalEndpoint{
		host: array[0],
		port: gconv.X取整数(array[1]),
	}
}

// Host 返回服务的 IPv4/IPv6 地址。 md5:c70938f835a0f6e4
func (e *LocalEndpoint) Host() string {
	return e.host
}

// Port 返回服务的端口。 md5:1650bc955f20ce4c
func (e *LocalEndpoint) Port() int {
	return e.port
}

// String 方法将Endpoint格式化为字符串，例如：192.168.1.100:80。 md5:b9ebe410fee82ac0
func (e *LocalEndpoint) String() string {
	return fmt.Sprintf(`%s:%d`, e.host, e.port)
}
