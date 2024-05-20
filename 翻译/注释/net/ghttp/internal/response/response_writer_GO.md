
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
// Writer wraps http.ResponseWriter for extra features.
<原文结束>

# <翻译开始>
// Writer 在 http.ResponseWriter 上添加了额外功能。. md5:204ac8c0cb436351
# <翻译结束>


<原文开始>
// The underlying ResponseWriter.
<原文结束>

# <翻译开始>
// 基础的ResponseWriter。. md5:1678e6fb48b792ff
# <翻译结束>


<原文开始>
// Mark this request is hijacked or not.
<原文结束>

# <翻译开始>
// 标记此请求是否已被劫持。. md5:80adeb664fa8ae97
# <翻译结束>


<原文开始>
// Is header wrote or not, avoiding error: superfluous/multiple response.WriteHeader call.
<原文结束>

# <翻译开始>
// 是否已经写入了响应头，避免出现错误：多余的或多个response.WriteHeader调用。. md5:59bda0050b534efa
# <翻译结束>


<原文开始>
// Bytes written to response.
<原文结束>

# <翻译开始>
// 写入响应的字节数。. md5:cc5fa1ce145684ed
# <翻译结束>


<原文开始>
// NewWriter creates and returns a new Writer.
<原文结束>

# <翻译开始>
// NewWriter 创建并返回一个新的 Writer。. md5:6fad96ecb42a0036
# <翻译结束>


<原文开始>
// WriteHeader implements the interface of http.ResponseWriter.WriteHeader.
// Note that the underlying `WriteHeader` can only be called once in a http response.
<原文结束>

# <翻译开始>
// WriteHeader 实现了 http.ResponseWriter.WriteHeader 接口的方法。
// 注意，底层的 `WriteHeader` 方法在一个http响应中只能被调用一次。
// md5:7158450c7ec7fc1a
# <翻译结束>


<原文开始>
// BytesWritten returns the length that was written to response.
<原文结束>

# <翻译开始>
// BytesWritten 返回写入响应的长度。. md5:2bc5d732217ae6e4
# <翻译结束>


<原文开始>
// Write implements the interface function of http.ResponseWriter.Write.
<原文结束>

# <翻译开始>
// Write实现了http.ResponseWriter.Write接口函数。. md5:7078e0a4eee107f7
# <翻译结束>


<原文开始>
// Hijack implements the interface function of http.Hijacker.Hijack.
<原文结束>

# <翻译开始>
// Hijack 实现了 http.Hijacker.Hijack 接口函数。. md5:7ef9ff81765b052e
# <翻译结束>


<原文开始>
// IsHeaderWrote returns if the header status is written.
<原文结束>

# <翻译开始>
// IsHeaderWrote 返回头部状态是否已写入。. md5:7785f14e4d061fc9
# <翻译结束>


<原文开始>
// IsHijacked returns if the connection is hijacked.
<原文结束>

# <翻译开始>
// IsHijacked 返回连接是否已被劫持。. md5:11468dbc47bf2400
# <翻译结束>


<原文开始>
// Flush sends any buffered data to the client.
<原文结束>

# <翻译开始>
// Flush 将缓冲区中的任何数据发送到客户端。. md5:38eb50b527a1bfc5
# <翻译结束>

