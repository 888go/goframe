// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/utils"
)

// dumpTextFormat 是转储原始字符串的格式. md5:050761651fa95433
const dumpTextFormat = `+---------------------------------------------+
|                   %s                  |
+---------------------------------------------+
%s
%s
`

// getResponseBody 返回响应体的文本。 md5:9b98c6205b0cfba9
func getResponseBody(res *http.Response) string {
	if res.Body == nil {
		return ""
	}
	bodyContent, _ := io.ReadAll(res.Body)
	res.Body = utils.NewReadCloser(bodyContent, true)
	return string(bodyContent)
}

// RawRequest 返回请求的原始内容。 md5:75945739a746b6fa
func (r *Response) RawRequest() string {
	// Response can be nil.
	if r == nil || r.request == nil {
		return ""
	}
		// DumpRequestOut 与 DumpRequest 相比，会写出更多的请求头信息，比如 User-Agent。 md5:acf8e2e8787c2534
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

// RawResponse 返回响应的原始内容。 md5:a3d4faef4d056f70
func (r *Response) RawResponse() string {
	// Response might be nil.
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

// Raw返回请求和响应的原始文本。 md5:68a1f59b34b9e33a
func (r *Response) Raw() string {
	return fmt.Sprintf("%s\n%s", r.RawRequest(), r.RawResponse())
}

// RawDump 将请求和响应的原始文本输出到stdout。 md5:aa9a7a2cc5e60970
func (r *Response) RawDump() {
	fmt.Println(r.Raw())
}
