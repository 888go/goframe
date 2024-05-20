
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
// DoRequestObj does HTTP request using standard request/response object.
// The request object `req` is defined like:
//
//	type UseCreateReq struct {
//	    g.Meta `path:"/user" method:"put"`
//	    // other fields....
//	}
//
// The response object `res` should be a pointer type. It automatically converts result
// to given object `res` is success.
//
// Example:
// var (
//
//	req = UseCreateReq{}
//	res *UseCreateRes
//
// )
//
// err := DoRequestObj(ctx, req, &res)
<原文结束>

# <翻译开始>
// DoRequestObj 使用标准的请求/响应对象进行HTTP请求。
// 请求对象 `req` 的定义如下：
//
//	类型 UseCreateReq 结构体 {
//	    g.Meta `path:"/user" method:"put"`
//	    // 其他字段...
//	}
//
// 响应对象 `res` 应该是指针类型。如果请求成功，它会自动将结果转换为给定的对象 `res`。
//
// 示例：
// 变量：
//
//	req = UseCreateReq{}
//	res *UseCreateRes
//
// 使用 DoRequestObj 函数，传入 ctx、req 和 res 指针，执行请求并获取可能的错误： 
//
//	err := DoRequestObj(ctx, req, &res)
// md5:a9b1690353dd26b2
# <翻译结束>


<原文开始>
// handlePathForObjRequest replaces parameters in `path` with parameters from request object.
// Eg:
// /order/{id}  -> /order/1
// /user/{name} -> /order/john
<原文结束>

# <翻译开始>
// handlePathForObjRequest 将请求对象中的参数替换到`path`中的占位符。
// 例如：
// /order/{id}  -> /order/1
// /user/{name} -> /user/john
// md5:96a0939362ee6d87
# <翻译结束>

