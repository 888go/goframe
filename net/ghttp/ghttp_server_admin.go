// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"os"
	"strings"
	"time"

	gfile "github.com/888go/goframe/os/gfile"
	gproc "github.com/888go/goframe/os/gproc"
	gtimer "github.com/888go/goframe/os/gtimer"
	gview "github.com/888go/goframe/os/gview"
)

// utilAdmin是管理的控制器。 md5:5be89cb222f1793c
type utilAdmin struct{}

// X显示管理页面 显示管理页面。 md5:768d1e5c44b4ae28
func (p *utilAdmin) X显示管理页面(r *Request) {
	data := map[string]interface{}{
		"pid":  gproc.Pid(),
		"path": gfile.X取当前进程路径(),
		"uri":  strings.TrimRight(r.URL.Path, "/"),
	}
	buffer, _ := gview.ParseContent(r.Context别名(), `
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

// X重启所有服务 重启进程中所有的服务器。 md5:5987ee764222ab8f
func (p *utilAdmin) X重启所有服务(r *Request) {
	var (
		ctx = r.Context别名()
		err error
	)
		// 自定义此进程退出时的自启动二进制路径。 md5:8198fd7edf44314f
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

// X关闭所有服务 关闭所有服务器。 md5:baedfe2ea04060bf
func (p *utilAdmin) X关闭所有服务(r *Request) {
	gtimer.SetTimeout别名(r.Context别名(), time.Second, func(ctx context.Context) {
		// 它在1秒后关闭服务器，这不由系统信号触发，
		// 以确保成功响应客户端。
		// md5:c6fb3a1932e2bc99
		_ = r.X服务.X关闭当前服务()
	})
	r.X响应.X写响应缓冲区并退出("server shutdown")
}

// X平滑重启服务开启 为进程启用管理功能。
// 可选参数 `pattern` 指定了管理页面的 URI。
// md5:edecba6f65e585e5
func (s *X服务) X平滑重启服务开启(管理页URI ...string) {
	p := "/debug/admin"
	if len(管理页URI) > 0 {
		p = 管理页URI[0]
	}
	s.X绑定对象(p, &utilAdmin{})
}

// X关闭当前服务 关闭当前服务器。 md5:b58984e69184996f
func (s *X服务) X关闭当前服务() error {
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
