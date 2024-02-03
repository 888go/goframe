// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtcp

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

const (
	// FreePortAddress 表示服务器使用随机空闲端口进行监听。
	FreePortAddress = ":0"
)

const (
	defaultServer = "default"
)

// Server 是一个 TCP 服务器。
type Server struct {
	mu        sync.Mutex   // 用于Server.listen方法的并发安全。-- 此段golang代码会通过数据竞争检测进行测试。
	listen    net.Listener // TCP地址监听器。
	address   string       // 服务监听地址。
	handler   func(*Conn)  // 连接处理器
	tlsConfig *tls.Config  // TLS 配置
}

// 用于单例目的，映射名称到服务器的Map。
var serverMapping = gmap.NewStrAnyMap(true)

// GetServer 函数返回指定名称 `name` 的 TCP 服务器，
// 如果该服务器不存在，则返回一个新的普通 TCP 服务器并命名为 `name`。
// 参数 `name` 用于指定要获取的 TCP 服务器。
func GetServer(name ...interface{}) *Server {
	serverName := defaultServer
	if len(name) > 0 && name[0] != "" {
		serverName = gconv.String(name[0])
	}
	return serverMapping.GetOrSetFuncLock(serverName, func() interface{} {
		return NewServer("", nil)
	}).(*Server)
}

// NewServer 创建并返回一个新的普通TCP服务器。
// 参数`name`是可选的，用于指定服务器实例的名称。
func NewServer(address string, handler func(*Conn), name ...string) *Server {
	s := &Server{
		address: address,
		handler: handler,
	}
	if len(name) > 0 && name[0] != "" {
		serverMapping.Set(name[0], s)
	}
	return s
}

// NewServerTLS 创建并返回一个带有 TLS 支持的新 TCP 服务器。
// 参数 `name` 是可选的，用于指定服务器实例的名称。
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*Conn), name ...string) *Server {
	s := NewServer(address, handler, name...)
	s.SetTLSConfig(tlsConfig)
	return s
}

// NewServerKeyCrt 创建并返回一个带有 TLS 支持的新 TCP 服务器。
// 参数 `name` 是可选的，用于指定服务器实例名称。
func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*Conn), name ...string) (*Server, error) {
	s := NewServer(address, handler, name...)
	if err := s.SetTLSKeyCrt(crtFile, keyFile); err != nil {
		return nil, err
	}
	return s, nil
}

// SetAddress 设置服务器的监听地址。
func (s *Server) SetAddress(address string) {
	s.address = address
}

// GetAddress 获取服务器的监听地址。
func (s *Server) GetAddress() string {
	return s.address
}

// SetHandler 设置服务器的连接处理器。
func (s *Server) SetHandler(handler func(*Conn)) {
	s.handler = handler
}

// SetTLSKeyCrt 用于设置服务器TLS配置所需的证书和密钥文件。
func (s *Server) SetTLSKeyCrt(crtFile, keyFile string) error {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return err
	}
	s.tlsConfig = tlsConfig
	return nil
}

// SetTLSConfig 设置服务器的 TLS 配置。
func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {
	s.tlsConfig = tlsConfig
}

// Close 关闭监听器并关闭服务器。
func (s *Server) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listen == nil {
		return nil
	}
	return s.listen.Close()
}

// Run 开始运行 TCP 服务器。
func (s *Server) Run() (err error) {
	if s.handler == nil {
		err = gerror.NewCode(gcode.CodeMissingConfiguration, "start running failed: socket handler not defined")
		return
	}
	if s.tlsConfig != nil {
		// TLS Server
		s.mu.Lock()
		s.listen, err = tls.Listen("tcp", s.address, s.tlsConfig)
		s.mu.Unlock()
		if err != nil {
			err = gerror.Wrapf(err, `tls.Listen failed for address "%s"`, s.address)
			return
		}
	} else {
		// Normal Server
		var tcpAddr *net.TCPAddr
		if tcpAddr, err = net.ResolveTCPAddr("tcp", s.address); err != nil {
			err = gerror.Wrapf(err, `net.ResolveTCPAddr failed for address "%s"`, s.address)
			return err
		}
		s.mu.Lock()
		s.listen, err = net.ListenTCP("tcp", tcpAddr)
		s.mu.Unlock()
		if err != nil {
			err = gerror.Wrapf(err, `net.ListenTCP failed for address "%s"`, s.address)
			return err
		}
	}
	// Listening loop.
	for {
		var conn net.Conn
		if conn, err = s.listen.Accept(); err != nil {
			err = gerror.Wrapf(err, `Listener.Accept failed`)
			return err
		} else if conn != nil {
			go s.handler(NewConnByNetConn(conn))
		}
	}
}

// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
func (s *Server) GetListenedAddress() string {
	if !gstr.Contains(s.address, FreePortAddress) {
		return s.address
	}
	var (
		address      = s.address
		listenedPort = s.GetListenedPort()
	)
	address = gstr.Replace(address, FreePortAddress, fmt.Sprintf(`:%d`, listenedPort))
	return address
}

// GetListenedPort 获取并返回当前服务器正在监听的一个端口。
func (s *Server) GetListenedPort() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if ln := s.listen; ln != nil {
		return ln.Addr().(*net.TCPAddr).Port
	}
	return -1
}
