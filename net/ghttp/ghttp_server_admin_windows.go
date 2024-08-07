// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

//go:build windows
// +build windows

package http类

import (
	"context"
	"os"

	gproc "github.com/888go/goframe/os/gproc"
)

// handleProcessSignal 以阻塞方式处理来自系统的所有信号。 md5:822a0f0abd5d924a
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})

	gproc.Listen()
}
