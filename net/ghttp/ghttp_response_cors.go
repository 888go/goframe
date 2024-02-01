// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package ghttp
import (
	"net/http"
	"net/url"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// CORSOptions 是CORS功能的选项。
// 参见 https://www.w3.org/TR/cors/ 。
type CORSOptions struct {
	AllowDomain      []string // 用于允许来自自定义域名的请求
	AllowOrigin      string   // Access-Control-Allow-Origin：允许跨域请求的源，这是HTTP响应头的一部分，用于指示服务器允许哪些源发起跨域请求。在Go语言中设置该响应头可以实现跨域资源共享（CORS）。
	AllowCredentials string   // Access-Control-Allow-Credentials: 允许跨域请求时携带验证凭据（cookies, Authorization 头等）
	ExposeHeaders    string   // Access-Control-Expose-Headers：
	MaxAge           int      // Access-Control-Max-Age： 
// （该注释表示HTTP响应头中的Access-Control-Max-Age字段，用于指示预检请求（OPTIONS）的结果能够被缓存多久。）
// 设置浏览器对跨域资源共享（CORS）中间结果的最大缓存时间（单位为秒），即预检请求的有效期。
// 当值设为非零时，在此期间内再次进行相同的跨域请求将不再发送预检请求，直接使用第一次预检请求的结果。
	AllowMethods     string   // Access-Control-Allow-Methods：允许跨域请求的方法列表
	AllowHeaders     string   // Access-Control-Allow-Headers: 允许跨域请求中携带的自定义请求头
}

var (
// defaultAllowHeaders 是 CORS 的默认允许头信息。
// 它定义了另一个映射，以便提高查找头部键的性能。
	defaultAllowHeaders    = "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With"
	defaultAllowHeadersMap = make(map[string]struct{})
)

func init() {
	array := gstr.SplitAndTrim(defaultAllowHeaders, ",")
	for _, header := range array {
		defaultAllowHeadersMap[header] = struct{}{}
	}
}

// DefaultCORSOptions 返回默认的 CORS 选项，
// 这些选项允许任何跨域请求。
func (r *Response) DefaultCORSOptions() CORSOptions {
	options := CORSOptions{
		AllowOrigin:      "*",
		AllowMethods:     supportedHttpMethods,
		AllowCredentials: "true",
		AllowHeaders:     defaultAllowHeaders,
		MaxAge:           3628800,
	}
	// 默认情况下，允许所有客户端自定义头部。
	if headers := r.Request.Header.Get("Access-Control-Request-Headers"); headers != "" {
		array := gstr.SplitAndTrim(headers, ",")
		for _, header := range array {
			if _, ok := defaultAllowHeadersMap[header]; !ok {
				options.AllowHeaders += "," + header
			}
		}
	}
	// 默认情况下允许所有来源的任何请求
	if origin := r.Request.Header.Get("Origin"); origin != "" {
		options.AllowOrigin = origin
	} else if referer := r.Request.Referer(); referer != "" {
		if p := gstr.PosR(referer, "/", 6); p != -1 {
			options.AllowOrigin = referer[:p]
		} else {
			options.AllowOrigin = referer
		}
	}
	return options
}

// CORS 设置自定义 CORS 选项。
// 参见 https://www.w3.org/TR/cors/ 。
func (r *Response) CORS(options CORSOptions) {
	if r.CORSAllowedOrigin(options) {
		r.Header().Set("Access-Control-Allow-Origin", options.AllowOrigin)
	}
	if options.AllowCredentials != "" {
		r.Header().Set("Access-Control-Allow-Credentials", options.AllowCredentials)
	}
	if options.ExposeHeaders != "" {
		r.Header().Set("Access-Control-Expose-Headers", options.ExposeHeaders)
	}
	if options.MaxAge != 0 {
		r.Header().Set("Access-Control-Max-Age", gconv.String(options.MaxAge))
	}
	if options.AllowMethods != "" {
		r.Header().Set("Access-Control-Allow-Methods", options.AllowMethods)
	}
	if options.AllowHeaders != "" {
		r.Header().Set("Access-Control-Allow-Headers", options.AllowHeaders)
	}
// 如果请求为OPTIONS，则不进行继续服务处理。
// 注意，之前在路由搜索中有特殊的检查，
// 所以如果执行到这里，意味着已存在正在服务的处理程序。
	if gstr.Equal(r.Request.Method, "OPTIONS") {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}
		// 不再继续服务。
		r.Request.ExitAll()
	}
}

// CORSAllowedOrigin CORSAllowedOrigin函数检查当前请求的来源是否允许跨域。
func (r *Response) CORSAllowedOrigin(options CORSOptions) bool {
	if options.AllowDomain == nil {
		return true
	}
	origin := r.Request.Header.Get("Origin")
	if origin == "" {
		return true
	}
	parsed, err := url.Parse(origin)
	if err != nil {
		return false
	}
	for _, v := range options.AllowDomain {
		if gstr.IsSubDomain(parsed.Host, v) {
			return true
		}
	}
	return false
}

// CORSDefault 使用默认CORS选项设置CORS，
// 这将允许任何跨域请求。
func (r *Response) CORSDefault() {
	r.CORS(r.DefaultCORSOptions())
}
