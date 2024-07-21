// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtcp

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	// FreePortAddress 标记服务器使用随机的空闲端口进行监听。 md5:16e8ca0633c4a135
	FreePortAddress = ":0"
)

const (
	defaultServer = "default"
)

// Server 是一个 TCP 服务器。 md5:15abc757287261ea
type Server struct {
	mu        sync.Mutex   // 用于 Server.listen 的并发安全。-- Go 语言的竞态条件检测会检查这个。 md5:d330fd21b35ec6b2
	listen    net.Listener // TCP address listener.
	address   string       // 服务器监听地址。 md5:c8adda00f51a60d8
	handler   func(*Conn)  // Connection handler.
	tlsConfig *tls.Config  // TLS configuration.
}

// 用于单例目的，存储名称到服务器的映射。 md5:8e877c386766a97c
var serverMapping = gmap.NewStrAnyMap(true)

// GetServer 返回指定名称的 TCP 服务器，如果不存在，则返回一个新创建的默认名为 `name` 的 TCP 服务器。参数 `name` 用于指定 TCP 服务器的名称。
// md5:f6bb57410cf2ca98
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
// 参数 `name` 是可选的，用于指定服务器的实例名称。
// md5:ce4abdc7a25f75da
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

// NewServerTLS 创建并返回一个支持TLS的TCP服务器。
// 参数`name`是可选的，用于指定服务器的实例名称。
// md5:102d98ca307029b3
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*Conn), name ...string) *Server {
	s := NewServer(address, handler, name...)
	s.SetTLSConfig(tlsConfig)
	return s
}

// NewServerKeyCrt 创建并返回一个支持TLS的TCP服务器。
// 参数 `name` 是可选的，用于指定服务器的实例名称。
// md5:65a6856829628fe8
func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*Conn), name ...string) (*Server, error) {
	s := NewServer(address, handler, name...)
	if err := s.SetTLSKeyCrt(crtFile, keyFile); err != nil {
		return nil, err
	}
	return s, nil
}

// SetAddress 设置服务器的监听地址。 md5:35306d25b7cbc244
func (s *Server) SetAddress(address string) {
	s.address = address
}

// GetAddress 获取服务器的监听地址。 md5:6085c2f0086d87f9
func (s *Server) GetAddress() string {
	return s.address
}

// SetHandler 设置服务器的连接处理器。 md5:10bacdc88ff59cee
func (s *Server) SetHandler(handler func(*Conn)) {
	s.handler = handler
}

// SetTLSKeyCrt 设置服务器TLS配置的证书和密钥文件。 md5:dd19415f9056b27d
func (s *Server) SetTLSKeyCrt(crtFile, keyFile string) error {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return err
	}
	s.tlsConfig = tlsConfig
	return nil
}

// SetTLSConfig 设置服务器的TLS配置。 md5:02f67dcfad23906c
func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {
	s.tlsConfig = tlsConfig
}

// Close 方法关闭监听器并停止服务器。 md5:494fcac465675910
func (s *Server) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listen == nil {
		return nil
	}
	return s.listen.Close()
}

// Run 开始运行TCP服务器。 md5:b107bdcd45f1ccdc
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

// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
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

// GetListenedPort 获取并返回当前服务器监听的其中一个端口。 md5:98e33a51d8d8309c
func (s *Server) GetListenedPort() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if ln := s.listen; ln != nil {
		return ln.Addr().(*net.TCPAddr).Port
	}
	return -1
}
