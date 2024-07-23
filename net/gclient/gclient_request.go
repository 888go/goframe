// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/httputil"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Get 发送GET请求并返回响应对象。
// 请注意，如果响应对象将永远不会被使用，必须关闭它。
// md5:bf82e1e2c38506f6
func (c *Client) Get(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodGet, url, data...)
}

// 发送PUT请求并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。
// md5:44e5f3e5edebbb91
func (c *Client) Put(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodPut, url, data...)
}

// Post 使用HTTP POST方法发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。
// md5:9ba8d1283ba032cb
func (c *Client) Post(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodPost, url, data...)
}

// Delete 发送DELETE请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。
// md5:4dde007718fff7a6
func (c *Client) Delete(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodDelete, url, data...)
}

// Head 发送HEAD请求并返回响应对象。
// 请注意，如果响应对象不会被使用，必须关闭它。
// md5:400dd3a80c3a0ccb
func (c *Client) Head(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodHead, url, data...)
}

// Patch 发送PATCH请求并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。
// md5:4e530560a87457a1
func (c *Client) Patch(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodPatch, url, data...)
}

// Connect 发送CONNECT请求并返回响应对象。
// 注意，如果响应对象永远不会使用，必须关闭它。
// md5:cb5555f2c2a7a29d
func (c *Client) Connect(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodConnect, url, data...)
}

// Options 发送OPTIONS请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。
// md5:3a2d4fbe5e9f5e31
func (c *Client) Options(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodOptions, url, data...)
}

// Trace 发送TRACE请求并返回响应对象。
// 请注意，如果响应对象将永远不会被使用，必须关闭它。
// md5:82834b110d843156
func (c *Client) Trace(ctx context.Context, url string, data ...interface{}) (*Response, error) {
	return c.DoRequest(ctx, http.MethodTrace, url, data...)
}

