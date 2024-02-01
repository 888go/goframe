// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils
import (
	"io"
	)
// ReadCloser 实现了 io.ReadCloser 接口，
// 这个接口常用于多次读取请求正文内容。
//
// 注意，它不能被关闭。
type ReadCloser struct {
	index      int    // 当前读取位置。
	content    []byte // Content.
	repeatable bool   // 标记内容可以进行可重复读取。
}

// NewReadCloser 创建并返回一个 RepeatReadCloser 对象。
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read 实现了 io.ReadCloser 接口。
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	// 使其可重复读取。
	if b.index >= len(b.content) && b.repeatable {
		b.index = 0
	}
	n = copy(p, b.content[b.index:])
	b.index += n
	if b.index >= len(b.content) {
		return n, io.EOF
	}
	return n, nil
}

// Close 实现了 io.ReadCloser 接口中的 Close 方法。
func (b *ReadCloser) Close() error {
	return nil
}
