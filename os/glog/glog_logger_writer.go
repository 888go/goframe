// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"bytes"
	"context"
)

// Write 实现了 io.Writer 接口。
// 它只是使用 Print 函数打印内容。
// md5:da123f9fe994f171
func (l *Logger) Write(p []byte) (n int, err error) {
	l.Header(false).Print(context.TODO(), string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
