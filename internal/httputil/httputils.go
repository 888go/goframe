// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// httputil 包提供内部使用的HTTP功能。
package httputil

import (
	"net/http"
	"strings"
	
	"github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

const (
	fileUploadingKey = "@file:"
)

// BuildParams 为 http 客户端构建请求字符串。`params` 参数可以是以下类型：
// string/[]byte/map/struct/*struct。
//
// 可选参数 `noUrlEncode` 指定是否忽略数据的 URL 编码。
// 详细解释：
// 此 Go 语言注释描述了一个名为 `BuildParams` 的函数，该函数用于根据给定的参数构建一个 HTTP 客户端的请求字符串。`params` 参数支持多种数据类型，包括但不限于：字符串(string)、字节切片([]byte)、映射(map)、结构体(struct)以及指向结构体的指针(*struct)。
// 另外，该函数接受一个可选参数 `noUrlEncode`，这个布尔类型的参数决定了在构建请求字符串时是否跳过对数据进行 URL 编码处理。如果设为 `true`，则表示不进行 URL 编码；否则（默认情况或设为 `false`），将对数据进行 URL 编码。
func BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string) {
	// 如果给定的是字符串或[]byte，直接将其转换并返回为字符串。
	switch v := params.(type) {
	case string, []byte:
		return gconv.String(params)
	case []interface{}:
		if len(v) > 0 {
			params = v[0]
		} else {
			params = nil
		}
	}
	// Else 将其转换为 map，并执行 URL 编码。
	m, urlEncode := gconv.Map(params), true
	if len(m) == 0 {
		return gconv.String(params)
	}
	if len(noUrlEncode) == 1 {
		urlEncode = !noUrlEncode[0]
	}
	// 如果存在文件上传，则忽略URL编码。
	if urlEncode {
		for k, v := range m {
			if gstr.Contains(k, fileUploadingKey) || gstr.Contains(gconv.String(v), fileUploadingKey) {
				urlEncode = false
				break
			}
		}
	}
	s := ""
	for k, v := range m {
		// 忽略nil属性。
		if empty.IsNil(v) {
			continue
		}
		if len(encodedParamStr) > 0 {
			encodedParamStr += "&"
		}
		s = gconv.String(v)
		if urlEncode {
			if strings.HasPrefix(s, fileUploadingKey) && len(s) > len(fileUploadingKey) {
				// 如果上传文件，则不进行URL编码
			} else {
				s = gurl.Encode(s)
			}
		}
		encodedParamStr += k + "=" + s
	}
	return
}

// HeaderToMap 将请求头转换为映射（map）。
func HeaderToMap(header http.Header) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range header {
		if len(v) > 1 {
			m[k] = v
		} else if len(v) == 1 {
			m[k] = v[0]
		}
	}
	return m
}
