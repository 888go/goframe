// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一份。
//

// Package gipv4 提供了用于处理 IPv4 地址的有用 API。
package gipv4
import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	
	"github.com/888go/goframe/text/gregex"
	)
// Ip2long 将IP地址转换为一个uint32整数。
func Ip2long(ip string) uint32 {
	netIp := net.ParseIP(ip)
	if netIp == nil {
		return 0
	}
	return binary.BigEndian.Uint32(netIp.To4())
}

// Long2ip 将一个 uint32 类型的整数形式 IP 地址转换为其字符串类型地址。
func Long2ip(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	return net.IP(ipByte).String()
}

// Validate用于检查给定的`ip`是否为有效的IPv4地址。
func Validate(ip string) bool {
	return gregex.IsMatchString(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`, ip)
}

// ParseAddress 将 `address` 解析为其 IP 地址和端口号。
// 例如：192.168.1.1:80 -> 192.168.1.1, 80
func ParseAddress(address string) (string, int) {
	match, err := gregex.MatchString(`^(.+):(\d+)$`, address)
	if err == nil {
		i, _ := strconv.Atoi(match[2])
		return match[1], i
	}
	return "", 0
}

// GetSegment 返回给定IP地址的段。例如：192.168.2.102 -> 192.168.2
// ```go
// 获取指定IP地址的段。
// 例如：192.168.2.102，将返回192.168.2
// func GetSegment(ipAddress string) string {
    // 实现代码...
// }
func GetSegment(ip string) string {
	match, err := gregex.MatchString(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`, ip)
	if err != nil || len(match) < 4 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s", match[1], match[2], match[3])
}
