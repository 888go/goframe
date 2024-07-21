// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/os/gview"
)

// utilAdmin是管理的控制器。 md5:5be89cb222f1793c
type utilAdmin struct{}

// Index 显示管理页面。 md5:768d1e5c44b4ae28
// ff:显示管理页面
// p:
// r:
func (p *utilAdmin) Index(r *Request) {
	data := map[string]interface{}{
		"pid":  gproc.Pid(),
		"path": gfile.SelfPath(),
		"uri":  strings.TrimRight(r.URL.Path, "/"),
	}
	buffer, _ := gview.ParseContent(r.Context(), `
            <html>
            <head>
                <title>GoFrame Web Server Admin</title>
            </head>
            <body>
                <p>Pid: {{.pid}}</p>
                <p>File Path: {{.path}}</p>
                <p><a href="{{$.uri}}/restart">Restart</a></p>
                <p><a href="{{$.uri}}/shutdown">Shutdown</a></p>
            </body>
            </html>
    `, data)
	r.Response.Write(buffer)
}

// Restart 重启进程中所有的服务器。 md5:5987ee764222ab8f
// ff:重启所有服务
// p:
// r:
func (p *utilAdmin) Restart(r *Request) {
	var (
		ctx = r.Context()
		err error
	)
	// 自定义此进程退出时的自启动二进制路径。 md5:8198fd7edf44314f
	path := r.GetQuery("newExeFilePath").String()
	if path == "" {
		path = os.Args[0]
	}
	if err = RestartAllServer(ctx, path); err == nil {
		r.Response.WriteExit("server restarted")
	} else {
		r.Response.WriteExit(err.Error())
	}
}

// Shutdown 关闭所有服务器。 md5:baedfe2ea04060bf
// ff:关闭所有服务
// p:
// r:
func (p *utilAdmin) Shutdown(r *Request) {
	gtimer.SetTimeout(r.Context(), time.Second, func(ctx context.Context) {
// 它在1秒后关闭服务器，这不由系统信号触发，
// 以确保成功响应客户端。
// md5:c6fb3a1932e2bc99
		_ = r.Server.Shutdown()
	})
	r.Response.WriteExit("server shutdown")
}

// EnableAdmin 为进程启用管理功能。
// 可选参数 `pattern` 指定了管理页面的 URI。
// md5:edecba6f65e585e5
// ff:平滑重启服务开启
// s:
// pattern:管理页URI
func (s *Server) EnableAdmin(pattern ...string) {
	p := "/debug/admin"
	if len(pattern) > 0 {
		p = pattern[0]
	}
	s.BindObject(p, &utilAdmin{})
}

// Shutdown 关闭当前服务器。 md5:b58984e69184996f
// ff:关闭当前服务
// s:
func (s *Server) Shutdown() error {
	var ctx = context.TODO()
	s.doServiceDeregister()
	// 只关闭当前的服务器。
	// 它可能有多个底层HTTP服务器。
	// md5:da82a73df4e8a814
	for _, v := range s.servers {
		v.close(ctx)
	}
	return nil
}
