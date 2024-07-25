// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtcp

import (
	"crypto/rand"
	"crypto/tls"
	"net"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	defaultConnTimeout    = 30 * time.Second       // 默认的连接超时时间。 md5:f32319a8522e8f90
	defaultRetryInternal  = 100 * time.Millisecond // 默认重试间隔。 md5:d53e6b260a9e594d
	defaultReadBufferSize = 128                    // 读取的字节缓冲区大小。 md5:3bb21d80469c9916
)

type Retry struct {
	Count    int           // Retry count.
	Interval time.Duration // Retry interval.
}

// NewNetConn 创建并返回一个具有给定地址（如 "127.0.0.1:80"）的net.Conn。
// 可选参数 `timeout` 指定了建立连接的超时时间。 md5:2e0124ac2d5ba04b
func NewNetConn(address string, timeout ...time.Duration) (net.Conn, error) {
	var (
		network  = `tcp`
		duration = defaultConnTimeout
	)
	if len(timeout) > 0 {
		duration = timeout[0]
	}
	conn, err := net.DialTimeout(network, address, duration)
	if err != nil {
		err = gerror.Wrapf(
			err,
			`net.DialTimeout failed with network "%s", address "%s", timeout "%s"`,
			network, address, duration,
		)
	}
	return conn, err
}

// NewNetConnTLS 创建并返回一个具有给定地址（如 "127.0.0.1:80"）的 TLS net.Conn。
// 可选参数 `timeout` 指定了建立连接时的超时时间。 md5:5eb25eb4d9f5078a
func NewNetConnTLS(address string, tlsConfig *tls.Config, timeout ...time.Duration) (net.Conn, error) {
	var (
		network = `tcp`
		dialer  = &net.Dialer{
			Timeout: defaultConnTimeout,
		}
	)
	if len(timeout) > 0 {
		dialer.Timeout = timeout[0]
	}
	conn, err := tls.DialWithDialer(dialer, network, address, tlsConfig)
	if err != nil {
		err = gerror.Wrapf(
			err,
			`tls.DialWithDialer failed with network "%s", address "%s", timeout "%s", tlsConfig "%v"`,
			network, address, dialer.Timeout, tlsConfig,
		)
	}
	return conn, err
}

// NewNetConnKeyCrt 创建并返回一个带有给定TLS证书和密钥文件的TLS net.Conn，地址类似于"127.0.0.1:80"。可选参数`timeout`指定了连接超时时间。 md5:232eecc025740731
func NewNetConnKeyCrt(addr, crtFile, keyFile string, timeout ...time.Duration) (net.Conn, error) {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return nil, err
	}
	return NewNetConnTLS(addr, tlsConfig, timeout...)
}

// Send 建立连接到 `address`，向连接写入 `data`，然后关闭连接。
// 可选参数 `retry` 指定在写入数据失败时的重试策略。 md5:657cbdf2b2958d6f
func Send(address string, data []byte, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Send(data, retry...)
}

// SendRecv 会创建到 `address` 的连接，向该连接写入 `data`，接收响应，
// 然后关闭连接。
//
// 参数 `length` 指定等待接收的字节数量。如果 `length` 为 -1，则接收缓冲区的所有内容并返回。
//
// 可选参数 `retry` 指定了在写入数据失败时的重试策略。 md5:2f0794c80f81d806
func SendRecv(address string, data []byte, length int, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecv(data, length, retry...)
}

// SendWithTimeout 在发送逻辑中添加了写入超时的限制。 md5:3ede704cb632bc5e
func SendWithTimeout(address string, data []byte, timeout time.Duration, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendWithTimeout(data, timeout, retry...)
}

// SendRecvWithTimeout 在限制读取超时的情况下执行SendRecv逻辑。 md5:a0b595ec27ab2abf
func SendRecvWithTimeout(address string, data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvWithTimeout(data, receive, timeout, retry...)
}

// isTimeout 检查给定的 `err` 是否是超时错误。 md5:c277cc8323b1a413
func isTimeout(err error) bool {
	if err == nil {
		return false
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return false
}

// LoadKeyCrt 根据给定的证书和密钥文件创建并返回一个 TLS 配置对象。 md5:e31385756c06b0a4
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error) {
	crtPath, err := gfile.Search(crtFile)
	if err != nil {
		return nil, err
	}
	keyPath, err := gfile.Search(keyFile)
	if err != nil {
		return nil, err
	}
	crt, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		return nil, gerror.Wrapf(err,
			`tls.LoadX509KeyPair failed for certFile "%s" and keyFile "%s"`,
			crtPath, keyPath,
		)
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	return tlsConfig, nil
}

// MustGetFreePort 的行为与 GetFreePort 相同，但如果发生任何错误，它会直接 panic。 md5:a1ae43bc1faffe59
func MustGetFreePort() int {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePort 获取并返回一个空闲的端口号。 md5:52dbf7a2d6e71da6
func GetFreePort() (port int, err error) {
	var (
		network = `tcp`
		address = `:0`
	)
	resolvedAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return 0, gerror.Wrapf(
			err,
			`net.ResolveTCPAddr failed for network "%s", address "%s"`,
			network, address,
		)
	}
	l, err := net.ListenTCP(network, resolvedAddr)
	if err != nil {
		return 0, gerror.Wrapf(
			err,
			`net.ListenTCP failed for network "%s", address "%s"`,
			network, resolvedAddr.String(),
		)
	}
	port = l.Addr().(*net.TCPAddr).Port
	if err = l.Close(); err != nil {
		err = gerror.Wrapf(
			err,
			`close listening failed for network "%s", address "%s", port "%d"`,
			network, resolvedAddr.String(), port,
		)
	}
	return
}

// GetFreePorts 获取并返回指定数量的空闲端口。 md5:ea99fb15b5bbc0fb
func GetFreePorts(count int) (ports []int, err error) {
	var (
		network = `tcp`
		address = `:0`
	)
	for i := 0; i < count; i++ {
		resolvedAddr, err := net.ResolveTCPAddr(network, address)
		if err != nil {
			return nil, gerror.Wrapf(
				err,
				`net.ResolveTCPAddr failed for network "%s", address "%s"`,
				network, address,
			)
		}
		l, err := net.ListenTCP(network, resolvedAddr)
		if err != nil {
			return nil, gerror.Wrapf(
				err,
				`net.ListenTCP failed for network "%s", address "%s"`,
				network, resolvedAddr.String(),
			)
		}
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
		_ = l.Close()
	}
	return ports, nil
}
