// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtcp

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Conn是TCP连接对象。. md5:d539be0916b2b1b8
type Conn struct {
	net.Conn                     // 基础TCP连接对象。. md5:b7b1c0f618f0d5c9
	reader         *bufio.Reader // 连接的缓冲读取器。. md5:490834adc66f0794
	deadlineRecv   time.Time     // 读取的超时点。. md5:c3a775342f7a3b1e
	deadlineSend   time.Time     // 写入的超时点。. md5:27ca5c03a339a72e
	bufferWaitRecv time.Duration // 读取缓冲区的间隔持续时间。. md5:373de61f6fb41d4d
}

const (
	// 读取缓冲区的默认间隔时间。. md5:f0502056d8bc6f66
	receiveAllWaitTimeout = time.Millisecond
)

// NewConn 创建并返回一个与给定地址的新连接。. md5:b8d7c0b5ae5f53f0
func NewConn(addr string, timeout ...time.Duration) (*Conn, error) {
	if conn, err := NewNetConn(addr, timeout...); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnTLS 创建并返回一个新的TLS连接，使用给定的地址和TLS配置。
// md5:a21dcb1cad67caa6
func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error) {
	if conn, err := NewNetConnTLS(addr, tlsConfig); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnKeyCrt 创建并返回一个新的带有给定地址和TLS证书及密钥文件的TLS连接。
// md5:b79c43de9a5f13ce
func NewConnKeyCrt(addr, crtFile, keyFile string) (*Conn, error) {
	if conn, err := NewNetConnKeyCrt(addr, crtFile, keyFile); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnByNetConn 根据给定的net.Conn对象创建并返回一个TCP连接对象。. md5:88aad787fa32f138
func NewConnByNetConn(conn net.Conn) *Conn {
	return &Conn{
		Conn:           conn,
		reader:         bufio.NewReader(conn),
		deadlineRecv:   time.Time{},
		deadlineSend:   time.Time{},
		bufferWaitRecv: receiveAllWaitTimeout,
	}
}

// Send 将数据写入远程地址。. md5:445009019bd4a1a8
func (c *Conn) Send(data []byte, retry ...Retry) error {
	for {
		if _, err := c.Write(data); err != nil {
			// Connection closed.
			if err == io.EOF {
				return err
			}
			// 重试后仍然失败。. md5:b819d69935ab7496
			if len(retry) == 0 || retry[0].Count == 0 {
				err = gerror.Wrap(err, `Write data failed`)
				return err
			}
			if len(retry) > 0 {
				retry[0].Count--
				if retry[0].Interval == 0 {
					retry[0].Interval = defaultRetryInternal
				}
				time.Sleep(retry[0].Interval)
			}
		} else {
			return nil
		}
	}
}

// Recv 从连接中接收并返回数据。
// 
// 注意，
// 1. 如果长度为0，表示它会从当前缓冲区接收数据并立即返回。
// 2. 如果长度小于0，表示它会接收连接中的所有数据，并返回直到没有更多的数据。如果决定接收缓冲区中的所有数据，开发者需要注意自行解析数据包。
// 3. 如果长度大于0，表示它会阻塞，直到接收到指定长度的数据。这是最常用的用于接收数据的长度值。
// md5:75d42f229725a3f7
func (c *Conn) Recv(length int, retry ...Retry) ([]byte, error) {
	var (
		err        error  // Reading error.
		size       int    // Reading size.
		index      int    // Received size.
		buffer     []byte // Buffer object.
		bufferWait bool   // 是否设置了缓冲区读取超时。. md5:7a71357a4bf2a1e8
	)
	if length > 0 {
		buffer = make([]byte, length)
	} else {
		buffer = make([]byte, defaultReadBufferSize)
	}

	for {
		if length < 0 && index > 0 {
			bufferWait = true
			if err = c.SetReadDeadline(time.Now().Add(c.bufferWaitRecv)); err != nil {
				err = gerror.Wrap(err, `SetReadDeadline for connection failed`)
				return nil, err
			}
		}
		size, err = c.reader.Read(buffer[index:])
		if size > 0 {
			index += size
			if length > 0 {
				// 如果指定了`length`，则读取`length`大小的内容。. md5:1216f8afe5006977
				if index == length {
					break
				}
			} else {
				if index >= defaultReadBufferSize {
					// 如果它超过了缓冲区的大小，那么它会自动增加其缓冲区的大小。. md5:2dfec7c7ce3557ab
					buffer = append(buffer, make([]byte, defaultReadBufferSize)...)
				} else {
					// 如果接收到的大小小于缓冲区大小，它会立即返回。. md5:e98695cece9485a3
					if !bufferWait {
						break
					}
				}
			}
		}
		if err != nil {
			// Connection closed.
			if err == io.EOF {
				break
			}
			// 重新设置读取数据时的超时时间。. md5:a04ae0a806c5a3c6
			if bufferWait && isTimeout(err) {
				if err = c.SetReadDeadline(c.deadlineRecv); err != nil {
					err = gerror.Wrap(err, `SetReadDeadline for connection failed`)
					return nil, err
				}
				err = nil
				break
			}
			if len(retry) > 0 {
				// 即使重试也会失败。. md5:7f32623c1f255555
				if retry[0].Count == 0 {
					break
				}
				retry[0].Count--
				if retry[0].Interval == 0 {
					retry[0].Interval = defaultRetryInternal
				}
				time.Sleep(retry[0].Interval)
				continue
			}
			break
		}
		// 从缓冲区只读一次。. md5:26f6ad162c136702
		if length == 0 {
			break
		}
	}
	return buffer[:index], err
}

// RecvLine 从连接中读取数据，直到读取到字符 '\n'。注意，返回的结果不包含最后一个字符 '\n'。
// md5:e8f4d38a9d0e03e2
func (c *Conn) RecvLine(retry ...Retry) ([]byte, error) {
	var (
		err    error
		buffer []byte
		data   = make([]byte, 0)
	)
	for {
		buffer, err = c.Recv(1, retry...)
		if len(buffer) > 0 {
			if buffer[0] == '\n' {
				data = append(data, buffer[:len(buffer)-1]...)
				break
			} else {
				data = append(data, buffer...)
			}
		}
		if err != nil {
			break
		}
	}
	return data, err
}

// RecvTill 从连接中读取数据，直到读取到字节 `til` 为止。
// 注意，返回的结果中包含最后一个字节 `til`。
// md5:3d5a6b2420bd7164
func (c *Conn) RecvTill(til []byte, retry ...Retry) ([]byte, error) {
	var (
		err    error
		buffer []byte
		data   = make([]byte, 0)
		length = len(til)
	)
	for {
		buffer, err = c.Recv(1, retry...)
		if len(buffer) > 0 {
			if length > 0 &&
				len(data) >= length-1 &&
				buffer[0] == til[length-1] &&
				bytes.EqualFold(data[len(data)-length+1:], til[:length-1]) {
				data = append(data, buffer...)
				break
			} else {
				data = append(data, buffer...)
			}
		}
		if err != nil {
			break
		}
	}
	return data, err
}

// RecvWithTimeout 在连接上读取数据，带有超时设置。. md5:9c30fddfcddde9a2
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

// SendWithTimeout 在超时时间内向连接写入数据。. md5:776c26aa00723dd4
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

// SendRecv 向连接写入数据，并阻塞读取响应。. md5:a92dbf9e10bfe35b
func (c *Conn) SendRecv(data []byte, length int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.Recv(length, retry...)
	} else {
		return nil, err
	}
}

// SendRecvWithTimeout 向连接写入数据并带有超时时间地读取响应。. md5:154815490fa55262
func (c *Conn) SendRecvWithTimeout(data []byte, length int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.RecvWithTimeout(length, timeout, retry...)
	} else {
		return nil, err
	}
}

