// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package ghttp
import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	
	"github.com/888go/goframe/net/ghttp/internal/response"
	)
// ResponseWriter 是用于 HTTP 响应的自定义编写器。
// 2024-01-07备注, X基础响应器对象, 此处不汉化, liteide有bug,会导致部分字段没有一起重命名
type ResponseWriter struct {
	Status int              // HTTP status.
	writer *response.Writer // 基础的 ResponseWriter。
	buffer *bytes.Buffer    // 输出缓冲区。
}

// RawWriter 返回底层的 ResponseWriter。
func (w *ResponseWriter) RawWriter() http.ResponseWriter {
	return w.writer
}

// Header 实现了 http.ResponseWriter 接口中的 Header 方法。
func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

// Write 实现了 http.ResponseWriter 接口中的 Write 函数。
func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.buffer.Write(data)
	return len(data), nil
}

// WriteHeader 实现了 http.ResponseWriter 接口中的 WriteHeader 方法。
func (w *ResponseWriter) WriteHeader(status int) {
	w.Status = status
}

// Hijack 实现了 http.Hijacker 接口中的 Hijack 函数。
func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.writer.Hijack()
}

// Flush 将缓冲区内容输出到客户端并清空缓冲区。
func (w *ResponseWriter) Flush() {
	if w.writer.IsHijacked() {
		return
	}

	if w.Status != 0 && !w.writer.IsHeaderWrote() {
		w.writer.WriteHeader(w.Status)
	}
	// 默认状态文本输出。
	if w.Status != http.StatusOK && w.buffer.Len() == 0 {
		w.buffer.WriteString(http.StatusText(w.Status))
	}
	if w.buffer.Len() > 0 {
		_, _ = w.writer.Write(w.buffer.Bytes())
		w.buffer.Reset()
		if flusher, ok := w.RawWriter().(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
