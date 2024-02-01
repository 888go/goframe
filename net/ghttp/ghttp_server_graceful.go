// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/text/gstr"
	)
// gracefulServer 将 net/http.Server 包装起来，为其提供优雅的重新加载/重启功能。
type gracefulServer struct {
	server      *Server      // Belonged server.
	fd          uintptr      // 用于在优雅重启时传递给子进程的文件描述符。
	address     string       // 监听地址格式如":80"、":8080"。
	httpServer  *http.Server // 底层的 http.Server.
	rawListener net.Listener // 基础的 net.Listener.
	rawLnMu     sync.RWMutex // `rawListener`的并发安全互斥锁。
	listener    net.Listener // 包装过的 net.Listener。
	isHttps     bool         // Is HTTPS.
	status      *gtype.Int   // 当前服务器状态，使用 `gtype` 以确保并发安全性。
}

// newGracefulServer 根据给定的地址创建并返回一个优雅的 HTTP 服务器。
// 可选参数 `fd` 指定了从父服务器传递过来的文件描述符。
func (s *Server) newGracefulServer(address string, fd ...int) *gracefulServer {
	// 将端口更改为地址形式，例如：80 -> :80
	if gstr.IsNumeric(address) {
		address = ":" + address
	}
	gs := &gracefulServer{
		server:     s,
		address:    address,
		httpServer: s.newHttpServer(address),
		status:     gtype.NewInt(),
	}
	if len(fd) > 0 && fd[0] > 0 {
		gs.fd = uintptr(fd[0])
	}
	if s.config.Listeners != nil {
		addrArray := gstr.SplitAndTrim(address, ":")
		addrPort, err := strconv.Atoi(addrArray[len(addrArray)-1])
		if err == nil {
			for _, v := range s.config.Listeners {
				if listenerPort := (v.Addr().(*net.TCPAddr)).Port; listenerPort == addrPort {
					gs.rawListener = v
					break
				}
			}
		}
	}
	return gs
}

// newHttpServer 根据给定的地址创建并返回一个底层的 http.Server。
func (s *Server) newHttpServer(address string) *http.Server {
	server := &http.Server{
		Addr:           address,
		Handler:        http.HandlerFunc(s.config.Handler),
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
		IdleTimeout:    s.config.IdleTimeout,
		MaxHeaderBytes: s.config.MaxHeaderBytes,
		ErrorLog:       log.New(&errorLogger{logger: s.config.Logger}, "", 0),
	}
	server.SetKeepAlivesEnabled(s.config.KeepAlive)
	return server
}

// Fd 获取并返回当前服务器的文件描述符。
// 该功能仅在类*nix操作系统中可用，如linux、unix、darwin。
func (s *gracefulServer) Fd() uintptr {
	if ln := s.getRawListener(); ln != nil {
		file, err := ln.(*net.TCPListener).File()
		if err == nil {
			return file.Fd()
		}
	}
	return 0
}

// CreateListener 在配置的地址上创建监听器。
func (s *gracefulServer) CreateListener() error {
	ln, err := s.getNetListener()
	if err != nil {
		return err
	}
	s.listener = ln
	s.setRawListener(ln)
	return nil
}

// CreateListenerTLS 在配置的地址上创建 HTTPS 侦听器。
// 参数 `certFile` 和 `keyFile` 指定用于 HTTPS 的必要证书和密钥文件。
// 可选参数 `tlsConfig` 指定自定义 TLS 配置。
func (s *gracefulServer) CreateListenerTLS(certFile, keyFile string, tlsConfig ...*tls.Config) error {
	var config *tls.Config
	if len(tlsConfig) > 0 && tlsConfig[0] != nil {
		config = tlsConfig[0]
	} else if s.httpServer.TLSConfig != nil {
		config = s.httpServer.TLSConfig
	} else {
		config = &tls.Config{}
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}
	var err error
	if len(config.Certificates) == 0 {
		config.Certificates = make([]tls.Certificate, 1)
		if gres.Contains(certFile) {
			config.Certificates[0], err = tls.X509KeyPair(
				gres.GetContent(certFile),
				gres.GetContent(keyFile),
			)
		} else {
			config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
		}
	}
	if err != nil {
		return gerror.Wrapf(err, `open certFile "%s" and keyFile "%s" failed`, certFile, keyFile)
	}
	ln, err := s.getNetListener()
	if err != nil {
		return err
	}

	s.listener = tls.NewListener(ln, config)
	s.setRawListener(ln)
	return nil
}