// SetDeadline 设置当前连接的截止时间。. md5:d3ab57d58a8a6b64
func (c *Conn) SetDeadline(t time.Time) (err error) {
	if err = c.Conn.SetDeadline(t); err == nil {
		c.deadlineRecv = t
		c.deadlineSend = t
	}
	if err != nil {
		err = gerror.Wrapf(err, `SetDeadline for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineRecv为当前连接设置接收截止期限。. md5:5cf236fbc63fc1f7
func (c *Conn) SetDeadlineRecv(t time.Time) (err error) {
	if err = c.SetReadDeadline(t); err == nil {
		c.deadlineRecv = t
	}
	if err != nil {
		err = gerror.Wrapf(err, `SetDeadlineRecv for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineSend 设置当前连接的发送截止期限。. md5:9f0d98d0e6beda95
func (c *Conn) SetDeadlineSend(t time.Time) (err error) {
	if err = c.SetWriteDeadline(t); err == nil {
		c.deadlineSend = t
	}
	if err != nil {
		err = gerror.Wrapf(err, `SetDeadlineSend for connection failed with "%s"`, t)
	}
	return err
}

// SetBufferWaitRecv 设置从连接读取所有数据时的缓冲等待超时时间。
// 等待时间不能过长，否则可能会延迟从远程地址接收数据。
// md5:54992dd21ce2360a
func (c *Conn) SetBufferWaitRecv(bufferWaitDuration time.Duration) {
	c.bufferWaitRecv = bufferWaitDuration
}
