// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtcp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Package_Basic(t *testing.T) {
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
	// SendPkg
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		for i := 0; i < 100; i++ {
			err := conn.SendPkg([]byte(gconv.String(i)))
			t.AssertNil(err)
		}
		for i := 0; i < 100; i++ {
			err := conn.SendPkgWithTimeout([]byte(gconv.String(i)), time.Second)
			t.AssertNil(err)
		}
	})
	// 发送大数据包 - 失败。 md5:363e253a410ac6a7
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65536)
		err = conn.SendPkg(data)
		t.AssertNE(err, nil)
	})
	// SendRecvPkg
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		for i := 100; i < 200; i++ {
			data := []byte(gconv.String(i))
			result, err := conn.SendRecvPkg(data)
			t.AssertNil(err)
			t.Assert(result, data)
		}
		for i := 100; i < 200; i++ {
			data := []byte(gconv.String(i))
			result, err := conn.SendRecvPkgWithTimeout(data, time.Second)
			t.AssertNil(err)
			t.Assert(result, data)
		}
	})
	// 使用大数据发送和接收数据包 - 失败。 md5:cd8900eaf76d0b6c
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65536)
		result, err := conn.SendRecvPkg(data)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg 处理大数据 - 成功。 md5:c826a3a68cb84d7b
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
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

func Test_Package_Basic_HeaderSize1(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg(gtcp.PkgOption{HeaderSize: 1})
			if err != nil {
				break
			}
			conn.SendPkg(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)

	// 发送接收空数据的包。 md5:597b0bab109c83d5
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0)
		result, err := conn.SendRecvPkg(data)
		t.AssertNil(err)
		t.AssertNil(result)
	})
}

func Test_Package_Timeout(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				break
			}
			time.Sleep(time.Second)
			gtest.Assert(conn.SendPkg(data), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Millisecond*500)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second*2)
		t.AssertNil(err)
		t.Assert(result, data)
	})
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second*2, gtcp.PkgOption{HeaderSize: 5})
		t.AssertNE(err, nil)
		t.AssertNil(result)
	})
}

func Test_Package_Option(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		option := gtcp.PkgOption{HeaderSize: 1}
		for {
			data, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
			gtest.Assert(conn.SendPkg(data, option), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// 使用大数据发送和接收数据包 - 失败。 md5:cd8900eaf76d0b6c
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF+1)
		result, err := conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 1})
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg 处理大数据 - 成功。 md5:c826a3a68cb84d7b
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 1})
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Package_Option_HeadSize3(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		option := gtcp.PkgOption{HeaderSize: 3}
		for {
			data, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
			gtest.Assert(conn.SendPkg(data, option), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 3})
		t.AssertNil(err)
		t.Assert(result, data)
	})
}

func Test_Package_Option_HeadSize4(t *testing.T) {
	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		option := gtcp.PkgOption{HeaderSize: 4}
		for {
			data, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
			gtest.Assert(conn.SendPkg(data, option), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// 使用大数据发送和接收数据包 - 失败。 md5:cd8900eaf76d0b6c
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFFFF+1)
		_, err = conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 4})
		t.AssertNil(err)
	})
	// SendRecvPkg 处理大数据 - 成功。 md5:c826a3a68cb84d7b
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 4})
		t.AssertNil(err)
		t.Assert(result, data)
	})
				// pkgOption.HeaderSize 头部溢出. md5:3ffe7aeda48c1a82
	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		_, err = conn.SendRecvPkg(data, gtcp.PkgOption{HeaderSize: 5})
		t.AssertNE(err, nil)
	})
}

func Test_Server_NewServerKeyCrt(t *testing.T) {
	var (
		noCrtFile = "noCrtFile"
		noKeyFile = "noKeyFile"
		crtFile   = gfile.Dir(gdebug.CallerFilePath()) + gfile.Separator + "testdata/crtFile"
		keyFile   = gfile.Dir(gdebug.CallerFilePath()) + gfile.Separator + "testdata/keyFile"
	)
	gtest.C(t, func(t *gtest.T) {
		addr := "127.0.0.1:%d"
		freePort, _ := gtcp.GetFreePort()
		addr = fmt.Sprintf(addr, freePort)
		s, err := gtcp.NewServerKeyCrt(addr, noCrtFile, noKeyFile, func(conn *gtcp.Conn) {
		})
		if err != nil {
			t.AssertNil(s)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		addr := "127.0.0.1:%d"
		freePort, _ := gtcp.GetFreePort()
		addr = fmt.Sprintf(addr, freePort)
		s, err := gtcp.NewServerKeyCrt(addr, crtFile, noKeyFile, func(conn *gtcp.Conn) {
		})
		if err != nil {
			t.AssertNil(s)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		addr := "127.0.0.1:%d"
		freePort, _ := gtcp.GetFreePort()
		addr = fmt.Sprintf(addr, freePort)
		s, err := gtcp.NewServerKeyCrt(addr, crtFile, keyFile, func(conn *gtcp.Conn) {
		})
		if err != nil {
			t.AssertNil(s)
		}
	})
}

func Test_Conn_RecvPkgError(t *testing.T) {

	s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
		defer conn.Close()
		option := gtcp.PkgOption{HeaderSize: 5}
		for {
			_, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		conn, err := gtcp.NewConn(s.GetListenedAddress())
		t.AssertNil(err)
		defer conn.Close()
		data := make([]byte, 65536)
		result, err := conn.SendRecvPkg(data)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
}
