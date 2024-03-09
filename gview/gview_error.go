// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
// commandEnvKeyForErrorPrint 用于指定控制错误信息打印到标准输出的键。
// 这种错误设计上不会被函数返回。
	commandEnvKeyForErrorPrint = "gf.gview.errorprint"
)

// errorPrint 检查是否将错误输出到标准输出（stdout）
func errorPrint() bool {
	return gcmd.GetOptWithEnv(commandEnvKeyForErrorPrint, true).Bool()
}
