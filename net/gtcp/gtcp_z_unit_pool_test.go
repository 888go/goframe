// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtcp_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_Pool_Basic1(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
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
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.SendPkg(data)
		t.AssertNil(err)
		err = conn.SendPkgWithTimeout(data, time.Second)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		_, err := gtcp.NewPoolConn("127.0.0.1:80")
		t.AssertNE(err, nil)
	})
}

func Test_Pool_Basic2(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		conn.Close()
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.SendPkg(data)
		t.AssertNil(err)
// 使用1秒超时时间发送数据包，并将错误结果赋值给err变量
// 断言err为nil（即无错误）

		_, err = conn.SendRecv(data, -1)
		t.AssertNE(err, nil)
	})
}

func Test_Pool_Send(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.Send(data)
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Pool_Recv(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.Send(data)
		t.AssertNil(err)
		time.Sleep(100 * time.Millisecond)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Pool_RecvLine(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999\n")
		err = conn.Send(data)
		t.AssertNil(err)
		time.Sleep(100 * time.Millisecond)
		result, err := conn.RecvLine()
		t.AssertNil(err)
		splitData := gstr.Split(string(data), "\n")
		t.Assert(result, splitData[0])
	})
}

func Test_Pool_RecvTill(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999\n")
		err = conn.Send(data)
		t.AssertNil(err)
		time.Sleep(100 * time.Millisecond)
		result, err := conn.RecvTill([]byte("\n"))
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Pool_RecvWithTimeout(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.Send(data)
		t.AssertNil(err)
		time.Sleep(100 * time.Millisecond)
		result, err := conn.RecvWithTimeout(-1, time.Millisecond*500)
		t.AssertNil(err)
		t.Assert(data, result)
	})
}

func Test_Pool_SendWithTimeout(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		err = conn.SendWithTimeout(data, time.Millisecond*500)
		t.AssertNil(err)
		result, err := conn.Recv(-1)
		t.AssertNil(err)
		t.Assert(data, result)
	})
}

func Test_Pool_SendRecvWithTimeout(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("9999")
		result, err := conn.SendRecvWithTimeout(data, -1, time.Millisecond*500)
		t.AssertNil(err)
		t.Assert(data, result)
	})
}
