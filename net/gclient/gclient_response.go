// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"bytes"
	"io"
	"net/http"
	
	"github.com/888go/goframe/internal/intlog"
)

// Response 是用于客户端请求响应的结构体。
type Response struct {
	*http.Response                   // Response 是某个请求的底层 http.Response 对象。
	request        *http.Request     // Request 是某个请求的底层 http.Request 对象。
	requestBody    []byte            // 这是某些请求的主体字节，仅在Dump功能中可用。
	cookies        map[string]string // 响应cookies，这些只解析一次。
}

// initCookie 初始化 Response 结构体中的 cookie map 属性。
func (r *Response) initCookie() {
	if r.cookies == nil {
		r.cookies = make(map[string]string)
		// 响应可能为 nil。
		if r.Response != nil {
			for _, v := range r.Cookies() {
				r.cookies[v.Name] = v.Value
			}
		}
	}
}

// GetCookie 根据指定的 `key` 获取并返回 cookie 的值。
func (r *Response) X取Cookie(名称 string) string {
	r.initCookie()
	return r.cookies[名称]
}

// GetCookieMap 获取并返回当前cookie值的映射副本。
func (r *Response) X取CookieMap() map[string]string {
	r.initCookie()
	m := make(map[string]string, len(r.cookies))
	for k, v := range r.cookies {
		m[k] = v
	}
	return m
}

// ReadAll 方法获取并返回响应内容作为 []byte 类型的切片。
func (r *Response) X取响应字节集() []byte {
	// 响应可能为 nil。
	if r == nil || r.Response == nil {
		return []byte{}
	}
	body, err := io.ReadAll(r.Response.Body)
	if err != nil {
		intlog.Errorf(r.request.Context(), `%+v`, err)
		return nil
	}
	return body
}

// ReadAllString 获取并返回响应内容作为字符串。
func (r *Response) X取响应文本() string {
	return string(r.X取响应字节集())
}

// SetBodyContent 函数用于覆盖并替换响应内容为自定义内容。
func (r *Response) X覆盖响应内容(字节集 []byte) {
	buffer := bytes.NewBuffer(字节集)
	r.Body = io.NopCloser(buffer)
	r.ContentLength = int64(buffer.Len())
}

// Close在响应不再需要使用时关闭该响应。
func (r *Response) X关闭() error {
	if r == nil || r.Response == nil {
		return nil
	}
	return r.Response.Body.Close()
}
