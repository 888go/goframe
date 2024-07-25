// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

// LockUpdate 为当前操作设置更新锁。 md5:ecffaffee1e7b1df
func (m *Model) LockUpdate() *Model {
	model := m.getModel()
	model.lockInfo = "FOR UPDATE"
	return model
}

// LockShared 将当前操作的锁设置为共享模式。 md5:d3afc426055403b9
func (m *Model) LockShared() *Model {
	model := m.getModel()
	model.lockInfo = "LOCK IN SHARE MODE"
	return model
}
