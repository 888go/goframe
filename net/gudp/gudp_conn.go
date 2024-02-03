// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gudp

import (
	"io"
	"net"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
)

// Conn 处理UDP连接。
type Conn struct {
	*net.UDPConn                 // 基础的UDP连接。
	remoteAddr     *net.UDPAddr  // Remote address.
	deadlineRecv   time.Time     // 读取数据的超时点。
	deadlineSend   time.Time     // 数据写入的超时点。
	bufferWaitRecv time.Duration // 读取缓冲区的间隔时长。
}

const (
	defaultRetryInterval  = 100 * time.Millisecond // Retry interval.
	defaultReadBufferSize = 1024                   // (字节)缓冲区大小。
	receiveAllWaitTimeout = time.Millisecond       // 默认读取缓冲区的间隔时间。
)

type Retry struct {
	Count    int           // Max retry count.
	Interval time.Duration // Retry interval.
}

// NewConn创建一个到`remoteAddress`的UDP连接。
// 可选参数`localAddress`用于指定连接的本地地址。
func NewConn(remoteAddress string, localAddress ...string) (*Conn, error) {
	if conn, err := NewNetConn(remoteAddress, localAddress...); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnByNetConn 通过给定的 *net.UDPConn 对象创建一个 UDP 连接对象。
func NewConnByNetConn(udp *net.UDPConn) *Conn {
	return &Conn{
		UDPConn:        udp,
		deadlineRecv:   time.Time{},
		deadlineSend:   time.Time{},
		bufferWaitRecv: receiveAllWaitTimeout,
	}
}

// Send 将数据写入远程地址。
func (c *Conn) Send(data []byte, retry ...Retry) (err error) {
	for {
		if c.remoteAddr != nil {
			_, err = c.WriteToUDP(data, c.remoteAddr)
		} else {
			_, err = c.Write(data)
		}
		if err != nil {
			// 连接已关闭。
			if err == io.EOF {
				return err
			}
			// 即便重试后仍然失败。
			if len(retry) == 0 || retry[0].Count == 0 {
				err = gerror.Wrap(err, `Write data failed`)
				return err
			}
			if len(retry) > 0 {
				retry[0].Count--
				if retry[0].Interval == 0 {
					retry[0].Interval = defaultRetryInterval
				}
				time.Sleep(retry[0].Interval)
			}
		} else {
			return nil
		}
	}
}

// Recv 从远程地址接收并返回数据。
// 参数`buffer`用于自定义接收缓冲区大小。如果`buffer` <= 0，
// 则使用默认的缓冲区大小，即1024字节。
//
// 在UDP协议中有包边界的限制，如果指定的缓冲区大小足够大，我们可以一次性接收一个完整的包。
// 非常注意的是，我们应该一次性接收完整个包，否则剩余的包数据会被丢弃。
func (c *Conn) Recv(buffer int, retry ...Retry) ([]byte, error) {
	var (
		err        error        // Reading error
		size       int          // Reading size
		data       []byte       // Buffer object
		remoteAddr *net.UDPAddr // 当前用于读取的远程地址
	)
	if buffer > 0 {
		data = make([]byte, buffer)
	} else {
		data = make([]byte, defaultReadBufferSize)
	}
	for {
		size, remoteAddr, err = c.ReadFromUDP(data)
		if err == nil {
			c.remoteAddr = remoteAddr
		}
		if err != nil {
			// 连接已关闭。
			if err == io.EOF {
				break
			}
			if len(retry) > 0 {
				// 即使重试了也会失败。
				if retry[0].Count == 0 {
					break
				}
				retry[0].Count--
				if retry[0].Interval == 0 {
					retry[0].Interval = defaultRetryInterval
				}
				time.Sleep(retry[0].Interval)
				continue
			}
			err = gerror.Wrap(err, `ReadFromUDP failed`)
			break
		}
		break
	}
	return data[:size], err
}

// SendRecv 向连接写入数据，并阻塞等待读取响应。
func (c *Conn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err != nil {
		return nil, err
	}
	return c.Recv(receive, retry...)
}

// RecvWithTimeout在指定超时时间内从远程地址读取数据。
func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error) {
	if err = c.SetDeadlineRecv(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		_ = c.SetDeadlineRecv(time.Time{})
	}()
	data, err = c.Recv(length, retry...)
	return
}

// SendWithTimeout 在设定的超时时间内向连接写入数据。
func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error) {
	if err = c.SetDeadlineSend(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		_ = c.SetDeadlineSend(time.Time{})
	}()
	err = c.Send(data, retry...)
	return
}

// SendRecvWithTimeout 将数据写入连接并在超时时间内读取响应。
func (c *Conn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err != nil {
		return nil, err
	}
	return c.RecvWithTimeout(receive, timeout, retry...)
}

// SetDeadline为连接设置读取和写入的截止时间。
func (c *Conn) SetDeadline(t time.Time) (err error) {
	if err = c.UDPConn.SetDeadline(t); err == nil {
		c.deadlineRecv = t
		c.deadlineSend = t
	} else {
		err = gerror.Wrapf(err, `SetDeadline for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineRecv 设置与连接关联的读取截止时间。
func (c *Conn) SetDeadlineRecv(t time.Time) (err error) {
	if err = c.SetReadDeadline(t); err == nil {
		c.deadlineRecv = t
	} else {
		err = gerror.Wrapf(err, `SetDeadlineRecv for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineSend 设置当前连接的发送截止时间。
func (c *Conn) SetDeadlineSend(t time.Time) (err error) {
	if err = c.SetWriteDeadline(t); err == nil {
		c.deadlineSend = t
	} else {
		err = gerror.Wrapf(err, `SetDeadlineSend for connection failed with "%s"`, t)
	}
	return err
}

// SetBufferWaitRecv 设置从连接读取所有数据时的缓冲等待超时时间。
// 等待时长不能过长，否则可能延迟从远程地址接收数据。
func (c *Conn) SetBufferWaitRecv(d time.Duration) {
	c.bufferWaitRecv = d
}

// RemoteAddr 返回当前 UDP 连接的远程地址。
// 注意，由于 c.conn.RemoteAddr() 为 nil，所以不能使用它。
func (c *Conn) RemoteAddr() net.Addr {
	return c.remoteAddr
}
