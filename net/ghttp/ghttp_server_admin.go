// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

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

// utilAdmin is the controller for administration.
type utilAdmin struct{}

// Index shows the administration page.

// ff:显示管理页面
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

// Restart restarts all the servers in the process.

// ff:重启所有服务
// r:
func (p *utilAdmin) Restart(r *Request) {
	var (
		ctx = r.Context()
		err error
	)
	// Custom start binary path when this process exits.
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

// Shutdown shuts down all the servers.

// ff:关闭所有服务
// r:
func (p *utilAdmin) Shutdown(r *Request) {
	gtimer.SetTimeout(r.Context(), time.Second, func(ctx context.Context) {
		// It shuts down the server after 1 second, which is not triggered by system signal,
		// to ensure the response successfully to the client.
		_ = r.Server.Shutdown()
	})
	r.Response.WriteExit("server shutdown")
}

// EnableAdmin enables the administration feature for the process.
// The optional parameter `pattern` specifies the URI for the administration page.

// ff:平滑重启服务开启
// pattern:管理页URI
func (s *Server) EnableAdmin(pattern ...string) {
	p := "/debug/admin"
	if len(pattern) > 0 {
		p = pattern[0]
	}
	s.BindObject(p, &utilAdmin{})
}

// Shutdown shuts down current server.

// ff:关闭当前服务
func (s *Server) Shutdown() error {
	var ctx = context.TODO()
	s.doServiceDeregister()
	// Only shut down current servers.
	// It may have multiple underlying http servers.
	for _, v := range s.servers {
		v.close(ctx)
	}
	return nil
}
