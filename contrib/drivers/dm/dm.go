// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包 dm 实现了 gdb.Driver，支持 DM 数据库的操作。 md5:7ee3e2c1c7faa19a
package dm

import (
	_ "gitee.com/chunanyong/dm"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"dm"}
	)
	for _, driverName := range driverNames {
		if err = gdb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New 创建并返回一个实现了gdb.Driver的驱动器，该驱动器支持dm的操作。 md5:eaafc794fe673f0d
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个针对 dm 的数据库对象。 md5:503a0529287d7ea6
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars 返回这种类型的数据库的安全字符。 md5:8a01432c4ed14729
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
