// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
)

// X事务 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
func (m *Model) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, tx TX) error) (错误 error) {
	if 上下文 == nil {
		上下文 = m.X取上下文对象()
	}
	if m.tx != nil {
		return m.tx.X事务(上下文, 回调函数)
	}
	return m.db.X事务(上下文, 回调函数)
}
