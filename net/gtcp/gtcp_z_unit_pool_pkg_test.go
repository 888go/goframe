// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Pool_Package_Basic(t *testing.T) {
	s := tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
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
	// SendPkg
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		for i := 0; i < 100; i++ {
			err := conn.SendPkg([]byte(转换类.String(i)))
			t.AssertNil(err)
		}
		for i := 0; i < 100; i++ {
			err := conn.SendPkgWithTimeout([]byte(转换类.String(i)), time.Second)
			t.AssertNil(err)
		}
	})
	// SendPkg 传输大数据 - 失败
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65536)
		err = conn.SendPkg(data)
		t.AssertNE(err, nil)
	})
	// SendRecvPkg
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		for i := 100; i < 200; i++ {
			data := []byte(转换类.String(i))
			result, err := conn.SendRecvPkg(data)
			t.AssertNil(err)
			t.Assert(result, data)
		}
		for i := 100; i < 200; i++ {
			data := []byte(转换类.String(i))
			result, err := conn.SendRecvPkgWithTimeout(data, time.Second)
			t.AssertNil(err)
			t.Assert(result, data)
		}
	})
	// SendRecvPkg 在大数据量下的发送与接收 - 失败情况
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65536)
		result, err := conn.SendRecvPkg(data)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg 处理大数据 - 成功。
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65500)
		data[100] = byte(65)
		data[65400] = byte(85)
		result, err := conn.SendRecvPkg(data)
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Pool_Package_Timeout(t *testing.T) {
	s := tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				break
			}
			time.Sleep(time.Second)
			单元测试类.Assert(conn.SendPkg(data), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Millisecond*500)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second*2)
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Pool_Package_Option(t *testing.T) {
	s := tcp类.NewServer(tcp类.FreePortAddress, func(conn *tcp类.Conn) {
		defer conn.Close()
		option := tcp类.PkgOption{HeaderSize: 1}
		for {
			data, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
			单元测试类.Assert(conn.SendPkg(data, option), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// SendRecvPkg 在大数据量下的发送与接收 - 失败情况
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF+1)
		result, err := conn.SendRecvPkg(data, tcp类.PkgOption{HeaderSize: 1})
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg 处理大数据 - 成功。
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkg(data, tcp类.PkgOption{HeaderSize: 1})
		t.AssertNil(err)
		t.Assert(result, data)
	})
	// SendRecvPkgWithTimeout 处理大数据 - 发送接收数据包并设置超时 - 失败情况
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF+1)
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second, tcp类.PkgOption{HeaderSize: 1})
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkgWithTimeout 在大数据量下 - 成功
// 此函数用于在设置超时的情况下，发送和接收大数据包，并且在此测试或示例中已成功完成。
	单元测试类.C(t, func(t *单元测试类.T) {
		conn, err := tcp类.NewPoolConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second, tcp类.PkgOption{HeaderSize: 1})
		t.AssertNil(err)
		t.Assert(result, data)
	})
}
