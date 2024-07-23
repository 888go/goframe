// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtimer

// Exit 在内部用于定时任务，它会退出并从计时器中标记为已关闭。后续定时任务会自动从计时器中移除。它内部使用了“panic-recover”机制来实现这个功能，这种设计旨在简化和方便操作。
// md5:f86628e24baaeeef
func Exit() {
	panic(panicExit)
}
