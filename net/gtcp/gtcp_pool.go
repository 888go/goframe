// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtcp

import (
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gpool"
)

// PoolConn 是一个具有连接池特性的TCP连接。
// 注意，它本身并不是一个连接池或连接管理器，而只是一个TCP连接对象。 md5:ba1f65d3a4240a38
type PoolConn struct {
	*Conn              // 底层连接对象。 md5:b967bfe4f6e1fc27
	pool   *gpool.Pool // 连接池，其实不是一个真正的连接池，而是一个连接的重用池。 md5:a26d386f822f05df
	status int         // 当前连接的状态，用于标记此连接是否可用。 md5:12cd6637c98f2ef8
}

const defaultPoolExpire = 10 * time.Second // 连接池中连接的默认TTL（生存时间）。 md5:0bf58836ef7b5d32

const (
	connStatusUnknown = iota // Means it is unknown it's connective or not.
	connStatusActive         // 意味着它现在是连接状态。 md5:fc77bf21979d8581
	connStatusError          // 表示应该关闭并从池中移除。 md5:0a8b237e0bb2ab9a
)

var (
	// addressPoolMap 是一个将地址映射到其池对象的映射。 md5:8e4ae3e1f1fdc0a6
	addressPoolMap = gmap.NewStrAnyMap(true)
)

// NewPoolConn 创建并返回一个具有连接池特性的连接。 md5:ee2281aa2be8c181
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

// Close将活动的连接放回池中，如果连接未活动，则将其关闭。
//
// 请注意，如果`c`调用Close函数关闭自身，那么`c`将无法再次使用。 md5:8596872730e65b10
func (c *PoolConn) Close() error {
	if c.pool != nil && c.status == connStatusActive {
		c.status = connStatusUnknown
		return c.pool.Put(c)
	}
	return c.Conn.Close()
}

// Send 将数据写入连接。如果写数据失败，它会从其池中获取一个新的连接。 md5:a5cfc10ec76d87e1
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

// Recv 从连接中接收数据。 md5:d32a0574f5be517a
func (c *PoolConn) Recv(length int, retry ...Retry) ([]byte, error) {
	data, err := c.Conn.Recv(length, retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvLine 从连接中读取数据，直到读取到字符 '\n'。注意，返回的结果不包含最后一个字符 '\n'。 md5:e8f4d38a9d0e03e2
func (c *PoolConn) RecvLine(retry ...Retry) ([]byte, error) {
	data, err := c.Conn.RecvLine(retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvTill 从连接中读取数据，直到读取到字节 `til` 为止。
// 注意，返回的结果中包含最后一个字节 `til`。 md5:3d5a6b2420bd7164
func (c *PoolConn) RecvTill(til []byte, retry ...Retry) ([]byte, error) {
	data, err := c.Conn.RecvTill(til, retry...)
	if err != nil {
		c.status = connStatusError
	} else {
		c.status = connStatusActive
	}
	return data, err
}

// RecvWithTimeout 在连接上读取数据，带有超时设置。 md5:9c30fddfcddde9a2
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

// SendWithTimeout 在超时时间内向连接写入数据。 md5:776c26aa00723dd4
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

// SendRecv 向连接写入数据，并阻塞读取响应。 md5:a92dbf9e10bfe35b
func (c *PoolConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.Recv(receive, retry...)
	} else {
		return nil, err
	}
}

// SendRecvWithTimeout 向连接写入数据并带有超时时间地读取响应。 md5:154815490fa55262
func (c *PoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.RecvWithTimeout(receive, timeout, retry...)
	} else {
		return nil, err
	}
}
