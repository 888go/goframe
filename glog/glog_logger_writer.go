// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"bytes"
	"context"
)

// Write 实现了 io.Writer 接口。
// 它只是使用 Print 打印内容。
func (l *Logger) Write(p []byte) (n int, err error) {
	l.Header(false).Print(context.TODO(), string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
