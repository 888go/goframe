// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package http类

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	
	"github.com/888go/goframe/net/ghttp/internal/response"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
)

// Response 是HTTP响应管理器。
// 请注意，它实现了带有缓冲功能的http.ResponseWriter接口。
type X响应 struct {
	*ResponseWriter                 // 基础的 ResponseWriter。
	Server          *X服务         // Parent server.
	Writer          *ResponseWriter // ResponseWriter的别名。
	Request         *X请求        // 根据请求。
}

// newResponse 创建并返回一个新的 Response 对象。
func newResponse(s *X服务, w http.ResponseWriter) *X响应 {
	r := &X响应{
		Server: s,
		ResponseWriter: &ResponseWriter{
			writer: response.NewWriter(w),
			buffer: bytes.NewBuffer(nil),
		},
	}
	r.Writer = r.ResponseWriter
	return r
}

// ServeFile 将文件发送至响应。
// 会自动识别文件格式，如果是目录或者文本内容将会直接展示文件内容。
// 如果path参数为目录，那么第二个参数allowIndex控制是否可以展示目录下的文件列表。
func (r *X响应) X发送文件(文件路径 string, 是否展示目录文件列表 ...bool) {
	var (
		serveFile *staticFile
	)
	if file := 资源类.Get(文件路径); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
	} else {
		文件路径, _ = 文件类.X查找(文件路径)
		if 文件路径 == "" {
			r.X写响应缓冲区与HTTP状态码(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: 文件路径}
	}
	r.Server.serveFile(r.Request, serveFile, 是否展示目录文件列表...)
}

// ServeFileDownload 为响应提供文件下载服务。
// 用于直接引导客户端下载指定路径的文件，并可以重新给定下载的文件名称。
// ServeFileDownload方法采用的是流式下载控制，对内存占用较少。
// 使用示例，我们把示例中的ServeFile方法改为ServeFileDownload方法：
// func main() {
// 	s := g.Server()
// 	s.BindHandler("/", func(r *ghttp.Request) {
// 		r.Response.ServeFileDownload("test.txt")
// 	})
// 	s.SetPort(8999)
// 	s.Run()
// }
func (r *X响应) X下载文件(路径 string, 文件名 ...string) {
	var (
		serveFile    *staticFile
		downloadName = ""
	)

	if len(文件名) > 0 {
		downloadName = 文件名[0]
	}
	if file := 资源类.Get(路径); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
		if downloadName == "" {
			downloadName = 文件类.X路径取文件名(file.Name())
		}
	} else {
		路径, _ = 文件类.X查找(路径)
		if 路径 == "" {
			r.X写响应缓冲区与HTTP状态码(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: 路径}
		if downloadName == "" {
			downloadName = 文件类.X路径取文件名(路径)
		}
	}
	r.Header().Set("Content-Type", "application/force-download")
	r.Header().Set("Accept-Ranges", "bytes")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(downloadName)))
	r.Server.serveFile(r.Request, serveFile)
}

// RedirectTo 重定向客户端到另一个位置。
// 可选参数 `code` 指定了用于重定向的 HTTP 状态码，
// 通常可以是 301 或 302，默认为 302。
func (r *X响应) X重定向(url地址 string, 重定向状态码 ...int) {
	r.Header().Set("Location", url地址)
	if len(重定向状态码) > 0 {
		r.WriteHeader(重定向状态码[0])
	} else {
		r.WriteHeader(http.StatusFound)
	}
	r.Request.X退出当前()
}

// RedirectBack 重定向客户端返回到referer页面。
// 可选参数 `code` 指定用于重定向的http状态码，通常可以是301或302，默认为302。
func (r *X响应) X重定向到来源页面(重定向状态码 ...int) {
	r.X重定向(r.Request.X取引用来源URL(), 重定向状态码...)
}

// Buffer返回缓冲区中的内容作为[]byte。
func (r *X响应) X取缓冲区字节集() []byte {
	return r.buffer.Bytes()
}

// BufferString 返回缓冲区中的内容作为字符串。
func (r *X响应) X取缓冲区文本() string {
	return r.buffer.String()
}

// BufferLength 返回缓冲区内容的长度。
func (r *X响应) X取缓冲区长度() int {
	return r.buffer.Len()
}

// SetBuffer 将`data`覆盖写入缓冲区。
func (r *X响应) X设置缓冲区字节集(字节集值 []byte) {
	r.buffer.Reset()
	r.buffer.Write(字节集值)
}

// 清空缓冲区 ClearBuffer 用于清空响应缓冲区。
func (r *X响应) X清空缓冲区() {
	r.buffer.Reset()
}

// ServeContent 函数通过提供的 ReadSeeker 中的内容回复请求。与 io.Copy 相比，ServeContent 的主要优点在于它能妥善处理 Range 请求，设置 MIME 类型，并正确处理 If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since 以及 If-Range 等请求。
//
// 参见 http.ServeContent
func (r *X响应) ServeContent(name string, modTime time.Time, content io.ReadSeeker) {
	http.ServeContent(r.Writer.RawWriter(), r.Request.Request, name, modTime, content)
}

// Flush 将缓冲区内容输出到客户端并清空缓冲区。
func (r *X响应) X输出缓存区() {
	r.Header().Set(responseHeaderTraceID, gtrace.GetTraceID(r.Request.Context别名()))
	if r.Server.config.X服务器代理 != "" {
		r.Header().Set("Server", r.Server.config.X服务器代理)
	}
	r.Writer.Flush()
}
