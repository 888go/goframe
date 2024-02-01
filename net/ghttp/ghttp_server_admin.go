// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"context"
	"os"
	"strings"
	"time"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/os/gview"
	)
// utilAdmin 是管理控制的控制器。
type utilAdmin struct{}

// Index显示管理页面。
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

// 重启 restarts 各个服务器进程中的所有服务器。
func (p *utilAdmin) Restart(r *Request) {
	var (
		ctx = r.Context()
		err error
	)
	// 当该进程退出时，自定义启动二进制文件路径。
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

// Shutdown 关闭所有服务器。
func (p *utilAdmin) Shutdown(r *Request) {
	gtimer.SetTimeout(r.Context(), time.Second, func(ctx context.Context) {
// 它在1秒后关闭服务器，这不是由系统信号触发的，
// 以确保成功向客户端发送响应。
		_ = r.Server.Shutdown()
	})
	r.Response.WriteExit("server shutdown")
}

// EnableAdmin 启用进程的管理功能。
// 可选参数 `pattern` 指定管理页面的 URI。
func (s *Server) EnableAdmin(pattern ...string) {
	p := "/debug/admin"
	if len(pattern) > 0 {
		p = pattern[0]
	}
	s.BindObject(p, &utilAdmin{})
}

// Shutdown 关闭当前服务器。
func (s *Server) Shutdown() error {
	var ctx = context.TODO()
	s.doServiceDeregister()
// 仅关闭当前服务器。
// 它可能有多个底层HTTP服务器。
	for _, v := range s.servers {
		v.close(ctx)
	}
	return nil
}
