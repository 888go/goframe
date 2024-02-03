// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfpool_test

import (
	"context"
	"os"
	"testing"
	"time"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfpool"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

// TestOpen 测试打开文件缓存
func TestOpen(t *testing.T) {
	testFile := start("TestOpen.txt")

	gtest.C(t, func(t *gtest.T) {
		f, err := gfpool.Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertEQ(err, nil)
		f.Close()

		f2, err1 := gfpool.Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertEQ(err1, nil)
		t.AssertEQ(f, f2)
		f2.Close()
	})

	stop(testFile)
}

// TestOpenErr 测试打开文件错误
func TestOpenErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testErrFile := "errorPath"
		_, err := gfpool.Open(testErrFile, os.O_RDWR, 0666)
		t.AssertNE(err, nil)

		// delete file error
		testFile := start("TestOpenDeleteErr.txt")
		pool := gfpool.New(testFile, os.O_RDWR, 0666)
		_, err1 := pool.File()
		t.AssertEQ(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		t.AssertNE(err1, nil)

		// 如果文件存在则以追加模式删除并重新创建
		testFile = start("TestOpenCreateErr.txt")
		pool = gfpool.New(testFile, os.O_CREATE, 0666)
		_, err1 = pool.File()
		t.AssertEQ(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		t.AssertEQ(err1, nil)

		// append mode delete file error
		testFile = start("TestOpenAppendErr.txt")
		pool = gfpool.New(testFile, os.O_APPEND, 0666)
		_, err1 = pool.File()
		t.AssertEQ(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		t.AssertNE(err1, nil)

		// "trunc"模式下删除文件错误
		testFile = start("TestOpenTruncErr.txt")
		pool = gfpool.New(testFile, os.O_TRUNC, 0666)
		_, err1 = pool.File()
		t.AssertEQ(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		t.AssertNE(err1, nil)
	})
}

// TestOpenExpire 测试打开文件缓存过期
func TestOpenExpire(t *testing.T) {
	testFile := start("TestOpenExpire.txt")

	gtest.C(t, func(t *gtest.T) {
		f, err := gfpool.Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666, 100*time.Millisecond)
		t.AssertEQ(err, nil)
		f.Close()

		time.Sleep(150 * time.Millisecond)
		f2, err1 := gfpool.Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666, 100*time.Millisecond)
		t.AssertEQ(err1, nil)
		//t.AssertNE(f, f2)
		f2.Close()
	})

	stop(testFile)
}

// TestNewPool 测试 gfpool 新建函数
func TestNewPool(t *testing.T) {
	testFile := start("TestNewPool.txt")

	gtest.C(t, func(t *gtest.T) {
		f, err := gfpool.Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertEQ(err, nil)
		f.Close()

		pool := gfpool.New(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		f2, err1 := pool.File()
		// pool not equal
		t.AssertEQ(err1, nil)
		//t.AssertNE(f, f2)
		f2.Close()

		pool.Close()
	})

	stop(testFile)
}

// test before
func start(name string) string {
	testFile := os.TempDir() + string(os.PathSeparator) + name
	if gfile.Exists(testFile) {
		gfile.Remove(testFile)
	}
	content := "123"
	gfile.PutContents(testFile, content)
	return testFile
}

// test after
func stop(testFile string) {
	if gfile.Exists(testFile) {
		err := gfile.Remove(testFile)
		if err != nil {
			glog.Error(context.TODO(), err)
		}
	}
}

func Test_ConcurrentOS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f2.Close()

		for i := 0; i < 100; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		for i := 0; i < 100; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2200)
	})

	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f2.Close()

		s1 := ""
		for i := 0; i < 1000; i++ {
			s1 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s1))
		t.AssertNil(err)

		s2 := ""
		for i := 0; i < 1000; i++ {
			s2 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s2))
		t.AssertNil(err)

		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
