// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包mysql实现了gdb.Driver，它支持MySQL数据库的操作。 md5:87c26760917d504d
package mysql

import (
	_ "github.com/go-sql-driver/mysql"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
)

// Driver 是用于 MySQL 数据库的驱动程序。 md5:db674980450242e3
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = "`"
)

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = gdb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New 创建并返回一个实现了gdb.Driver接口的驱动器，该驱动器支持对MySQL的操作。 md5:95ca0306bee3a521
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个针对 MySQL 的数据库对象。它实现了 gdb.Driver 接口，以便于额外的数据库驱动程序安装。
// md5:e61df629828efeff
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
