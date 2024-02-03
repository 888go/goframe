// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	netpprof "net/http/pprof"
	runpprof "runtime/pprof"
	"strings"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gview"
)

// utilPProf 是 PProf 接口的实现者。
type utilPProf struct{}

const (
	defaultPProfServerName = "pprof-server"
	defaultPProfPattern    = "/debug/pprof"
)

// StartPProfServer 启动并运行一个新的 pprof 服务端。
func StartPProfServer(port int, pattern ...string) {
	s := GetServer(defaultPProfServerName)
	s.EnablePProf()
	s.SetPort(port)
	s.Run()
}

// EnablePProf 启用服务器的 PProf 功能。
func (s *Server) EnablePProf(pattern ...string) {
	s.Domain(DefaultDomainName).EnablePProf(pattern...)
}

// EnablePProf 启用指定域名服务器的 PProf 功能。
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

// Index 显示 PProf 索引页面。
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

// Cmdline 函数响应运行程序的命令行参数，其中各个参数由 NUL 字节分隔。
// 包初始化时会将其注册为 /debug/pprof/cmdline 路径。
func (p *utilPProf) Cmdline(r *Request) {
	netpprof.Cmdline(r.Response.Writer, r.Request)
}

// Profile 函数响应 pprof 格式的 CPU 分析报告。
// 分析的持续时间由 GET 参数中指定的秒数决定，如果未指定，则默认为 30 秒。
// 包初始化时会将其注册为 /debug/pprof/profile 路径。
func (p *utilPProf) Profile(r *Request) {
	netpprof.Profile(r.Response.Writer, r.Request)
}

// Symbol 函数通过查询请求中列出的程序计数器，
// 并以程序计数器到函数名称的映射表作为响应。
// 包初始化时将其注册为 /debug/pprof/symbol 路径。
func (p *utilPProf) Symbol(r *Request) {
	netpprof.Symbol(r.Response.Writer, r.Request)
}

// Trace 以二进制形式响应执行跟踪。
// 跟踪持续时间由GET参数中指定的秒数决定，如果未指定，则持续1秒。
// 包初始化时将其注册为/debug/pprof/trace。
func (p *utilPProf) Trace(r *Request) {
	netpprof.Trace(r.Response.Writer, r.Request)
}
