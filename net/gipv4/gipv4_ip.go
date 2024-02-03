// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一份。
//

package gipv4

import (
	"net"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/errors/gerror"
)

// GetIpArray 获取并返回当前主机的所有IP地址。
func GetIpArray() (ips []string, err error) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		err = gerror.Wrap(err, `net.InterfaceAddrs failed`)
		return nil, err
	}
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips, nil
}

// MustGetIntranetIp 的行为与 GetIntranetIp 相同，但当发生任何错误时，它会触发 panic。
func MustGetIntranetIp() string {
	ip, err := GetIntranetIp()
	if err != nil {
		panic(err)
	}
	return ip
}

// GetIntranetIp 获取并返回当前机器的第一个内网IP地址。
func GetIntranetIp() (ip string, err error) {
	ips, err := GetIntranetIpArray()
	if err != nil {
		return "", err
	}
	if len(ips) == 0 {
		return "", gerror.New("no intranet ip found")
	}
	return ips[0], nil
}

// GetIntranetIpArray 获取并返回当前机器的内网IP列表。
func GetIntranetIpArray() (ips []string, err error) {
	var (
		addresses  []net.Addr
		interFaces []net.Interface
	)
	interFaces, err = net.Interfaces()
	if err != nil {
		err = gerror.Wrap(err, `net.Interfaces failed`)
		return ips, err
	}
	for _, interFace := range interFaces {
		if interFace.Flags&net.FlagUp == 0 {
			// interface down
			continue
		}
		if interFace.Flags&net.FlagLoopback != 0 {
			// 回环接口
			continue
		}
		// 忽略守卫桥接
		if strings.HasPrefix(interFace.Name, "w-") {
			continue
		}
		addresses, err = interFace.Addrs()
		if err != nil {
			err = gerror.Wrap(err, `interFace.Addrs failed`)
			return ips, err
		}
		for _, addr := range addresses {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				// 不是IPv4地址
				continue
			}
			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// IsIntranet 检查并返回给定的IP是否为内网IP。
//
// 本地（Local）：127.0.0.1
// A类：10.0.0.0--10.255.255.255
// B类：172.16.0.0--172.31.255.255
// C类：192.168.0.0--192.168.255.255
func IsIntranet(ip string) bool {
	if ip == "127.0.0.1" {
		return true
	}
	array := strings.Split(ip, ".")
	if len(array) != 4 {
		return false
	}
	// A
	if array[0] == "10" || (array[0] == "192" && array[1] == "168") {
		return true
	}
	// C
	if array[0] == "192" && array[1] == "168" {
		return true
	}
	// B
	if array[0] == "172" {
		second, err := strconv.ParseInt(array[1], 10, 64)
		if err != nil {
			return false
		}
		if second >= 16 && second <= 31 {
			return true
		}
	}
	return false
}
