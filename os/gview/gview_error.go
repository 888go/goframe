// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gview

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
	// commandEnvKeyForErrorPrint 用于指定控制将错误打印到 stdout 的键。这个错误设计上不被函数返回。
	// md5:45fbe9796576f681
	commandEnvKeyForErrorPrint = "gf.gview.errorprint"
)

// errorPrint 检查是否将错误打印到标准输出。 md5:9791c350cd960b88
func errorPrint() bool {
	return gcmd.GetOptWithEnv(commandEnvKeyForErrorPrint, true).Bool()
}
