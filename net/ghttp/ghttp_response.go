// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package http类

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/888go/goframe/net/ghttp/internal/response"
	"github.com/888go/goframe/net/gtrace"
	gfile "github.com/888go/goframe/os/gfile"
	gres "github.com/888go/goframe/os/gres"
)

// Response 是HTTP响应管理器。
// 注意它实现了带有缓冲功能的http.ResponseWriter接口。
// md5:897398e62eaf56fc
type Response struct {
	*response.BufferWriter          // 基础的 ResponseWriter。 md5:edecebd8a0d4cf02
	Server                 *X服务  // Parent server.
	Request                *Request // According request.
}

// newResponse 创建并返回一个新的 Response 对象。 md5:b2d8b0e3f410571c
func newResponse(s *X服务, w http.ResponseWriter) *Response {
	r := &Response{
		Server:       s,
		BufferWriter: response.NewBufferWriter(w),
	}
	return r
}

// X发送文件 向响应中发送文件。 md5:e5a83a4dd0cadaf6
func (r *Response) X发送文件(文件路径 string, 是否展示目录文件列表 ...bool) {
	var (
		serveFile *staticFile
	)
	if file := gres.Get(文件路径); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
	} else {
		文件路径, _ = gfile.X查找(文件路径)
		if 文件路径 == "" {
			r.X写响应缓冲区与HTTP状态码(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: 文件路径}
	}
	r.Server.serveFile(r.Request, serveFile, 是否展示目录文件列表...)
}

// X下载文件 用于将文件下载服务响应到请求。 md5:b5e9e8b76f0afca0
func (r *Response) X下载文件(路径 string, 文件名 ...string) {
	var (
		serveFile    *staticFile
		downloadName = ""
	)

	if len(文件名) > 0 {
		downloadName = 文件名[0]
	}
	if file := gres.Get(路径); file != nil {
		serveFile = &staticFile{
			File:  file,
			IsDir: file.FileInfo().IsDir(),
		}
		if downloadName == "" {
			downloadName = gfile.X路径取文件名(file.Name())
		}
	} else {
		路径, _ = gfile.X查找(路径)
		if 路径 == "" {
			r.X写响应缓冲区与HTTP状态码(http.StatusNotFound)
			return
		}
		serveFile = &staticFile{Path: 路径}
		if downloadName == "" {
			downloadName = gfile.X路径取文件名(路径)
		}
	}
	r.Header().Set("Content-Type", "application/force-download")
	r.Header().Set("Accept-Ranges", "bytes")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(downloadName)))
	r.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	r.Server.serveFile(r.Request, serveFile)
}

// X重定向 将客户端重定向到另一个位置。
// 可选参数 `code` 指定重定向的HTTP状态码，通常可以是301或302。默认为302。
// md5:ba008c02151efa61
func (r *Response) X重定向(url地址 string, 重定向状态码 ...int) {
	r.Header().Set("Location", url地址)
	if len(重定向状态码) > 0 {
		r.WriteHeader(重定向状态码[0])
	} else {
		r.WriteHeader(http.StatusFound)
	}
	r.Request.X退出当前()
}

// X重定向到来源页面 将客户端重定向回引荐来源。
// 可选参数 `code` 指定了用于重定向的HTTP状态码，
// 常见的可选值有301或302，默认情况下使用302。
// md5:b52d05fd1d742c11
func (r *Response) X重定向到来源页面(重定向状态码 ...int) {
	r.X重定向(r.Request.X取引用来源URL(), 重定向状态码...)
}

// ServeContent 使用提供的 ReadSeeker 中的内容回复请求。ServeContent 相较于 io.Copy 的主要优点是它能正确处理范围请求，设置 MIME 类型，并处理 If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since 和 If-Range 请求。
//
// 参考 http.ServeContent
// md5:935db9add8e4232c
func (r *Response) ServeContent(name string, modTime time.Time, content io.ReadSeeker) {
	http.ServeContent(r.RawWriter(), r.Request.Request, name, modTime, content)
}

// X输出缓存区 将缓冲区的内容输出到客户端并清空缓冲区。 md5:16e9c330d696be4e
func (r *Response) X输出缓存区() {
	r.Header().Set(responseHeaderTraceID, gtrace.GetTraceID(r.Request.Context别名()))
	if r.Server.config.ServerAgent != "" {
		r.Header().Set("Server", r.Server.config.ServerAgent)
	}
	r.BufferWriter.Flush()
}
