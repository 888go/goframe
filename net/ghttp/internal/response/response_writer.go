// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package response

import (
	"bufio"
	"net"
	"net/http"
)

// Writer 在 http.ResponseWriter 上添加了额外功能。 md5:204ac8c0cb436351
type Writer struct {
	http.ResponseWriter       // 基础的ResponseWriter。 md5:1678e6fb48b792ff
	hijacked            bool  // 标记此请求是否已被劫持。 md5:80adeb664fa8ae97
	wroteHeader         bool  // 是否已经写入了响应头，避免出现错误：多余的或多个response.WriteHeader调用。 md5:59bda0050b534efa
	bytesWritten        int64 // 写入响应的字节数。 md5:cc5fa1ce145684ed
}

// NewWriter 创建并返回一个新的 Writer。 md5:6fad96ecb42a0036
func NewWriter(writer http.ResponseWriter) *Writer {
	return &Writer{
		ResponseWriter: writer,
	}
}

// WriteHeader 实现了 http.ResponseWriter.WriteHeader 接口的方法。
// 注意，底层的 `WriteHeader` 方法在一个http响应中只能被调用一次。 md5:7158450c7ec7fc1a
func (w *Writer) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(status)
	w.wroteHeader = true
}

// BytesWritten 返回写入响应的长度。 md5:2bc5d732217ae6e4
func (w *Writer) BytesWritten() int64 {
	return w.bytesWritten
}

// Write实现了http.ResponseWriter.Write接口函数。 md5:7078e0a4eee107f7
func (w *Writer) Write(data []byte) (int, error) {
	n, err := w.ResponseWriter.Write(data)
	w.bytesWritten += int64(n)
	w.wroteHeader = true
	return n, err
}

// Hijack 实现了 http.Hijacker.Hijack 接口函数。 md5:7ef9ff81765b052e
func (w *Writer) Hijack() (conn net.Conn, writer *bufio.ReadWriter, err error) {
	conn, writer, err = w.ResponseWriter.(http.Hijacker).Hijack()
	w.hijacked = true
	return
}

// IsHeaderWrote 返回头部状态是否已写入。 md5:7785f14e4d061fc9
func (w *Writer) IsHeaderWrote() bool {
	return w.wroteHeader
}

// IsHijacked 返回连接是否已被劫持。 md5:11468dbc47bf2400
func (w *Writer) IsHijacked() bool {
	return w.hijacked
}

// Flush 将缓冲区中的任何数据发送到客户端。 md5:38eb50b527a1bfc5
func (w *Writer) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
		w.wroteHeader = true
	}
}
