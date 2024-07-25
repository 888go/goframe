
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
// It posts data along with file uploading.
// It does not url-encodes the parameters.
<原文结束>

# <翻译开始>
// 它在上传文件的同时发送数据。
// 它不会对参数进行URL编码。 md5:e7d22bb43988cf7d
# <翻译结束>


<原文开始>
// test abort, abort will not send
<原文结束>

# <翻译开始>
		// 测试中断，中断不会发送. md5:08b4e656520be948
# <翻译结束>


<原文开始>
	// No closing in case of DATA RACE due to keep alive connection of WebSocket.
	// defer s.Shutdown()
<原文结束>

# <翻译开始>
	// 注意：由于WebSocket保持活动连接，存在数据竞争风险，因此此处不使用关闭defer语句。
	// 	// s.Shutdown() 应在适当的地方手动调用以确保资源正确释放。 md5:e59a0c81d7768e8f
# <翻译结束>

