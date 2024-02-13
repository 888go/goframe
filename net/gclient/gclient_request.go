// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/httputil"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Get 发送GET请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，那么它必须被关闭。
func (c *Client) Get响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodGet, url, 参数...)
}

// Put 发送PUT请求并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭它。
func (c *Client) Put响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodPut, url, 参数...)
}

// Post 使用HTTP方法POST发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
func (c *Client) Post响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodPost, url, 参数...)
}

// Delete 发送DELETE请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，那么它必须被关闭。
func (c *Client) Delete响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodDelete, url, 参数...)
}

// Head 发送HEAD请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
func (c *Client) Head响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodHead, url, 参数...)
}

// Patch 发送 PATCH 请求并返回响应对象。
// 注意：如果响应对象将永不被使用，那么它必须被关闭。
func (c *Client) Patch响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodPatch, url, 参数...)
}

// Connect 发送 CONNECT 请求并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭它。
func (c *Client) Connect响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodConnect, url, 参数...)
}

// Options 发送 OPTIONS 请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
func (c *Client) Options响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodOptions, url, 参数...)
}

// Trace 发送 TRACE 请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，则必须关闭它。
func (c *Client) Trace响应对象(上下文 context.Context, url string, 参数 ...interface{}) (*Response, error) {
	return c.X请求响应对象(上下文, http.MethodTrace, url, 参数...)
}

// PostForm 与 net/http.PostForm 不同。
// 它是 Post 方法的一个包装器，会将 Content-Type 设置为 "multipart/form-data;"。
// 并且它会自动为请求体和 Content-Type 设置边界字符。
//
// 其效果类似于以下情况：
//
// Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN
//
// 表单数据格式如下：
// ------Boundarye4Ghaog6giyQ9ncN
// Content-Disposition: form-data; name="checkType"
//
// none
//
// 该方法用于发送表单数据。
// 注意，如果响应对象不再使用，则必须关闭它。
func (c *Client) Post表单响应对象(上下文 context.Context, url string, Map参数 map[string]string) (响应 *Response, 错误 error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range Map参数 {
		err := w.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}
	错误 = w.Close()
	if 错误 != nil {
		return nil, 错误
	}
	return c.X内容类型(w.FormDataContentType()).Post响应对象(上下文, url, body)
}

// DoRequest 使用给定的HTTP方法和数据发送请求，并返回响应对象。
// 注意，如果响应对象将不再使用，则必须关闭它。
//
// 注意，如果请求包含文件上传，则它使用"multipart/form-data"作为其Content-Type，
// 否则使用"application/x-www-form-urlencoded"。它还会自动检测POST内容的JSON格式，
// 对于JSON格式的内容，会自动将其Content-Type设置为"application/json"。
func (c *Client) X请求响应对象(上下文 context.Context, method, url string, 参数 ...interface{}) (响应 *Response, 错误 error) {
	req, 错误 := c.prepareRequest(上下文, method, url, 参数...)
	if 错误 != nil {
		return nil, 错误
	}

	// 客户端中间件。
	if len(c.middlewareHandler) > 0 {
		mdlHandlers := make([]X中间件函数, 0, len(c.middlewareHandler)+1)
		mdlHandlers = append(mdlHandlers, c.middlewareHandler...)
		mdlHandlers = append(mdlHandlers, func(cli *Client, r *http.Request) (*Response, error) {
			return cli.callRequest(r)
		})
		上下文 = context.WithValue(req.Context(), clientMiddlewareKey, &clientMiddleware{
			client:       c,
			handlers:     mdlHandlers,
			handlerIndex: -1,
		})
		req = req.WithContext(上下文)
		响应, 错误 = c.Next(req)
	} else {
		响应, 错误 = c.callRequest(req)
	}
	return 响应, 错误
}

