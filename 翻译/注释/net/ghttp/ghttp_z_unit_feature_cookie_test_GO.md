
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
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
// t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
// 中文注释：断言通过上下文ctx，使用client获取"/get?k=key2"的请求内容，并判断其响应状态码应为"200"
// t.Assert(client.GetContent(ctx, "/get?k=key3"), "")
// 中文注释：断言通过上下文ctx，使用client获取"/get?k=key3"的请求内容，并判断其响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key1"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key1"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key3"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key3"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/remove?k=key4"), "")
// 中文注释：断言通过上下文ctx，使用client发送"/remove?k=key4"的删除请求后，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/get?k=key1"), "")
// 中文注释：断言通过上下文ctx，使用client获取已删除的"/get?k=key1"的请求内容，判断响应内容为空字符串
// t.Assert(client.GetContent(ctx, "/get?k=key2"), "200")
// 中文注释：再次断言通过上下文ctx，使用client获取"/get?k=key2"的请求内容，并判断其响应状态码仍为"200"
// 上述代码是在进行HTTP接口测试，通过Assert方法验证请求和响应的结果是否符合预期。
# <翻译结束>


<原文开始>
//github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d
<原文结束>

# <翻译开始>
// 这是Go语言标准库GitHub仓库中的一次提交记录，对应的commit（提交）哈希值为542693e00529fbb4248fac614ece68b127a5ec4d。
// 由于没有提供具体的代码片段，这里无法给出详细的代码注释翻译。通常这种形式的引用用于指向Go语言源码在GitHub上的某一次具体提交，该提交可能包含对某个功能的修复、优化或新功能的添加等。若需要了解这次提交的具体内容，可前往GitHub对应仓库查看此次commit的详细信息和改动内容。
# <翻译结束>


<原文开始>
// For go < 1.16 cookie always output "SameSite", see: https://github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d
<原文结束>

# <翻译开始>
// 对于Go版本小于1.16的情况，cookie始终输出"SameSite"属性，参见：https://github.com/golang/go/commit/542693e00529fbb4248fac614ece68b127a5ec4d
// 这段注释是说明在Go语言版本1.16之前，处理HTTP cookie时，默认会始终添加"SameSite"属性。关于这个问题的更多详细信息，可以参考提供的GitHub链接。
# <翻译结束>

