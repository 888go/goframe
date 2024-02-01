// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtcp
import (
	"time"
	)
// SendPkg 将包含 `data` 的数据包发送到 `address`，并关闭连接。
// 可选参数 `option` 指定了发送数据包时的选项。
func SendPkg(address string, data []byte, option ...PkgOption) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkg(data, option...)
}

// SendRecvPkg 将包含 `data` 的数据包发送到 `address`，接收响应并关闭连接。可选参数 `option` 指定了发送数据包时的选项。
func SendRecvPkg(address string, data []byte, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkg(data, option...)
}

// SendPkgWithTimeout 函数向 `address` 发送包含 `data` 的数据包，并设置超时限制，同时关闭连接。可选参数 `option` 指定了发送数据包的相关选项。
func SendPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkgWithTimeout(data, timeout, option...)
}

// SendRecvPkgWithTimeout 函数向 `address` 发送包含 `data` 的数据包，在有限的超时时间内接收响应并关闭连接。
// 可选参数 `option` 用于指定发送数据包的相关选项。
func SendRecvPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkgWithTimeout(data, timeout, option...)
}
