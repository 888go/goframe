// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"net/http"
	"os"
	"sort"
	"strings"
	
	"github.com/888go/goframe/encoding/ghtml"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gspath"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gstr"
)

// ServeHTTP 是处理 HTTP 请求的默认处理器。
// 由于它是由 http.Server 已经创建的新 goroutine 调用的，所以不应在此函数中创建新的处理请求的 goroutine。
//
// 此外，这个函数实现了 http.Handler 接口。
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 最大请求体大小限制。
	if s.config.ClientMaxBodySize > 0 {
		r.Body = http.MaxBytesReader(w, r.Body, s.config.ClientMaxBodySize)
	}
	// 重写特性检查。
	if len(s.config.Rewrites) > 0 {
		if rewrite, ok := s.config.Rewrites[r.URL.Path]; ok {
			r.URL.Path = rewrite
		}
	}

	// 创建一个新的请求对象。
	request := newRequest(s, r, w)

	// 在用户处理器之前获取sessionId
	sessionId := request.X取SessionId()

	defer func() {
		request.LeaveTime = 时间类.X取时间戳毫秒()
		// 错误日志处理。
		if request.error != nil {
			s.handleErrorLog(request.error, request)
		} else {
			if exception := recover(); exception != nil {
				request.Response.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError)
				if v, ok := exception.(error); ok {
					if code := 错误类.X取错误码(v); code != 错误码类.CodeNil {
						s.handleErrorLog(v, request)
					} else {
						s.handleErrorLog(错误类.X多层错误码并跳过堆栈(错误码类.CodeInternalPanic, 1, v, ""), request)
					}
				} else {
					s.handleErrorLog(错误类.X创建错误码并跳过堆栈与格式化(错误码类.CodeInternalPanic, 1, "%+v", exception), request)
				}
			}
		}
		// 访问日志处理。
		s.handleAccessLog(request)
// 关闭会话，如果会话存在，会自动更新其TTL（生存时间）
		if err := request.Session.Close(); err != nil {
			intlog.Errorf(request.Context别名(), `%+v`, err)
		}

// 关闭请求和响应体
// 以便及时释放文件描述符。
		err := request.Request.Body.Close()
		if err != nil {
			intlog.Errorf(request.Context别名(), `%+v`, err)
		}
		if request.Request.Response != nil {
			err = request.Request.Response.Body.Close()
			if err != nil {
				intlog.Errorf(request.Context别名(), `%+v`, err)
			}
		}
	}()

// ============================================================
// 优先级：
// 静态文件 > 动态服务 > 静态目录
// ============================================================

// 搜索优先级最高的静态文件，并处理索引文件特性。
	if s.config.FileServerEnabled {
		request.StaticFile = s.searchStaticFile(r.URL.Path)
		if request.StaticFile != nil {
			request.isFileRequest = true
		}
	}

	// 搜索动态服务处理器。
	request.handlers,
		request.serveHandler,
		request.hasHookHandler,
		request.hasServeHandler = s.getHandlersWithCache(request)

	// 检查当前请求的服务类型是静态还是动态。
	if request.StaticFile != nil && request.StaticFile.IsDir && request.hasServeHandler {
		request.isFileRequest = false
	}

	// HOOK - 服务启动前
	s.callHookHandler(HookBeforeServe, request)

	// 核心服务处理。
	if !request.X是否已退出() {
		if request.isFileRequest {
			// 静态文件服务。
			s.serveFile(request, request.StaticFile)
		} else {
			if len(request.handlers) > 0 {
				// Dynamic service.
				request.Middleware.Next()
			} else {
				if request.StaticFile != nil && request.StaticFile.IsDir {
					// 服务目录（提供目录内容）
					s.serveFile(request, request.StaticFile)
				} else {
					if len(request.Response.Header()) == 0 &&
						request.Response.Status == 0 &&
						request.Response.X取缓冲区长度() == 0 {
						request.Response.WriteHeader(http.StatusNotFound)
					}
				}
			}
		}
	}

	// HOOK - AfterServe
	if !request.X是否已退出() {
		s.callHookHandler(HookAfterServe, request)
	}

	// HOOK - 输出之前
	if !request.X是否已退出() {
		s.callHookHandler(HookBeforeOutput, request)
	}

	// HTTP状态检查。
	if request.Response.Status == 0 {
		if request.StaticFile != nil || request.Middleware.served || request.Response.buffer.Len() > 0 {
			request.Response.WriteHeader(http.StatusOK)
		} else if err := request.X取错误信息(); err != nil {
			if request.Response.X取缓冲区长度() == 0 {
				request.Response.X写响应缓冲区(err.Error())
			}
			request.Response.WriteHeader(http.StatusInternalServerError)
		} else {
			request.Response.WriteHeader(http.StatusNotFound)
		}
	}
	// HTTP状态处理器。
	if request.Response.Status != http.StatusOK {
		statusFuncArray := s.getStatusHandler(request.Response.Status, request)
		for _, f := range statusFuncArray {
			// 调用自定义状态处理器。
			niceCallFunc(func() {
				f(request)
			})
			if request.X是否已退出() {
				break
			}
		}
	}

