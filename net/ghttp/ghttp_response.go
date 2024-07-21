// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package ghttp

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gogf/gf/v2/net/ghttp/internal/response"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

// Response 是HTTP响应管理器。
// 注意它实现了带有缓冲功能的http.ResponseWriter接口。
// md5:897398e62eaf56fc
type Response struct {
	*response.BufferWriter          // 基础的 ResponseWriter。 md5:edecebd8a0d4cf02
	Server                 *Server  // Parent server.
	Request                *Request // According request.
}

// newResponse 创建并返回一个新的 Response 对象。 md5:b2d8b0e3f410571c
func newResponse(s *Server, w http.ResponseWriter) *Response {
	r := &Response{
		Server:       s,
		BufferWriter: response.NewBufferWriter(w),
	}
	return r
}

// ServeFile 向响应中发送文件。 md5:e5a83a4dd0cadaf6
func (r *Response) ServeFile(path string, allowIndex ...bool) {
	var (
		serveFile *staticFile
	)
	if file := gres.Get(path); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
	} else {
		path, _ = gfile.Search(path)
		if path == "" {
			r.WriteStatus(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: path}
	}
	r.Server.serveFile(r.Request, serveFile, allowIndex...)
}

// ServeFileDownload 用于将文件下载服务响应到请求。 md5:b5e9e8b76f0afca0
func (r *Response) ServeFileDownload(path string, name ...string) {
	var (
		serveFile    *staticFile
		downloadName = ""
	)

	if len(name) > 0 {
		downloadName = name[0]
	}
	if file := gres.Get(path); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
		if downloadName == "" {
			downloadName = gfile.Basename(file.Name())
		}
	} else {
		path, _ = gfile.Search(path)
		if path == "" {
			r.WriteStatus(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: path}
		if downloadName == "" {
			downloadName = gfile.Basename(path)
		}
	}
	r.Header().Set("Content-Type", "application/force-download")
	r.Header().Set("Accept-Ranges", "bytes")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(downloadName)))
	r.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	r.Server.serveFile(r.Request, serveFile)
}

// RedirectTo 将客户端重定向到另一个位置。
// 可选参数 `code` 指定重定向的HTTP状态码，通常可以是301或302。默认为302。
// md5:ba008c02151efa61
func (r *Response) RedirectTo(location string, code ...int) {
	r.Header().Set("Location", location)
	if len(code) > 0 {
		r.WriteHeader(code[0])
	} else {
		r.WriteHeader(http.StatusFound)
	}
	r.Request.Exit()
}

// RedirectBack 将客户端重定向回引荐来源。
// 可选参数 `code` 指定了用于重定向的HTTP状态码，
// 常见的可选值有301或302，默认情况下使用302。
// md5:b52d05fd1d742c11
func (r *Response) RedirectBack(code ...int) {
	r.RedirectTo(r.Request.GetReferer(), code...)
}

// ServeContent 使用提供的 ReadSeeker 中的内容回复请求。ServeContent 相较于 io.Copy 的主要优点是它能正确处理范围请求，设置 MIME 类型，并处理 If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since 和 If-Range 请求。
//
// 参考 http.ServeContent
// md5:935db9add8e4232c
func (r *Response) ServeContent(name string, modTime time.Time, content io.ReadSeeker) {
	http.ServeContent(r.RawWriter(), r.Request.Request, name, modTime, content)
}

// Flush 将缓冲区的内容输出到客户端并清空缓冲区。 md5:16e9c330d696be4e
func (r *Response) Flush() {
	r.Header().Set(responseHeaderTraceID, gtrace.GetTraceID(r.Request.Context()))
	if r.Server.config.ServerAgent != "" {
		r.Header().Set("Server", r.Server.config.ServerAgent)
	}
	r.BufferWriter.Flush()
}
