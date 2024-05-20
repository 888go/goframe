// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package response

import (
	"bytes"
	"net/http"
)

// BufferWriter 是为带有缓冲的HTTP响应定制的写入器。. md5:07a94c0738608bd9
type BufferWriter struct {
	*Writer               // 底层的BufferWriter。. md5:6aff7656df3aed22
	Status  int           // HTTP status.
	buffer  *bytes.Buffer // The output buffer.
}

func NewBufferWriter(writer http.ResponseWriter) *BufferWriter {
	return &BufferWriter{
		Writer: NewWriter(writer),
		buffer: bytes.NewBuffer(nil),
	}
}

// RawWriter返回底层的BufferWriter。. md5:0f7f231b624df3bc
func (w *BufferWriter) RawWriter() http.ResponseWriter {
	return w.Writer
}

// Write实现了http.BufferWriter.Write接口函数。. md5:076fb1232497e47d
func (w *BufferWriter) Write(data []byte) (int, error) {
	return w.buffer.Write(data)
}

// WriteString 将字符串内容写入内部缓冲区。. md5:e617b9348a1616a2
func (w *BufferWriter) WriteString(data string) (int, error) {
	return w.buffer.WriteString(data)
}

// Buffer 返回缓冲区的内容为 []byte。. md5:e90c5097e6207aef
func (w *BufferWriter) Buffer() []byte {
	return w.buffer.Bytes()
}

// BufferString 返回缓冲区中的内容作为字符串。. md5:ae7d63ad64b703c4
func (w *BufferWriter) BufferString() string {
	return w.buffer.String()
}

// BufferLength 返回缓冲内容的长度。. md5:812f88e103bc03f5
func (w *BufferWriter) BufferLength() int {
	return w.buffer.Len()
}

// SetBuffer 使用 `data` 覆盖缓冲区。. md5:0e9e56a518a98342
func (w *BufferWriter) SetBuffer(data []byte) {
	w.buffer.Reset()
	w.buffer.Write(data)
}

// ClearBuffer 清空响应缓冲区。. md5:5309f2f2892d42bd
func (w *BufferWriter) ClearBuffer() {
	w.buffer.Reset()
}

// WriteHeader 实现了 http.BufferWriter.WriteHeader 接口。. md5:2d5e12d7a9f2b05d
func (w *BufferWriter) WriteHeader(status int) {
	w.Status = status
}

// Flush 将缓冲区的内容发送给客户端，并清空缓冲区。. md5:19043b7f30d54a02
func (w *BufferWriter) Flush() {
	if w.Writer.IsHijacked() {
		return
	}

	if w.Status != 0 && !w.Writer.IsHeaderWrote() {
		w.Writer.WriteHeader(w.Status)
	}
	// 默认状态文本输出。. md5:1cfe5c562c5550e1
	if w.Status != http.StatusOK && w.buffer.Len() == 0 {
		w.buffer.WriteString(http.StatusText(w.Status))
	}
	if w.buffer.Len() > 0 {
		_, _ = w.Writer.Write(w.buffer.Bytes())
		w.buffer.Reset()
		if flusher, ok := w.RawWriter().(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
