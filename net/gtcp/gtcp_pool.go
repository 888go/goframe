// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtcp

import (
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gpool"
)

// PoolConn 是具有连接池特性的 TCP 连接。
// 注意，它不是一个连接池或连接管理器，它仅仅是一个 TCP 连接对象。
type PoolConn struct {
	*Conn              // 基础连接对象。
	pool   *gpool.Pool // 连接池，虽然并非真正的连接池，但实现了连接复用的功能。
	status int         // 当前连接的状态，用于标记此连接是否可用。
}

const (
	defaultPoolExpire = 10 * time.Second // Default TTL for connection in the pool.
// 连接池中连接的默认生存时间（TTL）。
	connStatusUnknown = 0                // Means it is unknown it's connective or not.
	connStatusActive  = 1                // 表示它现在已经连接。
	connStatusError   = 2                // 这意味着它应该被关闭并从池中移除。
)

var (
	// addressPoolMap 是一个映射，用于将地址与其对应的池对象关联。
	addressPoolMap = gmap.NewStrAnyMap(true)
)

// NewPoolConn 创建并返回一个具有连接池功能的连接。
func NewPoolConn(addr string, timeout ...time.Duration) (*PoolConn, error) {
	v := addressPoolMap.GetOrSetFuncLock(addr, func() interface{} {
		var pool *gpool.Pool
		pool = gpool.New(defaultPoolExpire, func() (interface{}, error) {
			if conn, err := NewConn(addr, timeout...); err == nil {
				return &PoolConn{conn, pool, connStatusActive}, nil
			} else {
				return nil, err
			}
		})
		return pool
	})
	value, err := v.(*gpool.Pool).Get()
	if err != nil {
		return nil, err
	}
	return value.(*PoolConn), nil
}

// Close函数会将活跃的连接归还给连接池，如果该连接处于非活跃状态，则关闭该连接。
//
// 注意，如果通过`c`调用Close函数来关闭自身，则此后不能再使用`c`。
func (c *PoolConn) Close() error {
	if c.pool != nil && c.status == connStatusActive {
		c.status = connStatusUnknown
		return c.pool.Put(c)
	}
	return c.Conn.Close()
}

// Send 将数据写入连接。如果写入数据失败，它将从其连接池中获取一个新的连接。
func (c *PoolConn) Send(data []byte, retry ...Retry) error {
	err := c.Conn.Send(data, retry...)
	if err != nil && c.status == connStatusUnknown {
		if v, e := c.pool.Get(); e == nil {
			c.Conn = v.(*PoolConn).Conn
			err = c.Send(data, retry...)
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

// Recv 从连接中接收数据。
func (c *PoolConn) Recv(length int, retry ...Retry) ([]byte, error) {
	data, err := c.Conn.Recv(length, retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvLine 从连接中读取数据，直到读取到字符 '\n'。
// 注意，返回的结果不包含最后的字符 '\n'。
func (c *PoolConn) RecvLine(retry ...Retry) ([]byte, error) {
	data, err := c.Conn.RecvLine(retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvTill从连接中读取数据，直到读取到指定字节序列`til`为止。
// 注意，返回的结果中包含最后的字节序列`til`。
func (c *PoolConn) RecvTill(til []byte, retry ...Retry) ([]byte, error) {
	data, err := c.Conn.RecvTill(til, retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvWithTimeout函数以超时方式从连接中读取数据。
func (c *PoolConn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error) {
	if err := c.SetDeadlineRecv(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		_ = c.SetDeadlineRecv(time.Time{})
	}()
	data, err = c.Recv(length, retry...)
	return
}

// SendWithTimeout 函数在设定的超时时间内向连接写入数据。
func (c *PoolConn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error) {
	if err := c.SetDeadlineSend(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		_ = c.SetDeadlineSend(time.Time{})
	}()
	err = c.Send(data, retry...)
	return
}

// SendRecv 向连接写入数据，并阻塞等待读取响应。
func (c *PoolConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.Recv(receive, retry...)
	} else {
		return nil, err
	}
}

// SendRecvWithTimeout 函数向连接写入数据并设定超时读取响应。
func (c *PoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.RecvWithTimeout(receive, timeout, retry...)
	} else {
		return nil, err
	}
}
