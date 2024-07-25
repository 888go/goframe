// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfpool_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfpool"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

// TestOpen 测试打开文件缓存. md5:81eb9fe8e499fb8d
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

// TestOpenErr 测试打开文件错误. md5:b4c462ff3925e6a9
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

		// 追加模式下删除文件错误并重新创建. md5:8fe557a65dc84332
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

		// 剪切模式下删除文件出错. md5:7f5ba36c787b7f38
		testFile = start("TestOpenTruncErr.txt")
		pool = gfpool.New(testFile, os.O_TRUNC, 0666)
		_, err1 = pool.File()
		t.AssertEQ(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		t.AssertNE(err1, nil)
	})
}

// TestOpenExpire 测试打开文件缓存过期. md5:b650603dfb7db830
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

// TestNewPool 测试 gfpool 新建功能. md5:8e67fa187c59e60d
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
	// ```go
	// 数据竞争
	// gtest.C(t, func(t *gtest.T) {
	// 临时文件路径 := gfile.Temp(gtime.TimestampNanoStr())
	// defer 删除文件(临时文件路径)
	// f1, 错误 := os.OpenFile(临时文件路径, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	// t 断言错误为 nil
	// defer 关闭 f1
	//
	// f2, 错误 := os.OpenFile(临时文件路径, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	// t 断言错误为 nil
	// defer 关闭 f2
	//
	// wg := sync.WaitGroup{}
	// ch := make(chan struct{})
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		<-ch
	// 		_, 写入错误 = f1.Write([]byte("@1234567890#"))
	// 		t 断言错误为 nil
	// 	}()
	// }
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		<-ch
	// 		_, 写入错误 = f2.Write([]byte("@1234567890#"))
	// 		t 断言错误为 nil
	// 	}()
	// }
	// close(ch)
	// wg.Wait()
	// t 断言(gstr.Count(gfile.GetContents(临时文件路径), "@1234567890#"), 2000)
	// })
	// ```
	//
	// 这段Go代码的注释描述了一个数据竞争的例子。它创建了两个文件句柄`f1`和`f2`，并尝试并发地向同一个文件写入内容，通过1000个goroutine交替使用这两个句柄。在所有写入完成后，检查文件内容中特定字符串`@1234567890#`出现的次数是否为2000次。由于未进行同步控制，这可能导致数据竞争问题。 md5:0ab85d1fb1789860
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
	// gtest.C(t, 函数(t *gtest.T)) {
	// 创建一个临时文件，文件名包含当前时间戳
	// path := gfile.Temp(gtime.TimestampNanoStr())
	// 延迟删除临时文件
	// f1, 错误 := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	// t.AssertNil(错误) 	// 断言打开文件没有错误
	// defer f1.Close()
	//
	// 同样方式打开第二个文件
	// f2, 错误 := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	// t.AssertNil(错误)
	// defer f2.Close()
	//
	// 使用同步等待组和通道处理并发写入
	// 定义等待组和通道
	// wg := sync.WaitGroup{}
	// ch := make(chan struct{})
	//
	// 循环1000次，每个循环添加一个协程
	// for i := 0; i < 1000; i++ {
	//     wg.Add(1)
	//     go func() {
	//         defer wg.Done()
	//         <-ch
	//         _, 错误 = f1.Write([]byte("@1234567890#"))
	//         t.AssertNil(错误)
	//     }()
	// }
	//
	// 同样循环1000次，添加写入到f2的协程
	// for i := 0; i < 1000; i++ {
	//     wg.Add(1)
	//     go func() {
	//         defer wg.Done()
	//         <-ch
	//         _, 错误 = f2.Write([]byte("@1234567890#"))
	//         t.AssertNil(错误)
	//     }()
	// }
	//
	// 关闭通道
	// close(ch)
	//
	// 等待所有协程完成
	// wg.Wait()
	//
	// 断言文件内容中 "@1234567890#" 的数量为2000
	// t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	// } md5:c0dad68a7d55185f
}
