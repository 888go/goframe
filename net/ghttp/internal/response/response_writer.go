// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package response
import (
	"bufio"
	"net"
	"net/http"
	)
// Writer 封装了 http.ResponseWriter，提供了额外的功能。
type Writer struct {
	http.ResponseWriter      // 基础的 ResponseWriter。
	hijacked            bool // 标记该请求是否已被劫持
	wroteHeader         bool // IsHeaderWroteOrNot 判断头部是否已写入，避免出现“superfluous/multiple response.WriteHeader call”错误。
}

// NewWriter 创建并返回一个新的 Writer。
func NewWriter(writer http.ResponseWriter) *Writer {
	return &Writer{
		ResponseWriter: writer,
	}
}

// WriteHeader 实现了 http.ResponseWriter 接口中的 WriteHeader 方法。
func (w *Writer) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.wroteHeader = true
}

// Hijack 实现了 http.Hijacker 接口中的 Hijack 函数。
func (w *Writer) Hijack() (conn net.Conn, writer *bufio.ReadWriter, err error) {
	conn, writer, err = w.ResponseWriter.(http.Hijacker).Hijack()
	w.hijacked = true
	return
}

// IsHeaderWrote 返回是否已写入头部状态。
func (w *Writer) IsHeaderWrote() bool {
	return w.wroteHeader
}

// IsHijacked 返回连接是否已被劫持。
func (w *Writer) IsHijacked() bool {
	return w.hijacked
}

// Flush 将任何缓冲的数据发送到客户端。
func (w *Writer) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
		w.wroteHeader = true
	}
}
