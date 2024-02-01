
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
// Eg:
// var (
//
//	req = UseCreateReq{}
//	res *UseCreateRes
//
// )
// DoRequestObj(ctx, req, &res)
<原文结束>

# <翻译开始>
// DoRequestObj 使用标准请求/响应对象执行HTTP请求。
// 请求对象`req`的定义类似如下：
//
//	type UseCreateReq struct {
//	    g.Meta `path:"/user" method:"put"`
//// 其他字段...
//	}
//
// 响应对象`res`应该是指针类型。如果成功，它会自动将结果转换为给定的对象`res`。
// 例如：
// 
//	var (
//	
//		req = UseCreateReq{}
//		res *UseCreateRes
//	
//	)
// DoRequestObj(ctx, req, &res)
// 这段代码注释翻译成中文的大致意思是：该函数DoRequestObj利用标准的请求/响应对象进行HTTP请求操作。其中，请求对象`req`是一个结构体，其定义中包含了用于指定HTTP路径和方法的元数据字段。响应对象`res`应当是指针类型，当HTTP请求成功时，函数会自动将响应结果转换并填充到这个`res`对象中。示例代码展示了如何初始化请求与响应对象，并调用DoRequestObj函数执行请求。
# <翻译结束>


<原文开始>
// handlePathForObjRequest replaces parameters in `path` with parameters from request object.
// Eg:
// /order/{id}  -> /order/1
// /user/{name} -> /order/john
<原文结束>

# <翻译开始>
// handlePathForObjRequest 将 `path` 中的参数替换为请求对象中的参数。
// 例如：
// /order/{id}   -> /order/1
// /user/{name} -> /order/john
# <翻译结束>

