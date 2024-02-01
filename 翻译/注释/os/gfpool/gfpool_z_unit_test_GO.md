
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// TestOpen test open file cache
<原文结束>

# <翻译开始>
// TestOpen 测试打开文件缓存
# <翻译结束>


<原文开始>
// TestOpenErr test open file error
<原文结束>

# <翻译开始>
// TestOpenErr 测试打开文件错误
# <翻译结束>


<原文开始>
// append mode delete file error and create again
<原文结束>

# <翻译开始>
// 如果文件存在则以追加模式删除并重新创建
# <翻译结束>


<原文开始>
// trunc mode delete file error
<原文结束>

# <翻译开始>
// "trunc"模式下删除文件错误
# <翻译结束>


<原文开始>
// TestOpenExpire test open file cache expire
<原文结束>

# <翻译开始>
// TestOpenExpire 测试打开文件缓存过期
# <翻译结束>


<原文开始>
// TestNewPool test gfpool new function
<原文结束>

# <翻译开始>
// TestNewPool 测试 gfpool 新建函数
# <翻译结束>


<原文开始>
	// DATA RACE
	// gtest.C(t, func(t *gtest.T) {
	//	path := gfile.Temp(gtime.TimestampNanoStr())
	//	defer gfile.Remove(path)
	//	f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.AssertNil(err)
	//	defer f1.Close()
	//
	//	f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.AssertNil(err)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.AssertNil(err)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.AssertNil(err)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	// })
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
	// DATA RACE
	// gtest.C(t, func(t *gtest.T) {
	//	path := gfile.Temp(gtime.TimestampNanoStr())
	//	defer gfile.Remove(path)
	//	f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.AssertNil(err)
	//	defer f1.Close()
	//
	//	f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.AssertNil(err)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.AssertNil(err)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.AssertNil(err)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	// })
<原文结束>

# <翻译开始>
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
# <翻译结束>