// 如果在本次请求中创建了新的会话ID，并且启用了SessionCookieOutput，则自动将会话ID设置到cookie中。
	if s.config.SessionCookieOutput && request.Session.IsDirty() {
// 在初始化session之前，可以通过r.Session.SetId("")来改变
// 也可以通过r.Cookie.SetSessionId("")来改变
		sidFromSession, sidFromRequest := request.Session.MustId(), request.X取SessionId()
		if sidFromSession != sidFromRequest {
			if sidFromSession != sessionId {
				request.Cookie.X设置SessionId到Cookie(sidFromSession)
			} else {
				request.Cookie.X设置SessionId到Cookie(sidFromRequest)
			}
		}
	}
	// 将cookie内容输出到客户端。
	request.Cookie.X输出()
	// 将缓冲区内容输出到客户端。
	request.Response.X输出缓存区()
	// HOOK - 输出后
	if !request.X是否已退出() {
		s.callHookHandler(HookAfterOutput, request)
	}
}

// searchStaticFile 通过给定的 URI 搜索文件。
// 它返回一个文件结构体，该结构体指定了文件信息。
func (s *Server) searchStaticFile(uri string) *staticFile {
	var (
		file *资源类.File
		path string
		dir  bool
	)
	// 首先搜索 StaticPaths 映射。
	if len(s.config.StaticPaths) > 0 {
		for _, item := range s.config.StaticPaths {
			if len(uri) >= len(item.Prefix) && strings.EqualFold(item.Prefix, uri[0:len(item.Prefix)]) {
				// 为避免出现类似这种情况：/static/style -> /static/style.css 的情况
				if len(uri) > len(item.Prefix) && uri[len(item.Prefix)] != '/' {
					continue
				}
				file = 资源类.GetWithIndex(item.Path+uri[len(item.Prefix):], s.config.IndexFiles)
				if file != nil {
					return &staticFile{
						File:  file,
						IsDir: file.FileInfo().IsDir(),
					}
				}
				path, dir = 文件搜索类.Search(item.Path, uri[len(item.Prefix):], s.config.IndexFiles...)
				if path != "" {
					return &staticFile{
						Path:  path,
						IsDir: dir,
					}
				}
			}
		}
	}
	// 其次，搜索根目录和搜索路径。
	if len(s.config.SearchPaths) > 0 {
		for _, p := range s.config.SearchPaths {
			file = 资源类.GetWithIndex(p+uri, s.config.IndexFiles)
			if file != nil {
				return &staticFile{
					File:  file,
					IsDir: file.FileInfo().IsDir(),
				}
			}
			if path, dir = 文件搜索类.Search(p, uri, s.config.IndexFiles...); path != "" {
				return &staticFile{
					Path:  path,
					IsDir: dir,
				}
			}
		}
	}
	// 最后在资源管理器中进行搜索。
	if len(s.config.StaticPaths) == 0 && len(s.config.SearchPaths) == 0 {
		if file = 资源类.GetWithIndex(uri, s.config.IndexFiles); file != nil {
			return &staticFile{
				File:  file,
				IsDir: file.FileInfo().IsDir(),
			}
		}
	}
	return nil
}

