// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

//go:build windows

package 进程类

import (
	"syscall"

	gstr "github.com/888go/goframe/text/gstr"
)

// 在Windows平台直接设置底层参数. md5:418ca44ebddf20f0
func joinProcessArgs(p *Process) {
	p.SysProcAttr = &syscall.SysProcAttr{}
	p.SysProcAttr.CmdLine = gstr.X连接(p.Args, " ")
}
