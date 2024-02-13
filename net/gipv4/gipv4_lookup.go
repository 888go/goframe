// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一份。
//

package ipv4类

import (
	"net"
	"strings"
)

// GetHostByName 返回给定互联网主机名对应的IPv4地址。
func GetHostByName(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// GetHostsByName 函数根据给定的互联网主机名返回对应的 IPv4 地址列表。
func GetHostsByName(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		var ipStrings []string
		for _, v := range ips {
			if v.To4() != nil {
				ipStrings = append(ipStrings, v.String())
			}
		}
		return ipStrings, nil
	}
	return nil, err
}

// GetNameByAddr 根据给定的IP地址返回相应的Internet主机名。
func GetNameByAddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}
