// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/utils"
)

// dumpTextFormat is the format of the dumped raw string
const dumpTextFormat = `+---------------------------------------------+
|                   %s                  |
+---------------------------------------------+
%s
%s
`

// getResponseBody returns the text of the response body.
func getResponseBody(res *http.Response) string {
	if res.Body == nil {
		return ""
	}
	bodyContent, _ := io.ReadAll(res.Body)
	res.Body = utils.NewReadCloser(bodyContent, true)
	return string(bodyContent)
}

// RawRequest returns the raw content of the request.

// ff:取请求原始文本
func (r *Response) RawRequest() string {
	// Response can be nil.
	if r == nil || r.request == nil {
		return ""
	}
	// DumpRequestOut writes more request headers than DumpRequest, such as User-Agent.
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

// RawResponse returns the raw content of the response.

// ff:取响应原始文本
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

// Raw returns the raw text of the request and the response.

// ff:取请求和响应原始文本
func (r *Response) Raw() string {
	return fmt.Sprintf("%s\n%s", r.RawRequest(), r.RawResponse())
}

// RawDump outputs the raw text of the request and the response to stdout.

// ff:请求和响应输出终端
func (r *Response) RawDump() {
	fmt.Println(r.Raw())
}
