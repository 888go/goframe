// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一份。
//

package gipv4
import (
	"net"
	
	"github.com/888go/goframe/errors/gerror"
	)
// GetMac 获取并返回当前主机的第一个MAC地址。
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

// GetMacArray 获取并返回当前主机的所有MAC地址。
func GetMacArray() (macs []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		err = gerror.Wrap(err, `net.Interfaces failed`)
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
