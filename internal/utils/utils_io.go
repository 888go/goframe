// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package utils

import (
	"io"
)

// ReadCloser 实现了 io.ReadCloser 接口，
// 该接口用于多次读取请求体内容。
//
// 注意，它不能被关闭。 md5:dc906d3f78dd2393
type ReadCloser struct {
	index      int    // Current read position.
	content    []byte // Content.
	repeatable bool   // 标记内容可以被重复读取。 md5:795fd563924f3c07
}

// NewReadCloser 创建并返回一个 RepeatReadCloser 对象。 md5:5b08470c05886c6e
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read implements the io.ReadCloser接口。 md5:a139bded6161151f
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	// 使其可重复读取。 md5:51c43c26e5e5404d
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

// Close 实现了 io.ReadCloser 接口。 md5:597e2e893ae16680
func (b *ReadCloser) Close() error {
	return nil
}
