// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net/http"
	"os"
	"sort"
	"strings"

	ghtml "github.com/888go/goframe/encoding/ghtml"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gfile "github.com/888go/goframe/os/gfile"
	gres "github.com/888go/goframe/os/gres"
	gspath "github.com/888go/goframe/os/gspath"
	gtime "github.com/888go/goframe/os/gtime"
	gstr "github.com/888go/goframe/text/gstr"
)

// ServeHTTP 是 http 请求的默认处理器。
// 它不应该为处理请求创建新的goroutine，因为http.Server已经为此创建了一个新的goroutine。
//
// 这个函数还实现了http.Handler接口。
// md5:82dd5f4475c291db
func (s *X服务) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Max body size limit.
	if s.config.ClientMaxBodySize > 0 {
		r.Body = http.MaxBytesReader(w, r.Body, s.config.ClientMaxBodySize)
	}
		// 重写特征检查。 md5:9dab4befbfff965b
	if len(s.config.Rewrites) > 0 {
		if rewrite, ok := s.config.Rewrites[r.URL.Path]; ok {
			r.URL.Path = rewrite
		}
	}

	var (
		request   = newRequest(s, r, w)    // 创建一个新的请求对象。 md5:e4e00eb82696932c
		sessionId = request.X取SessionId() // 在用户处理器之前获取sessionId. md5:d1a5359c34fed9f9
	)
	defer s.handleAfterRequestDone(request)

	// ============================================================
	// 优先级：
	// 静态文件 > 动态服务 > 静态目录
	// ============================================================
	// md5:9514a47b66a76f01

	// 搜索具有最高优先级的静态文件，同时也处理索引文件功能。
	// md5:f618b1fa06ea7acb
	if s.config.FileServerEnabled {
		request.StaticFile = s.searchStaticFile(r.URL.Path)
		if request.StaticFile != nil {
			request.isFileRequest = true
		}
	}

		// 搜索动态服务处理器。 md5:0cbcd4f2d3569e55
	request.handlers,
		request.serveHandler,
		request.hasHookHandler,
		request.hasServeHandler = s.getHandlersWithCache(request)

		// 检查当前请求的服务类型是静态还是动态。 md5:642ac02f364c85bc
	if request.StaticFile != nil && request.StaticFile.IsDir && request.hasServeHandler {
		request.isFileRequest = false
	}

	// Metrics.
	s.handleMetricsBeforeRequest(request)

	// HOOK - BeforeServe
	s.callHookHandler(HookBeforeServe, request)

	// Core serving handling.
	if !request.X是否已退出() {
		if request.isFileRequest {
			// Static file service.
			s.serveFile(request, request.StaticFile)
		} else {
			if len(request.handlers) > 0 {
				// Dynamic service.
				request.X中间件管理器.Next()
			} else {
				if request.StaticFile != nil && request.StaticFile.IsDir {
					// Serve the directory.
					s.serveFile(request, request.StaticFile)
				} else {
					if len(request.X响应.Header()) == 0 &&
						request.X响应.Status == 0 &&
						request.X响应.BufferLength() == 0 {
						request.X响应.WriteHeader(http.StatusNotFound)
					}
				}
			}
		}
	}

	// HOOK - AfterServe
	if !request.X是否已退出() {
		s.callHookHandler(HookAfterServe, request)
	}

	// HOOK - BeforeOutput
	if !request.X是否已退出() {
		s.callHookHandler(HookBeforeOutput, request)
	}

	// Response handling.
	s.handleResponse(request, sessionId)

	// HOOK - AfterOutput
	if !request.X是否已退出() {
		s.callHookHandler(HookAfterOutput, request)
	}
}

