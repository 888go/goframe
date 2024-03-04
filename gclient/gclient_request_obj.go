// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"context"
	"net/http"
	"reflect"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/gogf/gf/v2/util/gutil"
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
		method = gmeta.Get(req, gtag.Method).String()
		path   = gmeta.Get(req, gtag.Path).String()
	)
	if method == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Method, reflect.TypeOf(req).String(),
		)
	}
	if path == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Path, reflect.TypeOf(req).String(),
		)
	}
	path = c.handlePathForObjRequest(path, req)
	switch gstr.ToUpper(method) {
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
		if result := c.RequestVar(ctx, method, path, req); res != nil && !result.IsEmpty() {
			return result.Scan(res)
		}
		return nil

	default:
		return gerror.Newf(`invalid HTTP method "%s"`, method)
	}
}

// handlePathForObjRequest 将 `path` 中的参数替换为请求对象中的参数。
// 例如：
// /order/{id}   -> /order/1
// /user/{name} -> /order/john
func (c *Client) handlePathForObjRequest(path string, req interface{}) string {
	if gstr.Contains(path, "{") {
		requestParamsMap := gconv.Map(req)
		if len(requestParamsMap) > 0 {
			path, _ = gregex.ReplaceStringFuncMatch(`\{(\w+)\}`, path, func(match []string) string {
				foundKey, foundValue := gutil.MapPossibleItemByKey(requestParamsMap, match[1])
				if foundKey != "" {
					return gconv.String(foundValue)
				}
				return match[0]
			})
		}
	}
	return path
}
