// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gudp

import (
	"io"
	"net"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Conn 处理 UDP 连接。 md5:3d72ff914b3663e1
type Conn struct {
	*net.UDPConn                 // 底层UDP连接。 md5:a4de01bc082c3b97
	remoteAddr     *net.UDPAddr  // Remote address.
	deadlineRecv   time.Time     // 读取数据的超时点。 md5:115720ebb2ad57ae
	deadlineSend   time.Time     // 写入数据的超时点。 md5:950ee5a16a0f64fc
	bufferWaitRecv time.Duration // 读取缓冲区的间隔持续时间。 md5:373de61f6fb41d4d
}

const (
	defaultRetryInterval  = 100 * time.Millisecond // Retry interval.
	defaultReadBufferSize = 1024                   // (Byte)Buffer size.
	receiveAllWaitTimeout = time.Millisecond       // 读取缓冲区的默认间隔时间。 md5:f0502056d8bc6f66
)

type Retry struct {
	Count    int           // Max retry count.
	Interval time.Duration // Retry interval.
}

// NewConn 创建到 `remoteAddress` 的 UDP 连接。
// 可选参数 `localAddress` 指定连接的本地地址。
// md5:d5e06df7ea2ee28d
func NewConn(remoteAddress string, localAddress ...string) (*Conn, error) {
	if conn, err := NewNetConn(remoteAddress, localAddress...); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnByNetConn 使用给定的 *net.UDPConn 对象创建一个UDP连接对象。 md5:8cbe128848b49656
func NewConnByNetConn(udp *net.UDPConn) *Conn {
	return &Conn{
		UDPConn:        udp,
		deadlineRecv:   time.Time{},
		deadlineSend:   time.Time{},
		bufferWaitRecv: receiveAllWaitTimeout,
	}
}

// Send 将数据写入远程地址。 md5:445009019bd4a1a8
func (c *Conn) Send(data []byte, retry ...Retry) (err error) {
	for {
		if c.remoteAddr != nil {
			_, err = c.WriteToUDP(data, c.remoteAddr)
		} else {
			_, err = c.Write(data)
		}
		if err != nil {
			// Connection closed.
			if err == io.EOF {
				return err
			}
			// 重试后仍然失败。 md5:b819d69935ab7496
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
// 参数 `buffer` 用于自定义接收缓冲区大小。如果 `buffer` <= 0，将使用默认的缓冲区大小，即1024字节。
//
// UDP协议存在分包边界，如果指定的缓冲区大小足够大，我们可以接收到一个完整的数据包。非常重要的是，必须一次性接收完整个包，否则剩下的包数据将会丢失。
// md5:190b81cc02f9d449
func (c *Conn) Recv(buffer int, retry ...Retry) ([]byte, error) {
	var (
		err        error        // Reading error
		size       int          // Reading size
		data       []byte       // Buffer object
		remoteAddr *net.UDPAddr // 当前用于读取的远程地址. md5:7bff72f1a6b8b788
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
			// Connection closed.
			if err == io.EOF {
				break
			}
			if len(retry) > 0 {
				// 即使重试也会失败。 md5:7f32623c1f255555
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

// SendRecv 向连接写入数据并阻塞读取响应。 md5:8416afe592163603
func (c *Conn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err != nil {
		return nil, err
	}
	return c.Recv(receive, retry...)
}

// RecvWithTimeout 带超时限制地从远程地址读取数据。 md5:9e229854a65f6de2
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

// SendWithTimeout 在连接中写入数据，并设置超时时间。 md5:d15d51d6004b2a6a
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

// SendRecvWithTimeout 向连接写入数据，并在超时时间内读取响应。 md5:6aa7751868598fb2
func (c *Conn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err != nil {
		return nil, err
	}
	return c.RecvWithTimeout(receive, timeout, retry...)
}

// SetDeadline 设置与连接相关的读写超时截止时间。 md5:e0438bc956760556
func (c *Conn) SetDeadline(t time.Time) (err error) {
	if err = c.UDPConn.SetDeadline(t); err == nil {
		c.deadlineRecv = t
		c.deadlineSend = t
	} else {
		err = gerror.Wrapf(err, `SetDeadline for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineRecv 设置与连接关联的读取截止时间。 md5:763094b16fe580fe
func (c *Conn) SetDeadlineRecv(t time.Time) (err error) {
	if err = c.SetReadDeadline(t); err == nil {
		c.deadlineRecv = t
	} else {
		err = gerror.Wrapf(err, `SetDeadlineRecv for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineSend 设置当前连接的发送截止期限。 md5:9f0d98d0e6beda95
func (c *Conn) SetDeadlineSend(t time.Time) (err error) {
	if err = c.SetWriteDeadline(t); err == nil {
		c.deadlineSend = t
	} else {
		err = gerror.Wrapf(err, `SetDeadlineSend for connection failed with "%s"`, t)
	}
	return err
}

// SetBufferWaitRecv 设置从连接读取所有数据时的缓冲等待超时时间。
// 等待时间不能过长，否则可能会延迟从远程地址接收数据。
// md5:54992dd21ce2360a
func (c *Conn) SetBufferWaitRecv(d time.Duration) {
	c.bufferWaitRecv = d
}

// RemoteAddr 返回当前UDP连接的远程地址。
// 请注意，它不能使用c.conn.RemoteAddr()，因为该值为nil。
// md5:0a785ae4cb967a81
func (c *Conn) RemoteAddr() net.Addr {
	return c.remoteAddr
}
