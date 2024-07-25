
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
// BindObject registers object to server routes with a given pattern.
//
// The optional parameter `method` is used to specify the method to be registered, which
// supports multiple method names; multiple methods are separated by char ',', case-sensitive.
<原文结束>

# <翻译开始>
// BindObject 将对象绑定到具有给定模式的服务器路由。
//
// 可选参数 `method` 用于指定要注册的方法，支持多个方法名称；多个方法名称之间用字符 `,` 分隔，区分大小写。 md5:224eaf0adfd81c84
# <翻译结束>


<原文开始>
// BindObjectMethod registers specified method of the object to server routes with a given pattern.
//
// The optional parameter `method` is used to specify the method to be registered, which
// does not support multiple method names but only one, case-sensitive.
<原文结束>

# <翻译开始>
// BindObjectMethod 将指定对象的特定方法与给定模式的服务器路由绑定。
//
// 可选参数 `method` 用于指定要注册的方法，它不支持多个方法名，仅支持一个，且区分大小写。 md5:badb3f7323abfd11
# <翻译结束>


<原文开始>
// BindObjectRest registers object in REST API styles to server with a specified pattern.
<原文结束>

# <翻译开始>
// BindObjectRest 使用指定的模式将对象以REST API风格注册到服务器。 md5:e071850c88eb6751
# <翻译结束>


<原文开始>
// Convert input method to map for convenience and high performance searching purpose.
<原文结束>

# <翻译开始>
	// 将输入方法转换为映射，以便于进行高效便捷的搜索。 md5:116ad79ef3003f65
# <翻译结束>


<原文开始>
	// If the `method` in `pattern` is `defaultMethod`,
	// it removes for convenience for next statement control.
<原文结束>

# <翻译开始>
	// 如果`pattern`中的`method`为`defaultMethod`，为了方便后续语句的控制，它会移除。 md5:08bf69a00eee9caa
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `reflectValue`.
	// It then can retrieve all the methods both of struct/*struct.
<原文结束>

# <翻译开始>
	// 如果给定的`object`不是指针，那么它会创建一个临时的，其值为`reflectValue`。
	// 然后它可以获取结构体/`*struct`的所有方法。 md5:1e216cd9c7839ef2
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
		// 如果存在"Index"方法，则会自动添加一个额外的路由来匹配主URI，例如：
		// 如果模式是"/user"，那么"/user"和"/user/index"都会被自动
		// 注册。
		//
		// 请注意，如果模式中包含内置变量，这条路由将不会被自动添加。 md5:96b4d9eca149582c
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `v`.
<原文结束>

# <翻译开始>
	// 如果给定的`object`不是指针，那么它会创建一个临时的指针，
	// 其值为`v`。 md5:ea1cbad8bfbac476
# <翻译结束>

