// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gipv4

import (
	"net"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

// GetIpArray 获取并返回当前主机的所有IP地址。 md5:6828d92b1a684cd2
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

// MustGetIntranetIp 执行与 GetIntranetIp 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:f08d856493c3c333
func MustGetIntranetIp() string {
	ip, err := GetIntranetIp()
	if err != nil {
		panic(err)
	}
	return ip
}

// GetIntranetIp 获取并返回当前机器的第一个内网IP。 md5:2e53e5f6a86c1f3c
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

// GetIntranetIpArray 获取并返回当前机器的内网IP列表。 md5:48fe9964790750ba
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
			// loop back interface
			continue
		}
		// ignore warden bridge
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
				// not an ipv4 address
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

// IsIntranet 检查并返回给定IP地址是否为内部网络IP。
// 当地：127.0.0.1
// A类：10.0.0.0--10.255.255.255
// B类：172.16.0.0--172.31.255.255
// C类：192.168.0.0--192.168.255.255
// md5:1f4c3df8068af016
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
