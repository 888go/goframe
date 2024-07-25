// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

// Request 是请求的上下文对象。 md5:90c1e82eacf87b05
type Request struct {
	*http.Request
	Server     *Server           // Server.
	Cookie     *Cookie           // Cookie.
	Session    *gsession.Session // Session.
	Response   *Response         // 对应于此请求的响应。 md5:aca9c7c1b8de7fd3
	Router     *Router           // Matched Router for this request. Note that it's not available in HOOK handler.
	EnterTime  *gtime.Time       // 请求开始时间（以毫秒为单位）。 md5:c8ed1608de735520
	LeaveTime  *gtime.Time       // 结束时间（以毫秒为单位）的请求。 md5:d10eec1013ccf8ce
	Middleware *middleware       // Middleware manager.
	StaticFile *staticFile       // 用于静态文件服务的静态文件对象。 md5:c9ba568602380c10

	// =================================================================================================================
	// 用于内部使用的私有属性。
	// =================================================================================================================
	// md5:5edc202a134e48b1

	handlers        []*HandlerItemParsed   // 包含此请求的处理器、挂钩和中间件的所有匹配处理程序。 md5:42382b62d60509ae
	serveHandler    *HandlerItemParsed     // 为这次请求的实际处理程序，不是钩子或中间件。 md5:60c0eea59027a44f
	handlerResponse interface{}            // 用于请求/响应处理程序的处理器响应对象。 md5:adefb98d0b335dd6
	hasHookHandler  bool                   // A bool marking whether there's hook handler in the handlers for performance purpose.
	hasServeHandler bool                   // A bool marking whether there's serving handler in the handlers for performance purpose.
	parsedQuery     bool                   // 一个布尔值，表示是否已经解析了GET参数。 md5:08f143c93bb48593
	parsedBody      bool                   // 一个标记请求体是否已解析的布尔值。 md5:f14ff93f821a596e
	parsedForm      bool                   // 一个布尔值，表示HTTP方法为PUT、POST或PATCH时，请求表单是否已解析。 md5:b372893681fa6b1a
	paramsMap       map[string]interface{} // Custom parameters map.
	routerMap       map[string]string      // 路由参数映射，如果没有路由参数，则可能为nil。 md5:1f669e8d65cc7201
	queryMap        map[string]interface{} // Query parameters map, which is nil if there's no query string.
	formMap         map[string]interface{} // Form parameters map, which is nil if there's no form of data from the client.
	bodyMap         map[string]interface{} // Body参数映射，如果内容为空则可能为nil。 md5:318eb5ce8cbb633b
	error           error                  // 请求当前执行的错误。 md5:fb1258df50e5180f
	exitAll         bool                   // 一个标记当前请求是否已退出的布尔值。 md5:df96d943707fa105
	parsedHost      string                 // 由GetHost函数使用的当前主机的解析后的主机名。 md5:f8aa264e6a28b762
	clientIp        string                 // 当前主机中由GetClientIp函数使用的解析后的客户端IP。 md5:de8953e109404d72
	bodyContent     []byte                 // Request body content.
	isFileRequest   bool                   // 一个bool标记当前请求是否是文件服务。 md5:71668278a657d9c6
	viewObject      *gview.View            // 为这次响应自定义的模板视图引擎对象。 md5:ddd788a8e05477da
	viewParams      gview.Params           // 为此响应定制的模板视图变量。 md5:4138dc353e5f6967
	originUrlPath   string                 // 从客户端传递的原始URL路径。 md5:c7368bb1d20fddcd
}

// staticFile是静态文件服务的文件结构体。 md5:1a45356e421cf5d2
type staticFile struct {
	File  *gres.File // Resource file object.
	Path  string     // File path.
	IsDir bool       // Is directory.
}

// newRequest 创建并返回一个新的请求对象。 md5:bbe6326af48b82f2
func newRequest(s *Server, r *http.Request, w http.ResponseWriter) *Request {
	request := &Request{
		Server:        s,
		Request:       r,
		Response:      newResponse(s, w),
		EnterTime:     gtime.Now(),
		originUrlPath: r.URL.Path,
	}
	request.Cookie = GetCookie(request)
	request.Session = s.sessionManager.New(
		r.Context(),
		request.GetSessionId(),
	)
	request.Response.Request = request
	request.Middleware = &middleware{
		request: request,
	}
		// 自定义会话ID生成函数。 md5:1530052018d41784
	err := request.Session.SetIdFunc(func(ttl time.Duration) string {
		var (
			address = request.RemoteAddr
			header  = fmt.Sprintf("%v", request.Header)
		)
		intlog.Print(r.Context(), address, header)
		return guid.S([]byte(address), []byte(header))
	})
	if err != nil {
		panic(err)
	}
		// 删除URI尾部的'/'字符。 md5:5cd878ba748e3629
	if request.URL.Path != "/" {
		for len(request.URL.Path) > 0 && request.URL.Path[len(request.URL.Path)-1] == '/' {
			request.URL.Path = request.URL.Path[:len(request.URL.Path)-1]
		}
	}

		// 如果为空，默认的URI值。 md5:ba9367be3b98edbd
	if request.URL.Path == "" {
		request.URL.Path = "/"
	}
	return request
}

// 将当前请求升级为WebSocket请求。
// 如果成功，返回一个新的WebSocket对象，如果失败，则返回错误。
// 注意，请求必须是WebSocket请求，否则升级肯定会失败。
// 
// 警告：将来将被移除，请使用第三方WebSocket库代替。
// md5:68ad17d4740330b2
func (r *Request) WebSocket() (*WebSocket, error) {
	if conn, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil); err == nil {
		return &WebSocket{
			conn,
		}, nil
	} else {
		return nil, err
	}
}

