
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
// BindObject registers object to server routes with a given pattern.
//
// The optional parameter `method` is used to specify the method to be registered, which
// supports multiple method names; multiple methods are separated by char ',', case-sensitive.
<原文结束>

# <翻译开始>
// BindObject 将对象注册到服务器路由上，给定特定的模式。
//
// 可选参数 `method` 用于指定要注册的方法，该方法支持多个方法名；
// 多个方法之间用字符 ',' 分隔，大小写敏感。
# <翻译结束>


<原文开始>
// BindObjectMethod registers specified method of the object to server routes with a given pattern.
//
// The optional parameter `method` is used to specify the method to be registered, which
// does not support multiple method names but only one, case-sensitive.
<原文结束>

# <翻译开始>
// BindObjectMethod 将指定对象的方法注册到服务器路由中，使用给定的模式。
//
// 可选参数 `method` 用于指定要注册的方法，该参数不支持多个方法名，仅支持单个、大小写敏感的方法名。
# <翻译结束>


<原文开始>
// BindObjectRest registers object in REST API styles to server with a specified pattern.
<原文结束>

# <翻译开始>
// BindObjectRest 以指定模式将符合REST API风格的对象注册到服务器。
# <翻译结束>


<原文开始>
// Convert input method to map for convenience and high performance searching purpose.
<原文结束>

# <翻译开始>
// 将输入方法转换为映射以便于实现高效检索
# <翻译结束>


<原文开始>
	// If the `method` in `pattern` is `defaultMethod`,
	// it removes for convenience for next statement control.
<原文结束>

# <翻译开始>
// 如果`pattern`中的`method`是`defaultMethod`，
// 为了方便后续语句的控制，将其移除。
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `reflectValue`.
	// It then can retrieve all the methods both of struct/*struct.
<原文结束>

# <翻译开始>
// 如果给定的`object`不是指针，它会创建一个临时指针，
// 其指向值为`reflectValue`。
// 然后可以获取结构体（包括结构体指针）的所有方法。
// 这段代码注释的翻译如下：
// ```go
// 如果传入的`object`不是一个指针类型，
// 则会创建一个临时指针变量，该指针指向`reflectValue`。
// 这样就可以获取到结构体及其指针类型的全部方法。
# <翻译结束>


<原文开始>
		// If there's "Index" method, then an additional route is automatically added
		// to match the main URI, for example:
		// If pattern is "/user", then "/user" and "/user/index" are both automatically
		// registered.
		//
		// Note that if there's built-in variables in pattern, this route will not be added
		// automatically.
<原文结束>

# <翻译开始>
// 如果存在"Index"方法，则会自动添加一个附加路由以匹配主URI，例如：
// 如果模式是"/user"，那么"/user"和"/user/index"都会被自动注册。
//
// 注意，如果模式中存在内置变量，则此路由不会被自动添加。
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `v`.
<原文结束>

# <翻译开始>
// 如果给定的`object`不是指针，则创建一个临时指针，
// 其值为`v`。
# <翻译结束>

