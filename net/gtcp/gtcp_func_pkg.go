// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtcp

import "time"

// SendPkg 将包含 `data` 的包发送到 `address`，并关闭连接。可选参数 `option` 指定了发送包的选项。
// md5:45d7e35d28ec87b6
// ff:
// address:
// data:
// option:
func SendPkg(address string, data []byte, option ...PkgOption) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkg(data, option...)
}

// SendRecvPkg 向 `address` 发送一个包含 `data` 的数据包，接收响应并关闭连接。可选参数 `option` 用于指定发送数据包时的选项。
// md5:0bc7256049faf99e
// ff:
// address:
// data:
// option:
func SendRecvPkg(address string, data []byte, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkg(data, option...)
}

// SendPkgWithTimeout 向 `address` 发送包含 `data` 的包裹，并设置超时限制，然后关闭连接。可选参数 `option` 用于指定发送包裹的选项。
// md5:cf66e180792b75b0
// ff:
// address:
// data:
// timeout:
// option:
func SendPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkgWithTimeout(data, timeout, option...)
}

	// SendRecvPkgWithTimeout 向 `address` 发送包含 `data` 的包，同时设置超时限制来接收响应，并关闭连接。可选参数 `option` 指定发送包的选项。
	// md5:576d58e144643a35
// ff:
// address:
// data:
// timeout:
// option:
func SendRecvPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkgWithTimeout(data, timeout, option...)
}
