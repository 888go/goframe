// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包 httputil 提供仅供内部使用的 HTTP 函数。 md5:68a87514ccfd9190
package httputil

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	fileUploadingKey = "@file:"
)

// BuildParams 为http客户端构建请求字符串。`params`可以是以下类型：
// 字符串/字节切片/映射/结构体/结构体指针。
//
// 可选参数`noUrlEncode`指定是否忽略数据的URL编码。 md5:664ad104f4b3f610
func BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string) {
	// 如果给定字符串/字节切片，会直接转换并返回它作为字符串。 md5:80d9827515b7e847
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
	// 否则，它会将它转换为map并进行URL编码。 md5:932b0b216ae84f60
	m, urlEncode := gconv.Map(params), true
	if len(m) == 0 {
		return gconv.String(params)
	}
	if len(noUrlEncode) == 1 {
		urlEncode = !noUrlEncode[0]
	}
	// 如果有文件上传，它将忽略URL编码。 md5:e349803af0cef3a3
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
		// Ignore nil attributes.
		if empty.IsNil(v) {
			continue
		}
		if len(encodedParamStr) > 0 {
			encodedParamStr += "&"
		}
		s = gconv.String(v)
		if urlEncode {
			if strings.HasPrefix(s, fileUploadingKey) && len(s) > len(fileUploadingKey) {
				// 如果正在上传文件，则不进行URL编码。 md5:1d89b2d337a7a0e9
			} else {
				s = gurl.Encode(s)
			}
		}
		encodedParamStr += k + "=" + s
	}
	return
}

// HeaderToMap 将请求头转换为映射。 md5:d7b057a672ffda30
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
