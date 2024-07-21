// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package pgsql implements gdb.Driver, which supports operations for database PostgreSQL.
//
// 1. It does not support Replace features.
// 2. It does not support Insert Ignore features.
package pgsql

import (
	_ "github.com/lib/pq"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
)

// Driver 是用于 postgresql 数据库的驱动程序。 md5:4abf0752f49a3cfc
type Driver struct {
	*gdb.Core
}

const (
	internalPrimaryKeyInCtx gctx.StrKey = "primary_key"
	defaultSchema           string      = "public"
	quoteChar               string      = `"`
)

func init() {
	if err := gdb.Register(`pgsql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver的驱动器，该驱动器支持PostgreSQL操作。 md5:183551a5c197dfc4
// ff:
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于postgresql的数据库对象。
// 它实现了gdb.Driver接口，以便安装额外的数据库驱动。
// md5:05f196cdca4e65a1
// ff:
// d:
// core:
// node:
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
// ff:底层取数据库安全字符
// d:
// charLeft:左字符
// charRight:右字符
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
