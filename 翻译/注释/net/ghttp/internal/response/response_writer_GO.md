
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
// Writer wraps http.ResponseWriter for extra features.
<原文结束>

# <翻译开始>
// Writer 封装了 http.ResponseWriter，提供了额外的功能。
# <翻译结束>


<原文开始>
// The underlying ResponseWriter.
<原文结束>

# <翻译开始>
// 基础的 ResponseWriter。
# <翻译结束>


<原文开始>
// Mark this request is hijacked or not.
<原文结束>

# <翻译开始>
// 标记该请求是否已被劫持
# <翻译结束>


<原文开始>
// Is header wrote or not, avoiding error: superfluous/multiple response.WriteHeader call.
<原文结束>

# <翻译开始>
// IsHeaderWroteOrNot 判断头部是否已写入，避免出现“superfluous/multiple response.WriteHeader call”错误。
# <翻译结束>


<原文开始>
// NewWriter creates and returns a new Writer.
<原文结束>

# <翻译开始>
// NewWriter 创建并返回一个新的 Writer。
# <翻译结束>


<原文开始>
// WriteHeader implements the interface of http.ResponseWriter.WriteHeader.
<原文结束>

# <翻译开始>
// WriteHeader 实现了 http.ResponseWriter 接口中的 WriteHeader 方法。
# <翻译结束>


<原文开始>
// Hijack implements the interface function of http.Hijacker.Hijack.
<原文结束>

# <翻译开始>
// Hijack 实现了 http.Hijacker 接口中的 Hijack 函数。
# <翻译结束>


<原文开始>
// IsHeaderWrote returns if the header status is written.
<原文结束>

# <翻译开始>
// IsHeaderWrote 返回是否已写入头部状态。
# <翻译结束>


<原文开始>
// IsHijacked returns if the connection is hijacked.
<原文结束>

# <翻译开始>
// IsHijacked 返回连接是否已被劫持。
# <翻译结束>


<原文开始>
// Flush sends any buffered data to the client.
<原文结束>

# <翻译开始>
// Flush 将任何缓冲的数据发送到客户端。
# <翻译结束>

