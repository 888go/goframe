
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
// BindHandler registers a handler function to server with a given pattern.
//
// Note that the parameter `handler` can be type of:
// 1. func(*ghttp.Request)
// 2. func(context.Context, BizRequest)(BizResponse, error)
<原文结束>

# <翻译开始>
// BindHandler 将一个处理器函数注册到服务器，使用给定的模式。
// 
// 注意，参数 `handler` 可以是以下两种类型之一：
// 1. func(*ghttp.Request)
// 2. func(context.Context, BizRequest) (BizResponse, error)
// md5:245b5139c4d933ad
# <翻译结束>


<原文开始>
// doBindHandler registers a handler function to server with given pattern.
//
// The parameter `pattern` is like:
// /user/list, put:/user, delete:/user, post:/user@goframe.org
<原文结束>

# <翻译开始>
// doBindHandler 使用给定的模式向服务器注册一个处理函数。
//
// 参数 `pattern` 的格式如下：
// /user/list, put:/user, delete:/user, post:/user@goframe.org
// md5:d71f121a1c2830d3
# <翻译结束>


<原文开始>
// bindHandlerByMap registers handlers to server using map.
<原文结束>

# <翻译开始>
// bindHandlerByMap 使用映射注册处理器到服务器。 md5:15729f837b1bc875
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
// mergeBuildInNameToPattern 将内置名称合并到模式中，根据以下规则进行操作，内置名称的命名方式为"{.xxx}"。
// 规则 1：如果模式中的URI包含{.struct}关键字，它将替换该关键字为结构体名称；
// 规则 2：如果模式中的URI包含{.method}关键字，它将替换该关键字为方法名称；
// 规则 3：如果没有满足规则 1，那么直接在模式中的URI后添加方法名称。
// 
// 参数 `allowAppend` 指定是否允许将方法名称追加到模式的末尾。
// md5:1c79af7afc57b081
# <翻译结束>


<原文开始>
// Check domain parameter.
<原文结束>

# <翻译开始>
	// 检查域名参数。 md5:1a963c36e4fee004
# <翻译结束>


<原文开始>
// Append the domain parameter to URI.
<原文结束>

# <翻译开始>
	// 将域名参数追加到URI。 md5:f94214453c1409c8
# <翻译结束>


<原文开始>
// nameToUri converts the given name to the URL format using the following rules:
// Rule 0: Convert all method names to lowercase, add char '-' between words.
// Rule 1: Do not convert the method name, construct the URI with the original method name.
// Rule 2: Convert all method names to lowercase, no connecting symbols between words.
// Rule 3: Use camel case naming.
<原文结束>

# <翻译开始>
// nameToUri 使用以下规则将给定的名称转换为URL格式：
// 规则0：将所有方法名转换为小写，单词间添加字符'-'。
// 规则1：不转换方法名，使用原始方法名构建URI。
// 规则2：将所有方法名转换为小写，单词间不添加连接符号。
// 规则3：使用驼峰式命名。
// md5:c9f350c3c6635668
# <翻译结束>


<原文开始>
	// Do not enable this logic, as many users are already using none struct pointer type
	// as the first output parameter.
<原文结束>

# <翻译开始>
	// 不要启用此逻辑，因为许多用户已经将非结构指针类型作为第一个输出参数使用。
	// md5:46785e26d27207d1
# <翻译结束>


<原文开始>
// The request struct should be named as `xxxReq`.
<原文结束>

# <翻译开始>
	// 请求结构体应该命名为 `xxxReq`。 md5:f366399bf3de35a1
# <翻译结束>


<原文开始>
// The response struct should be named as `xxxRes`.
<原文结束>

# <翻译开始>
	// 响应结构体应当命名为 `xxxRes`。 md5:0e837067ff972f27
# <翻译结束>


<原文开始>
// It retrieves and returns the request struct fields.
<原文结束>

# <翻译开始>
	// 该函数获取并返回请求结构体的字段。 md5:25b3db67b1969d01
# <翻译结束>


<原文开始>
// Call handler with dynamic created parameter values.
<原文结束>

# <翻译开始>
		// 使用动态创建的参数值调用处理器。 md5:991efec71cdcc95a
# <翻译结束>


<原文开始>
// trimGeneric removes type definitions string from response type name if generic
<原文结束>

# <翻译开始>
// trimGeneric 如果响应类型名称包含泛型定义，删除字符串类型的定义. md5:3c6ea03dfa650b71
# <翻译结束>


<原文开始>
// for generic, it is faster to start at the end than at the beginning
<原文结束>

# <翻译开始>
// 对于泛型来说，从末尾开始比从开头开始更快. md5:9e8730bfe1647d52
# <翻译结束>


<原文开始>
		// may be a slice, because generic is '[X]', not '[]'
		// to be compatible with bad return parameter type: []XxxRes
<原文结束>

# <翻译开始>
		// 可能是一个切片，因为泛型是'[X]'而不是'[]'
		// 以兼容不良的返回参数类型：[]XxxRes
		// md5:a521893d3e187a1a
# <翻译结束>

