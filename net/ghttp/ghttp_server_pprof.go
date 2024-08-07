// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	netpprof "net/http/pprof"
	runpprof "runtime/pprof"
	"strings"

	"github.com/888go/goframe/internal/intlog"
	gview "github.com/888go/goframe/os/gview"
)

// utilPProf是实现PProf接口的结构。 md5:61c1485646c2e81a
type utilPProf struct{}

const (
	defaultPProfServerName = "pprof-server"
	defaultPProfPattern    = "/debug/pprof"
)

// PProf服务端创建 启动并运行一个新的pprof服务器。 md5:4c0c47dfda03a84b
func PProf服务端创建(监听端口 int, 作废参数 ...string) {
	s := X取服务对象(defaultPProfServerName)
	s.PProf开启(作废参数...)
	s.X设置监听端口(监听端口)
	s.X启动服务()
}

// PProf开启 启用服务器的PProf功能。 md5:5603a60f147574d1
func (s *X服务) PProf开启(路由地址 ...string) {
	s.X创建域名路由(DefaultDomainName).PProf开启(路由地址...)
}

// PProf开启 为指定域名的服务器启用 PProf 功能。 md5:46c19e5f1d55beb1
func (d *Domain) PProf开启(路由地址 ...string) {
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

// X显示页面 显示 PProf 的索引页面。 md5:606e9224f8418b6e
func (p *utilPProf) X显示页面(r *Request) {
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
		buffer, _ := gview.ParseContent(r.Context别名(), `
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

// Cmdline 响应正在运行程序的命令行，参数之间用 NULL 字节分隔。包初始化时将其注册为 /debug/pprof/cmdline。
// md5:35f5d246119cca43
func (p *utilPProf) Cmdline(r *Request) {
	netpprof.Cmdline(r.X响应.Writer, r.Request)
}

// Profile 使用pprof格式返回CPU profiling信息。
// 如果GET参数指定了持续时间，那么 profiling 将持续该秒数；如果没有指定，则默认为30秒。
// 在包初始化时，它会注册为 "/debug/pprof/profile"。
// md5:11bd281949c0ba3c
func (p *utilPProf) Profile(r *Request) {
	netpprof.Profile(r.X响应.Writer, r.Request)
}

// Symbol 查找请求中列出的程序计数器，
// 并以映射表的形式响应，该映射表将程序计数器与函数名称关联起来。
// 包初始化时将其注册为 /debug/pprof/symbol 路由。
// md5:2944ed5cfe9e0c52
func (p *utilPProf) Symbol(r *Request) {
	netpprof.Symbol(r.X响应.Writer, r.Request)
}

// Trace 返回执行跟踪的二进制形式。
// 跟踪将持续指定的GET参数中的秒数，如果没有指定，则为1秒。
// 包初始化时将其注册为/debug/pprof/trace。
// md5:02830b4c9b48681f
func (p *utilPProf) Trace(r *Request) {
	netpprof.Trace(r.X响应.Writer, r.Request)
}
