// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
)

// Transaction 通过函数 `f` 包装事务逻辑。如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。若函数 `f` 返回空（nil）错误，它将提交事务并返回空。
// 注意：在函数 `f` 中不应手动调用 Commit 或 Rollback 方法处理事务，因为这些操作在此函数中已自动完成。
func (m *Model) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, 事务对象 TX) error) (错误 error) {
	if 上下文 == nil {
		上下文 = m.X取上下文对象()
	}
	if m.tx != nil {
		return m.tx.X事务(上下文, 回调函数)
	}
	return m.db.X事务(上下文, 回调函数)
}