// 数据竞争
// gtest.C(t, func(t *gtest.T) {
// 	// 创建一个临时文件路径，使用当前时间戳（纳秒级）作为文件名，并在函数结束时删除该文件
// 	path := gfile.Temp(gtime.TimestampNanoStr())
// 	defer gfile.Remove(path)
//  
// 	// 打开文件 f1，设置读写、创建、清空和追加模式，权限为0666
// 	f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
// 	t.AssertNil(err) // 断言错误为空
// 	defer f1.Close() // 在函数结束时关闭文件 f1
// 	// 同样方式打开文件 f2
// 	f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
// 	t.AssertNil(err) // 断言错误为空
// 	defer f2.Close() // 在函数结束时关闭文件 f2
// 	// 创建一个 WaitGroup 和一个无缓冲的 channel
// 	wg := sync.WaitGroup{}
// 	ch := make(chan struct{})
// 	// 对于前1000次循环：
// 	for i := 0; i < 1000; i++ {
//  	wg.Add(1) // 增加 WaitGroup 的计数器
//  	go func() {
//  		defer wg.Done() // 当协程完成时减少 WaitGroup 计数器
//  		<-ch          // 阻塞等待 channel 发送信号
//  		_, err = f1.Write([]byte("@1234567890#")) // 将字符串写入文件 f1
//  		t.AssertNil(err)                          // 断言写入操作没有错误
//  	}()
// 	}
// 	// 对于后1000次循环，执行相同的操作，但向文件 f2 写入数据
// 	for i := 0; i < 1000; i++ {
//  	wg.Add(1)
//  	go func() {
//  		defer wg.Done()
//  		<-ch
//  		_, err = f2.Write([]byte("@1234567890#"))
//  		t.AssertNil(err)
//  	}()
// 	}
// 	close(ch) // 关闭 channel，触发所有阻塞的接收操作
// 	wg.Wait() // 等待所有协程完成
// 	// 检查临时文件中特定字符串出现的次数，期望值为2000次
// 	t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
// })
}

func Test_ConcurrentGFPool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f1.Close()

		f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.AssertNil(err)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.AssertNil(err)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
// 数据竞争
// gtest.C(t, func(t *gtest.T) {
//   // 创建一个临时文件路径，格式为当前时间戳（纳秒）
//   path := gfile.Temp(gtime.TimestampNanoStr())
//   // 在测试结束后删除该临时文件
//   defer gfile.Remove(path)
//   // 打开文件 f1，设置读写、创建、清空和追加模式，并赋予0666权限
//   f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
//   // 断言错误为空
//   t.AssertNil(err)
//   // 在函数结束时关闭文件 f1
//   defer f1.Close()
//
//   // 再次以相同方式打开同一个文件 f2
//   f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
//   // 断言错误为空
//   t.AssertNil(err)
//   // 在函数结束时关闭文件 f2
//   defer f2.Close()
//
//   // 创建一个 WaitGroup 实例 wg 用于同步并发操作
//   wg := sync.WaitGroup{}
//   // 创建一个无缓冲的 channel ch，用于控制并发执行
//   ch := make(chan struct{})
//   // 并发执行写入操作：启动1000个 goroutine，每个 goroutine 都向文件 f1 写入 "@1234567890#" 字符串
//   for i := 0; i < 1000; i++ {
//       wg.Add(1)
//       go func() {
//           defer wg.Done()
//           <-ch
//           _, err = f1.Write([]byte("@1234567890#"))
//           t.AssertNil(err)
//       }()
//   }
//   // 同样并发执行写入操作：启动1000个 goroutine，每个 goroutine 都向文件 f2 写入 "@1234567890#" 字符串
//   for i := 0; i < 1000; i++ {
//       wg.Add(1)
//       go func() {
//           defer wg.Done()
//           <-ch
//           _, err = f2.Write([]byte("@1234567890#"))
//           t.AssertNil(err)
//       }()
//   }
//   // 关闭 channel ch，使得所有等待的 goroutine 开始执行
//   close(ch)
//   // 等待所有并发任务完成
//   wg.Wait()
//   // 断言文件内容中 "@1234567890#" 字符串出现次数为 2000 次
//   t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
// })
}
