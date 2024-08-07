// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"bytes"
	"context"

	glog "github.com/888go/goframe/os/glog"
)

// errorLogger是底层net/http.Server的错误日志记录器。 md5:6405822f309730c2
type errorLogger struct {
	logger *glog.Logger
}

// Write 实现了 io.Writer 接口。 md5:6464c47cfa35b955
func (l *errorLogger) Write(p []byte) (n int, err error) {
	l.logger.X堆栈偏移量(1).Error(context.TODO(), string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
