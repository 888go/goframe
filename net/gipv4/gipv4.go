// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

// gipv4 包提供了用于处理IPv4地址的有用API。 md5:dc7fb957be20c17f
package ipv4类

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"

	gregex "github.com/888go/goframe/text/gregex"
)

// Ip2long 将IP地址转换为一个uint32整数。 md5:160031646344d859
func Ip2long(ip string) uint32 {
	netIp := net.ParseIP(ip)
	if netIp == nil {
		return 0
	}
	return binary.BigEndian.Uint32(netIp.To4())
}

// Long2ip 将无符号 32 位整数形式的 IP 地址转换为字符串类型的地址。 md5:de7a5a15d74ae9b6
func Long2ip(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	return net.IP(ipByte).String()
}

// Validate 检查给定的 `ip` 是否为有效的IPv4地址。 md5:d26ab457dd3beb9e
func Validate(ip string) bool {
	return gregex.IsMatchString(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`, ip)
}

// ParseAddress 将 `地址` 解析为它的IP和端口。
// 例如：192.168.1.1:80 -> 192.168.1.1, 80
// md5:224991801d25eab5
func ParseAddress(address string) (string, int) {
	match, err := gregex.MatchString(`^(.+):(\d+)$`, address)
	if err == nil {
		i, _ := strconv.Atoi(match[2])
		return match[1], i
	}
	return "", 0
}

// GetSegment 返回给定IP地址的段。
// 例如：192.168.2.102 -> 192.168.2
// md5:6b442ab0a95dc737
func GetSegment(ip string) string {
	match, err := gregex.MatchString(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`, ip)
	if err != nil || len(match) < 4 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s", match[1], match[2], match[3])
}
