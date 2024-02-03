// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

//go:build windows
// +build windows

package ghttp

import (
	"context"
	"os"
	
	"github.com/888go/goframe/os/gproc"
)

// handleProcessSignal 以阻塞方式处理来自系统的所有信号。
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})

	gproc.Listen()
}
