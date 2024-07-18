// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gudp

import (
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

// Server是UDP服务器。 md5:34c72ea6deda36f9
type Server struct {
	mu      sync.Mutex  // 用于 Server.listen 的并发安全。-- Go 语言的竞态条件检测会检查这个。 md5:d330fd21b35ec6b2
	conn    *Conn       // UDP服务器连接对象。 md5:eb4722aff16908ab
	address string      // UDP服务器监听地址。 md5:a7756994f6ef60d7
	handler func(*Conn) // 处理UDP连接的处理器。 md5:7ad03bfb1dfd96cd
}

var (
	// serverMapping 用于实例名称到其 UDP 服务器映射。 md5:f02a58894dbf986b
	serverMapping = gmap.NewStrAnyMap(true)
)

// GetServer 创建并返回一个给定名称的UDP服务器实例。 md5:c822bb20e355a198
// ff:
// name:
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

// NewServer 创建并返回一个UDP服务器。
// 可选参数`name`用于指定服务器的名称，该名称可以用于
// GetServer 函数来检索其实例。
// md5:752020b7ca7ce4b2
// ff:
// address:
// handler:
// name:
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

// SetAddress 设置UDP服务器的地址。 md5:7159be88401e01c8
// ff:
// s:
// address:
func (s *Server) SetAddress(address string) {
	s.address = address
}

// SetHandler 设置UDP服务器的连接处理器。 md5:734c7ee9adee69b0
// ff:
// s:
// handler:
func (s *Server) SetHandler(handler func(*Conn)) {
	s.handler = handler
}

// Close 关闭连接。
// 它将使服务器立即关闭。
// md5:251649bd57732e67
// ff:
// s:
// err:
func (s *Server) Close() (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err = s.conn.Close()
	if err != nil {
		err = gerror.Wrap(err, "connection failed")
	}
	return
}

// Run 开始监听UDP连接。 md5:582eb8bc9f8281c9
// ff:
// s:
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

// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
// ff:
// s:
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
// ff:
// s:
func (s *Server) GetListenedPort() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if ln := s.conn; ln != nil {
		return ln.LocalAddr().(*net.UDPAddr).Port
	}
	return -1
}
