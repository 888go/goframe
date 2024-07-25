
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
//t.Assert(string(result), fmt.Sprintf(`> %d`, i))
<原文结束>

# <翻译开始>
			// 使用t.Assert断言result转换为字符串后，应该大于指定的整数i，并格式化输出`> %d`。 md5:a472ffa3da0404c5
# <翻译结束>


<原文开始>
// gudp.Conn.SendWithTimeout
<原文结束>

# <翻译开始>
	// gudp.Conn.SendWithTimeout. md5:5bddfc1c824abcc9
# <翻译结束>


<原文开始>
// gudp.Conn.RecvWithTimeout
<原文结束>

# <翻译开始>
	// gudp.Conn.RecvWithTimeout. md5:230cfff3316a9a8e
# <翻译结束>


<原文开始>
// gudp.Conn.SendRecvWithTimeout
<原文结束>

# <翻译开始>
	// gudp.Conn.SendRecvWithTimeout 是一个方法，用于在连接上同时发送和接收数据，并带有超时设置。 md5:61b16b2e37fcaedb
# <翻译结束>


<原文开始>
// If the read buffer size is less than the sent package size,
// the rest data would be dropped.
//func Test_Buffer(t *testing.T) {
//	var ctx = context.TODO()
//	s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.Conn) {
//		defer conn.Close()
//		for {
//			data, err := conn.Recv(1)
//			if len(data) > 0 {
//				if err := conn.Send(data); err != nil {
//					glog.Error(ctx, err)
//				}
//			}
//			if err != nil {
//				break
//			}
//		}
//	})
//	go s.Run()
//	defer s.Close()
//	time.Sleep(100 * time.Millisecond)
//	gtest.C(t, func(t *gtest.T) {
//		result, err := gudp.SendRecv(s.GetListenedAddress(), []byte("123"), -1)
//		t.AssertNil(err)
//		t.Assert(string(result), "1")
//	})
//	gtest.C(t, func(t *gtest.T) {
//		result, err := gudp.SendRecv(s.GetListenedAddress(), []byte("456"), -1)
//		t.AssertNil(err)
//		t.Assert(string(result), "4")
//	})
//}
<原文结束>

# <翻译开始>
// 如果接收缓冲区的大小小于发送的数据包大小，剩余的数据会被丢弃。
// 测试函数：Buffer
// 使用上下文管理器获取一个空闲的UDP服务器地址，并设置一个处理连接的回调函数：
// 在回调函数中，循环接收数据，如果接收到数据，则发送回去。当接收到错误时，退出循环。
// 启动服务器并等待一段时间后，进行以下测试：
// 1. 发送字符串"123"到服务器地址，期望返回结果为"1"，并验证错误是否为nil。
// 2. 再次发送字符串"456"到服务器地址，期望返回结果为"4"，同样验证错误是否为nil。
//
// 最后确保关闭服务器。 md5:e1d59962b87c120a
# <翻译结束>

