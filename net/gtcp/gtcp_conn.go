// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
)

// Conn 是 TCP 连接对象。
type Conn struct {
	net.Conn                     // 底层TCP连接对象。
	reader         *bufio.Reader // 连接的缓冲读取器
	deadlineRecv   time.Time     // 读取操作的超时点。
	deadlineSend   time.Time     // 写入操作的超时点。
	bufferWaitRecv time.Duration // 读取缓冲区的间隔时长。
}

const (
	// 默认读取缓冲区的间隔时间
	receiveAllWaitTimeout = time.Millisecond
)

// NewConn 根据给定的地址创建并返回一个新的连接。
func NewConn(addr string, timeout ...time.Duration) (*Conn, error) {
	if conn, err := NewNetConn(addr, timeout...); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnTLS 根据给定的地址和 TLS 配置创建并返回一个新的 TLS 连接。
func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error) {
	if conn, err := NewNetConnTLS(addr, tlsConfig); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnKeyCrt 根据给定的地址和 TLS 证书及密钥文件创建并返回一个新的 TLS 连接。
func NewConnKeyCrt(addr, crtFile, keyFile string) (*Conn, error) {
	if conn, err := NewNetConnKeyCrt(addr, crtFile, keyFile); err == nil {
		return NewConnByNetConn(conn), nil
	} else {
		return nil, err
	}
}

// NewConnByNetConn 根据给定的 net.Conn 对象创建并返回一个 TCP 连接对象。
func NewConnByNetConn(conn net.Conn) *Conn {
	return &Conn{
		Conn:           conn,
		reader:         bufio.NewReader(conn),
		deadlineRecv:   time.Time{},
		deadlineSend:   time.Time{},
		bufferWaitRecv: receiveAllWaitTimeout,
	}
}

// Send 将数据写入远程地址。
func (c *Conn) Send(data []byte, retry ...Retry) error {
	for {
		if _, err := c.Write(data); err != nil {
			// 连接已关闭。
			if err == io.EOF {
				return err
			}
			// 即使重试后仍然失败。
			if len(retry) == 0 || retry[0].Count == 0 {
				err = 错误类.X多层错误(err, `Write data failed`)
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
// 注意：
//  1. 如果 length 等于 0，表示它会从当前缓冲区接收数据并立即返回。
//  2. 如果 length 小于 0，表示它会从连接中接收所有数据并返回，直到没有更多的数据从连接中获取为止。开发者需要注意，
//     如果决定从缓冲区接收所有数据，则需要自行处理包解析问题。
//  3. 如果 length 大于 0，表示它会阻塞读取连接中的数据，直到接收到长度为 length 的数据为止。这是最常见的用于接收数据的长度值。
func (c *Conn) Recv(length int, retry ...Retry) ([]byte, error) {
	var (
		err        error  // Reading error.
		size       int    // Reading size.
		index      int    // Received size.
		buffer     []byte // Buffer object.
		bufferWait bool   // 是否设置了缓冲区读取超时。
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
				err = 错误类.X多层错误(err, `SetReadDeadline for connection failed`)
				return nil, err
			}
		}
		size, err = c.reader.Read(buffer[index:])
		if size > 0 {
			index += size
			if length > 0 {
				// 如果指定了`length`，则读取到指定长度为止。
				if index == length {
					break
				}
			} else {
				if index >= defaultReadBufferSize {
					// 如果超过缓冲区大小，它会自动增加其缓冲区大小。
					buffer = append(buffer, make([]byte, defaultReadBufferSize)...)
				} else {
					// 如果接收到的大小小于缓冲区大小，则它会立即返回。
					if !bufferWait {
						break
					}
				}
			}
		}
		if err != nil {
			// 连接已关闭。
			if err == io.EOF {
				break
			}
			// 在读取数据时重新设置超时时间。
			if bufferWait && isTimeout(err) {
				if err = c.SetReadDeadline(c.deadlineRecv); err != nil {
					err = 错误类.X多层错误(err, `SetReadDeadline for connection failed`)
					return nil, err
				}
				err = nil
				break
			}
			if len(retry) > 0 {
				// 即使重试了也会失败。
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
		// 仅从缓冲区读取一次。
		if length == 0 {
			break
		}
	}
	return buffer[:index], err
}

// RecvLine 从连接中读取数据，直到读取到字符 '\n'。
// 注意，返回的结果不包含最后的字符 '\n'。
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

// RecvTill从连接中读取数据，直到读取到指定字节序列`til`为止。
// 注意，返回的结果中包含最后的字节序列`til`。
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

// RecvWithTimeout函数以超时方式从连接中读取数据。
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

// SendWithTimeout 函数在设定的超时时间内向连接写入数据。
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

// SendRecv 向连接写入数据，并阻塞等待读取响应。
func (c *Conn) SendRecv(data []byte, length int, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.Recv(length, retry...)
	} else {
		return nil, err
	}
}

// SendRecvWithTimeout 函数向连接写入数据并设定超时读取响应。
func (c *Conn) SendRecvWithTimeout(data []byte, length int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.RecvWithTimeout(length, timeout, retry...)
	} else {
		return nil, err
	}
}

// SetDeadline 设置当前连接的截止时间。
func (c *Conn) SetDeadline(t time.Time) (err error) {
	if err = c.Conn.SetDeadline(t); err == nil {
		c.deadlineRecv = t
		c.deadlineSend = t
	}
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `SetDeadline for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineRecv为当前连接设置接收的截止时间。
func (c *Conn) SetDeadlineRecv(t time.Time) (err error) {
	if err = c.SetReadDeadline(t); err == nil {
		c.deadlineRecv = t
	}
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `SetDeadlineRecv for connection failed with "%s"`, t)
	}
	return err
}

// SetDeadlineSend 为当前连接设置发送的截止时间。
func (c *Conn) SetDeadlineSend(t time.Time) (err error) {
	if err = c.SetWriteDeadline(t); err == nil {
		c.deadlineSend = t
	}
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `SetDeadlineSend for connection failed with "%s"`, t)
	}
	return err
}

// SetBufferWaitRecv 设置从连接中读取所有数据时的缓冲等待超时时间。
// 等待时长不能过长，否则可能导致接收远程地址数据延迟。
func (c *Conn) SetBufferWaitRecv(bufferWaitDuration time.Duration) {
	c.bufferWaitRecv = bufferWaitDuration
}
