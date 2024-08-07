// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package http类

import (
	"net/http"
	"net/url"

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// CORSOptions 是 CORS 功能的选项。
// 参见 https://www.w3.org/TR/cors/ 。
// md5:86678849c932cd8d
type CORSOptions struct {
	AllowDomain      []string // 用于允许来自自定义域名的请求. md5:3050713aeb6de06c
	AllowOrigin      string   // Access-Control-Allow-Origin 是一个HTTP响应头字段，用于指定浏览器在进行跨域请求时可以访问资源的源（Origin）。 md5:64b1bac364c85a72
	AllowCredentials string   // 允许凭证访问控制. md5:9413744affc62151
	ExposeHeaders    string   // Access-Control-Expose-Headers：这是HTTP响应头的一部分，用于指定哪些响应头可以被CORS（跨源资源共享）策略允许从服务器传递到客户端。 md5:edb32baedc37800f
	MaxAge           int      // Access-Control-Max-Age
	AllowMethods     string   // 跨域资源共享允许的方法. md5:c78ddd1745514f4a
	AllowHeaders     string   // Access-Control-Allow-Headers：允许的头部字段. md5:9812fd3132d19ca6
}

var (
	// defaultAllowHeaders 是CORS默认允许的头信息。
	// 为了提高头部键搜索性能，我们定义了另一个映射。
	// md5:e6a13ea98879b3e6
	defaultAllowHeaders    = "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With"
	defaultAllowHeadersMap = make(map[string]struct{})
)

func init() {
	array := gstr.X分割并忽略空值(defaultAllowHeaders, ",")
	for _, header := range array {
		defaultAllowHeadersMap[header] = struct{}{}
	}
}

// X取跨域默认选项 返回默认的 CORS 选项，它允许任何跨域请求。
// md5:ed45ce5e88088eac
func (r *Response) X取跨域默认选项() CORSOptions {
	options := CORSOptions{
		AllowOrigin:      "*",
		AllowMethods:     supportedHttpMethods,
		AllowCredentials: "true",
		AllowHeaders:     defaultAllowHeaders,
		MaxAge:           3628800,
	}
		// 默认允许客户端的所有自定义头部。 md5:5aa0a6d974ed81b6
	if headers := r.Request.Header.Get("Access-Control-Request-Headers"); headers != "" {
		array := gstr.X分割并忽略空值(headers, ",")
		for _, header := range array {
			if _, ok := defaultAllowHeadersMap[header]; !ok {
				options.AllowHeaders += "," + header
			}
		}
	}
		// 默认允许所有来源的访问。 md5:bd5e36856694e82f
	if origin := r.Request.Header.Get("Origin"); origin != "" {
		options.AllowOrigin = origin
	} else if referer := r.Request.Referer(); referer != "" {
		if p := gstr.X倒找(referer, "/", 6); p != -1 {
			options.AllowOrigin = referer[:p]
		} else {
			options.AllowOrigin = referer
		}
	}
	return options
}

// X跨域请求设置 设置自定义X跨域请求设置选项。
// 参见 https://www.w3.org/TR/cors/ 。
// md5:5ace1c84086a260a
func (r *Response) X跨域请求设置(跨域选项 CORSOptions) {
	if r.X是否允许跨域(跨域选项) {
		r.Header().Set("Access-Control-Allow-Origin", 跨域选项.AllowOrigin)
	}
	if 跨域选项.AllowCredentials != "" {
		r.Header().Set("Access-Control-Allow-Credentials", 跨域选项.AllowCredentials)
	}
	if 跨域选项.ExposeHeaders != "" {
		r.Header().Set("Access-Control-Expose-Headers", 跨域选项.ExposeHeaders)
	}
	if 跨域选项.MaxAge != 0 {
		r.Header().Set("Access-Control-Max-Age", gconv.String(跨域选项.MaxAge))
	}
	if 跨域选项.AllowMethods != "" {
		r.Header().Set("Access-Control-Allow-Methods", 跨域选项.AllowMethods)
	}
	if 跨域选项.AllowHeaders != "" {
		r.Header().Set("Access-Control-Allow-Headers", 跨域选项.AllowHeaders)
	}
	// 如果请求是OPTIONS类型，不继续服务处理。
	// 注意，之前的路由器搜索中已经有特殊检查，
	// 所以如果到达这里，意味着已经存在正在处理的处理器。
	// md5:178e6bee651f512f
	if gstr.X相等比较并忽略大小写(r.Request.Method, "OPTIONS") {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}
		// No continue serving.
		r.Request.X退出全部()
	}
}

// X是否允许跨域 CORSAllowed 检查当前请求的来源是否被允许进行跨域。 md5:599a140b617c5c1c
func (r *Response) X是否允许跨域(跨域选项 CORSOptions) bool {
	if 跨域选项.AllowDomain == nil {
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
	for _, v := range 跨域选项.AllowDomain {
		if gstr.X是否为子域名(parsed.Host, v) {
			return true
		}
	}
	return false
}

// X跨域请求全允许 使用默认的 CORS 选项设置 CORS，
// 允许任何跨域请求。
// md5:2808119e534c338a
func (r *Response) X跨域请求全允许() {
	r.X跨域请求设置(r.X取跨域默认选项())
}