// Serve 以阻塞方式启动服务。
func (s *gracefulServer) Serve(ctx context.Context) error {
	if s.rawListener == nil {
		return gerror.NewCode(gcode.CodeInvalidOperation, `call CreateListener/CreateListenerTLS before Serve`)
	}

	action := "started"
	if s.fd != 0 {
		action = "reloaded"
	}
	s.server.Logger().Infof(
		ctx,
		`pid[%d]: %s server %s listening on [%s]`,
		gproc.Pid(), s.getProto(), action, s.GetListenedAddress(),
	)
	s.status.Set(ServerStatusRunning)
	err := s.httpServer.Serve(s.listener)
	s.status.Set(ServerStatusStopped)
	return err
}

// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
func (s *gracefulServer) GetListenedAddress() string {
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

// GetListenedPort 获取并返回当前服务器监听的端口。
// 注意：只有当服务器正在监听单个端口时，此方法才可用。
func (s *gracefulServer) GetListenedPort() int {
	if ln := s.getRawListener(); ln != nil {
		return ln.Addr().(*net.TCPAddr).Port
	}
	return -1
}

// getProto 获取并返回当前服务器的协议字符串。
func (s *gracefulServer) getProto() string {
	proto := "http"
	if s.isHttps {
		proto = "https"
	}
	return proto
}

// getNetListener 获取并返回封装后的 net.Listener。
func (s *gracefulServer) getNetListener() (net.Listener, error) {
	if s.rawListener != nil {
		return s.rawListener, nil
	}
	var (
		ln  net.Listener
		err error
	)
	if s.fd > 0 {
		f := os.NewFile(s.fd, "")
		ln, err = net.FileListener(f)
		if err != nil {
			err = gerror.Wrap(err, "net.FileListener failed")
			return nil, err
		}
	} else {
		ln, err = net.Listen("tcp", s.httpServer.Addr)
		if err != nil {
			err = gerror.Wrapf(err, `net.Listen address "%s" failed`, s.httpServer.Addr)
		}
	}
	return ln, err
}

// shutdown优雅地关闭服务器。
func (s *gracefulServer) shutdown(ctx context.Context) {
	if s.status.Val() == ServerStatusStopped {
		return
	}
	timeoutCtx, cancelFunc := context.WithTimeout(
		ctx,
		time.Duration(s.server.config.GracefulShutdownTimeout)*time.Second,
	)
	defer cancelFunc()
	if err := s.httpServer.Shutdown(timeoutCtx); err != nil {
		s.server.Logger().Errorf(
			ctx,
			"%d: %s server [%s] shutdown error: %v",
			gproc.Pid(), s.getProto(), s.address, err,
		)
	}
}

// setRawListener 将给定的 net.Listener 设置到 `rawListener`。
func (s *gracefulServer) setRawListener(ln net.Listener) {
	s.rawLnMu.Lock()
	defer s.rawLnMu.Unlock()
	s.rawListener = ln
}

// setRawListener 返回当前服务器的 `原始监听器`。
func (s *gracefulServer) getRawListener() net.Listener {
	s.rawLnMu.RLock()
	defer s.rawLnMu.RUnlock()
	return s.rawListener
}

// close 强制关闭服务器。
func (s *gracefulServer) close(ctx context.Context) {
	if s.status.Val() == ServerStatusStopped {
		return
	}
	if err := s.httpServer.Close(); err != nil {
		s.server.Logger().Errorf(
			ctx,
			"%d: %s server [%s] closed error: %v",
			gproc.Pid(), s.getProto(), s.address, err,
		)
	}
}
