// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gudp_test

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/gogf/gf/v2/net/gudp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	simpleTimeout = time.Millisecond * 100
	sendData      = []byte("hello")
)

func startUDPServer(addr string) *gudp.Server {
	s := gudp.NewServer(addr, func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if err != nil {
				if err != io.EOF {
					glog.Error(context.TODO(), err)
				}
				break
			}
			if err = conn.Send(data); err != nil {
				glog.Error(context.TODO(), err)
			}
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func Test_Basic(t *testing.T) {
	var ctx = context.TODO()
	s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					glog.Error(ctx, err)
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// gudp.Conn.Send
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(s.GetListenedAddress())
			t.AssertNil(err)
			t.Assert(conn.Send([]byte(gconv.String(i))), nil)
			t.AssertNil(conn.RemoteAddr())
			result, err := conn.Recv(-1)
			t.AssertNil(err)
			t.AssertNE(conn.RemoteAddr(), nil)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})
	// gudp.Conn.SendRecv
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(s.GetListenedAddress())
			t.AssertNil(err)
			_, err = conn.SendRecv([]byte(gconv.String(i)), -1)
			t.AssertNil(err)
			// 使用t.Assert断言result转换为字符串后，应该大于指定的整数i，并格式化输出`> %d`。 md5:a472ffa3da0404c5
			conn.Close()
		}
	})
	// gudp.Conn.SendWithTimeout. md5:5bddfc1c824abcc9
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(s.GetListenedAddress())
			t.AssertNil(err)
			err = conn.SendWithTimeout([]byte(gconv.String(i)), time.Second)
			t.AssertNil(err)
			conn.Close()
		}
	})
	// gudp.Conn.RecvWithTimeout. md5:230cfff3316a9a8e
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(s.GetListenedAddress())
			t.AssertNil(err)
			err = conn.Send([]byte(gconv.String(i)))
			t.AssertNil(err)
			conn.SetBufferWaitRecv(time.Millisecond * 100)
			result, err := conn.RecvWithTimeout(-1, time.Second)
			t.AssertNil(err)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})
	// gudp.Conn.SendRecvWithTimeout 是一个方法，用于在连接上同时发送和接收数据，并带有超时设置。 md5:61b16b2e37fcaedb
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(s.GetListenedAddress())
			t.AssertNil(err)
			result, err := conn.SendRecvWithTimeout([]byte(gconv.String(i)), -1, time.Second)
			t.AssertNil(err)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})
	// gudp.Send
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			err := gudp.Send(s.GetListenedAddress(), []byte(gconv.String(i)))
			t.AssertNil(err)
		}
	})
	// gudp.SendRecv
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			result, err := gudp.SendRecv(s.GetListenedAddress(), []byte(gconv.String(i)), -1)
			t.AssertNil(err)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
		}
	})
}

	// 如果接收缓冲区的大小小于发送的数据包大小，剩余的数据会被丢弃。
	// 测试函数：Buffer
	// 使用上下文管理器获取一个空闲的UDP服务器地址，并设置一个处理连接的回调函数：
	// 在回调函数中，循环接收数据，如果接收到数据，则发送回去。当接收到错误时，退出循环。
	// 启动服务器并等待一段时间后，进行以下测试：
	// 1. 发送字符串"123"到服务器地址，期望返回结果为"1"，并验证错误是否为nil。
	// 2. 再次发送字符串"456"到服务器地址，期望返回结果为"4"，同样验证错误是否为nil。
	// 
	// 最后确保关闭服务器。
	// md5:e1d59962b87c120a

func Test_NewConn(t *testing.T) {
	s := startUDPServer(gudp.FreePortAddress)

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", gudp.MustGetFreePort()))
		t.AssertNil(err)
		conn.SetDeadline(time.Now().Add(time.Second))
		t.Assert(conn.Send(sendData), nil)
		conn.Close()
	})

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", 99999))
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewConn(fmt.Sprintf("127.0.0.1:%d", 99999))
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})
}

func Test_GetFreePorts(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ports, err := gudp.GetFreePorts(2)
		t.AssertNil(err)
		t.AssertEQ(len(ports), 2)
	})
}

func Test_Server(t *testing.T) {
	gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(1)
			if len(data) > 0 {
				conn.Send(data)
			}
			if err != nil {
				break
			}
		}
	}, "GoFrameUDPServer")

	gtest.C(t, func(t *gtest.T) {
		server := gudp.GetServer("GoFrameUDPServer")
		t.AssertNE(server, nil)
		server = gudp.GetServer("TestUDPServer")
		t.AssertNE(server, nil)
		server.SetAddress("127.0.0.1:8888")
		server.SetHandler(func(conn *gudp.Conn) {
			defer conn.Close()
			for {
				conn.Send([]byte("OtherHandle"))
			}
		})
	})
}