// Exit 退出当前HTTP处理器的执行。 md5:3a3298adda39cc74
func (r *Request) Exit() {
	panic(exceptionExit)
}

// ExitAll 退出当前及后续的HTTP处理器执行。 md5:53932e5e1bdd10d5
func (r *Request) ExitAll() {
	r.exitAll = true
	panic(exceptionExitAll)
}

// ExitHook 退出当前及后续HTTP钩子处理器的执行。 md5:ef92857b0e046888
func (r *Request) ExitHook() {
	panic(exceptionExitHook)
}

// IsExited 检查并返回当前请求是否已退出。 md5:9198deaaaf14733a
func (r *Request) IsExited() bool {
	return r.exitAll
}

// GetHeader 获取并返回给定`key`对应的头值。 md5:3088bb7beaf8a754
func (r *Request) GetHeader(key string) string {
	return r.Header.Get(key)
}

// GetHost 返回当前请求的主机名，可能是不带端口的域名或IP。 md5:3a06fa36ddefd149
func (r *Request) GetHost() string {
	if len(r.parsedHost) == 0 {
		array, _ := gregex.MatchString(`(.+):(\d+)`, r.Host)
		if len(array) > 1 {
			r.parsedHost = array[1]
		} else {
			r.parsedHost = r.Host
		}
	}
	return r.parsedHost
}

// IsFileRequest 检查并返回当前请求是否是为文件服务的。 md5:a849769abec62994
func (r *Request) IsFileRequest() bool {
	return r.isFileRequest
}

// IsAjaxRequest 检查并返回当前请求是否为AJAX请求。 md5:c4e5c9eb4c13dae7
func (r *Request) IsAjaxRequest() bool {
	return strings.EqualFold(r.Header.Get("X-Requested-With"), "XMLHttpRequest")
}

// GetClientIp 返回此请求的客户端IP（不包含端口号）。
// 注意，此IP地址可能已被客户端头部信息修改。
// md5:54dc4a1d4744646a
func (r *Request) GetClientIp() string {
	if r.clientIp != "" {
		return r.clientIp
	}
	realIps := r.Header.Get("X-Forwarded-For")
	if realIps != "" && len(realIps) != 0 && !strings.EqualFold("unknown", realIps) {
		ipArray := strings.Split(realIps, ",")
		r.clientIp = ipArray[0]
	}
	if r.clientIp == "" {
		r.clientIp = r.Header.Get("Proxy-Client-IP")
	}
	if r.clientIp == "" {
		r.clientIp = r.Header.Get("WL-Proxy-Client-IP")
	}
	if r.clientIp == "" {
		r.clientIp = r.Header.Get("HTTP_CLIENT_IP")
	}
	if r.clientIp == "" {
		r.clientIp = r.Header.Get("HTTP_X_FORWARDED_FOR")
	}
	if r.clientIp == "" {
		r.clientIp = r.Header.Get("X-Real-IP")
	}
	if r.clientIp == "" {
		r.clientIp = r.GetRemoteIp()
	}
	return r.clientIp
}

// GetRemoteIp 从 RemoteAddr 中返回 IP 地址。 md5:fca642ffe8c25d8c
func (r *Request) GetRemoteIp() string {
	array, _ := gregex.MatchString(`(.+):(\d+)`, r.RemoteAddr)
	if len(array) > 1 {
		return strings.Trim(array[1], "[]")
	}
	return r.RemoteAddr
}

// GetSchema 返回此请求的架构。 md5:7bbb51fb51117978
func (r *Request) GetSchema() string {
	var (
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	return scheme
}

// GetUrl 返回本次请求的当前URL。 md5:8be855812fe4346f
func (r *Request) GetUrl() string {
	var (
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
	)

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	return fmt.Sprintf(`%s://%s%s`, scheme, r.Host, r.URL.String())
}

// GetSessionId 从cookie或header中检索并返回会话ID。 md5:06fb7350f9f5597f
func (r *Request) GetSessionId() string {
	id := r.Cookie.GetSessionId()
	if id == "" {
		id = r.Header.Get(r.Server.GetSessionIdName())
	}
	return id
}

// GetReferer 返回此请求的引荐来源。 md5:4684519c6f7dc2c0
func (r *Request) GetReferer() string {
	return r.Header.Get("Referer")
}

// GetError 返回请求过程中发生的错误。如果没有错误，它将返回 nil。
// md5:035e49a2662f9c04
func (r *Request) GetError() error {
	return r.error
}

// SetError 为当前请求设置自定义错误。 md5:025f3a0817a4be99
func (r *Request) SetError(err error) {
	r.error = err
}

// ReloadParam 用于修改请求参数。
// 有时，我们希望通过中间件来修改请求参数，但直接修改 Request.Body 是无效的，
// 因此它会清除 Request 的已解析标记，使得参数能够被重新解析。
// md5:af7ff26c797683ef
func (r *Request) ReloadParam() {
	r.parsedBody = false
	r.parsedForm = false
	r.parsedQuery = false
	r.bodyContent = nil
}

// GetHandlerResponse 获取并返回处理器响应对象及其错误信息。 md5:d6ef2cb1d4fef297
func (r *Request) GetHandlerResponse() interface{} {
	return r.handlerResponse
}

// GetServeHandler 获取并返回用于处理此请求的用户定义的处理器。 md5:6aef7d779ee52097
func (r *Request) GetServeHandler() *HandlerItemParsed {
	return r.serveHandler
}
