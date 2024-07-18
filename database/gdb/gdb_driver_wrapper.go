// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

// DriverWrapper是一个驱动程序包装器，用于通过嵌入式驱动程序扩展功能。 md5:7294ac58de5aa606
type DriverWrapper struct {
	driver Driver
}

// New 创建并返回一个针对 MySQL 的数据库对象。它实现了 gdb.Driver 接口，以便于额外的数据库驱动程序安装。
// md5:e61df629828efeff
// ff:
// d:
// core:
// node:
// DB:
func (d *DriverWrapper) New(core *Core, node *ConfigNode) (DB, error) {
	db, err := d.driver.New(core, node)
	if err != nil {
		return nil, err
	}
	return &DriverWrapperDB{
		DB: db,
	}, nil
}

// newDriverWrapper 创建并返回一个驱动包装器。 md5:4bc742bfe28b9706
func newDriverWrapper(driver Driver) Driver {
	return &DriverWrapper{
		driver: driver,
	}
}
