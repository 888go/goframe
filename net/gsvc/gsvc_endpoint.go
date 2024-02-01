// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsvc 提供了服务注册与发现的定义。
package gsvc
import (
	"fmt"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// LocalEndpoint 实现了接口 Endpoint。
type LocalEndpoint struct {
	host string // host可以是IPv4或IPv6地址。
	port int    // port 是我们通常所说的端口。
}

// NewEndpoint 从形如 "host:port"（例如："192.168.1.100:80"）的地址字符串创建并返回一个 Endpoint 对象。
func NewEndpoint(address string) Endpoint {
	array := gstr.SplitAndTrim(address, EndpointHostPortDelimiter)
	if len(array) != 2 {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid address "%s" for creating endpoint, endpoint address is like "ip:port"`,
			address,
		))
	}
	return &LocalEndpoint{
		host: array[0],
		port: gconv.Int(array[1]),
	}
}

// Host 返回一个服务的 IPv4/IPv6 地址。
func (e *LocalEndpoint) Host() string {
	return e.host
}

// Port 返回一个服务的端口号。
func (e *LocalEndpoint) Port() int {
	return e.port
}

// String 将Endpoint格式化并以字符串形式返回，例如：192.168.1.100:80。
func (e *LocalEndpoint) String() string {
	return fmt.Sprintf(`%s:%d`, e.host, e.port)
}