// serveFile 为客户端提供静态文件服务。
// 可选参数 `allowIndex` 指定当 `f` 是目录时，是否允许目录列表展示。
func (s *Server) serveFile(r *Request, f *staticFile, allowIndex ...bool) {
	// 从内存中使用资源文件。
	if f.File != nil {
		if f.IsDir {
			if s.config.IndexFolder || (len(allowIndex) > 0 && allowIndex[0]) {
				s.listDir(r, f.File)
			} else {
				r.Response.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
			}
		} else {
			info := f.File.FileInfo()
			r.Response.ServeContent(info.Name(), info.ModTime(), f.File)
		}
		return
	}
	// 使用来自dist目录的文件。
	file, err := os.Open(f.Path)
	if err != nil {
		r.Response.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		return
	}
	defer file.Close()

// 在文件服务之前清空响应缓冲区。
// 它会忽略所有自定义缓冲区内容，并使用文件内容。
	r.Response.X清空缓冲区()

	info, _ := file.Stat()
	if info.IsDir() {
		if s.config.IndexFolder || (len(allowIndex) > 0 && allowIndex[0]) {
			s.listDir(r, file)
		} else {
			r.Response.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		}
	} else {
		r.Response.ServeContent(info.Name(), info.ModTime(), file)
	}
}

// listDir 将指定目录下的子文件列表以HTML内容的形式发送给客户端。
func (s *Server) listDir(r *Request, f http.File) {
	files, err := f.Readdir(-1)
	if err != nil {
		r.Response.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, "Error reading directory")
		return
	}
	// 文件夹类型比文件类型具有更高的优先级。
	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() && !files[j].IsDir() {
			return true
		}
		if !files[i].IsDir() && files[j].IsDir() {
			return false
		}
		return files[i].Name() < files[j].Name()
	})
	if r.Response.Header().Get("Content-Type") == "" {
		r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	r.Response.X写响应缓冲区(`<html>`)
	r.Response.X写响应缓冲区(`<head>`)
	r.Response.X写响应缓冲区(`<style>`)
	r.Response.X写响应缓冲区(`body {font-family:Consolas, Monaco, "Andale Mono", "Ubuntu Mono", monospace;}`)
	r.Response.X写响应缓冲区(`</style>`)
	r.Response.X写响应缓冲区(`</head>`)
	r.Response.X写响应缓冲区(`<body>`)
	r.Response.X写响应缓冲区并格式化(`<h1>Index of %s</h1>`, r.URL.Path)
	r.Response.X写响应缓冲区并格式化(`<hr />`)
	r.Response.X写响应缓冲区(`<table>`)
	if r.URL.Path != "/" {
		r.Response.X写响应缓冲区(`<tr>`)
		r.Response.X写响应缓冲区并格式化(`<td><a href="%s">..</a></td>`, 文件类.X路径取父目录(r.URL.Path))
		r.Response.X写响应缓冲区(`</tr>`)
	}
	name := ""
	size := ""
	prefix := 文本类.X过滤尾字符并含空白(r.URL.Path, "/")
	for _, file := range files {
		name = file.Name()
		size = 文件类.X字节长度转易读格式(file.Size())
		if file.IsDir() {
			name += "/"
			size = "-"
		}
		r.Response.X写响应缓冲区(`<tr>`)
		r.Response.X写响应缓冲区并格式化(`<td><a href="%s/%s">%s</a></td>`, prefix, name, html类.X编码特殊字符(name))
		r.Response.X写响应缓冲区并格式化(`<td style="width:300px;text-align:center;">%s</td>`, 时间类.X创建(file.ModTime()).X取文本时间ISO8601())
		r.Response.X写响应缓冲区并格式化(`<td style="width:80px;text-align:right;">%s</td>`, size)
		r.Response.X写响应缓冲区(`</tr>`)
	}
	r.Response.X写响应缓冲区(`</table>`)
	r.Response.X写响应缓冲区(`</body>`)
	r.Response.X写响应缓冲区(`</html>`)
}
