// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gogf/gf/v2/internal/intlog"
)

// Response is the struct for client request response.
type Response struct {
	*http.Response                   // Response is the underlying http.Response object of certain request.
	request        *http.Request     // Request is the underlying http.Request object of certain request.
	requestBody    []byte            // The body bytes of certain request, only available in Dump feature.
	cookies        map[string]string // Response cookies, which are only parsed once.
}

// initCookie initializes the cookie map attribute of Response.
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

// GetCookie retrieves and returns the cookie value of specified `key`.
// ff:取Cookie
// r:
// key:名称
func (r *Response) GetCookie(key string) string {
	r.initCookie()
	return r.cookies[key]
}

// GetCookieMap retrieves and returns a copy of current cookie values map.
// ff:取CookieMap
// r:
func (r *Response) GetCookieMap() map[string]string {
	r.initCookie()
	m := make(map[string]string, len(r.cookies))
	for k, v := range r.cookies {
		m[k] = v
	}
	return m
}

// ReadAll retrieves and returns the response content as []byte.
// ff:取响应字节集
// r:
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

// ReadAllString retrieves and returns the response content as string.
// ff:取响应文本
// r:
func (r *Response) ReadAllString() string {
	return string(r.ReadAll())
}

// SetBodyContent overwrites response content with custom one.
// ff:覆盖响应内容
// r:
// content:字节集
func (r *Response) SetBodyContent(content []byte) {
	buffer := bytes.NewBuffer(content)
	r.Body = io.NopCloser(buffer)
	r.ContentLength = int64(buffer.Len())
}

// Close closes the response when it will never be used.
// ff:关闭
// r:
func (r *Response) Close() error {
	if r == nil || r.Response == nil {
		return nil
	}
	return r.Response.Body.Close()
}
