// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"fmt"
	"net/http"
	"strings"
	"time"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gsession"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
	)
// Request 是一个请求的上下文对象。
type Request struct {
	*http.Request
	Server     *Server           // Server.
	Cookie     *Cookie           // Cookie.
	Session    *gsession.Session // Session.
	Response   *Response         // 此请求对应的响应。
	Router     *Router           // Matched Router for this request. Note that it's not available in HOOK handler.
	EnterTime  int64             // 请求开始时间（毫秒）
	LeaveTime  int64             // 请求结束时的时间（毫秒）。
	Middleware *middleware       // 中间件管理器。
	StaticFile *staticFile       // 静态文件对象，用于静态文件服务。

// =================================================================================================================
// 用于内部使用的私有属性。
// =================================================================================================================

	handlers        []*HandlerItemParsed   // 此请求所匹配的所有包含处理器、钩子和中间件的处理程序集合。
	serveHandler    *HandlerItemParsed     // 此请求的实际处理程序，非钩子或中间件。
	handlerResponse interface{}            // Handler响应对象，用于Request/Response处理器。
	hasHookHandler  bool                   // A bool marking whether there's hook handler in the handlers for performance purpose.
	hasServeHandler bool                   // A bool marking whether there's serving handler in the handlers for performance purpose.
	parsedQuery     bool                   // 一个布尔值，标记是否已解析GET参数。
	parsedBody      bool                   // 一个布尔值，标记请求体是否已解析。
	parsedForm      bool                   // A bool 标记是否已解析HTTP方法 PUT, POST, PATCH 的请求 Form。
	paramsMap       map[string]interface{} // 自定义参数映射。
	routerMap       map[string]string      // 路由参数映射表，如果没有路由参数，则可能为nil。
	queryMap        map[string]interface{} // Query parameters map, which is nil if there's no query string.
	formMap         map[string]interface{} // Form parameters map, which is nil if there's no form of data from the client.
	bodyMap         map[string]interface{} // Body参数映射，如果请求体没有内容，则可能为nil。
	error           error                  // 当前请求正在执行的错误
	exitAll         bool                   // 一个布尔值，标记当前请求是否已退出。
	parsedHost      string                 // GetHost 函数使用的当前主机解析后的主机名。
	clientIp        string                 // GetClientIp 函数当前使用的、已解析的客户端 IP 地址（针对当前主机）。
	bodyContent     []byte                 // 请求体内容。
	isFileRequest   bool                   // 一个布尔值，表示当前请求是否正在提供文件服务。
	viewObject      *gview.View            // 自定义模板视图引擎对象，用于此次响应。
	viewParams      gview.Params           // 为本次响应自定义模板视图变量。
	originUrlPath   string                 // 从客户端传递过来的原始URL路径。
}

// staticFile 是用于静态文件服务的文件结构体。
type staticFile struct {
	File  *gres.File // 资源文件对象。
	Path  string     // File path.
	IsDir bool       // Is directory.
}

// newRequest 创建并返回一个新的请求对象。
func newRequest(s *Server, r *http.Request, w http.ResponseWriter) *Request {
	request := &Request{
		Server:        s,
		Request:       r,
		Response:      newResponse(s, w),
		EnterTime:     gtime.TimestampMilli(),
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
	// 自定义会话ID生成函数。
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
	// 删除URI尾部的字符'/'
	if request.URL.Path != "/" {
		for len(request.URL.Path) > 0 && request.URL.Path[len(request.URL.Path)-1] == '/' {
			request.URL.Path = request.URL.Path[:len(request.URL.Path)-1]
		}
	}

	// 如果URI为空，则使用此默认URI值。
	if request.URL.Path == "" {
		request.URL.Path = "/"
	}
	return request
}

// WebSocket将当前请求升级为websocket请求。
// 如果升级成功，返回一个新的WebSocket对象；如果失败，则返回错误信息。
// 注意，该请求必须是websocket请求，否则升级必定会失败。
func (r *Request) WebSocket() (*WebSocket, error) {
	if conn, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil); err == nil {
		return &WebSocket{
			conn,
		}, nil
	} else {
		return nil, err
	}
}

// Exit 中断当前HTTP处理器的执行。
func (r *Request) Exit() {
	panic(exceptionExit)
}

// ExitAll 退出当前及后续HTTP处理器的执行。
func (r *Request) ExitAll() {
	r.exitAll = true
	panic(exceptionExitAll)
}

// ExitHook 结束当前及后续 HTTP HOOK 处理器的执行。
func (r *Request) ExitHook() {
	panic(exceptionExitHook)
}

// IsExited 检查并返回当前请求是否已退出。
func (r *Request) IsExited() bool {
	return r.exitAll
}

// GetHeader根据给定的`key`检索并返回头部值。
func (r *Request) GetHeader(key string) string {
	return r.Header.Get(key)
}

// GetHost 返回当前请求的主机名，该主机名可能是域名或不带端口号的IP地址。
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

// IsFileRequest 检查并返回当前请求是否正在提供文件服务。
func (r *Request) IsFileRequest() bool {
	return r.isFileRequest
}

// IsAjaxRequest 检查并返回当前请求是否为 AJAX 请求。
func (r *Request) IsAjaxRequest() bool {
	return strings.EqualFold(r.Header.Get("X-Requested-With"), "XMLHttpRequest")
}

// GetClientIp 返回该请求的客户端IP地址，不包含端口号。
// 注意：此IP地址可能已被客户端头部信息修改。
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

// GetRemoteIp 从 RemoteAddr 返回 IP 地址。
func (r *Request) GetRemoteIp() string {
	array, _ := gregex.MatchString(`(.+):(\d+)`, r.RemoteAddr)
	if len(array) > 1 {
		return strings.Trim(array[1], "[]")
	}
	return r.RemoteAddr
}

// GetUrl 返回当前请求的URL。
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

// GetSessionId 从cookie或header中检索并返回会话ID。
func (r *Request) GetSessionId() string {
	id := r.Cookie.GetSessionId()
	if id == "" {
		id = r.Header.Get(r.Server.GetSessionIdName())
	}
	return id
}

// GetReferer 返回该请求的引用来源。
func (r *Request) GetReferer() string {
	return r.Header.Get("Referer")
}

// GetError 返回在请求过程中发生的错误。
// 如果没有错误，它将返回 nil。
func (r *Request) GetError() error {
	return r.error
}

// SetError为当前请求设置自定义错误。
func (r *Request) SetError(err error) {
	r.error = err
}

// ReloadParam 用于修改请求参数。
// 有时，我们希望通过中间件来修改请求参数，但直接修改 Request.Body 是无效的，
// 所以它会清除 Request 中已解析标志，以使参数重新解析。
func (r *Request) ReloadParam() {
	r.parsedBody = false
	r.parsedForm = false
	r.parsedQuery = false
	r.bodyContent = nil
}

// GetHandlerResponse 获取并返回处理器响应对象及其错误信息。
func (r *Request) GetHandlerResponse() interface{} {
	return r.handlerResponse
}

// GetServeHandler 获取并返回用户自定义的用于处理当前请求的处理器。
func (r *Request) GetServeHandler() *HandlerItemParsed {
	return r.serveHandler
}