// 准备请求：验证请求参数，构建并返回HTTP请求。
func (c *Client) prepareRequest(ctx context.Context, method, url string, data ...interface{}) (req *http.Request, err error) {
	method = strings.ToUpper(method)
	if len(c.prefix) > 0 {
		url = c.prefix + 文本类.X过滤首尾符并含空白(url)
	}
	if !文本类.X是否包含并忽略大小写(url, httpProtocolName) {
		url = httpProtocolName + `://` + url
	}
	var params string
	if len(data) > 0 {
		switch c.header[httpHeaderContentType] {
		case httpHeaderContentTypeJson:
			switch data[0].(type) {
			case string, []byte:
				params = 转换类.String(data[0])
			default:
				if b, err := json.Marshal(data[0]); err != nil {
					return nil, err
				} else {
					params = string(b)
				}
			}

		case httpHeaderContentTypeXml:
			switch data[0].(type) {
			case string, []byte:
				params = 转换类.String(data[0])
			default:
				if b, err := json类.X创建(data[0]).X取xml字节集(); err != nil {
					return nil, err
				} else {
					params = string(b)
				}
			}
		default:
			params = httputil.BuildParams(data[0], c.noUrlEncode)
		}
	}
	if method == http.MethodGet {
		var bodyBuffer *bytes.Buffer
		if params != "" {
			switch c.header[httpHeaderContentType] {
			case
				httpHeaderContentTypeJson,
				httpHeaderContentTypeXml:
				bodyBuffer = bytes.NewBuffer([]byte(params))
			default:
// 如果HTTP方法为GET且未指定Content-Type，则将参数追加到URL中。
				if 文本类.X是否包含(url, "?") {
					url = url + "&" + params
				} else {
					url = url + "?" + params
				}
				bodyBuffer = bytes.NewBuffer(nil)
			}
		} else {
			bodyBuffer = bytes.NewBuffer(nil)
		}
		if req, err = http.NewRequest(method, url, bodyBuffer); err != nil {
			err = 错误类.X多层错误并格式化(err, `http.NewRequest failed with method "%s" and URL "%s"`, method, url)
			return nil, err
		}
	} else {
		if strings.Contains(params, httpParamFileHolder) {
			// 文件上传请求。
			var (
				buffer = bytes.NewBuffer(nil)
				writer = multipart.NewWriter(buffer)
			)
			for _, item := range strings.Split(params, "&") {
				array := strings.Split(item, "=")
				if len(array[1]) > 6 && strings.Compare(array[1][0:6], httpParamFileHolder) == 0 {
					path := array[1][6:]
					if !文件类.X是否存在(path) {
						return nil, 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `"%s" does not exist`, path)
					}
					var (
						file          io.Writer
						formFileName  = 文件类.X路径取文件名(path)
						formFieldName = array[0]
					)
					if file, err = writer.CreateFormFile(formFieldName, formFileName); err != nil {
						err = 错误类.X多层错误并格式化(err, `CreateFormFile failed with "%s", "%s"`, formFieldName, formFileName)
						return nil, err
					} else {
						var f *os.File
						if f, err = 文件类.X打开并按只读模式(path); err != nil {
							return nil, err
						}
						if _, err = io.Copy(file, f); err != nil {
							err = 错误类.X多层错误并格式化(err, `io.Copy failed from "%s" to form "%s"`, path, formFieldName)
							_ = f.Close()
							return nil, err
						}
						_ = f.Close()
					}
				} else {
					var (
						fieldName  = array[0]
						fieldValue = array[1]
					)
					if err = writer.WriteField(fieldName, fieldValue); err != nil {
						err = 错误类.X多层错误并格式化(err, `write form field failed with "%s", "%s"`, fieldName, fieldValue)
						return nil, err
					}
				}
			}
// Close 结束多部分消息，并将尾部边界结束行写入输出。
			if err = writer.Close(); err != nil {
				err = 错误类.X多层错误并格式化(err, `form writer close failed`)
				return nil, err
			}

			if req, err = http.NewRequest(method, url, buffer); err != nil {
				err = 错误类.X多层错误并格式化(err, `http.NewRequest failed for method "%s" and URL "%s"`, method, url)
				return nil, err
			} else {
				req.Header.Set(httpHeaderContentType, writer.FormDataContentType())
			}
		} else {
			// Normal request.
			paramBytes := []byte(params)
			if req, err = http.NewRequest(method, url, bytes.NewReader(paramBytes)); err != nil {
				err = 错误类.X多层错误并格式化(err, `http.NewRequest failed for method "%s" and URL "%s"`, method, url)
				return nil, err
			} else {
				if v, ok := c.header[httpHeaderContentType]; ok {
					// 自定义内容类型。
					req.Header.Set(httpHeaderContentType, v)
				} else if len(paramBytes) > 0 {
					if (paramBytes[0] == '[' || paramBytes[0] == '{') && json.Valid(paramBytes) {
						// 自动检测并设置帖子内容格式：JSON。
						req.Header.Set(httpHeaderContentType, httpHeaderContentTypeJson)
					} else if 正则类.X是否匹配文本(httpRegexParamJson, params) {
						// 如果传入的参数类似 "name=value" 形式，则使用表单类型。
						req.Header.Set(httpHeaderContentType, httpHeaderContentTypeForm)
					}
				}
			}
		}
	}

	// Context.
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	// Custom header.
	if len(c.header) > 0 {
		for k, v := range c.header {
			req.Header.Set(k, v)
		}
	}
// 如果你想自定义请求的host值，设置req.Host是必要的。
// 如果header中的"Host"值不为空，则会使用该值。
	if reqHeaderHost := req.Header.Get(httpHeaderHost); reqHeaderHost != "" {
		req.Host = reqHeaderHost
	}
	// Custom Cookie.
	if len(c.cookies) > 0 {
		headerCookie := ""
		for k, v := range c.cookies {
			if len(headerCookie) > 0 {
				headerCookie += ";"
			}
			headerCookie += k + "=" + v
		}
		if len(headerCookie) > 0 {
			req.Header.Set(httpHeaderCookie, headerCookie)
		}
	}
	// HTTP基础认证
	if len(c.authUser) > 0 {
		req.SetBasicAuth(c.authUser, c.authPass)
	}
	return req, nil
}

// callRequest 函数使用给定的 http.Request 发送请求，并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭该响应对象。
func (c *Client) callRequest(req *http.Request) (resp *Response, err error) {
	resp = &Response{
		request: req,
	}
// Dump 功能.
// 请求体可用于转储
// 原始 HTTP 请求-响应过程.
	reqBodyContent, _ := io.ReadAll(req.Body)
	resp.requestBody = reqBodyContent
	for {
		req.Body = utils.NewReadCloser(reqBodyContent, false)
		if resp.Response, err = c.Do(req); err != nil {
			err = 错误类.X多层错误并格式化(err, `request failed`)
			// 当err不为nil时，response可能也不会为nil。
			if resp.Response != nil {
				_ = resp.Response.Body.Close()
			}
			if c.retryCount > 0 {
				c.retryCount--
				time.Sleep(c.retryInterval)
			} else {
				// return resp, err
				break
			}
		} else {
			break
		}
	}
	return resp, err
}
