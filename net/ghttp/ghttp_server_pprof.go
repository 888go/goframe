// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	netpprof "net/http/pprof"
	runpprof "runtime/pprof"
	"strings"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gview"
)

// utilPProf是实现PProf接口的结构。. md5:61c1485646c2e81a
type utilPProf struct{}

const (
	defaultPProfServerName = "pprof-server"
	defaultPProfPattern    = "/debug/pprof"
)

// StartPProfServer 启动并运行一个新的pprof服务器。. md5:4c0c47dfda03a84b
func StartPProfServer(port int, pattern ...string) {
	s := GetServer(defaultPProfServerName)
	s.EnablePProf(pattern...)
	s.SetPort(port)
	s.Run()
}

// EnablePProf 启用服务器的PProf功能。. md5:5603a60f147574d1
func (s *Server) EnablePProf(pattern ...string) {
	s.Domain(DefaultDomainName).EnablePProf(pattern...)
}

// EnablePProf 为指定域名的服务器启用 PProf 功能。. md5:46c19e5f1d55beb1
func (d *Domain) EnablePProf(pattern ...string) {
	p := defaultPProfPattern
	if len(pattern) > 0 && pattern[0] != "" {
		p = pattern[0]
	}
	up := &utilPProf{}
	_, _, uri, _ := d.server.parsePattern(p)
	uri = strings.TrimRight(uri, "/")
	d.Group(uri, func(group *RouterGroup) {
		group.ALL("/*action", up.Index)
		group.ALL("/cmdline", up.Cmdline)
		group.ALL("/profile", up.Profile)
		group.ALL("/symbol", up.Symbol)
		group.ALL("/trace", up.Trace)
	})
}

// Index 显示 PProf 的索引页面。. md5:606e9224f8418b6e
func (p *utilPProf) Index(r *Request) {
	var (
		ctx      = r.Context()
		profiles = runpprof.Profiles()
		action   = r.Get("action").String()
		data     = map[string]interface{}{
			"uri":      strings.TrimRight(r.URL.Path, "/") + "/",
			"profiles": profiles,
		}
	)
	if len(action) == 0 {
		buffer, _ := gview.ParseContent(r.Context(), `
            <html>
            <head>
                <title>GoFrame PProf</title>
            </head>
            {{$uri := .uri}}
            <body>
                profiles:<br>
                <table>
                    {{range .profiles}}
						<tr>
							<td align=right>{{.Count}}</td>
							<td><a href="{{$uri}}{{.Name}}?debug=1">{{.Name}}</a></td>
						<tr>
					{{end}}
                </table>
                <br><a href="{{$uri}}goroutine?debug=2">full goroutine stack dump</a><br>
            </body>
            </html>
            `, data)
		r.Response.Write(buffer)
		return
	}
	for _, p := range profiles {
		if p.Name() == action {
			if err := p.WriteTo(r.Response.Writer, r.GetRequest("debug").Int()); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			break
		}
	}
}

// Cmdline 响应正在运行程序的命令行，参数之间用 NULL 字节分隔。包初始化时将其注册为 /debug/pprof/cmdline。
// md5:35f5d246119cca43
func (p *utilPProf) Cmdline(r *Request) {
	netpprof.Cmdline(r.Response.Writer, r.Request)
}

// Profile 使用pprof格式返回CPU profiling信息。
// 如果GET参数指定了持续时间，那么 profiling 将持续该秒数；如果没有指定，则默认为30秒。
// 在包初始化时，它会注册为 "/debug/pprof/profile"。
// md5:11bd281949c0ba3c
func (p *utilPProf) Profile(r *Request) {
	netpprof.Profile(r.Response.Writer, r.Request)
}

// Symbol 查找请求中列出的程序计数器，
// 并以映射表的形式响应，该映射表将程序计数器与函数名称关联起来。
// 包初始化时将其注册为 /debug/pprof/symbol 路由。
// md5:2944ed5cfe9e0c52
func (p *utilPProf) Symbol(r *Request) {
	netpprof.Symbol(r.Response.Writer, r.Request)
}

// Trace 返回执行跟踪的二进制形式。
// 跟踪将持续指定的GET参数中的秒数，如果没有指定，则为1秒。
// 包初始化时将其注册为/debug/pprof/trace。
// md5:02830b4c9b48681f
func (p *utilPProf) Trace(r *Request) {
	netpprof.Trace(r.Response.Writer, r.Request)
}
