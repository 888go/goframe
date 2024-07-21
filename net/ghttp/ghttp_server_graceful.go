// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

// gracefulServer 是一个包装了 net/http.Server 的结构，添加了优雅的重新加载/重启功能。 md5:8d812c91a33cd2a2
type gracefulServer struct {
	server      *Server      // Belonged server.
	fd          uintptr      // 用于在优雅重启时传递给子进程的文件描述符。 md5:72ea9b448b106b41
	address     string       // 监听地址，例如":80"，":8080"。 md5:c746ec22043cf3e0
	httpServer  *http.Server // 底层的http.Server。 md5:3b44f2da7272f7f3
	rawListener net.Listener // 底层的net.Listener。 md5:95d2f6c4d9084a5b
	rawLnMu     sync.RWMutex // 为`rawListener`提供并发安全的互斥锁。 md5:7b358a2cf029baae
	listener    net.Listener // Wrapped net.Listener.
	isHttps     bool         // Is HTTPS.
	status      *gtype.Int   // 当前服务器的状态。使用 `gtype` 确保并发安全。 md5:d11344d5afa40f3a
}

// newGracefulServer 创建并返回一个给定地址的优雅HTTP服务器。
// 可选参数 `fd` 指定了从父服务器传递过来的文件描述符。
// md5:e7000c344ed0446f
func (s *Server) newGracefulServer(address string, fd ...int) *gracefulServer {
	// 将端口转换为地址形式，如：80 -> :80. md5:71e59572a00dec96
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

// newHttpServer 创建并返回一个带有给定地址的底层 http.Server。 md5:12a45a5b95a4e7c3
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

// Fd获取并返回当前服务器的文件描述符。它只在*nix类操作系统中可用，如Linux、Unix和Darwin。
// md5:40546fed24d791cd
// ff:
// s:
func (s *gracefulServer) Fd() uintptr {
	if ln := s.getRawListener(); ln != nil {
		file, err := ln.(*net.TCPListener).File()
		if err == nil {
			return file.Fd()
		}
	}
	return 0
}

// CreateListener 在配置的地址上创建监听器。 md5:89f8795cf6b796f9
// ff:
// s:
func (s *gracefulServer) CreateListener() error {
	ln, err := s.getNetListener()
	if err != nil {
		return err
	}
	s.listener = ln
	s.setRawListener(ln)
	return nil
}

// CreateListenerTLS 在配置的地址上创建使用HTTPS的监听器。
// 参数`certFile`和`keyFile`指定了HTTPS所需的证书和密钥文件。
// 可选参数`tlsConfig`指定自定义的TLS配置。
// md5:04f46f61853037ca
// ff:
// s:
// certFile:
// keyFile:
// tlsConfig:
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

// Serve以阻塞方式启动服务。 md5:230e5731ffa3d482
// ff:
// s:
// ctx:
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

// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
// ff:
// s:
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

// GetListenedPort 获取并返回当前服务器正在监听的其中一个端口。
// 注意，如果服务器只监听一个端口，则此方法才可用。
// md5:2fe5eae2317fe8f9
// ff:
// s:
func (s *gracefulServer) GetListenedPort() int {
	if ln := s.getRawListener(); ln != nil {
		return ln.Addr().(*net.TCPAddr).Port
	}
	return -1
}

// getProto 获取并返回当前服务器的proto字符串。 md5:7860227f594f2ca9
func (s *gracefulServer) getProto() string {
	proto := "http"
	if s.isHttps {
		proto = "https"
	}
	return proto
}

// getNetListener 获取并返回包装的net.Listener。 md5:36d0b8cf9a591408
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

// shutdown 停止服务器，优雅地关闭。 md5:6befce727da40eb9
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

// 设置RawListener，将给定的net.Listener设置为`rawListener`。 md5:0fe9b7938ed0a876
func (s *gracefulServer) setRawListener(ln net.Listener) {
	s.rawLnMu.Lock()
	defer s.rawLnMu.Unlock()
	s.rawListener = ln
}

// setRawListener 返回当前服务器的 `rawListener`。 md5:e7b9cd54708d26f8
func (s *gracefulServer) getRawListener() net.Listener {
	s.rawLnMu.RLock()
	defer s.rawLnMu.RUnlock()
	return s.rawListener
}

// close 强制关闭服务器。 md5:46634188c0dbdf78
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
