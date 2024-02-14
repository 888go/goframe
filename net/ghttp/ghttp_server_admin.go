// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

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
func (p *utilAdmin) X显示管理页面(r *X请求) {
	data := map[string]interface{}{
		"pid":  进程类.Pid(),
		"path": 文件类.X取当前进程路径(),
		"uri":  strings.TrimRight(r.URL.Path, "/"),
	}
	buffer, _ := 模板类.ParseContent(r.Context别名(), `
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
	r.X响应.X写响应缓冲区(buffer)
}

// 重启 restarts 各个服务器进程中的所有服务器。
func (p *utilAdmin) X重启所有服务(r *X请求) {
	var (
		ctx = r.Context别名()
		err error
	)
	// 当该进程退出时，自定义启动二进制文件路径。
	path := r.X取查询参数到泛型类("newExeFilePath").String()
	if path == "" {
		path = os.Args[0]
	}
	if err = X平滑重启所有服务(ctx, path); err == nil {
		r.X响应.X写响应缓冲区并退出("server restarted")
	} else {
		r.X响应.X写响应缓冲区并退出(err.Error())
	}
}

// Shutdown 关闭所有服务器。
func (p *utilAdmin) X关闭所有服务(r *X请求) {
	定时类.SetTimeout别名(r.Context别名(), time.Second, func(ctx context.Context) {
// 它在1秒后关闭服务器，这不是由系统信号触发的，
// 以确保成功向客户端发送响应。
		_ = r.X服务.X关闭当前服务()
	})
	r.X响应.X写响应缓冲区并退出("server shutdown")
}

// EnableAdmin 启用进程的管理功能。
// 可选参数 `pattern` 指定管理页面的 URI。
func (s *X服务) X平滑重启服务开启(管理页URI ...string) {
	p := "/debug/admin"
	if len(管理页URI) > 0 {
		p = 管理页URI[0]
	}
	s.X绑定对象(p, &utilAdmin{})
}

// Shutdown 关闭当前服务器。
func (s *X服务) X关闭当前服务() error {
	var ctx = context.TODO()
	s.doServiceDeregister()
// 仅关闭当前服务器。
// 它可能有多个底层HTTP服务器。
	for _, v := range s.servers {
		v.close(ctx)
	}
	return nil
}
