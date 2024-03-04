// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtimer

// Exit 在内部被用于定时任务，它会退出并标记该任务从计时器中关闭。
// 之后，该定时任务将自动从计时器中移除。它内部使用了“panic-recover”机制
// 实现此功能，旨在简化并方便使用。
func Exit() {
	panic(panicExit)
}
