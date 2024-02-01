// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"bytes"
	"context"
	
	"github.com/888go/goframe/os/glog"
	)
// errorLogger 是用于底层 net/http.Server 的错误日志记录器。
type errorLogger struct {
	logger *glog.Logger
}

// Write 实现了 io.Writer 接口。
func (l *errorLogger) Write(p []byte) (n int, err error) {
	l.logger.Skip(1).Error(context.TODO(), string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
