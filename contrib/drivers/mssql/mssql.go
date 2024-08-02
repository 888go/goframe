// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包mssql实现了gdb.Driver，它支持MSSql数据库的操作。
//
// 注意：
// 1. 它不支持Replace功能。
// 2. 它不支持LastInsertId。
// md5:fb4f98bf77e784a6
package mssql

import (
	_ "github.com/denisenkom/go-mssqldb"

	gdb "github.com/888go/goframe/database/gdb"
)

// Driver是SQL服务器数据库的驱动程序。 md5:5c5e74aefaf2ae3d
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.Register(`mssql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver的驱动，该驱动支持Mssql的操作。 md5:4893e60c8841c569
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于SQL服务器的数据库对象。
// 它实现了gdb.Driver接口，以便于额外的数据库驱动安装。
// md5:ce3ce027533c5bb6
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
