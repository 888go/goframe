// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package clickhouse 实现了 gdb.Driver，它支持 ClickHouse 数据库的操作。 md5:8c421a92a767ff72
package clickhouse

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
)

// Driver 是用于 postgresql 数据库的驱动程序。 md5:4abf0752f49a3cfc
type Driver struct {
	*gdb.Core
}

var (
	errUnsupportedInsertIgnore = errors.New("unsupported method:InsertIgnore")
	errUnsupportedInsertGetId  = errors.New("unsupported method:InsertGetId")
	errUnsupportedReplace      = errors.New("unsupported method:Replace")
	errUnsupportedBegin        = errors.New("unsupported method:Begin")
	errUnsupportedTransaction  = errors.New("unsupported method:Transaction")
)

const (
	updateFilterPattern              = `(?i)UPDATE[\s]+?(\w+[\.]?\w+)[\s]+?SET`
	deleteFilterPattern              = `(?i)DELETE[\s]+?FROM[\s]+?(\w+[\.]?\w+)`
	filterTypePattern                = `(?i)^UPDATE|DELETE`
	replaceSchemaPattern             = `@(.+?)/([\w\.\-]+)+`
	needParsedSqlInCtx   gctx.StrKey = "NeedParsedSql"
	driverName                       = "clickhouse"
)

func init() {
	if err := gdb.Register(`clickhouse`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver接口的驱动器，该驱动器支持对Clickhouse的操作。 md5:e191d797c82bf046
// ff:
func New() gdb.Driver {
	return &Driver{}
}

		// New 创建并返回一个用于ClickHouse的数据库对象。它实现了gdb.Driver接口，以便于额外的数据库驱动程序安装。
		// md5:79dabf2eba06bc88
// ff:
// d:
// core:
// node:
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

func (d *Driver) injectNeedParsedSql(ctx context.Context) context.Context {
	if ctx.Value(needParsedSqlInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, needParsedSqlInCtx, true)
}
