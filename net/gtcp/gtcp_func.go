// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtcp
import (
	"crypto/rand"
	"crypto/tls"
	"net"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfile"
	)
const (
	defaultConnTimeout    = 30 * time.Second       // 默认连接超时时间
	defaultRetryInternal  = 100 * time.Millisecond // 默认重试间隔。
	defaultReadBufferSize = 128                    // (字节) 读取时的缓冲区大小。
)

type Retry struct {
	Count    int           // Retry count.
	Interval time.Duration // Retry interval.
}

// NewNetConn 创建并返回一个 net.Conn，其地址格式如 "127.0.0.1:80"。
// 可选参数 `timeout` 指定了拨号连接的超时时间。
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

// NewNetConnTLS 根据给定的地址（如 "127.0.0.1:80"）创建并返回一个 TLS 安全连接 net.Conn。
// 可选参数 `timeout` 指定了建立连接时的超时时间。
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

// NewNetConnKeyCrt 根据给定的 TLS 证书和密钥文件以及类似 "127.0.0.1:80" 的地址创建并返回一个 TLS 网络连接（net.Conn）。可选参数 `timeout` 指定了建立连接时的超时时间。
func NewNetConnKeyCrt(addr, crtFile, keyFile string, timeout ...time.Duration) (net.Conn, error) {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return nil, err
	}
	return NewNetConnTLS(addr, tlsConfig, timeout...)
}

// Send 函数创建到 `address` 的连接，将 `data` 数据写入该连接，然后关闭连接。
// 可选参数 `retry` 指定了在写入数据失败时的重试策略。
func Send(address string, data []byte, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Send(data, retry...)
}

// SendRecv 创建到 `address` 的连接，将 `data` 写入连接，接收响应，然后关闭连接。
//
// 参数 `length` 指定等待接收的字节数量。如果 `length` 为 -1，则接收所有缓冲区内容并返回。
//
// 可选参数 `retry` 指定了在写入数据失败时重试策略。
func SendRecv(address string, data []byte, length int, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecv(data, length, retry...)
}

// SendWithTimeout 在写入超时限制下执行Send逻辑。
func SendWithTimeout(address string, data []byte, timeout time.Duration, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendWithTimeout(data, timeout, retry...)
}

// SendRecvWithTimeout 在读取超时限制下执行SendRecv逻辑。
func SendRecvWithTimeout(address string, data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvWithTimeout(data, receive, timeout, retry...)
}

// isTimeout 检查给定的 `err` 是否为超时错误。
func isTimeout(err error) bool {
	if err == nil {
		return false
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return false
}

// LoadKeyCrt 通过给定的证书和密钥文件创建并返回一个 TLS 配置对象。
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

// MustGetFreePort 的行为与 GetFreePort 相同，但是，如果发生任何错误，它会触发 panic。
func MustGetFreePort() int {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePort 获取并返回一个可用的端口。
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

// GetFreePorts 获取并返回指定数量的空闲端口。
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
