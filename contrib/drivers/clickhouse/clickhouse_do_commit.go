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

// DoCommit 将当前SQL和参数提交给底层SQL驱动程序。 md5:7cf9b1f6f4d9d2cb
func (d *Driver) DoCommit(ctx context.Context, in gdb.DoCommitInput) (out gdb.DoCommitOutput, err error) {
	ctx = d.InjectIgnoreResult(ctx)
	return d.Core.DoCommit(ctx, in)
}
