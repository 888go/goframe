// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"context"
	"net/http"
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
	"github.com/888go/goframe/util/gutil"
)

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
func (c *Client) DoRequestObj(ctx context.Context, req, res interface{}) error {
	var (
		method = 元数据类.Get(req, gtag.Method).String()
		path   = 元数据类.Get(req, gtag.Path).String()
	)
	if method == "" {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Method, reflect.TypeOf(req).String(),
		)
	}
	if path == "" {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Path, reflect.TypeOf(req).String(),
		)
	}
	path = c.handlePathForObjRequest(path, req)
	switch 文本类.X到大写(method) {
	case
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace:
		if result := c.X请求泛型类(ctx, method, path, req); res != nil && !result.X是否为空() {
			return result.X取结构体指针(res)
		}
		return nil

	default:
		return 错误类.X创建并格式化(`invalid HTTP method "%s"`, method)
	}
}

// handlePathForObjRequest 将 `path` 中的参数替换为请求对象中的参数。
// 例如：
// /order/{id}   -> /order/1
// /user/{name} -> /order/john
func (c *Client) handlePathForObjRequest(path string, req interface{}) string {
	if 文本类.X是否包含(path, "{") {
		requestParamsMap := 转换类.X取Map(req)
		if len(requestParamsMap) > 0 {
			path, _ = 正则类.ReplaceStringFuncMatch(`\{(\w+)\}`, path, func(match []string) string {
				foundKey, foundValue := 工具类.MapPossibleItemByKey(requestParamsMap, match[1])
				if foundKey != "" {
					return 转换类.String(foundValue)
				}
				return match[0]
			})
		}
	}
	return path
}