// PostForm 与 net/http.PostForm 不同。它是一个 Post 方法的包装器，会将 Content-Type 设置为 "multipart/form-data;"。
// 它会自动为请求体和 Content-Type 设置边界字符。
// 
// 类似于下面的情况：
//
// Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN
//
// 表单数据如下：
// ------Boundarye4Ghaog6giyQ9ncN
// Content-Disposition: form-data; name="checkType"
//
// none
// 
// 它用于发送表单数据。请注意，如果响应对象永远不会使用，必须关闭它。
// md5:bd2237aaca8f2a89
func (c *Client) PostForm(ctx context.Context, url string, data map[string]string) (resp *Response, err error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range data {
		err := w.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return c.ContentType(w.FormDataContentType()).Post(ctx, url, body)
}

// DoRequest 发送给定HTTP方法和数据的请求，并返回响应对象。
// 注意，如果响应对象永远不会使用，必须关闭它。
// 
// 如果包含文件上传，它将使用"multipart/form-data"作为Content-Type，否则使用"application/x-www-form-urlencoded"。它还能自动检测POST内容的JSON格式，并为此自动设置Content-Type为"application/json"。
// md5:09c1fd65446e9a2e
func (c *Client) DoRequest(
	ctx context.Context, method, url string, data ...interface{},
) (resp *Response, err error) {
	var requestStartTime = gtime.Now()
	req, err := c.prepareRequest(ctx, method, url, data...)
	if err != nil {
		return nil, err
	}

	// Metrics.
	c.handleMetricsBeforeRequest(req)
	defer c.handleMetricsAfterRequestDone(req, requestStartTime)

	// Client middleware.
	if len(c.middlewareHandler) > 0 {
		mdlHandlers := make([]HandlerFunc, 0, len(c.middlewareHandler)+1)
		mdlHandlers = append(mdlHandlers, c.middlewareHandler...)
		mdlHandlers = append(mdlHandlers, func(cli *Client, r *http.Request) (*Response, error) {
			return cli.callRequest(r)
		})
		ctx = context.WithValue(req.Context(), clientMiddlewareKey, &clientMiddleware{
			client:       c,
			handlers:     mdlHandlers,
			handlerIndex: -1,
		})
		req = req.WithContext(ctx)
		resp, err = c.Next(req)
	} else {
		resp, err = c.callRequest(req)
	}
	if resp != nil && resp.Response != nil {
		req.Response = resp.Response
	}
	return resp, err
}

// prepareRequest 验证请求参数，构建并返回http请求。 md5:e955238a4d45cf59
func (c *Client) prepareRequest(ctx context.Context, method, url string, data ...interface{}) (req *http.Request, err error) {
	method = strings.ToUpper(method)
	if len(c.prefix) > 0 {
		url = c.prefix + gstr.Trim(url)
	}
	if !gstr.ContainsI(url, httpProtocolName) {
		url = httpProtocolName + `://` + url
	}
	var params string
	if len(data) > 0 {
		switch c.header[httpHeaderContentType] {
		case httpHeaderContentTypeJson:
			switch data[0].(type) {
			case string, []byte:
				params = gconv.String(data[0])
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
				params = gconv.String(data[0])
			default:
				if b, err := gjson.New(data[0]).ToXml(); err != nil {
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
				// 如果HTTP方法为GET且未指定Content-Type时，它将参数追加到URL中。
				// md5:a6325a5bd7f8b355
				if gstr.Contains(url, "?") {
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
			err = gerror.Wrapf(err, `http.NewRequest failed with method "%s" and URL "%s"`, method, url)
			return nil, err
		}
	} else {
		if strings.Contains(params, httpParamFileHolder) {
			// 上传文件请求。 md5:7975fe0b1475ea53
			var (
				buffer = bytes.NewBuffer(nil)
				writer = multipart.NewWriter(buffer)
			)
			for _, item := range strings.Split(params, "&") {
				array := strings.Split(item, "=")
				if len(array[1]) > 6 && strings.Compare(array[1][0:6], httpParamFileHolder) == 0 {
					path := array[1][6:]
					if !gfile.Exists(path) {
						return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `"%s" does not exist`, path)
					}
					var (
						file          io.Writer
						formFileName  = gfile.Basename(path)
						formFieldName = array[0]
					)
					if file, err = writer.CreateFormFile(formFieldName, formFileName); err != nil {
						err = gerror.Wrapf(err, `CreateFormFile failed with "%s", "%s"`, formFieldName, formFileName)
						return nil, err
					} else {
						var f *os.File
						if f, err = gfile.Open(path); err != nil {
							return nil, err
						}
						if _, err = io.Copy(file, f); err != nil {
							err = gerror.Wrapf(err, `io.Copy failed from "%s" to form "%s"`, path, formFieldName)
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
						err = gerror.Wrapf(err, `write form field failed with "%s", "%s"`, fieldName, fieldValue)
						return nil, err
					}
				}
			}
			// Close完成multipart消息并写入尾部边界结束行到输出。
			// md5:0a89f03d075fd9ee
			if err = writer.Close(); err != nil {
				err = gerror.Wrapf(err, `form writer close failed`)
				return nil, err
			}

			if req, err = http.NewRequest(method, url, buffer); err != nil {
				err = gerror.Wrapf(err, `http.NewRequest failed for method "%s" and URL "%s"`, method, url)
				return nil, err
			} else {
				req.Header.Set(httpHeaderContentType, writer.FormDataContentType())
			}
		} else {
			// Normal request.
			paramBytes := []byte(params)
			if req, err = http.NewRequest(method, url, bytes.NewReader(paramBytes)); err != nil {
				err = gerror.Wrapf(err, `http.NewRequest failed for method "%s" and URL "%s"`, method, url)
				return nil, err
			} else {
				if v, ok := c.header[httpHeaderContentType]; ok {
					// Custom Content-Type.
					req.Header.Set(httpHeaderContentType, v)
				} else if len(paramBytes) > 0 {
					if (paramBytes[0] == '[' || paramBytes[0] == '{') && json.Valid(paramBytes) {
						// 自动检测并设置帖子内容格式：JSON。 md5:735d9fcd3200585a
						req.Header.Set(httpHeaderContentType, httpHeaderContentTypeJson)
					} else if gregex.IsMatchString(httpRegexParamJson, params) {
						// 如果传递的参数形如 "name=value"，则使用表单类型。 md5:2f5188c0993569a1
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
	// 如果你想要自定义请求的主机值，那么设置 `req.Host` 是必要的。
	// 如果 `Host` 头部不为空，它会使用头部的 "Host" 值。
	// md5:e71cb70a52453d4c
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
	// HTTP基本身份验证。 md5:f6fdca448f00def0
	if len(c.authUser) > 0 {
		req.SetBasicAuth(c.authUser, c.authPass)
	}
	return req, nil
}

// callRequest 使用给定的http.Request发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。
// md5:d6e9d1e1953e082b
func (c *Client) callRequest(req *http.Request) (resp *Response, err error) {
	resp = &Response{
		request: req,
	}
	// Dump 功能。
	// 请求体可以用于转储
	// 原始HTTP请求-响应过程。
	// md5:57f6d6cec0adad22
	reqBodyContent, _ := io.ReadAll(req.Body)
	resp.requestBody = reqBodyContent
	for {
		req.Body = utils.NewReadCloser(reqBodyContent, false)
		if resp.Response, err = c.Do(req); err != nil {
			err = gerror.Wrapf(err, `request failed`)
			// 当err不为nil时，response可能不为nil。 md5:30e2b1a262fbd8ac
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
