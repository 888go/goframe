
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
		//t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
		//t.Assert(client.GetContent(ctx, "/get?k=key3"), "")
		//t.Assert(client.GetContent(ctx, "/remove?k=key1"), "")
		//t.Assert(client.GetContent(ctx, "/remove?k=key3"), "")
		//t.Assert(client.GetContent(ctx, "/remove?k=key4"), "")
		//t.Assert(client.GetContent(ctx, "/get?k=key1"), "")
		//t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
<原文结束>

# <翻译开始>
		// 测试断言：获取"/get?k=key2"的响应状态为200
		// 测试断言：获取"/get?k=key3"的响应为空字符串
		// 测试断言：获取"/remove?k=key1"的响应为空字符串
		// 测试断言：获取"/remove?k=key3"的响应为空字符串
		// 测试断言：获取"/remove?k=key4"的响应为空字符串
		// 测试断言：获取"/get?k=key1"的响应为空字符串
		// 测试断言：获取"/get?k=key2"的响应状态为200
		// md5:fa4c58c1c55bab25
# <翻译结束>


<原文开始>
// For go < 1.16 cookie always output "SameSite", see: https://github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d
<原文结束>

# <翻译开始>
// 对于 Go < 1.16 版本，Cookie 始终会输出 "SameSite"，详情请参阅：https://github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d. md5:6f46e9f2afc803ab
# <翻译结束>

