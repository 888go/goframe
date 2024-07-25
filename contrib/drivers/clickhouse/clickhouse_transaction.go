// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package clickhouse

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

// Begin 方法开始并返回事务对象。 md5:8fa258c0123d6569
func (d *Driver) Begin(ctx context.Context) (tx gdb.TX, err error) {
	return nil, errUnsupportedBegin
}

// Transaction 使用函数 `f` 包裹事务逻辑。 md5:e6842034a9a21a66
func (d *Driver) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) error {
	return errUnsupportedTransaction
}
