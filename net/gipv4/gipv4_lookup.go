// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package ipv4类

import (
	"net"
	"strings"
)

// GetHostByName 返回给定互联网主机名对应的 IPv4 地址。 md5:b6dcb532a1a7dd84
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

// GetHostsByName 根据给定的互联网主机名返回一个IPv4地址列表。
// md5:57770571b1c3af32
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

// GetNameByAddr 根据给定的IP地址返回对应的互联网主机名。 md5:9abc34f304140e79
func GetNameByAddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}
