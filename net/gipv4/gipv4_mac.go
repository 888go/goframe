// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package ipv4类

import (
	"net"

	gerror "github.com/888go/goframe/errors/gerror"
)

// GetMac 获取并返回当前主机的第一个MAC地址。 md5:59a524347ebb9dcc
func GetMac() (mac string, err error) {
	macs, err := GetMacArray()
	if err != nil {
		return "", err
	}
	if len(macs) > 0 {
		return macs[0], nil
	}
	return "", nil
}

// GetMacArray 获取并返回当前主机的所有MAC地址。 md5:18289e2c88c62bec
func GetMacArray() (macs []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		err = gerror.X多层错误(err, `net.Interfaces failed`)
		return nil, err
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macs = append(macs, macAddr)
	}
	return macs, nil
}
