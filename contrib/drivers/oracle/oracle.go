// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package oracle 实现了 gdb.Driver，支持 Oracle 数据库的操作。
//
// 注意：
// 1. 它不支持 Save/Replace 功能。
// 2. 它不支持 LastInsertId。
// md5:898bf7c99e282e5b
package oracle

import (
	gdb "github.com/888go/goframe/database/gdb"
)

// Driver 是用于 Oracle 数据库的驱动程序。 md5:dc4dd4c6d7cda96f
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.X注册驱动(`oracle`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver接口的驱动器，该驱动器支持对Oracle数据库的操作。 md5:c6dcda5e6735bff7
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于Oracle的数据库对象。它实现了gdb.Driver接口，以便于额外的数据库驱动程序安装。
// md5:7f4f55e7c45290be
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// X底层取数据库安全字符 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
func (d *Driver) X底层取数据库安全字符() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
