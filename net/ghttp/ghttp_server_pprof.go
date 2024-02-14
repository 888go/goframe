// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

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
func PProf服务端创建(监听端口 int, 作废参数 ...string) {
	s := X取服务对象(defaultPProfServerName)
	s.PProf开启()
	s.X设置监听端口(监听端口)
	s.X启动服务()
}

// EnablePProf 启用服务器的 PProf 功能。
func (s *X服务) PProf开启(路由地址 ...string) {
	s.X创建域名路由(DefaultDomainName).PProf开启(路由地址...)
}

// EnablePProf 启用指定域名服务器的 PProf 功能。
func (d *X域名路由) PProf开启(路由地址 ...string) {
	p := defaultPProfPattern
	if len(路由地址) > 0 && 路由地址[0] != "" {
		p = 路由地址[0]
	}
	up := &utilPProf{}
	_, _, uri, _ := d.server.parsePattern(p)
	uri = strings.TrimRight(uri, "/")
	d.X创建分组路由(uri, func(group *X分组路由) {
		group.X绑定所有类型("/*action", up.X显示页面)
		group.X绑定所有类型("/cmdline", up.Cmdline)
		group.X绑定所有类型("/profile", up.Profile)
		group.X绑定所有类型("/symbol", up.Symbol)
		group.X绑定所有类型("/trace", up.Trace)
	})
}

// Index 显示 PProf 索引页面。
func (p *utilPProf) X显示页面(r *X请求) {
	var (
		ctx      = r.Context别名()
		profiles = runpprof.Profiles()
		action   = r.Get别名("action").String()
		data     = map[string]interface{}{
			"uri":      strings.TrimRight(r.URL.Path, "/") + "/",
			"profiles": profiles,
		}
	)
	if len(action) == 0 {
		buffer, _ := 模板类.ParseContent(r.Context别名(), `
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
		r.X响应.X写响应缓冲区(buffer)
		return
	}
	for _, p := range profiles {
		if p.Name() == action {
			if err := p.WriteTo(r.X响应.Writer, r.X取参数("debug").X取整数()); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			break
		}
	}
}

// Cmdline 函数响应运行程序的命令行参数，其中各个参数由 NUL 字节分隔。
// 包初始化时会将其注册为 /debug/pprof/cmdline 路径。
func (p *utilPProf) Cmdline(r *X请求) {
	netpprof.Cmdline(r.X响应.Writer, r.Request)
}

// Profile 函数响应 pprof 格式的 CPU 分析报告。
// 分析的持续时间由 GET 参数中指定的秒数决定，如果未指定，则默认为 30 秒。
// 包初始化时会将其注册为 /debug/pprof/profile 路径。
func (p *utilPProf) Profile(r *X请求) {
	netpprof.Profile(r.X响应.Writer, r.Request)
}

// Symbol 函数通过查询请求中列出的程序计数器，
// 并以程序计数器到函数名称的映射表作为响应。
// 包初始化时将其注册为 /debug/pprof/symbol 路径。
func (p *utilPProf) Symbol(r *X请求) {
	netpprof.Symbol(r.X响应.Writer, r.Request)
}

// Trace 以二进制形式响应执行跟踪。
// 跟踪持续时间由GET参数中指定的秒数决定，如果未指定，则持续1秒。
// 包初始化时将其注册为/debug/pprof/trace。
func (p *utilPProf) Trace(r *X请求) {
	netpprof.Trace(r.X响应.Writer, r.Request)
}
