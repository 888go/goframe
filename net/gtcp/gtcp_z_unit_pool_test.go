// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package tcp类_test

import (
	"testing"
	"time"

	gtcp "github.com/888go/goframe/net/gtcp"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
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
		// 使用超时时间1秒发送数据包到conn
		// 使用断言检查错误是否为nil
		// md5:9d972f2dbbbeb674

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
		splitData := gstr.X分割(string(data), "\n")
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
