// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"context"
	"net/http"
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gmeta "github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
	gutil "github.com/888go/goframe/util/gutil"
)

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

// handlePathForObjRequest 将请求对象中的参数替换到`path`中的占位符。
// 例如：
// /order/{id}  -> /order/1
// /user/{name} -> /user/john
// md5:96a0939362ee6d87
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
