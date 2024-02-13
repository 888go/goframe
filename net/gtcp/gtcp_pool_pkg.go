// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类

import (
	"time"
)

// SendPkg 将包含 `data` 的数据包发送到连接。
// 可选参数 `option` 指定发送数据包的选项。
func (c *PoolConn) SendPkg(data []byte, option ...PkgOption) (err error) {
	if err = c.Conn.SendPkg(data, option...); err != nil && c.status == connStatusUnknown {
		if v, e := c.pool.NewFunc(); e == nil {
			c.Conn = v.(*PoolConn).Conn
			err = c.Conn.SendPkg(data, option...)
		} else {
			err = e
		}
	}
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return err
}

// RecvPkg 通过简单的包协议从连接中接收数据包。
// 可选参数 `option` 指定了接收数据包时的包选项。
func (c *PoolConn) RecvPkg(option ...PkgOption) ([]byte, error) {
	data, err := c.Conn.RecvPkg(option...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvPkgWithTimeout 使用简单包协议，以超时方式从连接中读取数据。
func (c *PoolConn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error) {
	if err := c.SetDeadlineRecv(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		_ = c.SetDeadlineRecv(time.Time{})
	}()
	data, err = c.RecvPkg(option...)
	return
}

// SendPkgWithTimeout 使用简单的包协议并设置超时，向连接写入数据。
func (c *PoolConn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error) {
	if err := c.SetDeadlineSend(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		_ = c.SetDeadlineSend(time.Time{})
	}()
	err = c.SendPkg(data, option...)
	return
}

// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞等待读取响应。
func (c *PoolConn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkg(option...)
	} else {
		return nil, err
	}
}

// SendRecvPkgWithTimeout 使用简单包协议并通过超时读取从连接中接收数据。
func (c *PoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkgWithTimeout(timeout, option...)
	} else {
		return nil, err
	}
}
