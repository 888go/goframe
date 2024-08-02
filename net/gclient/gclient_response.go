// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"bytes"
	"io"
	"net/http"

	"github.com/888go/goframe/internal/intlog"
)

// Response 是客户端请求响应的结构体。 md5:3dcf67f589d2cb7d
type Response struct {
	*http.Response                   // Response是某些请求的底层http.Response对象。 md5:a2ccaed0095f8487
	request        *http.Request     // Request是某些请求的底层http.Request对象。 md5:2ad9c57c9cf73fe5
	requestBody    []byte            // 某些请求的主体字节，仅在Dump功能中可用。 md5:6d14d65059cbc145
	cookies        map[string]string // 响应cookies，只解析一次。 md5:1769cabc50a055f0
}

// initCookie 初始化Response的cookie映射属性。 md5:c9b19789c8900efe
func (r *Response) initCookie() {
	if r.cookies == nil {
		r.cookies = make(map[string]string)
		// Response might be nil.
		if r.Response != nil {
			for _, v := range r.Cookies() {
				r.cookies[v.Name] = v.Value
			}
		}
	}
}

// GetCookie 从指定的`key`获取并返回 cookie 的值。 md5:59efa13c53c894a8
func (r *Response) GetCookie(key string) string {
	r.initCookie()
	return r.cookies[key]
}

// GetCookieMap 获取并返回当前cookie值映射的副本。 md5:b1a4ecf0af8f77bd
func (r *Response) GetCookieMap() map[string]string {
	r.initCookie()
	m := make(map[string]string, len(r.cookies))
	for k, v := range r.cookies {
		m[k] = v
	}
	return m
}

// ReadAll 读取并返回响应内容为 []byte。 md5:a94558987266f586
func (r *Response) ReadAll() []byte {
	// Response might be nil.
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

// ReadAllString 读取并返回响应内容作为字符串。 md5:5c05c71fc55a7f9d
func (r *Response) ReadAllString() string {
	return string(r.ReadAll())
}

// SetBodyContent 用自定义内容替换响应体。 md5:01e03245d3adc65f
func (r *Response) SetBodyContent(content []byte) {
	buffer := bytes.NewBuffer(content)
	r.Body = io.NopCloser(buffer)
	r.ContentLength = int64(buffer.Len())
}

// 当响应将永远不再被使用时，Close会关闭该响应。 md5:3c208e3775456196
func (r *Response) Close() error {
	if r == nil || r.Response == nil {
		return nil
	}
	return r.Response.Body.Close()
}