func (s *X服务) handleResponse(request *Request, sessionId string) {
	// HTTP status checking.
	if request.X响应.Status == 0 {
		if request.StaticFile != nil || request.X中间件管理器.served || request.X响应.BufferLength() > 0 {
			request.X响应.WriteHeader(http.StatusOK)
		} else if err := request.X取错误信息(); err != nil {
			if request.X响应.BufferLength() == 0 {
				request.X响应.X写响应缓冲区(err.Error())
			}
			request.X响应.WriteHeader(http.StatusInternalServerError)
		} else {
			request.X响应.WriteHeader(http.StatusNotFound)
		}
	}
	// HTTP status handler.
	if request.X响应.Status != http.StatusOK {
		statusFuncArray := s.getStatusHandler(request.X响应.Status, request)
		for _, f := range statusFuncArray {
						// 调用自定义状态处理器。 md5:8a7c4e0df133e717
			niceCallFunc(func() {
				f(request)
			})
			if request.X是否已退出() {
				break
			}
		}
	}

	// 如果在这个请求中生成了新的会话ID，并且启用了SessionCookieOutput，自动将会话ID设置为cookie。
	// md5:2c6864797c5d809f
	if s.config.SessionCookieOutput && request.Session.IsDirty() {
		// 初始化会话前，可以通过 r.Session.SetId("") 来更改
		// 也可以通过 r.Cookie.SetSessionId("") 来更改
		// md5:7175563db73b9a50
		sidFromSession, sidFromRequest := request.Session.MustId(), request.X取SessionId()
		if sidFromSession != sidFromRequest {
			if sidFromSession != sessionId {
				request.Cookie.X设置SessionId到Cookie(sidFromSession)
			} else {
				request.Cookie.X设置SessionId到Cookie(sidFromRequest)
			}
		}
	}
		// 将cookie内容输出到客户端。 md5:b9694a9aa06119db
	request.Cookie.X输出()
		// 将缓冲区内容输出到客户端。 md5:fe2997ba592b17ad
	request.X响应.X输出缓存区()
}

func (s *X服务) handleAfterRequestDone(request *Request) {
	request.LeaveTime = gtime.X创建并按当前时间()
	// error log handling.
	if request.error != nil {
		s.handleErrorLog(request.error, request)
	} else {
		if exception := recover(); exception != nil {
			request.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError)
			if v, ok := exception.(error); ok {
				if code := gerror.X取错误码(v); code != gcode.CodeNil {
					s.handleErrorLog(v, request)
				} else {
					s.handleErrorLog(
						gerror.X多层错误码并跳过堆栈(gcode.CodeInternalPanic, 1, v, ""),
						request,
					)
				}
			} else {
				s.handleErrorLog(
					gerror.X创建错误码并跳过堆栈与格式化(gcode.CodeInternalPanic, 1, "%+v", exception),
					request,
				)
			}
		}
	}
	// access log handling.
	s.handleAccessLog(request)
	// 关闭会话，如果会话存在，这将自动更新其TTL（超时时间）。
	// md5:a86a4db886c94158
	if err := request.Session.Close(); err != nil {
		intlog.Errorf(request.Context别名(), `%+v`, err)
	}

	// 关闭请求和响应体以及时释放文件描述符。
	// md5:aea97d230b2451b0
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

	// Metrics.
	s.handleMetricsAfterRequestDone(request)
}

// searchStaticFile 根据给定的URI搜索文件。
// 它返回一个file结构体，其中包含文件信息。
// md5:e5b76cc2b6c98a07
func (s *X服务) searchStaticFile(uri string) *staticFile {
	var (
		file *gres.File
		path string
		dir  bool
	)
		// 首先搜索StaticPaths映射。 md5:4f9c5afa25bf93dd
	if len(s.config.StaticPaths) > 0 {
		for _, item := range s.config.StaticPaths {
			if len(uri) >= len(item.Prefix) && strings.EqualFold(item.Prefix, uri[0:len(item.Prefix)]) {
								// 为了避免像这样的情况：/static/style -> /static/style.css. md5:74ccef8cd597d359
				if len(uri) > len(item.Prefix) && uri[len(item.Prefix)] != '/' {
					continue
				}
				file = gres.GetWithIndex(item.Path+uri[len(item.Prefix):], s.config.IndexFiles)
				if file != nil {
					return &staticFile{
						File:  file,
						IsDir: file.FileInfo().IsDir(),
					}
				}
				path, dir = gspath.Search(item.Path, uri[len(item.Prefix):], s.config.IndexFiles...)
				if path != "" {
					return &staticFile{
						Path:  path,
						IsDir: dir,
					}
				}
			}
		}
	}
		// 其次，搜索根目录和搜索路径。 md5:9b1b9aadf8478052
	if len(s.config.SearchPaths) > 0 {
		for _, p := range s.config.SearchPaths {
			file = gres.GetWithIndex(p+uri, s.config.IndexFiles)
			if file != nil {
				return &staticFile{
					File:  file,
					IsDir: file.FileInfo().IsDir(),
				}
			}
			if path, dir = gspath.Search(p, uri, s.config.IndexFiles...); path != "" {
				return &staticFile{
					Path:  path,
					IsDir: dir,
				}
			}
		}
	}
		// 最后搜索资源管理器。 md5:1ccc8123528fc4a4
	if len(s.config.StaticPaths) == 0 && len(s.config.SearchPaths) == 0 {
		if file = gres.GetWithIndex(uri, s.config.IndexFiles); file != nil {
			return &staticFile{
				File:  file,
				IsDir: file.FileInfo().IsDir(),
			}
		}
	}
	return nil
}

