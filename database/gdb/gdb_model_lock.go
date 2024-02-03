// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// LockUpdate 为当前操作设置更新锁。
//
// LockUpdate/锁读写 与 LockShared/锁写入 都是用于确保被选中的记录值不能被其它事务更新（上锁），
// 两者的区别在于 LockShared/锁写入 不会阻塞其它事务读取被锁定行记录的值，
// 而 LockUpdate/锁读写 会阻塞其他锁定性读对锁定行的读取（非锁定性读仍然可以读取这些记录，LockShared/锁写入 和 LockUpdate/锁读写 都是锁定性读）。
//
// 例子:
// g.Model("users").Ctx(ctx).Where("votes>?", 100).LockUpdate().All();
// 上面这个查询等价于下面这条 SQL 语句：
// SELECT * FROM `users` WHERE `votes` > 100 FOR UPDATE
func (m *Model) LockUpdate() *Model {
	model := m.getModel()
	model.lockInfo = "FOR UPDATE"
	return model
}

// LockShared 将锁设置为当前操作的共享模式。
//
// LockUpdate/锁读写 与 LockShared/锁写入 都是用于确保被选中的记录值不能被其它事务更新（上锁），
// 两者的区别在于 LockShared/锁写入 不会阻塞其它事务读取被锁定行记录的值，
// 而 LockUpdate/锁读写 会阻塞其他锁定性读对锁定行的读取（非锁定性读仍然可以读取这些记录，LockShared/锁写入 和 LockUpdate/锁读写 都是锁定性读）。
//
// 例子:
// g.Model("users").Ctx(ctx).Where("votes>?", 100).LockShared().All();
// 上面这个查询等价于下面这条 SQL 语句：
// SELECT * FROM `users` WHERE `votes` > 100 LOCK IN SHARE MODE
func (m *Model) LockShared() *Model {
	model := m.getModel()
	model.lockInfo = "LOCK IN SHARE MODE"
	return model
}
