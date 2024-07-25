// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtcp

import (
	"time"
)

// SendPkg 向连接发送一个包含 `data` 的数据包。
// 可选参数 `option` 用于指定发送数据包时的选项。 md5:b992093685758441
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

// RecvPkg 使用简单的包协议从连接接收包。
// 可选的`option`参数指定了接收包时的选项。 md5:6b708b1338c6bb8c
func (c *PoolConn) RecvPkg(option ...PkgOption) ([]byte, error) {
	data, err := c.Conn.RecvPkg(option...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvPkgWithTimeout 使用简单包协议，从连接中读取数据，同时设置超时。 md5:5e1d4882f4476862
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

// 使用简单包协议带超时时间地向连接发送数据。 md5:3f89f6011aed63bc
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

// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞读取响应。 md5:c157760431f11896
func (c *PoolConn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkg(option...)
	} else {
		return nil, err
	}
}

// SendRecvPkgWithTimeout 使用简单包协议，带超时读取连接上的数据。 md5:41984892fec65364
func (c *PoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkgWithTimeout(timeout, option...)
	} else {
		return nil, err
	}
}
