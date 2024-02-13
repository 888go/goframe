// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类_test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

var (
	simpleTimeout = time.Millisecond * 100
	sendData      = []byte("hello")
	invalidAddr   = "127.0.0.1:99999"
	crtFile       = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.crt"
	keyFile       = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.key"
)

func startTCPServer(addr string) *tcp类.Server {
	s := tcp类.NewServer(addr, func(conn *tcp类.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func startTCPPkgServer(addr string) *tcp类.Server {
	s := tcp类.NewServer(addr, func(conn *tcp类.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				break
			}
			conn.SendPkg(data)
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func startTCPTLSServer(addr string) *tcp类.Server {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates: []tls.Certificate{
			{},
		},
	}
	s := tcp类.NewServerTLS(addr, tlsConfig, func(conn *tcp类.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func startTCPKeyCrtServer(addr string) *tcp类.Server {
	s, _ := tcp类.NewServerKeyCrt(addr, crtFile, keyFile, func(conn *tcp类.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func TestGetFreePorts(t *testing.T) {
	ports, _ := tcp类.GetFreePorts(2)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT(ports[0], 0)
		t.AssertGT(ports[1], 0)
	})

	startTCPServer(fmt.Sprintf("%s:%d", "127.0.0.1", ports[0]))

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", ports[0]))
		t.AssertNil(err)
		defer conn.Close()
		result, err := conn.SendRecv(sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", 80))
		t.AssertNE(err, nil)
		t.AssertNil(conn)
	})
}

func TestMustGetFreePort(t *testing.T) {
	port := tcp类.MustGetFreePort()
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	startTCPServer(addr)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := tcp类.SendRecv(addr, sendData, -1)
		t.AssertNil(err)
		t.Assert(sendData, result)
	})
}

func TestNewConn(t *testing.T) {
	addr := tcp类.FreePortAddress

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(addr, simpleTimeout)
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := startTCPServer(tcp类.FreePortAddress)

		conn, err := tcp类.NewConn(s.GetListenedAddress(), simpleTimeout)
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		defer conn.Close()
		result, err := conn.SendRecv(sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

// TODO
func TestNewConnTLS(t *testing.T) {
	addr := tcp类.FreePortAddress

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConnTLS(addr, &tls.Config{})
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := startTCPTLSServer(addr)

		conn, err := tcp类.NewConnTLS(s.GetListenedAddress(), &tls.Config{
			InsecureSkipVerify: true,
			Certificates: []tls.Certificate{
				{},
			},
		})
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})
}

func TestNewConnKeyCrt(t *testing.T) {
	addr := tcp类.FreePortAddress

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConnKeyCrt(addr, crtFile, keyFile)
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := startTCPKeyCrtServer(addr)

		conn, err := tcp类.NewConnKeyCrt(s.GetListenedAddress(), crtFile, keyFile)
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})
}

func TestConn_Send(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		err = conn.Send(sendData, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_SendWithTimeout(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		err = conn.SendWithTimeout(sendData, time.Second, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_SendRecv(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		result, err := conn.SendRecv(sendData, -1, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_SendRecvWithTimeout(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		result, err := conn.SendRecvWithTimeout(sendData, -1, time.Second, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_RecvWithTimeout(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		conn.Send(sendData)
		result, err := conn.RecvWithTimeout(-1, time.Second, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_RecvLine(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		data := []byte("hello\n")
		conn.Send(data)
		result, err := conn.RecvLine(tcp类.Retry{Count: 1})
		t.AssertNil(err)
		splitData := 文本类.X分割(string(data), "\n")
		t.Assert(result, splitData[0])
	})
}

func TestConn_RecvTill(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		conn.Send(sendData)
		result, err := conn.RecvTill([]byte("hello"), tcp类.Retry{Count: 1})
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_SetDeadline(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		conn.SetDeadline(time.Time{})
		err = conn.Send(sendData, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestConn_SetReceiveBufferWait(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		t.AssertNE(conn, nil)
		conn.SetBufferWaitRecv(time.Millisecond * 100)
		err = conn.Send(sendData, tcp类.Retry{Count: 1})
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestNewNetConnKeyCrt(t *testing.T) {
	addr := tcp类.FreePortAddress

	startTCPKeyCrtServer(addr)

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewNetConnKeyCrt(addr, "crtFile", keyFile, time.Second)
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewNetConnKeyCrt(addr, crtFile, keyFile, time.Second)
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})
}

func TestSend(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.Send(invalidAddr, sendData, tcp类.Retry{Count: 1})
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.Send(s.GetListenedAddress(), sendData, tcp类.Retry{Count: 1})
		t.AssertNil(err)
	})
}

func TestSendRecv(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := tcp类.SendRecv(invalidAddr, sendData, -1)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestSendWithTimeout(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendWithTimeout(invalidAddr, sendData, time.Millisecond*500)
		t.AssertNE(err, nil)
		err = tcp类.SendWithTimeout(s.GetListenedAddress(), sendData, time.Millisecond*500)
		t.AssertNil(err)
	})
}

func TestSendRecvWithTimeout(t *testing.T) {
	s := startTCPServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		result, err := tcp类.SendRecvWithTimeout(invalidAddr, sendData, -1, time.Millisecond*500)
		t.AssertNil(result)
		t.AssertNE(err, nil)
		result, err = tcp类.SendRecvWithTimeout(s.GetListenedAddress(), sendData, -1, time.Millisecond*500)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestSendPkg(t *testing.T) {
	s := startTCPPkgServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		err = tcp类.SendPkg(invalidAddr, sendData)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData, tcp类.PkgOption{Retry: tcp类.Retry{Count: 3}})
		t.AssertNil(err)
		err = tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
	})
}

func TestSendRecvPkg(t *testing.T) {
	s := startTCPPkgServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		_, err = tcp类.SendRecvPkg(invalidAddr, sendData)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		result, err := tcp类.SendRecvPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestSendPkgWithTimeout(t *testing.T) {
	s := startTCPPkgServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		err = tcp类.SendPkgWithTimeout(invalidAddr, sendData, time.Second)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		err = tcp类.SendPkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
		t.AssertNil(err)
	})
}

func TestSendRecvPkgWithTimeout(t *testing.T) {
	s := startTCPPkgServer(tcp类.FreePortAddress)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		_, err = tcp类.SendRecvPkgWithTimeout(invalidAddr, sendData, time.Second)
		t.AssertNE(err, nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := tcp类.SendPkg(s.GetListenedAddress(), sendData)
		t.AssertNil(err)
		result, err := tcp类.SendRecvPkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestNewServer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
			defer conn.Close()
			for {
				data, err := conn.Recv(-1)
				if err != nil {
					break
				}
				conn.Send(data)
			}
		}, "NewServer")
		defer s.Close()
		go s.Run()

		time.Sleep(simpleTimeout)

		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestGetServer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.GetServer("GetServer")
		defer s.Close()
		go s.Run()

		t.Assert(s.GetAddress(), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
			defer conn.Close()
			for {
				data, err := conn.Recv(-1)
				if err != nil {
					break
				}
				conn.Send(data)
			}
		}, "NewServer")

		s := tcp类.GetServer("NewServer")
		defer s.Close()
		go s.Run()

		time.Sleep(simpleTimeout)

		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestServer_SetAddress(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.NewServer("", func(conn *tcp类.Conn) {
			defer conn.Close()
			for {
				data, err := conn.Recv(-1)
				if err != nil {
					break
				}
				conn.Send(data)
			}
		})
		defer s.Close()
		t.Assert(s.GetAddress(), "")
		s.SetAddress(tcp类.FreePortAddress)
		go s.Run()

		time.Sleep(simpleTimeout)

		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestServer_SetHandler(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.NewServer(tcp类.FreePortAddress, nil)
		defer s.Close()
		s.SetHandler(func(conn *tcp类.Conn) {
			defer conn.Close()
			for {
				data, err := conn.Recv(-1)
				if err != nil {
					break
				}
				conn.Send(data)
			}
		})
		go s.Run()

		time.Sleep(simpleTimeout)

		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})
}

func TestServer_Run(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
			defer conn.Close()
			for {
				data, err := conn.Recv(-1)
				if err != nil {
					break
				}
				conn.Send(data)
			}
		})
		defer s.Close()
		go s.Run()

		time.Sleep(simpleTimeout)

		result, err := tcp类.SendRecv(s.GetListenedAddress(), sendData, -1)
		t.AssertNil(err)
		t.Assert(result, sendData)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := tcp类.NewServer(tcp类.FreePortAddress, nil)
		defer s.Close()
		go func() {
			err := s.Run()
			t.AssertNE(err, nil)
		}()
	})
}
