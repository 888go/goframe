// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/utils"
)

// dumpTextFormat 是转储原始字符串的格式
const dumpTextFormat = `+---------------------------------------------+
|                   %s                  |
+---------------------------------------------+
%s
%s
`

// getResponseBody 函数返回响应体的文本内容。
func getResponseBody(res *http.Response) string {
	if res.Body == nil {
		return ""
	}
	bodyContent, _ := io.ReadAll(res.Body)
	res.Body = utils.NewReadCloser(bodyContent, true)
	return string(bodyContent)
}

// RawRequest 返回请求的原始内容。
func (r *Response) X取请求原始文本() string {
	// 响应可以为 nil。
	if r == nil || r.request == nil {
		return ""
	}
	// DumpRequestOut 会比 DumpRequest 写出更多的请求头信息，比如 User-Agent。
	bs, err := httputil.DumpRequestOut(r.request, false)
	if err != nil {
		intlog.Errorf(r.request.Context(), `%+v`, err)
		return ""
	}
	return fmt.Sprintf(
		dumpTextFormat,
		"REQUEST ",
		string(bs),
		r.requestBody,
	)
}

// RawResponse 返回响应的原始内容。
func (r *Response) X取响应原始文本() string {
	// 响应可能为 nil。
	if r == nil || r.Response == nil {
		return ""
	}
	bs, err := httputil.DumpResponse(r.Response, false)
	if err != nil {
		intlog.Errorf(r.request.Context(), `%+v`, err)
		return ""
	}

	return fmt.Sprintf(
		dumpTextFormat,
		"RESPONSE",
		string(bs),
		getResponseBody(r.Response),
	)
}

// Raw 返回请求和响应的原始文本。
func (r *Response) X取请求和响应原始文本() string {
	return fmt.Sprintf("%s\n%s", r.X取请求原始文本(), r.X取响应原始文本())
}

// RawDump 将请求和响应的原始文本输出到标准输出（stdout）。
func (r *Response) X请求和响应输出终端() {
	fmt.Println(r.X取请求和响应原始文本())
}
