
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
// BindHandler registers a handler function to server with a given pattern.
//
// Note that the parameter `handler` can be type of:
// 1. func(*ghttp.Request)
// 2. func(context.Context, BizRequest)(BizResponse, error)
<原文结束>

# <翻译开始>
// BindHandler 将一个处理函数注册到服务器，该函数与给定的模式关联。
//
// 注意参数 `handler` 可以是以下类型：
// 1. func(*ghttp.Request) // 类型为接收*ghttp.Request参数的函数
// 2. func(context.Context, BizRequest)(BizResponse, error) // 类型为接收context.Context和BizRequest参数，并返回BizResponse和error的函数
# <翻译结束>


<原文开始>
// doBindHandler registers a handler function to server with given pattern.
//
// The parameter `pattern` is like:
// /user/list, put:/user, delete:/user, post:/user@goframe.org
<原文结束>

# <翻译开始>
// doBindHandler 函数用于将指定模式的处理器函数注册到服务器。
//
// 参数 `pattern` 形如：
// /user/list, put:/user, delete:/user, post:/user@goframe.org
// 其中，这些模式用于定义HTTP请求的方法（如GET、PUT、DELETE等）以及对应的路由路径。
# <翻译结束>


<原文开始>
// bindHandlerByMap registers handlers to server using map.
<原文结束>

# <翻译开始>
// bindHandlerByMap 使用map将处理器注册到服务器。
# <翻译结束>


<原文开始>
// mergeBuildInNameToPattern merges build-in names into the pattern according to the following
// rules, and the built-in names are named like "{.xxx}".
// Rule 1: The URI in pattern contains the {.struct} keyword, it then replaces the keyword with the struct name;
// Rule 2: The URI in pattern contains the {.method} keyword, it then replaces the keyword with the method name;
// Rule 2: If Rule 1 is not met, it then adds the method name directly to the URI in the pattern;
//
// The parameter `allowAppend` specifies whether allowing appending method name to the tail of pattern.
<原文结束>

# <翻译开始>
// mergeBuildInNameToPattern 将内建名称按照以下规则合并到模式中，这些内建名称的命名格式为 "{.xxx}"。
// 规则1：若模式中的URI包含 {.struct} 关键字，则用结构体名称替换该关键字；
// 规则2：若模式中的URI包含 {.method} 关键字，则用方法名称替换该关键字；
// 规则3：如果未满足规则1，则将方法名称直接追加到模式中URI的末尾；
//
// 参数 `allowAppend` 指定是否允许将方法名称追加到模式末尾。
# <翻译结束>







<原文开始>
// Append the domain parameter to URI.
<原文结束>

# <翻译开始>
// 将domain参数追加到URI。
# <翻译结束>


<原文开始>
// nameToUri converts the given name to the URL format using the following rules:
// Rule 0: Convert all method names to lowercase, add char '-' between words.
// Rule 1: Do not convert the method name, construct the URI with the original method name.
// Rule 2: Convert all method names to lowercase, no connecting symbols between words.
// Rule 3: Use camel case naming.
<原文结束>

# <翻译开始>
// nameToUri 将给定名称转换为URL格式，遵循以下规则：
// 规则0：将所有方法名转为小写，并在单词间添加字符'-'。
// 规则1：不转换方法名，使用原始方法名构建URI。
// 规则2：将所有方法名转为小写，单词间无连接符号。
// 规则3：使用驼峰命名法。
# <翻译结束>


<原文开始>
	// Do not enable this logic, as many users are already using none struct pointer type
	// as the first output parameter.
<原文结束>

# <翻译开始>
// 不要启用这段逻辑，因为许多用户已经使用非结构体指针类型作为第一个输出参数。
# <翻译结束>


<原文开始>
// The request struct should be named as `xxxReq`.
<原文结束>

# <翻译开始>
// 请求结构体应命名为 `xxxReq`。
# <翻译结束>


<原文开始>
// The response struct should be named as `xxxRes`.
<原文结束>

# <翻译开始>
// 响应结构体应当命名为 `xxxRes`。
# <翻译结束>


<原文开始>
// It retrieves and returns the request struct fields.
<原文结束>

# <翻译开始>
// 它检索并返回请求结构体的字段。
# <翻译结束>


<原文开始>
// Call handler with dynamic created parameter values.
<原文结束>

# <翻译开始>
// 使用动态创建的参数值调用处理器。
# <翻译结束>


<原文开始>
// trimGeneric removes type definitions string from response type name if generic
<原文结束>

# <翻译开始>
// trimGeneric 从响应类型名称中移除泛型的类型定义字符串（如果存在的话）
# <翻译结束>


<原文开始>
// for generic, it is faster to start at the end than at the beginning
<原文结束>

# <翻译开始>
// 对于泛型，从结尾开始遍历比从开头开始更快
# <翻译结束>







<原文开始>
		// may be a slice, because generic is '[X]', not '[]'
		// to be compatible with bad return parameter type: []XxxRes
<原文结束>

# <翻译开始>
// 可能是一个切片，因为泛型是 '[X]'，而不是 '[]'
// 为了兼容不规范的返回参数类型：[]XxxRes
# <翻译结束>


<原文开始>
// Check domain parameter.
<原文结束>

# <翻译开始>
// 检查域名参数。
# <翻译结束>


<原文开始>
// not found '[' or ']'
<原文结束>

# <翻译开始>
// 未找到 '[' 或 ']'
# <翻译结束>

