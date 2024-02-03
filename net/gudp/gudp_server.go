// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gudp

import (
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

// Server 是 UDP 服务器。
type Server struct {
	mu      sync.Mutex  // 用于保证Server.listen方法的并发安全性。-- 该Go语言测试会通过数据竞争检查进行验证。
	conn    *Conn       // UDP服务器连接对象。
	address string      // UDP服务器监听地址。
	handler func(*Conn) // 处理UDP连接的处理器。
}

var (
	// serverMapping 用于实例名与其UDP服务器之间的映射。
	serverMapping = gmap.NewStrAnyMap(true)
)

// GetServer 根据给定名称创建并返回一个UDP服务器实例。
func GetServer(name ...interface{}) *Server {
	serverName := defaultServer
	if len(name) > 0 && name[0] != "" {
		serverName = gconv.String(name[0])
	}
	if s := serverMapping.Get(serverName); s != nil {
		return s.(*Server)
	}
	s := NewServer("", nil)
	serverMapping.Set(serverName, s)
	return s
}

// NewServer 创建并返回一个 UDP 服务器。
// 可选参数 `name` 用于指定其名称，可用于 GetServer 函数来获取其实例。
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

// SetAddress 设置 UDP 服务器的服务器地址。
func (s *Server) SetAddress(address string) {
	s.address = address
}

// SetHandler 设置 UDP 服务器的连接处理器。
func (s *Server) SetHandler(handler func(*Conn)) {
	s.handler = handler
}

// Close 关闭连接。
// 它将使服务器立即关闭。
func (s *Server) Close() (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err = s.conn.Close()
	if err != nil {
		err = gerror.Wrap(err, "connection failed")
	}
	return
}

// Run 开始监听 UDP 连接。
func (s *Server) Run() error {
	if s.handler == nil {
		err := gerror.NewCode(gcode.CodeMissingConfiguration, "start running failed: socket handler not defined")
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", s.address)
	if err != nil {
		err = gerror.Wrapf(err, `net.ResolveUDPAddr failed for address "%s"`, s.address)
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		err = gerror.Wrapf(err, `net.ListenUDP failed for address "%s"`, s.address)
		return err
	}
	s.mu.Lock()
	s.conn = NewConnByNetConn(conn)
	s.mu.Unlock()
	s.handler(s.conn)
	return nil
}

// GetListenedAddress 获取并返回当前服务器正在监听的地址字符串。
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
	if ln := s.conn; ln != nil {
		return ln.LocalAddr().(*net.UDPAddr).Port
	}
	return -1
}
