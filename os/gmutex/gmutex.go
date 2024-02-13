// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gmutex 继承并扩展了 sync.Mutex 和 sync.RWMutex，提供了更多的功能。
//
// 注意：从 GoFrame 版本 v2.5.2 开始，它是通过对标准库中 sync 包中的互斥锁进行重构实现的。
package 互斥锁类

// New 创建并返回一个新的互斥锁。
// 已弃用：请改用 Mutex 或 RWMutex。
func X创建() *RW互斥锁 {
	return &RW互斥锁{}
}
