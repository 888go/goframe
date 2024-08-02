// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包sqlite实现了gdb.Driver，支持SQLite数据库的操作。 md5:52f270ca46090ebc
package sqlite

import (
	_ "github.com/glebarez/go-sqlite"

	gdb "github.com/888go/goframe/database/gdb"
)

// Driver是sqlite数据库的驱动程序。 md5:afc6ba2ecab1097e
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = "`"
)

func init() {
	if err := gdb.Register(`sqlite`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了gdb.Driver的驱动器，该驱动器支持SQLite操作。 md5:7fb5a6475b83117f
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于SQLite的数据库对象。
// 它实现了gdb.Driver接口，以便作为额外的数据库驱动程序进行安装。
// md5:27a088b6550263f7
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
