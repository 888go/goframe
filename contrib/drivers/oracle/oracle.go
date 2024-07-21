// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package oracle implements gdb.Driver, which supports operations for database Oracle.
//
// 1. It does not support Save/Replace features.
// 2. It does not support LastInsertId.
package oracle

import (
	"github.com/gogf/gf/v2/database/gdb"
)

// Driver 是用于 Oracle 数据库的驱动程序。 md5:dc4dd4c6d7cda96f
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.Register(`oracle`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver接口的驱动器，该驱动器支持对Oracle数据库的操作。 md5:c6dcda5e6735bff7
// ff:
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于Oracle的数据库对象。它实现了gdb.Driver接口，以便于额外的数据库驱动程序安装。
// md5:7f4f55e7c45290be
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
// ff:
// d:
// charLeft:
// charRight:
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