// serveFile 为客户端服务静态文件。
// 可选参数 `allowIndex` 指定如果 `f` 是一个目录时是否允许目录列表。
// md5:1741c137e9fcf4cd
func (s *X服务) serveFile(r *Request, f *staticFile, allowIndex ...bool) {
		// 从内存中使用资源文件。 md5:eb37e3d39231ad74
	if f.File != nil {
		if f.IsDir {
			if s.config.IndexFolder || (len(allowIndex) > 0 && allowIndex[0]) {
				s.listDir(r, f.File)
			} else {
				r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
			}
		} else {
			info := f.File.FileInfo()
			r.X响应.ServeContent(info.Name(), info.ModTime(), f.File)
		}
		return
	}
	// Use file from dist.
	file, err := os.Open(f.Path)
	if err != nil {
		r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	// 在服务文件之前清空响应缓冲区。
	// 它忽略所有自定义的缓冲内容，转而使用文件内容。
	// md5:b7ae0cf8ef13c29c
	r.X响应.ClearBuffer()

	info, _ := file.Stat()
	if info.IsDir() {
		if s.config.IndexFolder || (len(allowIndex) > 0 && allowIndex[0]) {
			s.listDir(r, file)
		} else {
			r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		}
	} else {
		r.X响应.ServeContent(info.Name(), info.ModTime(), file)
	}
}

// listDir 将指定目录下的子文件以HTML内容的形式列出来发送给客户端。 md5:1648438b6fcd2bd5
func (s *X服务) listDir(r *Request, f http.File) {
	files, err := f.Readdir(-1)
	if err != nil {
		r.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, "Error reading directory")
		return
	}
		// 文件夹类型优先于文件。 md5:f5cc5a85f701d6c1
	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() && !files[j].IsDir() {
			return true
		}
		if !files[i].IsDir() && files[j].IsDir() {
			return false
		}
		return files[i].Name() < files[j].Name()
	})
	if r.X响应.Header().Get("Content-Type") == "" {
		r.X响应.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	r.X响应.X写响应缓冲区(`<html>`)
	r.X响应.X写响应缓冲区(`<head>`)
	r.X响应.X写响应缓冲区(`<style>`)
	r.X响应.X写响应缓冲区(`body {font-family:Consolas, Monaco, "Andale Mono", "Ubuntu Mono", monospace;}`)
	r.X响应.X写响应缓冲区(`</style>`)
	r.X响应.X写响应缓冲区(`</head>`)
	r.X响应.X写响应缓冲区(`<body>`)
	r.X响应.X写响应缓冲区并格式化(`<h1>Index of %s</h1>`, r.URL.Path)
	r.X响应.X写响应缓冲区并格式化(`<hr />`)
	r.X响应.X写响应缓冲区(`<table>`)
	if r.URL.Path != "/" {
		r.X响应.X写响应缓冲区(`<tr>`)
		r.X响应.X写响应缓冲区并格式化(`<td><a href="%s">..</a></td>`, gfile.X路径取父目录(r.URL.Path))
		r.X响应.X写响应缓冲区(`</tr>`)
	}
	name := ""
	size := ""
	prefix := gstr.X过滤尾字符并含空白(r.URL.Path, "/")
	for _, file := range files {
		name = file.Name()
		size = gfile.X字节长度转易读格式(file.Size())
		if file.IsDir() {
			name += "/"
			size = "-"
		}
		r.X响应.X写响应缓冲区(`<tr>`)
		r.X响应.X写响应缓冲区并格式化(`<td><a href="%s/%s">%s</a></td>`, prefix, name, ghtml.X编码特殊字符(name))
		r.X响应.X写响应缓冲区并格式化(`<td style="width:300px;text-align:center;">%s</td>`, gtime.X创建(file.ModTime()).X取文本时间ISO8601())
		r.X响应.X写响应缓冲区并格式化(`<td style="width:80px;text-align:right;">%s</td>`, size)
		r.X响应.X写响应缓冲区(`</tr>`)
	}
	r.X响应.X写响应缓冲区(`</table>`)
	r.X响应.X写响应缓冲区(`</body>`)
	r.X响应.X写响应缓冲区(`</html>`)
}
